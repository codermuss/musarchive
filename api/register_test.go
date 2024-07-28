package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/mustafayilmazdev/musarchive/db/mock"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

// Custom matcher for comparing InsertUserParams with hashed password
type eqCreateUserParamsMatcher struct {
	arg      db.InsertUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.InsertUserParams)
	if !ok {
		return false
	}

	// Check if the password matches the hashed version
	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	var leftValue db.InsertUserParams = e.arg
	var rightValue db.InsertUserParams = arg
	result := reflect.DeepEqual(leftValue, rightValue)
	return result
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

// Helper function to create the custom matcher
func EqCreateUserParams(arg db.InsertUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

// Test case for creating a user via the API
func TestCreateUserAPI(t *testing.T) {
	user, password := RandomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":   user.Username,
				"password":   password,
				"full_name":  user.FullName,
				"email":      user.Email,
				"avatar":     user.Avatar,
				"role":       user.Role,
				"birth_date": user.BirthDate,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.InsertUserParams{
					Username:  user.Username,
					FullName:  user.FullName,
					Email:     user.Email,
					Avatar:    user.Avatar,
					BirthDate: user.BirthDate,
					Role:      user.Role,
				}
				store.EXPECT().
					InsertUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
				require.Equal(t, user.Role, util.Standard)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username":  "invalid-user#1",
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, db.ErrUniqueViolation)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"username":  user.Username,
				"password":  "123",
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	// Run each test case
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/auth/register"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

// Helper function to generate a random user for testing
func RandomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		Avatar: pgtype.Text{
			Valid:  true,
			String: util.RandomImage(),
		},
		Role: util.Standard,
		BirthDate: pgtype.Date{
			Valid: true,
			Time:  util.DateFixed(),
		},
		PasswordChangedAt: util.DateFixed(),
		CreatedAt:         util.DateFixed(),
	}
	return
}

// Helper function to compare response body with the expected user
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, expectedUser db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var baseResp BaseResponse
	err = json.Unmarshal(data, &baseResp)
	require.NoError(t, err)

	// Marshal the Data field back to JSON
	dataBytes, err := json.Marshal(baseResp.Data)
	require.NoError(t, err)

	// Unmarshal the Data field into a db.User
	var gotUser db.User
	err = json.Unmarshal(dataBytes, &gotUser)
	require.NoError(t, err)

	// Compare the expected and actual user fields
	require.Equal(t, expectedUser.Username, gotUser.Username)
	require.Equal(t, expectedUser.FullName, gotUser.FullName)
	require.Equal(t, expectedUser.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}

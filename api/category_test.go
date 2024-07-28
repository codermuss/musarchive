package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/mustafayilmazdev/musarchive/db/mock"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	user, _ := RandomUser(t)

	categories := []db.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCategories(gomock.Any()).
					Times(1).
					Return(categories, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		{
			name: "InternalError",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCategories(gomock.Any()).
					Times(1).Return([]db.Category{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/v1/categories/index"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchCategories(t *testing.T, body *bytes.Buffer, categories []db.Category) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response BaseResponse
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.Code)

	categoriesJSON, err := json.Marshal(categories)
	require.NoError(t, err)

	responseDataJSON, err := json.Marshal(response.Data)
	require.NoError(t, err)

	require.JSONEq(t, string(categoriesJSON), string(responseDataJSON))
}

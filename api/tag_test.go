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

func TestGetTags(t *testing.T) {
	user, _ := RandomUser(t)

	tags := []db.Tag{
		{ID: 1, Name: "Tag 1"},
		{ID: 2, Name: "Tag 2"},
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
					GetTags(gomock.Any()).
					Times(1).
					Return(tags, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTags(t, recorder.Body, tags)
			},
		},
		{
			name: "InternalError",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTags(gomock.Any()).
					Times(1).Return([]db.Tag{}, sql.ErrConnDone)
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

			server := newTestServer(t, store, nil)
			recorder := httptest.NewRecorder()

			url := "/v1/tags/index"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchTags(t *testing.T, body *bytes.Buffer, Tags []db.Tag) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response BaseResponse
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.Code)

	TagsJSON, err := json.Marshal(Tags)
	require.NoError(t, err)

	responseDataJSON, err := json.Marshal(response.Data)
	require.NoError(t, err)

	require.JSONEq(t, string(TagsJSON), string(responseDataJSON))
}

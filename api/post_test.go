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

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/mustafayilmazdev/musarchive/db/mock"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePostAPI(t *testing.T) {
	user, _ := RandomUser(t)
	post := randomPost(user.ID)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"cover_image": post.CoverImage,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.InsertPostParams{
					UserID: pgtype.Int4{
						Valid: true,
						Int32: int32(user.ID),
					},
					Title:      post.Title,
					Content:    post.Content,
					CoverImage: post.CoverImage,
				}

				store.EXPECT().InsertPost(gomock.Any(), gomock.Eq(arg)).Times(1).Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"cover_image": post.CoverImage,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().InsertPost(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"cover_image": post.CoverImage,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().InsertPost(gomock.Any(), gomock.Any()).Times(1).Return(db.Post{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidTitle",
			body: gin.H{
				"title":       "",
				"content":     post.Content,
				"cover_image": post.CoverImage,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().InsertPost(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store, nil)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/posts/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetPostsAPI(t *testing.T) {
	user, _ := RandomUser(t)
	posts := []db.Post{
		randomPost(user.ID),
		randomPost(user.ID),
		randomPost(user.ID),
	}

	testCases := []struct {
		name          string
		queryParams   map[string]string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			queryParams: map[string]string{
				"page":   "1",
				"size":   "10",
				"locale": "en",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},

			body: gin.H{
				"categories": []int32{1, 2},
				"tags":       []int32{3, 4},
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPostsWithFilterParams{
					Column3: []int32{1, 2},
					Column4: []int32{3, 4},
					Limit:   10,
					Offset:  0,
				}

				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Eq(arg)).Times(1).Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPosts(t, recorder.Body, posts)
			},
		},
		{
			name: "DEFAULT",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPostsWithFilterParams{
					Limit:  10,
					Offset: 0,
				}

				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Eq(arg)).Times(1).Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPosts(t, recorder.Body, posts)
			},
		},
		{
			name: "InternalError",
			queryParams: map[string]string{
				"page":   "1",
				"size":   "10",
				"locale": "en",
			},

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			body: gin.H{
				"categories": []int32{1, 2},
				"tags":       []int32{3, 4},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},

		{
			name: "InvalidPage",
			queryParams: map[string]string{
				"page":   "-1",
				"size":   "10",
				"locale": "en",
			},

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			body: gin.H{
				"categories": "invalid_categories",
				"tags":       "invalid_tags",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidSize",
			queryParams: map[string]string{
				"page":   "1",
				"size":   "-10",
				"locale": "en",
			},

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			body: gin.H{
				"categories": "invalid_categories",
				"tags":       "invalid_tags",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidBody",
			queryParams: map[string]string{
				"page":   "1",
				"size":   "10",
				"locale": "en",
			},

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, int(user.ID), user.Role, time.Minute)
			},
			body: gin.H{
				"categories": "invalid_categories",
				"tags":       "invalid_tags",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetPostsWithFilter(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store, nil)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/posts/index"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)

			// Add query parameters
			q := request.URL.Query()
			for k, v := range tc.queryParams {
				q.Add(k, v)
			}
			request.URL.RawQuery = q.Encode()
			request.Body = io.NopCloser(bytes.NewReader(data))
		})
	}
}

func requireBodyMatchPosts(t *testing.T, body *bytes.Buffer, posts []db.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response struct {
		Code    int       `json:"code"`
		Data    []db.Post `json:"data"`
		Message *string   `json:"message"`
	}
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, 200, response.Code)
	require.Equal(t, posts, response.Data)
}

func randomPost(id int32) db.Post {
	return db.Post{
		UserID: pgtype.Int4{
			Valid: true,
			Int32: id,
		},
		Title:      util.RandomString(6),
		Content:    util.RandomString(50),
		CoverImage: pgtype.Text{Valid: true, String: util.RandomImage()},
	}
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mustafayilmazdev/musarchive/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteFeaturedStory mocks base method.
func (m *MockStore) DeleteFeaturedStory(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFeaturedStory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFeaturedStory indicates an expected call of DeleteFeaturedStory.
func (mr *MockStoreMockRecorder) DeleteFeaturedStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFeaturedStory", reflect.TypeOf((*MockStore)(nil).DeleteFeaturedStory), arg0, arg1)
}

// DeleteOnboarding mocks base method.
func (m *MockStore) DeleteOnboarding(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOnboarding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOnboarding indicates an expected call of DeleteOnboarding.
func (mr *MockStoreMockRecorder) DeleteOnboarding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOnboarding", reflect.TypeOf((*MockStore)(nil).DeleteOnboarding), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockStore) DeletePost(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockStoreMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockStore)(nil).DeletePost), arg0, arg1)
}

// DeletePostCategory mocks base method.
func (m *MockStore) DeletePostCategory(arg0 context.Context, arg1 db.DeletePostCategoryParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostCategory indicates an expected call of DeletePostCategory.
func (mr *MockStoreMockRecorder) DeletePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostCategory", reflect.TypeOf((*MockStore)(nil).DeletePostCategory), arg0, arg1)
}

// DeletePostTag mocks base method.
func (m *MockStore) DeletePostTag(arg0 context.Context, arg1 db.DeletePostTagParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostTag indicates an expected call of DeletePostTag.
func (mr *MockStoreMockRecorder) DeletePostTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostTag", reflect.TypeOf((*MockStore)(nil).DeletePostTag), arg0, arg1)
}

// DeleteProfile mocks base method.
func (m *MockStore) DeleteProfile(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockStoreMockRecorder) DeleteProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockStore)(nil).DeleteProfile), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0, arg1)
}

// DeleteTag mocks base method.
func (m *MockStore) DeleteTag(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockStoreMockRecorder) DeleteTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockStore)(nil).DeleteTag), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserFollower mocks base method.
func (m *MockStore) DeleteUserFollower(arg0 context.Context, arg1 db.DeleteUserFollowerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserFollower", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserFollower indicates an expected call of DeleteUserFollower.
func (mr *MockStoreMockRecorder) DeleteUserFollower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserFollower", reflect.TypeOf((*MockStore)(nil).DeleteUserFollower), arg0, arg1)
}

// DeleteUserPost mocks base method.
func (m *MockStore) DeleteUserPost(arg0 context.Context, arg1 db.DeleteUserPostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserPost indicates an expected call of DeleteUserPost.
func (mr *MockStoreMockRecorder) DeleteUserPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPost", reflect.TypeOf((*MockStore)(nil).DeleteUserPost), arg0, arg1)
}

// GetCategories mocks base method.
func (m *MockStore) GetCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockStoreMockRecorder) GetCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockStore)(nil).GetCategories), arg0)
}

// GetCategoriesForPost mocks base method.
func (m *MockStore) GetCategoriesForPost(arg0 context.Context, arg1 int32) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesForPost", arg0, arg1)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesForPost indicates an expected call of GetCategoriesForPost.
func (mr *MockStoreMockRecorder) GetCategoriesForPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesForPost", reflect.TypeOf((*MockStore)(nil).GetCategoriesForPost), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int32) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetCommentsForPost mocks base method.
func (m *MockStore) GetCommentsForPost(arg0 context.Context, arg1 int32) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsForPost", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsForPost indicates an expected call of GetCommentsForPost.
func (mr *MockStoreMockRecorder) GetCommentsForPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsForPost", reflect.TypeOf((*MockStore)(nil).GetCommentsForPost), arg0, arg1)
}

// GetFeaturedStory mocks base method.
func (m *MockStore) GetFeaturedStory(arg0 context.Context, arg1 int32) (db.FeaturedStory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeaturedStory", arg0, arg1)
	ret0, _ := ret[0].(db.FeaturedStory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeaturedStory indicates an expected call of GetFeaturedStory.
func (mr *MockStoreMockRecorder) GetFeaturedStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeaturedStory", reflect.TypeOf((*MockStore)(nil).GetFeaturedStory), arg0, arg1)
}

// GetFollowedPosts mocks base method.
func (m *MockStore) GetFollowedPosts(arg0 context.Context, arg1 db.GetFollowedPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowedPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowedPosts indicates an expected call of GetFollowedPosts.
func (mr *MockStoreMockRecorder) GetFollowedPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowedPosts", reflect.TypeOf((*MockStore)(nil).GetFollowedPosts), arg0, arg1)
}

// GetFollowersOfUser mocks base method.
func (m *MockStore) GetFollowersOfUser(arg0 context.Context, arg1 int32) ([]db.GetFollowersOfUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowersOfUser", arg0, arg1)
	ret0, _ := ret[0].([]db.GetFollowersOfUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowersOfUser indicates an expected call of GetFollowersOfUser.
func (mr *MockStoreMockRecorder) GetFollowersOfUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowersOfUser", reflect.TypeOf((*MockStore)(nil).GetFollowersOfUser), arg0, arg1)
}

// GetFollowingUsers mocks base method.
func (m *MockStore) GetFollowingUsers(arg0 context.Context, arg1 int32) ([]db.GetFollowingUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowingUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.GetFollowingUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowingUsers indicates an expected call of GetFollowingUsers.
func (mr *MockStoreMockRecorder) GetFollowingUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowingUsers", reflect.TypeOf((*MockStore)(nil).GetFollowingUsers), arg0, arg1)
}

// GetOnboarding mocks base method.
func (m *MockStore) GetOnboarding(arg0 context.Context, arg1 int32) (db.Onboarding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnboarding", arg0, arg1)
	ret0, _ := ret[0].(db.Onboarding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnboarding indicates an expected call of GetOnboarding.
func (mr *MockStoreMockRecorder) GetOnboarding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnboarding", reflect.TypeOf((*MockStore)(nil).GetOnboarding), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockStore) GetPost(arg0 context.Context, arg1 int32) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockStoreMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStore)(nil).GetPost), arg0, arg1)
}

// GetPosts mocks base method.
func (m *MockStore) GetPosts(arg0 context.Context, arg1 db.GetPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockStoreMockRecorder) GetPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockStore)(nil).GetPosts), arg0, arg1)
}

// GetPostsWithFilter mocks base method.
func (m *MockStore) GetPostsWithFilter(arg0 context.Context, arg1 db.GetPostsWithFilterParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsWithFilter", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsWithFilter indicates an expected call of GetPostsWithFilter.
func (mr *MockStoreMockRecorder) GetPostsWithFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithFilter", reflect.TypeOf((*MockStore)(nil).GetPostsWithFilter), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockStore) GetProfile(arg0 context.Context, arg1 int32) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockStoreMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockStore)(nil).GetProfile), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetTag mocks base method.
func (m *MockStore) GetTag(arg0 context.Context, arg1 int32) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockStoreMockRecorder) GetTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockStore)(nil).GetTag), arg0, arg1)
}

// GetTags mocks base method.
func (m *MockStore) GetTags(arg0 context.Context) ([]db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0)
	ret0, _ := ret[0].([]db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockStoreMockRecorder) GetTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockStore)(nil).GetTags), arg0)
}

// GetTagsForPost mocks base method.
func (m *MockStore) GetTagsForPost(arg0 context.Context, arg1 int32) ([]db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsForPost", arg0, arg1)
	ret0, _ := ret[0].([]db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsForPost indicates an expected call of GetTagsForPost.
func (mr *MockStoreMockRecorder) GetTagsForPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsForPost", reflect.TypeOf((*MockStore)(nil).GetTagsForPost), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserPost mocks base method.
func (m *MockStore) GetUserPost(arg0 context.Context, arg1 db.GetUserPostParams) (db.GetUserPostRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPost", arg0, arg1)
	ret0, _ := ret[0].(db.GetUserPostRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPost indicates an expected call of GetUserPost.
func (mr *MockStoreMockRecorder) GetUserPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPost", reflect.TypeOf((*MockStore)(nil).GetUserPost), arg0, arg1)
}

// GetUserPosts mocks base method.
func (m *MockStore) GetUserPosts(arg0 context.Context, arg1 int32) ([]db.GetUserPostsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.GetUserPostsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPosts indicates an expected call of GetUserPosts.
func (mr *MockStoreMockRecorder) GetUserPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPosts", reflect.TypeOf((*MockStore)(nil).GetUserPosts), arg0, arg1)
}

// InsertCategory mocks base method.
func (m *MockStore) InsertCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCategory indicates an expected call of InsertCategory.
func (mr *MockStoreMockRecorder) InsertCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCategory", reflect.TypeOf((*MockStore)(nil).InsertCategory), arg0, arg1)
}

// InsertComment mocks base method.
func (m *MockStore) InsertComment(arg0 context.Context, arg1 db.InsertCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertComment indicates an expected call of InsertComment.
func (mr *MockStoreMockRecorder) InsertComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertComment", reflect.TypeOf((*MockStore)(nil).InsertComment), arg0, arg1)
}

// InsertFeaturedStory mocks base method.
func (m *MockStore) InsertFeaturedStory(arg0 context.Context, arg1 db.InsertFeaturedStoryParams) (db.FeaturedStory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFeaturedStory", arg0, arg1)
	ret0, _ := ret[0].(db.FeaturedStory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertFeaturedStory indicates an expected call of InsertFeaturedStory.
func (mr *MockStoreMockRecorder) InsertFeaturedStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFeaturedStory", reflect.TypeOf((*MockStore)(nil).InsertFeaturedStory), arg0, arg1)
}

// InsertOnboarding mocks base method.
func (m *MockStore) InsertOnboarding(arg0 context.Context, arg1 db.InsertOnboardingParams) (db.Onboarding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOnboarding", arg0, arg1)
	ret0, _ := ret[0].(db.Onboarding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOnboarding indicates an expected call of InsertOnboarding.
func (mr *MockStoreMockRecorder) InsertOnboarding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOnboarding", reflect.TypeOf((*MockStore)(nil).InsertOnboarding), arg0, arg1)
}

// InsertPost mocks base method.
func (m *MockStore) InsertPost(arg0 context.Context, arg1 db.InsertPostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPost indicates an expected call of InsertPost.
func (mr *MockStoreMockRecorder) InsertPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPost", reflect.TypeOf((*MockStore)(nil).InsertPost), arg0, arg1)
}

// InsertPostCategory mocks base method.
func (m *MockStore) InsertPostCategory(arg0 context.Context, arg1 db.InsertPostCategoryParams) (db.PostCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPostCategory indicates an expected call of InsertPostCategory.
func (mr *MockStoreMockRecorder) InsertPostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPostCategory", reflect.TypeOf((*MockStore)(nil).InsertPostCategory), arg0, arg1)
}

// InsertPostTag mocks base method.
func (m *MockStore) InsertPostTag(arg0 context.Context, arg1 db.InsertPostTagParams) (db.PostTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPostTag", arg0, arg1)
	ret0, _ := ret[0].(db.PostTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPostTag indicates an expected call of InsertPostTag.
func (mr *MockStoreMockRecorder) InsertPostTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPostTag", reflect.TypeOf((*MockStore)(nil).InsertPostTag), arg0, arg1)
}

// InsertProfile mocks base method.
func (m *MockStore) InsertProfile(arg0 context.Context, arg1 db.InsertProfileParams) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProfile indicates an expected call of InsertProfile.
func (mr *MockStoreMockRecorder) InsertProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProfile", reflect.TypeOf((*MockStore)(nil).InsertProfile), arg0, arg1)
}

// InsertSession mocks base method.
func (m *MockStore) InsertSession(arg0 context.Context, arg1 db.InsertSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertSession indicates an expected call of InsertSession.
func (mr *MockStoreMockRecorder) InsertSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSession", reflect.TypeOf((*MockStore)(nil).InsertSession), arg0, arg1)
}

// InsertTag mocks base method.
func (m *MockStore) InsertTag(arg0 context.Context, arg1 string) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTag", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTag indicates an expected call of InsertTag.
func (mr *MockStoreMockRecorder) InsertTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTag", reflect.TypeOf((*MockStore)(nil).InsertTag), arg0, arg1)
}

// InsertUser mocks base method.
func (m *MockStore) InsertUser(arg0 context.Context, arg1 db.InsertUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockStoreMockRecorder) InsertUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockStore)(nil).InsertUser), arg0, arg1)
}

// InsertUserFollower mocks base method.
func (m *MockStore) InsertUserFollower(arg0 context.Context, arg1 db.InsertUserFollowerParams) (db.UserFollower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserFollower", arg0, arg1)
	ret0, _ := ret[0].(db.UserFollower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserFollower indicates an expected call of InsertUserFollower.
func (mr *MockStoreMockRecorder) InsertUserFollower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserFollower", reflect.TypeOf((*MockStore)(nil).InsertUserFollower), arg0, arg1)
}

// InsertUserPost mocks base method.
func (m *MockStore) InsertUserPost(arg0 context.Context, arg1 db.InsertUserPostParams) (db.UserPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserPost", arg0, arg1)
	ret0, _ := ret[0].(db.UserPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserPost indicates an expected call of InsertUserPost.
func (mr *MockStoreMockRecorder) InsertUserPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserPost", reflect.TypeOf((*MockStore)(nil).InsertUserPost), arg0, arg1)
}

// ListOnboarding mocks base method.
func (m *MockStore) ListOnboarding(arg0 context.Context) ([]db.Onboarding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOnboarding", arg0)
	ret0, _ := ret[0].([]db.Onboarding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOnboarding indicates an expected call of ListOnboarding.
func (mr *MockStoreMockRecorder) ListOnboarding(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOnboarding", reflect.TypeOf((*MockStore)(nil).ListOnboarding), arg0)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateFeaturedStory mocks base method.
func (m *MockStore) UpdateFeaturedStory(arg0 context.Context, arg1 db.UpdateFeaturedStoryParams) (db.FeaturedStory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeaturedStory", arg0, arg1)
	ret0, _ := ret[0].(db.FeaturedStory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFeaturedStory indicates an expected call of UpdateFeaturedStory.
func (mr *MockStoreMockRecorder) UpdateFeaturedStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeaturedStory", reflect.TypeOf((*MockStore)(nil).UpdateFeaturedStory), arg0, arg1)
}

// UpdateOnboarding mocks base method.
func (m *MockStore) UpdateOnboarding(arg0 context.Context, arg1 db.UpdateOnboardingParams) (db.Onboarding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOnboarding", arg0, arg1)
	ret0, _ := ret[0].(db.Onboarding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOnboarding indicates an expected call of UpdateOnboarding.
func (mr *MockStoreMockRecorder) UpdateOnboarding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOnboarding", reflect.TypeOf((*MockStore)(nil).UpdateOnboarding), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockStore) UpdateProfile(arg0 context.Context, arg1 db.UpdateProfileParams) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockStoreMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockStore)(nil).UpdateProfile), arg0, arg1)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(arg0 context.Context, arg1 db.UpdateSessionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0, arg1)
}

// UpdateTag mocks base method.
func (m *MockStore) UpdateTag(arg0 context.Context, arg1 db.UpdateTagParams) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTag indicates an expected call of UpdateTag.
func (mr *MockStoreMockRecorder) UpdateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockStore)(nil).UpdateTag), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

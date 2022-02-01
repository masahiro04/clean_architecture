// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecases/INTERACTOR.go

// Package mock is a generated GoMock package.
package mock

import (
	domains "clean_architecture/golang/domains"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogger) Log(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log.
func (mr *MockLoggerMockRecorder) Log(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), arg0...)
}

// MockPresenter is a mock of Presenter interface.
type MockPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockPresenterMockRecorder
}

// MockPresenterMockRecorder is the mock recorder for MockPresenter.
type MockPresenterMockRecorder struct {
	mock *MockPresenter
}

// NewMockPresenter creates a new mock instance.
func NewMockPresenter(ctrl *gomock.Controller) *MockPresenter {
	mock := &MockPresenter{ctrl: ctrl}
	mock.recorder = &MockPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresenter) EXPECT() *MockPresenterMockRecorder {
	return m.recorder
}

// CreateBlog mocks base method.
func (m *MockPresenter) CreateBlog(blog *domains.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBlog", blog)
}

// CreateBlog indicates an expected call of CreateBlog.
func (mr *MockPresenterMockRecorder) CreateBlog(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlog", reflect.TypeOf((*MockPresenter)(nil).CreateBlog), blog)
}

// GetBlog mocks base method.
func (m *MockPresenter) GetBlog(blog *domains.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBlog", blog)
}

// GetBlog indicates an expected call of GetBlog.
func (mr *MockPresenterMockRecorder) GetBlog(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlog", reflect.TypeOf((*MockPresenter)(nil).GetBlog), blog)
}

// GetBlogs mocks base method.
func (m *MockPresenter) GetBlogs(blogs *domains.Blogs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBlogs", blogs)
}

// GetBlogs indicates an expected call of GetBlogs.
func (mr *MockPresenterMockRecorder) GetBlogs(blogs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogs", reflect.TypeOf((*MockPresenter)(nil).GetBlogs), blogs)
}

// GetUser mocks base method.
func (m *MockPresenter) GetUser(user *domains.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", user)
}

// GetUser indicates an expected call of GetUser.
func (mr *MockPresenterMockRecorder) GetUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPresenter)(nil).GetUser), user)
}

// GetUsers mocks base method.
func (m *MockPresenter) GetUsers(users *domains.Users) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", users)
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockPresenterMockRecorder) GetUsers(users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockPresenter)(nil).GetUsers), users)
}

// Present mocks base method.
func (m *MockPresenter) Present() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Present")
	ret0, _ := ret[0].(error)
	return ret0
}

// Present indicates an expected call of Present.
func (mr *MockPresenterMockRecorder) Present() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Present", reflect.TypeOf((*MockPresenter)(nil).Present))
}

// Raise mocks base method.
func (m *MockPresenter) Raise(errorKind domains.ErrorKinds, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Raise", errorKind, err)
}

// Raise indicates an expected call of Raise.
func (mr *MockPresenterMockRecorder) Raise(errorKind, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raise", reflect.TypeOf((*MockPresenter)(nil).Raise), errorKind, err)
}

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidator) Validate(targetStruct interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", targetStruct)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorMockRecorder) Validate(targetStruct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), targetStruct)
}

// MockDBTransaction is a mock of DBTransaction interface.
type MockDBTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockDBTransactionMockRecorder
}

// MockDBTransactionMockRecorder is the mock recorder for MockDBTransaction.
type MockDBTransactionMockRecorder struct {
	mock *MockDBTransaction
}

// NewMockDBTransaction creates a new mock instance.
func NewMockDBTransaction(ctrl *gomock.Controller) *MockDBTransaction {
	mock := &MockDBTransaction{ctrl: ctrl}
	mock.recorder = &MockDBTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBTransaction) EXPECT() *MockDBTransactionMockRecorder {
	return m.recorder
}

// WithTx mocks base method.
func (m *MockDBTransaction) WithTx(runner func(*sql.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", runner)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockDBTransactionMockRecorder) WithTx(runner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockDBTransaction)(nil).WithTx), runner)
}

// MockBlogDao is a mock of BlogDao interface.
type MockBlogDao struct {
	ctrl     *gomock.Controller
	recorder *MockBlogDaoMockRecorder
}

// MockBlogDaoMockRecorder is the mock recorder for MockBlogDao.
type MockBlogDaoMockRecorder struct {
	mock *MockBlogDao
}

// NewMockBlogDao creates a new mock instance.
func NewMockBlogDao(ctrl *gomock.Controller) *MockBlogDao {
	mock := &MockBlogDao{ctrl: ctrl}
	mock.recorder = &MockBlogDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogDao) EXPECT() *MockBlogDaoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlogDao) Create(blog domains.Blog) (*domains.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", blog)
	ret0, _ := ret[0].(*domains.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlogDaoMockRecorder) Create(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlogDao)(nil).Create), blog)
}

// CreateTx mocks base method.
func (m *MockBlogDao) CreateTx(company domains.Blog, tx *sql.Tx) (*domains.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", company, tx)
	ret0, _ := ret[0].(*domains.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockBlogDaoMockRecorder) CreateTx(company, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockBlogDao)(nil).CreateTx), company, tx)
}

// Delete mocks base method.
func (m *MockBlogDao) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlogDaoMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlogDao)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockBlogDao) GetAll() (*domains.Blogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*domains.Blogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBlogDaoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBlogDao)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockBlogDao) GetById(id int) (*domains.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*domains.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBlogDaoMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBlogDao)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockBlogDao) Update(id int, company domains.Blog) (*domains.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, company)
	ret0, _ := ret[0].(*domains.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlogDaoMockRecorder) Update(id, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlogDao)(nil).Update), id, company)
}

// MockUserDao is a mock of UserDao interface.
type MockUserDao struct {
	ctrl     *gomock.Controller
	recorder *MockUserDaoMockRecorder
}

// MockUserDaoMockRecorder is the mock recorder for MockUserDao.
type MockUserDaoMockRecorder struct {
	mock *MockUserDao
}

// NewMockUserDao creates a new mock instance.
func NewMockUserDao(ctrl *gomock.Controller) *MockUserDao {
	mock := &MockUserDao{ctrl: ctrl}
	mock.recorder = &MockUserDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDao) EXPECT() *MockUserDaoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserDao) Create(user domains.User) (*domains.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*domains.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserDaoMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDao)(nil).Create), user)
}

// CreateTx mocks base method.
func (m *MockUserDao) CreateTx(user domains.User, tx *sql.Tx) (*domains.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", user, tx)
	ret0, _ := ret[0].(*domains.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockUserDaoMockRecorder) CreateTx(user, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockUserDao)(nil).CreateTx), user, tx)
}

// Delete mocks base method.
func (m *MockUserDao) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserDaoMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDao)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockUserDao) GetAll() (*domains.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*domains.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserDaoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserDao)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockUserDao) GetById(id int) (*domains.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*domains.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserDaoMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserDao)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockUserDao) Update(id int, user domains.User) (*domains.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, user)
	ret0, _ := ret[0].(*domains.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserDaoMockRecorder) Update(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDao)(nil).Update), id, user)
}

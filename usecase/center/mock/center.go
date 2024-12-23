// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/center/interface.go

// Package mock_center is a generated GoMock package.
package mock_center

import (
	reflect "reflect"
	entity "sudhagar/glad/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReader) Get(id entity.ID) (*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), id)
}

// GetCount mocks base method.
func (m *MockReader) GetCount(id entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockReaderMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockReader)(nil).GetCount), id)
}

// List mocks base method.
func (m *MockReader) List(tenantID entity.ID, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List), tenantID, page, limit)
}

// Search mocks base method.
func (m *MockReader) Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockReaderMockRecorder) Search(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockReader)(nil).Search), tenantID, query, page, limit)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriter) Create(e *entity.Center) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockWriter) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), id)
}

// Update mocks base method.
func (m *MockWriter) Update(e *entity.Center) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWriterMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriter)(nil).Update), e)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(e *entity.Center) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockRepository) Get(id entity.ID) (*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount(id entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount), id)
}

// List mocks base method.
func (m *MockRepository) List(tenantID entity.ID, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), tenantID, page, limit)
}

// Search mocks base method.
func (m *MockRepository) Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepositoryMockRecorder) Search(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), tenantID, query, page, limit)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Center) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateCenter mocks base method.
func (m *MockUseCase) CreateCenter(tenantID entity.ID, extID, extName, name string, mode entity.CenterMode, isEnabled bool) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCenter", tenantID, extID, extName, name, mode, isEnabled)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCenter indicates an expected call of CreateCenter.
func (mr *MockUseCaseMockRecorder) CreateCenter(tenantID, extID, extName, name, mode, isEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCenter", reflect.TypeOf((*MockUseCase)(nil).CreateCenter), tenantID, extID, extName, name, mode, isEnabled)
}

// DeleteCenter mocks base method.
func (m *MockUseCase) DeleteCenter(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCenter", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCenter indicates an expected call of DeleteCenter.
func (mr *MockUseCaseMockRecorder) DeleteCenter(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCenter", reflect.TypeOf((*MockUseCase)(nil).DeleteCenter), id)
}

// GetCenter mocks base method.
func (m *MockUseCase) GetCenter(id entity.ID) (*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCenter", id)
	ret0, _ := ret[0].(*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCenter indicates an expected call of GetCenter.
func (mr *MockUseCaseMockRecorder) GetCenter(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCenter", reflect.TypeOf((*MockUseCase)(nil).GetCenter), id)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount(id entity.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCount indicates an expected call of GetCount.
func (mr *MockUseCaseMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockUseCase)(nil).GetCount), id)
}

// ListCenters mocks base method.
func (m *MockUseCase) ListCenters(tenantID entity.ID, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCenters", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCenters indicates an expected call of ListCenters.
func (mr *MockUseCaseMockRecorder) ListCenters(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCenters", reflect.TypeOf((*MockUseCase)(nil).ListCenters), tenantID, page, limit)
}

// SearchCenters mocks base method.
func (m *MockUseCase) SearchCenters(tenantID entity.ID, query string, page, limit int) ([]*entity.Center, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCenters", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Center)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCenters indicates an expected call of SearchCenters.
func (mr *MockUseCaseMockRecorder) SearchCenters(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCenters", reflect.TypeOf((*MockUseCase)(nil).SearchCenters), tenantID, query, page, limit)
}

// UpdateCenter mocks base method.
func (m *MockUseCase) UpdateCenter(e *entity.Center) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCenter", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCenter indicates an expected call of UpdateCenter.
func (mr *MockUseCaseMockRecorder) UpdateCenter(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCenter", reflect.TypeOf((*MockUseCase)(nil).UpdateCenter), e)
}

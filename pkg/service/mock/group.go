// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockGroupService is a mock of GroupService interface.
type MockGroupService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceMockRecorder
}

// MockGroupServiceMockRecorder is the mock recorder for MockGroupService.
type MockGroupServiceMockRecorder struct {
	mock *MockGroupService
}

// NewMockGroupService creates a new mock instance.
func NewMockGroupService(ctrl *gomock.Controller) *MockGroupService {
	mock := &MockGroupService{ctrl: ctrl}
	mock.recorder = &MockGroupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupService) EXPECT() *MockGroupServiceMockRecorder {
	return m.recorder
}

// BulkCreateSubjectMembersWithTx mocks base method.
func (m *MockGroupService) BulkCreateSubjectMembersWithTx(tx *sqlx.Tx, parentPK int64, relations []types.SubjectRelation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateSubjectMembersWithTx", tx, parentPK, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateSubjectMembersWithTx indicates an expected call of BulkCreateSubjectMembersWithTx.
func (mr *MockGroupServiceMockRecorder) BulkCreateSubjectMembersWithTx(tx, parentPK, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateSubjectMembersWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkCreateSubjectMembersWithTx), tx, parentPK, relations)
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockGroupService) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, pks []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, pks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockGroupServiceMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteBySubjectPKsWithTx), tx, pks)
}

// BulkDeleteSubjectMembers mocks base method.
func (m *MockGroupService) BulkDeleteSubjectMembers(parentPK int64, userPKs, departmentPKs []int64) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjectMembers", parentPK, userPKs, departmentPKs)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteSubjectMembers indicates an expected call of BulkDeleteSubjectMembers.
func (mr *MockGroupServiceMockRecorder) BulkDeleteSubjectMembers(parentPK, userPKs, departmentPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjectMembers", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteSubjectMembers), parentPK, userPKs, departmentPKs)
}

// GetEffectThinSubjectGroups mocks base method.
func (m *MockGroupService) GetEffectThinSubjectGroups(pk int64) ([]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEffectThinSubjectGroups", pk)
	ret0, _ := ret[0].([]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEffectThinSubjectGroups indicates an expected call of GetEffectThinSubjectGroups.
func (mr *MockGroupServiceMockRecorder) GetEffectThinSubjectGroups(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEffectThinSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).GetEffectThinSubjectGroups), pk)
}

// GetMemberCount mocks base method.
func (m *MockGroupService) GetMemberCount(parentPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCount", parentPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCount indicates an expected call of GetMemberCount.
func (mr *MockGroupServiceMockRecorder) GetMemberCount(parentPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCount", reflect.TypeOf((*MockGroupService)(nil).GetMemberCount), parentPK)
}

// GetMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetMemberCountBeforeExpiredAt(parentPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCountBeforeExpiredAt", parentPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCountBeforeExpiredAt indicates an expected call of GetMemberCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetMemberCountBeforeExpiredAt(parentPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetMemberCountBeforeExpiredAt), parentPK, expiredAt)
}

// ListEffectThinSubjectGroups mocks base method.
func (m *MockGroupService) ListEffectThinSubjectGroups(pks []int64) (map[int64][]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroups", pks)
	ret0, _ := ret[0].(map[int64][]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroups indicates an expected call of ListEffectThinSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListEffectThinSubjectGroups(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListEffectThinSubjectGroups), pks)
}

// ListExistSubjectsBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListExistSubjectsBeforeExpiredAt(parentPKs []int64, expiredAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExistSubjectsBeforeExpiredAt", parentPKs, expiredAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExistSubjectsBeforeExpiredAt indicates an expected call of ListExistSubjectsBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListExistSubjectsBeforeExpiredAt(parentPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExistSubjectsBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListExistSubjectsBeforeExpiredAt), parentPKs, expiredAt)
}

// ListMember mocks base method.
func (m *MockGroupService) ListMember(parentPK int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMember", parentPK)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMember indicates an expected call of ListMember.
func (mr *MockGroupServiceMockRecorder) ListMember(parentPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMember", reflect.TypeOf((*MockGroupService)(nil).ListMember), parentPK)
}

// ListPagingMember mocks base method.
func (m *MockGroupService) ListPagingMember(parentPK, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMember", parentPK, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMember indicates an expected call of ListPagingMember.
func (mr *MockGroupServiceMockRecorder) ListPagingMember(parentPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMember", reflect.TypeOf((*MockGroupService)(nil).ListPagingMember), parentPK, limit, offset)
}

// ListPagingMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListPagingMemberBeforeExpiredAt(parentPK, expiredAt, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMemberBeforeExpiredAt", parentPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMemberBeforeExpiredAt indicates an expected call of ListPagingMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListPagingMemberBeforeExpiredAt(parentPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListPagingMemberBeforeExpiredAt), parentPK, expiredAt, limit, offset)
}

// ListSubjectGroups mocks base method.
func (m *MockGroupService) ListSubjectGroups(subjectPK, beforeExpiredAt int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectGroups", subjectPK, beforeExpiredAt)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectGroups indicates an expected call of ListSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListSubjectGroups(subjectPK, beforeExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListSubjectGroups), subjectPK, beforeExpiredAt)
}

// UpdateMembersExpiredAtWithTx mocks base method.
func (m *MockGroupService) UpdateMembersExpiredAtWithTx(tx *sqlx.Tx, parentPK int64, members []types.SubjectRelationPKPolicyExpiredAt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMembersExpiredAtWithTx", tx, parentPK, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMembersExpiredAtWithTx indicates an expected call of UpdateMembersExpiredAtWithTx.
func (mr *MockGroupServiceMockRecorder) UpdateMembersExpiredAtWithTx(tx, parentPK, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMembersExpiredAtWithTx", reflect.TypeOf((*MockGroupService)(nil).UpdateMembersExpiredAtWithTx), tx, parentPK, members)
}
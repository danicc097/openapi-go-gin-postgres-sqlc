// Code generated by counterfeiter. DO NOT EDIT.
package repostesting

import (
	"context"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

type FakeUser struct {
	ByAPIKeyStub        func(context.Context, db.DBTX, string) (*db.User, error)
	byAPIKeyMutex       sync.RWMutex
	byAPIKeyArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	byAPIKeyReturns struct {
		result1 *db.User
		result2 error
	}
	byAPIKeyReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	ByEmailStub        func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)
	byEmailMutex       sync.RWMutex
	byEmailArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}
	byEmailReturns struct {
		result1 *db.User
		result2 error
	}
	byEmailReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	ByExternalIDStub        func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)
	byExternalIDMutex       sync.RWMutex
	byExternalIDArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}
	byExternalIDReturns struct {
		result1 *db.User
		result2 error
	}
	byExternalIDReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	ByIDStub        func(context.Context, db.DBTX, uuid.UUID, ...db.UserSelectConfigOption) (*db.User, error)
	byIDMutex       sync.RWMutex
	byIDArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
		arg4 []db.UserSelectConfigOption
	}
	byIDReturns struct {
		result1 *db.User
		result2 error
	}
	byIDReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	ByUsernameStub        func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)
	byUsernameMutex       sync.RWMutex
	byUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}
	byUsernameReturns struct {
		result1 *db.User
		result2 error
	}
	byUsernameReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	CreateStub        func(context.Context, db.DBTX, *db.UserCreateParams) (*db.User, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 *db.UserCreateParams
	}
	createReturns struct {
		result1 *db.User
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	CreateAPIKeyStub        func(context.Context, db.DBTX, *db.User) (*db.UserAPIKey, error)
	createAPIKeyMutex       sync.RWMutex
	createAPIKeyArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 *db.User
	}
	createAPIKeyReturns struct {
		result1 *db.UserAPIKey
		result2 error
	}
	createAPIKeyReturnsOnCall map[int]struct {
		result1 *db.UserAPIKey
		result2 error
	}
	DeleteStub        func(context.Context, db.DBTX, uuid.UUID) (*db.User, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
	}
	deleteReturns struct {
		result1 *db.User
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UpdateStub        func(context.Context, db.DBTX, uuid.UUID, *db.UserUpdateParams) (*db.User, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
		arg4 *db.UserUpdateParams
	}
	updateReturns struct {
		result1 *db.User
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUser) ByAPIKey(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.byAPIKeyMutex.Lock()
	ret, specificReturn := fake.byAPIKeyReturnsOnCall[len(fake.byAPIKeyArgsForCall)]
	fake.byAPIKeyArgsForCall = append(fake.byAPIKeyArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ByAPIKeyStub
	fakeReturns := fake.byAPIKeyReturns
	fake.recordInvocation("ByAPIKey", []interface{}{arg1, arg2, arg3})
	fake.byAPIKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) ByAPIKeyCallCount() int {
	fake.byAPIKeyMutex.RLock()
	defer fake.byAPIKeyMutex.RUnlock()
	return len(fake.byAPIKeyArgsForCall)
}

func (fake *FakeUser) ByAPIKeyCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.byAPIKeyMutex.Lock()
	defer fake.byAPIKeyMutex.Unlock()
	fake.ByAPIKeyStub = stub
}

func (fake *FakeUser) ByAPIKeyArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.byAPIKeyMutex.RLock()
	defer fake.byAPIKeyMutex.RUnlock()
	argsForCall := fake.byAPIKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) ByAPIKeyReturns(result1 *db.User, result2 error) {
	fake.byAPIKeyMutex.Lock()
	defer fake.byAPIKeyMutex.Unlock()
	fake.ByAPIKeyStub = nil
	fake.byAPIKeyReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByAPIKeyReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.byAPIKeyMutex.Lock()
	defer fake.byAPIKeyMutex.Unlock()
	fake.ByAPIKeyStub = nil
	if fake.byAPIKeyReturnsOnCall == nil {
		fake.byAPIKeyReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.byAPIKeyReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByEmail(arg1 context.Context, arg2 db.DBTX, arg3 string, arg4 ...db.UserSelectConfigOption) (*db.User, error) {
	fake.byEmailMutex.Lock()
	ret, specificReturn := fake.byEmailReturnsOnCall[len(fake.byEmailArgsForCall)]
	fake.byEmailArgsForCall = append(fake.byEmailArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ByEmailStub
	fakeReturns := fake.byEmailReturns
	fake.recordInvocation("ByEmail", []interface{}{arg1, arg2, arg3, arg4})
	fake.byEmailMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) ByEmailCallCount() int {
	fake.byEmailMutex.RLock()
	defer fake.byEmailMutex.RUnlock()
	return len(fake.byEmailArgsForCall)
}

func (fake *FakeUser) ByEmailCalls(stub func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)) {
	fake.byEmailMutex.Lock()
	defer fake.byEmailMutex.Unlock()
	fake.ByEmailStub = stub
}

func (fake *FakeUser) ByEmailArgsForCall(i int) (context.Context, db.DBTX, string, []db.UserSelectConfigOption) {
	fake.byEmailMutex.RLock()
	defer fake.byEmailMutex.RUnlock()
	argsForCall := fake.byEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUser) ByEmailReturns(result1 *db.User, result2 error) {
	fake.byEmailMutex.Lock()
	defer fake.byEmailMutex.Unlock()
	fake.ByEmailStub = nil
	fake.byEmailReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByEmailReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.byEmailMutex.Lock()
	defer fake.byEmailMutex.Unlock()
	fake.ByEmailStub = nil
	if fake.byEmailReturnsOnCall == nil {
		fake.byEmailReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.byEmailReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByExternalID(arg1 context.Context, arg2 db.DBTX, arg3 string, arg4 ...db.UserSelectConfigOption) (*db.User, error) {
	fake.byExternalIDMutex.Lock()
	ret, specificReturn := fake.byExternalIDReturnsOnCall[len(fake.byExternalIDArgsForCall)]
	fake.byExternalIDArgsForCall = append(fake.byExternalIDArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ByExternalIDStub
	fakeReturns := fake.byExternalIDReturns
	fake.recordInvocation("ByExternalID", []interface{}{arg1, arg2, arg3, arg4})
	fake.byExternalIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) ByExternalIDCallCount() int {
	fake.byExternalIDMutex.RLock()
	defer fake.byExternalIDMutex.RUnlock()
	return len(fake.byExternalIDArgsForCall)
}

func (fake *FakeUser) ByExternalIDCalls(stub func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)) {
	fake.byExternalIDMutex.Lock()
	defer fake.byExternalIDMutex.Unlock()
	fake.ByExternalIDStub = stub
}

func (fake *FakeUser) ByExternalIDArgsForCall(i int) (context.Context, db.DBTX, string, []db.UserSelectConfigOption) {
	fake.byExternalIDMutex.RLock()
	defer fake.byExternalIDMutex.RUnlock()
	argsForCall := fake.byExternalIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUser) ByExternalIDReturns(result1 *db.User, result2 error) {
	fake.byExternalIDMutex.Lock()
	defer fake.byExternalIDMutex.Unlock()
	fake.ByExternalIDStub = nil
	fake.byExternalIDReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByExternalIDReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.byExternalIDMutex.Lock()
	defer fake.byExternalIDMutex.Unlock()
	fake.ByExternalIDStub = nil
	if fake.byExternalIDReturnsOnCall == nil {
		fake.byExternalIDReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.byExternalIDReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByID(arg1 context.Context, arg2 db.DBTX, arg3 uuid.UUID, arg4 ...db.UserSelectConfigOption) (*db.User, error) {
	fake.byIDMutex.Lock()
	ret, specificReturn := fake.byIDReturnsOnCall[len(fake.byIDArgsForCall)]
	fake.byIDArgsForCall = append(fake.byIDArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
		arg4 []db.UserSelectConfigOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ByIDStub
	fakeReturns := fake.byIDReturns
	fake.recordInvocation("ByID", []interface{}{arg1, arg2, arg3, arg4})
	fake.byIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) ByIDCallCount() int {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	return len(fake.byIDArgsForCall)
}

func (fake *FakeUser) ByIDCalls(stub func(context.Context, db.DBTX, uuid.UUID, ...db.UserSelectConfigOption) (*db.User, error)) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = stub
}

func (fake *FakeUser) ByIDArgsForCall(i int) (context.Context, db.DBTX, uuid.UUID, []db.UserSelectConfigOption) {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	argsForCall := fake.byIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUser) ByIDReturns(result1 *db.User, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	fake.byIDReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByIDReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	if fake.byIDReturnsOnCall == nil {
		fake.byIDReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.byIDReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByUsername(arg1 context.Context, arg2 db.DBTX, arg3 string, arg4 ...db.UserSelectConfigOption) (*db.User, error) {
	fake.byUsernameMutex.Lock()
	ret, specificReturn := fake.byUsernameReturnsOnCall[len(fake.byUsernameArgsForCall)]
	fake.byUsernameArgsForCall = append(fake.byUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 []db.UserSelectConfigOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ByUsernameStub
	fakeReturns := fake.byUsernameReturns
	fake.recordInvocation("ByUsername", []interface{}{arg1, arg2, arg3, arg4})
	fake.byUsernameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) ByUsernameCallCount() int {
	fake.byUsernameMutex.RLock()
	defer fake.byUsernameMutex.RUnlock()
	return len(fake.byUsernameArgsForCall)
}

func (fake *FakeUser) ByUsernameCalls(stub func(context.Context, db.DBTX, string, ...db.UserSelectConfigOption) (*db.User, error)) {
	fake.byUsernameMutex.Lock()
	defer fake.byUsernameMutex.Unlock()
	fake.ByUsernameStub = stub
}

func (fake *FakeUser) ByUsernameArgsForCall(i int) (context.Context, db.DBTX, string, []db.UserSelectConfigOption) {
	fake.byUsernameMutex.RLock()
	defer fake.byUsernameMutex.RUnlock()
	argsForCall := fake.byUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUser) ByUsernameReturns(result1 *db.User, result2 error) {
	fake.byUsernameMutex.Lock()
	defer fake.byUsernameMutex.Unlock()
	fake.ByUsernameStub = nil
	fake.byUsernameReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) ByUsernameReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.byUsernameMutex.Lock()
	defer fake.byUsernameMutex.Unlock()
	fake.ByUsernameStub = nil
	if fake.byUsernameReturnsOnCall == nil {
		fake.byUsernameReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.byUsernameReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Create(arg1 context.Context, arg2 db.DBTX, arg3 *db.UserCreateParams) (*db.User, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 *db.UserCreateParams
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUser) CreateCalls(stub func(context.Context, db.DBTX, *db.UserCreateParams) (*db.User, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeUser) CreateArgsForCall(i int) (context.Context, db.DBTX, *db.UserCreateParams) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) CreateReturns(result1 *db.User, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) CreateReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) CreateAPIKey(arg1 context.Context, arg2 db.DBTX, arg3 *db.User) (*db.UserAPIKey, error) {
	fake.createAPIKeyMutex.Lock()
	ret, specificReturn := fake.createAPIKeyReturnsOnCall[len(fake.createAPIKeyArgsForCall)]
	fake.createAPIKeyArgsForCall = append(fake.createAPIKeyArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 *db.User
	}{arg1, arg2, arg3})
	stub := fake.CreateAPIKeyStub
	fakeReturns := fake.createAPIKeyReturns
	fake.recordInvocation("CreateAPIKey", []interface{}{arg1, arg2, arg3})
	fake.createAPIKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) CreateAPIKeyCallCount() int {
	fake.createAPIKeyMutex.RLock()
	defer fake.createAPIKeyMutex.RUnlock()
	return len(fake.createAPIKeyArgsForCall)
}

func (fake *FakeUser) CreateAPIKeyCalls(stub func(context.Context, db.DBTX, *db.User) (*db.UserAPIKey, error)) {
	fake.createAPIKeyMutex.Lock()
	defer fake.createAPIKeyMutex.Unlock()
	fake.CreateAPIKeyStub = stub
}

func (fake *FakeUser) CreateAPIKeyArgsForCall(i int) (context.Context, db.DBTX, *db.User) {
	fake.createAPIKeyMutex.RLock()
	defer fake.createAPIKeyMutex.RUnlock()
	argsForCall := fake.createAPIKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) CreateAPIKeyReturns(result1 *db.UserAPIKey, result2 error) {
	fake.createAPIKeyMutex.Lock()
	defer fake.createAPIKeyMutex.Unlock()
	fake.CreateAPIKeyStub = nil
	fake.createAPIKeyReturns = struct {
		result1 *db.UserAPIKey
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) CreateAPIKeyReturnsOnCall(i int, result1 *db.UserAPIKey, result2 error) {
	fake.createAPIKeyMutex.Lock()
	defer fake.createAPIKeyMutex.Unlock()
	fake.CreateAPIKeyStub = nil
	if fake.createAPIKeyReturnsOnCall == nil {
		fake.createAPIKeyReturnsOnCall = make(map[int]struct {
			result1 *db.UserAPIKey
			result2 error
		})
	}
	fake.createAPIKeyReturnsOnCall[i] = struct {
		result1 *db.UserAPIKey
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Delete(arg1 context.Context, arg2 db.DBTX, arg3 uuid.UUID) (*db.User, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeUser) DeleteCalls(stub func(context.Context, db.DBTX, uuid.UUID) (*db.User, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeUser) DeleteArgsForCall(i int) (context.Context, db.DBTX, uuid.UUID) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) DeleteReturns(result1 *db.User, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) DeleteReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Update(arg1 context.Context, arg2 db.DBTX, arg3 uuid.UUID, arg4 *db.UserUpdateParams) (*db.User, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 uuid.UUID
		arg4 *db.UserUpdateParams
	}{arg1, arg2, arg3, arg4})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeUser) UpdateCalls(stub func(context.Context, db.DBTX, uuid.UUID, *db.UserUpdateParams) (*db.User, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeUser) UpdateArgsForCall(i int) (context.Context, db.DBTX, uuid.UUID, *db.UserUpdateParams) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUser) UpdateReturns(result1 *db.User, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UpdateReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byAPIKeyMutex.RLock()
	defer fake.byAPIKeyMutex.RUnlock()
	fake.byEmailMutex.RLock()
	defer fake.byEmailMutex.RUnlock()
	fake.byExternalIDMutex.RLock()
	defer fake.byExternalIDMutex.RUnlock()
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	fake.byUsernameMutex.RLock()
	defer fake.byUsernameMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createAPIKeyMutex.RLock()
	defer fake.createAPIKeyMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUser) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ repos.User = new(FakeUser)

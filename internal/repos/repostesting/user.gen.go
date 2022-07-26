// Code generated by counterfeiter. DO NOT EDIT.
package repostesting

import (
	"context"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type FakeUser struct {
	CreateStub        func(context.Context, db.DBTX, repos.UserCreateParams) (*db.User, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 repos.UserCreateParams
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
	DeleteStub        func(context.Context, db.DBTX, string) (*db.User, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	deleteReturns struct {
		result1 *db.User
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UpdateStub        func(context.Context, db.DBTX, string, repos.UserUpdateParams) (*db.User, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 repos.UserUpdateParams
	}
	updateReturns struct {
		result1 *db.User
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UserByAPIKeyStub        func(context.Context, db.DBTX, string) (*db.User, error)
	userByAPIKeyMutex       sync.RWMutex
	userByAPIKeyArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	userByAPIKeyReturns struct {
		result1 *db.User
		result2 error
	}
	userByAPIKeyReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UserByEmailStub        func(context.Context, db.DBTX, string) (*db.User, error)
	userByEmailMutex       sync.RWMutex
	userByEmailArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	userByEmailReturns struct {
		result1 *db.User
		result2 error
	}
	userByEmailReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UserByExternalIDStub        func(context.Context, db.DBTX, string) (*db.User, error)
	userByExternalIDMutex       sync.RWMutex
	userByExternalIDArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	userByExternalIDReturns struct {
		result1 *db.User
		result2 error
	}
	userByExternalIDReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UserByIDStub        func(context.Context, db.DBTX, string) (*db.User, error)
	userByIDMutex       sync.RWMutex
	userByIDArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	userByIDReturns struct {
		result1 *db.User
		result2 error
	}
	userByIDReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	UserByUsernameStub        func(context.Context, db.DBTX, string) (*db.User, error)
	userByUsernameMutex       sync.RWMutex
	userByUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}
	userByUsernameReturns struct {
		result1 *db.User
		result2 error
	}
	userByUsernameReturnsOnCall map[int]struct {
		result1 *db.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUser) Create(arg1 context.Context, arg2 db.DBTX, arg3 repos.UserCreateParams) (*db.User, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 repos.UserCreateParams
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

func (fake *FakeUser) CreateCalls(stub func(context.Context, db.DBTX, repos.UserCreateParams) (*db.User, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeUser) CreateArgsForCall(i int) (context.Context, db.DBTX, repos.UserCreateParams) {
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

func (fake *FakeUser) Delete(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
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

func (fake *FakeUser) DeleteCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeUser) DeleteArgsForCall(i int) (context.Context, db.DBTX, string) {
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

func (fake *FakeUser) Update(arg1 context.Context, arg2 db.DBTX, arg3 string, arg4 repos.UserUpdateParams) (*db.User, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
		arg4 repos.UserUpdateParams
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

func (fake *FakeUser) UpdateCalls(stub func(context.Context, db.DBTX, string, repos.UserUpdateParams) (*db.User, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeUser) UpdateArgsForCall(i int) (context.Context, db.DBTX, string, repos.UserUpdateParams) {
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

func (fake *FakeUser) UserByAPIKey(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.userByAPIKeyMutex.Lock()
	ret, specificReturn := fake.userByAPIKeyReturnsOnCall[len(fake.userByAPIKeyArgsForCall)]
	fake.userByAPIKeyArgsForCall = append(fake.userByAPIKeyArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UserByAPIKeyStub
	fakeReturns := fake.userByAPIKeyReturns
	fake.recordInvocation("UserByAPIKey", []interface{}{arg1, arg2, arg3})
	fake.userByAPIKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UserByAPIKeyCallCount() int {
	fake.userByAPIKeyMutex.RLock()
	defer fake.userByAPIKeyMutex.RUnlock()
	return len(fake.userByAPIKeyArgsForCall)
}

func (fake *FakeUser) UserByAPIKeyCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.userByAPIKeyMutex.Lock()
	defer fake.userByAPIKeyMutex.Unlock()
	fake.UserByAPIKeyStub = stub
}

func (fake *FakeUser) UserByAPIKeyArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.userByAPIKeyMutex.RLock()
	defer fake.userByAPIKeyMutex.RUnlock()
	argsForCall := fake.userByAPIKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) UserByAPIKeyReturns(result1 *db.User, result2 error) {
	fake.userByAPIKeyMutex.Lock()
	defer fake.userByAPIKeyMutex.Unlock()
	fake.UserByAPIKeyStub = nil
	fake.userByAPIKeyReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByAPIKeyReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.userByAPIKeyMutex.Lock()
	defer fake.userByAPIKeyMutex.Unlock()
	fake.UserByAPIKeyStub = nil
	if fake.userByAPIKeyReturnsOnCall == nil {
		fake.userByAPIKeyReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.userByAPIKeyReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByEmail(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.userByEmailMutex.Lock()
	ret, specificReturn := fake.userByEmailReturnsOnCall[len(fake.userByEmailArgsForCall)]
	fake.userByEmailArgsForCall = append(fake.userByEmailArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UserByEmailStub
	fakeReturns := fake.userByEmailReturns
	fake.recordInvocation("UserByEmail", []interface{}{arg1, arg2, arg3})
	fake.userByEmailMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UserByEmailCallCount() int {
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	return len(fake.userByEmailArgsForCall)
}

func (fake *FakeUser) UserByEmailCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = stub
}

func (fake *FakeUser) UserByEmailArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	argsForCall := fake.userByEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) UserByEmailReturns(result1 *db.User, result2 error) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = nil
	fake.userByEmailReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByEmailReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = nil
	if fake.userByEmailReturnsOnCall == nil {
		fake.userByEmailReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.userByEmailReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByExternalID(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.userByExternalIDMutex.Lock()
	ret, specificReturn := fake.userByExternalIDReturnsOnCall[len(fake.userByExternalIDArgsForCall)]
	fake.userByExternalIDArgsForCall = append(fake.userByExternalIDArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UserByExternalIDStub
	fakeReturns := fake.userByExternalIDReturns
	fake.recordInvocation("UserByExternalID", []interface{}{arg1, arg2, arg3})
	fake.userByExternalIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UserByExternalIDCallCount() int {
	fake.userByExternalIDMutex.RLock()
	defer fake.userByExternalIDMutex.RUnlock()
	return len(fake.userByExternalIDArgsForCall)
}

func (fake *FakeUser) UserByExternalIDCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.userByExternalIDMutex.Lock()
	defer fake.userByExternalIDMutex.Unlock()
	fake.UserByExternalIDStub = stub
}

func (fake *FakeUser) UserByExternalIDArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.userByExternalIDMutex.RLock()
	defer fake.userByExternalIDMutex.RUnlock()
	argsForCall := fake.userByExternalIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) UserByExternalIDReturns(result1 *db.User, result2 error) {
	fake.userByExternalIDMutex.Lock()
	defer fake.userByExternalIDMutex.Unlock()
	fake.UserByExternalIDStub = nil
	fake.userByExternalIDReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByExternalIDReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.userByExternalIDMutex.Lock()
	defer fake.userByExternalIDMutex.Unlock()
	fake.UserByExternalIDStub = nil
	if fake.userByExternalIDReturnsOnCall == nil {
		fake.userByExternalIDReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.userByExternalIDReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByID(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.userByIDMutex.Lock()
	ret, specificReturn := fake.userByIDReturnsOnCall[len(fake.userByIDArgsForCall)]
	fake.userByIDArgsForCall = append(fake.userByIDArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UserByIDStub
	fakeReturns := fake.userByIDReturns
	fake.recordInvocation("UserByID", []interface{}{arg1, arg2, arg3})
	fake.userByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UserByIDCallCount() int {
	fake.userByIDMutex.RLock()
	defer fake.userByIDMutex.RUnlock()
	return len(fake.userByIDArgsForCall)
}

func (fake *FakeUser) UserByIDCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.userByIDMutex.Lock()
	defer fake.userByIDMutex.Unlock()
	fake.UserByIDStub = stub
}

func (fake *FakeUser) UserByIDArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.userByIDMutex.RLock()
	defer fake.userByIDMutex.RUnlock()
	argsForCall := fake.userByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) UserByIDReturns(result1 *db.User, result2 error) {
	fake.userByIDMutex.Lock()
	defer fake.userByIDMutex.Unlock()
	fake.UserByIDStub = nil
	fake.userByIDReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByIDReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.userByIDMutex.Lock()
	defer fake.userByIDMutex.Unlock()
	fake.UserByIDStub = nil
	if fake.userByIDReturnsOnCall == nil {
		fake.userByIDReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.userByIDReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByUsername(arg1 context.Context, arg2 db.DBTX, arg3 string) (*db.User, error) {
	fake.userByUsernameMutex.Lock()
	ret, specificReturn := fake.userByUsernameReturnsOnCall[len(fake.userByUsernameArgsForCall)]
	fake.userByUsernameArgsForCall = append(fake.userByUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UserByUsernameStub
	fakeReturns := fake.userByUsernameReturns
	fake.recordInvocation("UserByUsername", []interface{}{arg1, arg2, arg3})
	fake.userByUsernameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) UserByUsernameCallCount() int {
	fake.userByUsernameMutex.RLock()
	defer fake.userByUsernameMutex.RUnlock()
	return len(fake.userByUsernameArgsForCall)
}

func (fake *FakeUser) UserByUsernameCalls(stub func(context.Context, db.DBTX, string) (*db.User, error)) {
	fake.userByUsernameMutex.Lock()
	defer fake.userByUsernameMutex.Unlock()
	fake.UserByUsernameStub = stub
}

func (fake *FakeUser) UserByUsernameArgsForCall(i int) (context.Context, db.DBTX, string) {
	fake.userByUsernameMutex.RLock()
	defer fake.userByUsernameMutex.RUnlock()
	argsForCall := fake.userByUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUser) UserByUsernameReturns(result1 *db.User, result2 error) {
	fake.userByUsernameMutex.Lock()
	defer fake.userByUsernameMutex.Unlock()
	fake.UserByUsernameStub = nil
	fake.userByUsernameReturns = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) UserByUsernameReturnsOnCall(i int, result1 *db.User, result2 error) {
	fake.userByUsernameMutex.Lock()
	defer fake.userByUsernameMutex.Unlock()
	fake.UserByUsernameStub = nil
	if fake.userByUsernameReturnsOnCall == nil {
		fake.userByUsernameReturnsOnCall = make(map[int]struct {
			result1 *db.User
			result2 error
		})
	}
	fake.userByUsernameReturnsOnCall[i] = struct {
		result1 *db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createAPIKeyMutex.RLock()
	defer fake.createAPIKeyMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.userByAPIKeyMutex.RLock()
	defer fake.userByAPIKeyMutex.RUnlock()
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	fake.userByExternalIDMutex.RLock()
	defer fake.userByExternalIDMutex.RUnlock()
	fake.userByIDMutex.RLock()
	defer fake.userByIDMutex.RUnlock()
	fake.userByUsernameMutex.RLock()
	defer fake.userByUsernameMutex.RUnlock()
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

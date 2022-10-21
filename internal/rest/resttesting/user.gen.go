// Code generated by counterfeiter. DO NOT EDIT.
package resttesting

import (
	"context"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

type FakeUserService struct {
	RegisterStub        func(context.Context, *crud.User) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 context.Context
		arg2 *crud.User
	}
	registerReturns struct {
		result1 error
	}
	registerReturnsOnCall map[int]struct {
		result1 error
	}
	UpsertStub        func(context.Context, *crud.User) error
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		arg1 context.Context
		arg2 *crud.User
	}
	upsertReturns struct {
		result1 error
	}
	upsertReturnsOnCall map[int]struct {
		result1 error
	}
	UserByEmailStub        func(context.Context, string) (*crud.User, error)
	userByEmailMutex       sync.RWMutex
	userByEmailArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	userByEmailReturns struct {
		result1 *crud.User
		result2 error
	}
	userByEmailReturnsOnCall map[int]struct {
		result1 *crud.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserService) Register(arg1 context.Context, arg2 *crud.User) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 context.Context
		arg2 *crud.User
	}{arg1, arg2})
	stub := fake.RegisterStub
	fakeReturns := fake.registerReturns
	fake.recordInvocation("Register", []interface{}{arg1, arg2})
	fake.registerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserService) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeUserService) RegisterCalls(stub func(context.Context, *crud.User) error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = stub
}

func (fake *FakeUserService) RegisterArgsForCall(i int) (context.Context, *crud.User) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) RegisterReturns(result1 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) RegisterReturnsOnCall(i int, result1 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) Upsert(arg1 context.Context, arg2 *crud.User) error {
	fake.upsertMutex.Lock()
	ret, specificReturn := fake.upsertReturnsOnCall[len(fake.upsertArgsForCall)]
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		arg1 context.Context
		arg2 *crud.User
	}{arg1, arg2})
	stub := fake.UpsertStub
	fakeReturns := fake.upsertReturns
	fake.recordInvocation("Upsert", []interface{}{arg1, arg2})
	fake.upsertMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserService) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *FakeUserService) UpsertCalls(stub func(context.Context, *crud.User) error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = stub
}

func (fake *FakeUserService) UpsertArgsForCall(i int) (context.Context, *crud.User) {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	argsForCall := fake.upsertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) UpsertReturns(result1 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	fake.upsertReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) UpsertReturnsOnCall(i int, result1 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	if fake.upsertReturnsOnCall == nil {
		fake.upsertReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.upsertReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) UserByEmail(arg1 context.Context, arg2 string) (*crud.User, error) {
	fake.userByEmailMutex.Lock()
	ret, specificReturn := fake.userByEmailReturnsOnCall[len(fake.userByEmailArgsForCall)]
	fake.userByEmailArgsForCall = append(fake.userByEmailArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.UserByEmailStub
	fakeReturns := fake.userByEmailReturns
	fake.recordInvocation("UserByEmail", []interface{}{arg1, arg2})
	fake.userByEmailMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) UserByEmailCallCount() int {
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	return len(fake.userByEmailArgsForCall)
}

func (fake *FakeUserService) UserByEmailCalls(stub func(context.Context, string) (*crud.User, error)) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = stub
}

func (fake *FakeUserService) UserByEmailArgsForCall(i int) (context.Context, string) {
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	argsForCall := fake.userByEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) UserByEmailReturns(result1 *crud.User, result2 error) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = nil
	fake.userByEmailReturns = struct {
		result1 *crud.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) UserByEmailReturnsOnCall(i int, result1 *crud.User, result2 error) {
	fake.userByEmailMutex.Lock()
	defer fake.userByEmailMutex.Unlock()
	fake.UserByEmailStub = nil
	if fake.userByEmailReturnsOnCall == nil {
		fake.userByEmailReturnsOnCall = make(map[int]struct {
			result1 *crud.User
			result2 error
		})
	}
	fake.userByEmailReturnsOnCall[i] = struct {
		result1 *crud.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	fake.userByEmailMutex.RLock()
	defer fake.userByEmailMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserService) recordInvocation(key string, args []interface{}) {
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

var _ rest.UserService = new(FakeUserService)

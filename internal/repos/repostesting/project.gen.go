// Code generated by counterfeiter. DO NOT EDIT.
package repostesting

import (
	"context"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type FakeProject struct {
	ByIDStub        func(context.Context, db.DBTX, db.ProjectID, ...db.ProjectSelectConfigOption) (*db.Project, error)
	byIDMutex       sync.RWMutex
	byIDArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 db.ProjectID
		arg4 []db.ProjectSelectConfigOption
	}
	byIDReturns struct {
		result1 *db.Project
		result2 error
	}
	byIDReturnsOnCall map[int]struct {
		result1 *db.Project
		result2 error
	}
	ByNameStub        func(context.Context, db.DBTX, models.Project, ...db.ProjectSelectConfigOption) (*db.Project, error)
	byNameMutex       sync.RWMutex
	byNameArgsForCall []struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 models.Project
		arg4 []db.ProjectSelectConfigOption
	}
	byNameReturns struct {
		result1 *db.Project
		result2 error
	}
	byNameReturnsOnCall map[int]struct {
		result1 *db.Project
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProject) ByID(arg1 context.Context, arg2 db.DBTX, arg3 db.ProjectID, arg4 ...db.ProjectSelectConfigOption) (*db.Project, error) {
	fake.byIDMutex.Lock()
	ret, specificReturn := fake.byIDReturnsOnCall[len(fake.byIDArgsForCall)]
	fake.byIDArgsForCall = append(fake.byIDArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 db.ProjectID
		arg4 []db.ProjectSelectConfigOption
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

func (fake *FakeProject) ByIDCallCount() int {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	return len(fake.byIDArgsForCall)
}

func (fake *FakeProject) ByIDCalls(stub func(context.Context, db.DBTX, db.ProjectID, ...db.ProjectSelectConfigOption) (*db.Project, error)) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = stub
}

func (fake *FakeProject) ByIDArgsForCall(i int) (context.Context, db.DBTX, db.ProjectID, []db.ProjectSelectConfigOption) {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	argsForCall := fake.byIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeProject) ByIDReturns(result1 *db.Project, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	fake.byIDReturns = struct {
		result1 *db.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProject) ByIDReturnsOnCall(i int, result1 *db.Project, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	if fake.byIDReturnsOnCall == nil {
		fake.byIDReturnsOnCall = make(map[int]struct {
			result1 *db.Project
			result2 error
		})
	}
	fake.byIDReturnsOnCall[i] = struct {
		result1 *db.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProject) ByName(arg1 context.Context, arg2 db.DBTX, arg3 models.Project, arg4 ...db.ProjectSelectConfigOption) (*db.Project, error) {
	fake.byNameMutex.Lock()
	ret, specificReturn := fake.byNameReturnsOnCall[len(fake.byNameArgsForCall)]
	fake.byNameArgsForCall = append(fake.byNameArgsForCall, struct {
		arg1 context.Context
		arg2 db.DBTX
		arg3 models.Project
		arg4 []db.ProjectSelectConfigOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ByNameStub
	fakeReturns := fake.byNameReturns
	fake.recordInvocation("ByName", []interface{}{arg1, arg2, arg3, arg4})
	fake.byNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProject) ByNameCallCount() int {
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	return len(fake.byNameArgsForCall)
}

func (fake *FakeProject) ByNameCalls(stub func(context.Context, db.DBTX, models.Project, ...db.ProjectSelectConfigOption) (*db.Project, error)) {
	fake.byNameMutex.Lock()
	defer fake.byNameMutex.Unlock()
	fake.ByNameStub = stub
}

func (fake *FakeProject) ByNameArgsForCall(i int) (context.Context, db.DBTX, models.Project, []db.ProjectSelectConfigOption) {
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	argsForCall := fake.byNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeProject) ByNameReturns(result1 *db.Project, result2 error) {
	fake.byNameMutex.Lock()
	defer fake.byNameMutex.Unlock()
	fake.ByNameStub = nil
	fake.byNameReturns = struct {
		result1 *db.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProject) ByNameReturnsOnCall(i int, result1 *db.Project, result2 error) {
	fake.byNameMutex.Lock()
	defer fake.byNameMutex.Unlock()
	fake.ByNameStub = nil
	if fake.byNameReturnsOnCall == nil {
		fake.byNameReturnsOnCall = make(map[int]struct {
			result1 *db.Project
			result2 error
		})
	}
	fake.byNameReturnsOnCall[i] = struct {
		result1 *db.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProject) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProject) recordInvocation(key string, args []interface{}) {
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

var _ repos.Project = new(FakeProject)

package repostesting

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type fakeUserStore struct {
	users map[string]*db.User

	mu sync.Mutex
}

func (f *fakeUserStore) get(id string) (*db.User, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[id]

	return user, ok
}

func (f *fakeUserStore) set(id string, user *db.User) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.users[id] = user
}

// NewFakeUser returns a stub for the User repository.
func NewFakeUser() *FakeUser {
	fks := &fakeUserStore{
		users: make(map[string]*db.User),
		mu:    sync.Mutex{},
	}

	fakeUserRepo := &FakeUser{}

	fakeUserRepo.UserByIDStub = func(ctx context.Context, d db.DBTX, id string) (*db.User, error) {
		user, ok := fks.get(id)
		if !ok {
			return &db.User{}, errors.New("could not get user by ID")
		}

		return user, nil
	}

	fakeUserRepo.UpdateStub = func(ctx context.Context, d db.DBTX, params repos.UserUpdateParams) (*db.User, error) {
		user, err := fakeUserRepo.UserByIDStub(ctx, d, params.ID)
		if err != nil {
			return &db.User{}, fmt.Errorf("UserByIDStub: %w", err)
		}

		if params.FirstName != nil {
			user.FirstName = params.FirstName
		}
		if params.LastName != nil {
			user.LastName = params.LastName
		}
		if params.Scopes != nil {
			user.Scopes = *params.Scopes
		}
		if params.Rank != nil {
			user.RoleRank = *params.Rank
		}

		fks.set(params.ID, user)

		return user, nil
	}

	return fakeUserRepo
}

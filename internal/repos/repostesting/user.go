package repostesting

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

type fakeUserStore struct {
	users map[uuid.UUID]db.User

	mu sync.Mutex
}

func (f *fakeUserStore) get(id uuid.UUID) (db.User, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	user, ok := f.users[id]

	return user, ok
}

func (f *fakeUserStore) set(id uuid.UUID, user *db.User) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.users[id] = *user
}

// NewFakeUser returns a mock for the User repository, initializing it with copies of
// the passed users.
// Deprecated: use postgres repo directly.
func NewFakeUser(users ...*db.User) *FakeUser {
	fks := &fakeUserStore{
		users: make(map[uuid.UUID]db.User),
		mu:    sync.Mutex{},
	}

	for _, u := range users {
		uc := *u
		fks.set(u.UserID, &uc)
	}

	fakeUserRepo := &FakeUser{}

	fakeUserRepo.ByIDStub = func(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error) {
		user, ok := fks.get(id)
		if !ok {
			return &db.User{}, errors.New("could not get user by ID")
		}

		return &user, nil
	}

	fakeUserRepo.UpdateStub = func(ctx context.Context, d db.DBTX, id uuid.UUID, params db.UserUpdateParams) (*db.User, error) {
		user, err := fakeUserRepo.ByID(ctx, d, id)
		if err != nil {
			return &db.User{}, fmt.Errorf("UserByIDStub: %w", err)
		}

		if params.FirstName != nil {
			user.FirstName = *params.FirstName
		}
		if params.LastName != nil {
			user.LastName = *params.LastName
		}
		if params.Scopes != nil {
			user.Scopes = *params.Scopes
		}
		if params.RoleRank != nil {
			user.RoleRank = *params.RoleRank
		}

		fks.set(id, user)

		return user, nil
	}

	return fakeUserRepo
}

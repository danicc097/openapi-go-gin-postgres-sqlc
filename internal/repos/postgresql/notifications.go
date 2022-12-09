package postgresql

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Notification represents the repository used for interacting with Notification records.
type Notification struct {
	q *db.Queries
}

// NewNotification instantiates the Notification repository.
func NewNotification() *Notification {
	return &Notification{
		q: db.New(),
	}
}

// var _ repos.Notification = (*Notification)(nil)

// func (u *Notification) Create(ctx context.Context, d db.DBTX, params repos.UserCreateParams) (*db.Notification, error) {
// 	user := &db.Notification{
// 		Username:   params.Username,
// 		Email:      params.Email,
// 		FirstName:  params.FirstName,
// 		LastName:   params.LastName,
// 		ExternalID: params.ExternalID,
// 		RoleRank:   params.RoleRank,
// 		Scopes:     params.Scopes,
// 	}

// 	if err := user.Save(ctx, d); err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (u *Notification) Update(ctx context.Context, d db.DBTX, id string, params repos.UserUpdateParams) (*db.Notification, error) {
// 	user, err := u.UserByID(ctx, d, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user by id %w", parseErrorDetail(err))
// 	}

// 	// distinguish keys not present in json body and zero valued ones
// 	if params.FirstName != nil {
// 		user.FirstName = params.FirstName
// 	}
// 	if params.LastName != nil {
// 		user.LastName = params.LastName
// 	}
// 	if params.Scopes != nil {
// 		user.Scopes = *params.Scopes
// 	}
// 	if params.Rank != nil {
// 		user.RoleRank = *params.Rank
// 	}

// 	err = user.Update(ctx, d)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
// 	}

// 	return user, err
// }

// func (u *Notification) Delete(ctx context.Context, d db.DBTX, id string) (*db.Notification, error) {
// 	user, err := u.UserByID(ctx, d, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user by id %w", parseErrorDetail(err))
// 	}

// 	user.DeletedAt = pointers.New(time.Now())

// 	err = user.Update(ctx, d)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not mark user as deleted: %w", parseErrorDetail(err))
// 	}

// 	return user, err
// }

// func (u *Notification) UserByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.Notification, error) {
// 	user, err := db.UserByExternalID(ctx, d, extID)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
// 	}

// 	return user, nil
// }

// func (u *Notification) UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.Notification, error) {
// 	user, err := db.UserByEmail(ctx, d, email)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
// 	}

// 	return user, nil
// }

// func (u *Notification) UserByUsername(ctx context.Context, d db.DBTX, username string) (*db.Notification, error) {
// 	user, err := db.UserByUsername(ctx, d, username)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
// 	}

// 	return user, nil
// }

// func (u *Notification) UserByID(ctx context.Context, d db.DBTX, id string) (*db.Notification, error) {
// 	uid, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not parse id as UUID: %w", parseErrorDetail(err))
// 	}

// 	user, err := db.UserByUserID(ctx, d, uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
// 	}

// 	return user, nil
// }

// func (u *Notification) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.Notification, error) {
// 	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{Notification: true}))
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get api key: %w", parseErrorDetail(err))
// 	}

// 	if uak.Notification == nil {
// 		return nil, fmt.Errorf("could not join user by api key")
// 	}

// 	return uak.Notification, nil
// }

// func (u *Notification) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.Notification) (*db.UserAPIKey, error) {
// 	uak := &db.UserAPIKey{
// 		APIKey:    uuid.NewString(),
// 		ExpiresOn: time.Now().AddDate(1, 0, 0),
// 		UserID:    user.UserID,
// 	}
// 	if err := uak.Save(ctx, d); err != nil {
// 		return nil, fmt.Errorf("could not save api key: %w", parseErrorDetail(err))
// 	}

// 	user.APIKeyID = pointers.New(uak.UserAPIKeyID)
// 	if err := user.Update(ctx, d); err != nil {
// 		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
// 	}

// 	return uak, nil
// }

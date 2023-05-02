package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
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

var _ repos.Notification = (*Notification)(nil)

// TODO see internal/repos/postgresql/TODO.md
// using created_at > @last_notification_created_at is all we need at the very least. add more parameters
// to ensure uniqueness in more complex cases
// TODO database sql not needed with jet. we generate raw sql and parameters from jet and then call pgx.
// so that way we can use transactions, same signatures, etc.
// func (u *Notification) LatestUserNotifications(ctx context.Context, d sql.Conn, userID string) ([]*db.UserNotification, error) {
// 	uid, err := uuid.Parse(userID)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not parse user id as UUID: %w", parseErrorDetail(err))
// 	}

// 	query, args := queries.GetUserNotificationsByUserID(uid).Sql()

// 	fmt.Println(query) // will print parameterized sql ($1, ...)
// 	fmt.Println(args)

// 	// we can execute the query with pgx now. We have query and args

// 	type Res []struct {
// 		model.UserNotifications

// 		Notification model.Notifications
// 	}

// 	dest := &Res{}

// 	// won't be able to use same transaction and also need a sql.DB pool apart from pgxpool opened with postgresql.New
// 	// https://github.com/go-jet/jet/issues/59
// 	// this will break our repo and service (d db.DBTX) param
// 	err = getUserNotificationsByUserID.QueryContext(context.Background(), sqlpool, dest)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("dest: %#v\n", dest)
// 	format.PrintJSON(dest)
// 	// joins := db.WithUserNotificationJoin(db.UserNotificationJoins{Notification: true})
// 	// orderby := db.WithUserNotificationOrderBy(db.NotificationCreatedAtDescNullsLast)
// 	// nn, err := db.UserNotificationsByUserID(ctx, d, uid, joins, orderby)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("could not get user notifications: %w", parseErrorDetail(err))
// 	// }

// 	return nn, nil
// }

func (u *Notification) LatestUserNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error) {
	nn, err := u.q.GetUserNotifications(ctx, d, *params)
	if err != nil {
		return nil, fmt.Errorf("could not get notifications for user: %w", parseErrorDetail(err))
	}

	return nn, nil
}

func (u *Notification) Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (*db.Notification, error) {
	notification := &db.Notification{}
	notification.SetCreateParams(params)
	if _, err := notification.Insert(ctx, d); err != nil {
		return nil, err
	}

	return notification, nil
}

func (u *Notification) Delete(ctx context.Context, d db.DBTX, id int) (*db.Notification, error) {
	notification, err := db.NotificationByNotificationID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get notification by id %w", parseErrorDetail(err))
	}

	err = notification.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete notification: %w", parseErrorDetail(err))
	}

	return notification, err
}

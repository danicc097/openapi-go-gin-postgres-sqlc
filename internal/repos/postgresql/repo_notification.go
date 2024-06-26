package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// Notification represents the repository used for interacting with Notification records.
type Notification struct {
	q models.Querier
}

// NewNotification instantiates the Notification repository.
func NewNotification() *Notification {
	return &Notification{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.Notification = (*Notification)(nil)

// TODO see internal/repos/postgresql/TODO.md to have xo generate paginated queries
// using created_at > @last_notification_created_at is all we need at the very least. add more parameters
// to ensure uniqueness in more complex cases
// TODO database sql not needed with jet. we generate raw sql and parameters from jet and then call pgx.
// so that way we can use transactions, same signatures, etc.
// func (u *Notification) LatestNotifications(ctx context.Context, d sql.Conn, userID string) ([]*db.UserNotification, error) {
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

func (u *Notification) LatestNotifications(ctx context.Context, d models.DBTX, params *models.GetUserNotificationsParams) ([]models.GetUserNotificationsRow, error) {
	nn, err := u.q.GetUserNotifications(ctx, d, *params)
	if err != nil {
		return nil, fmt.Errorf("could not get notifications for user: %w", ParseDBErrorDetail(err))
	}

	return nn, nil
}

func (u *Notification) Create(ctx context.Context, d models.DBTX, params *models.NotificationCreateParams) (*models.UserNotification, error) {
	notification, err := models.CreateNotification(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create notification: %w", ParseDBErrorDetail(err))
	}

	// only retrieve 1 user notification at most
	nn, err := models.UserNotificationsByNotificationID(ctx, d, notification.NotificationID, models.WithUserNotificationLimit(1))
	if len(nn) == 0 {
		return nil, fmt.Errorf("could not create notification fan out: %w", ParseDBErrorDetail(err))
	}

	return &nn[0], nil
}

func (u *Notification) PaginatedUserNotifications(ctx context.Context, d models.DBTX, userID models.UserID, params models.GetPaginatedNotificationsParams) ([]models.UserNotification, error) {
	opts := []models.UserNotificationSelectConfigOption{
		models.WithUserNotificationFilters(map[string][]any{
			"user_id = $i": {userID}, // further restrictions as desired
		}),
		models.WithUserNotificationJoin(models.UserNotificationJoins{Notification: true}),
	}
	if params.Limit > 0 {
		opts = append(opts, models.WithUserNotificationLimit(params.Limit))
	}

	cursor := models.PaginationCursor{Column: "userNotificationID", Value: pointers.New[interface{}](params.Cursor), Direction: params.Direction}
	notifications, err := models.UserNotificationPaginated(ctx, d, cursor, opts...)
	if err != nil {
		return nil, fmt.Errorf("could get paginated notifications: %w", ParseDBErrorDetail(err))
	}

	return notifications, nil
}

func (u *Notification) Delete(ctx context.Context, d models.DBTX, id models.NotificationID) (*models.Notification, error) {
	notification := &models.Notification{
		NotificationID: id,
	}

	err := notification.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete notification: %w", ParseDBErrorDetail(err))
	}

	return notification, err
}

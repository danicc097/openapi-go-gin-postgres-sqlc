package repostesting

import (
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

type fakeNotificationStore struct {
	notifications map[int]models.Notification

	mu sync.Mutex
}

func (f *fakeNotificationStore) get(id int) (models.Notification, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	notification, ok := f.notifications[id]

	return notification, ok
}

func (f *fakeNotificationStore) set(id int, notification *models.Notification) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.notifications[id] = *notification
}

// NewFakeNotification returns a mock for the Notification repository, initializing it with copies of
// the passed notifications.
func NewFakeNotification(notifications ...*models.Notification) *FakeNotification {
	fks := &fakeNotificationStore{
		notifications: make(map[int]models.Notification),
		mu:            sync.Mutex{},
	}

	for _, u := range notifications {
		uc := *u
		fks.set(int(u.NotificationID), &uc)
	}

	fakeNotificationRepo := &FakeNotification{}

	return fakeNotificationRepo
}

// Code generated by gowrap. DO NOT EDIT.
// template: ../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package repos

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// UserWithPrometheus implements User interface with all methods wrapped
// with Prometheus metrics
type UserWithPrometheus struct {
	base         User
	instanceName string
}

var userDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "user_duration_seconds",
		Help:       "user runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewUserWithPrometheus returns an instance of the User decorated with prometheus summary metric
func NewUserWithPrometheus(base User, instanceName string) UserWithPrometheus {
	return UserWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// Create implements User
func (_d UserWithPrometheus) Create(ctx context.Context, d db.DBTX, params UserCreateParams) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// CreateAPIKey implements User
func (_d UserWithPrometheus) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (up1 *db.UserAPIKey, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateAPIKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateAPIKey(ctx, d, user)
}

// Update implements User
func (_d UserWithPrometheus) Update(ctx context.Context, d db.DBTX, id string, params UserUpdateParams) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}

// UserByAPIKey implements User
func (_d UserWithPrometheus) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "UserByAPIKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UserByAPIKey(ctx, d, apiKey)
}

// UserByEmail implements User
func (_d UserWithPrometheus) UserByEmail(ctx context.Context, d db.DBTX, email string) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "UserByEmail", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UserByEmail(ctx, d, email)
}

// UserByExternalID implements User
func (_d UserWithPrometheus) UserByExternalID(ctx context.Context, d db.DBTX, extID string) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "UserByExternalID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UserByExternalID(ctx, d, extID)
}

// UserByID implements User
func (_d UserWithPrometheus) UserByID(ctx context.Context, d db.DBTX, id string) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "UserByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UserByID(ctx, d, id)
}

// UserByUsername implements User
func (_d UserWithPrometheus) UserByUsername(ctx context.Context, d db.DBTX, username string) (up1 *db.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userDurationSummaryVec.WithLabelValues(_d.instanceName, "UserByUsername", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UserByUsername(ctx, d, username)
}

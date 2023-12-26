// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// EntityNotificationWithPrometheus implements repos.EntityNotification interface with all methods wrapped
// with Prometheus metrics
type EntityNotificationWithPrometheus struct {
	base         repos.EntityNotification
	instanceName string
}

var entitynotificationDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "entitynotification_duration_seconds",
		Help:       "entitynotification runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewEntityNotificationWithPrometheus returns an instance of the repos.EntityNotification decorated with prometheus summary metric
func NewEntityNotificationWithPrometheus(base repos.EntityNotification, instanceName string) EntityNotificationWithPrometheus {
	return EntityNotificationWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.EntityNotification
func (_d EntityNotificationWithPrometheus) ByID(ctx context.Context, d db.DBTX, id db.EntityNotificationID, opts ...db.EntityNotificationSelectConfigOption) (ep1 *db.EntityNotification, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		entitynotificationDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id, opts...)
}

// Create implements repos.EntityNotification
func (_d EntityNotificationWithPrometheus) Create(ctx context.Context, d db.DBTX, params *db.EntityNotificationCreateParams) (ep1 *db.EntityNotification, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		entitynotificationDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Delete implements repos.EntityNotification
func (_d EntityNotificationWithPrometheus) Delete(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (ep1 *db.EntityNotification, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		entitynotificationDurationSummaryVec.WithLabelValues(_d.instanceName, "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Delete(ctx, d, id)
}

// Update implements repos.EntityNotification
func (_d EntityNotificationWithPrometheus) Update(ctx context.Context, d db.DBTX, id db.EntityNotificationID, params *db.EntityNotificationUpdateParams) (ep1 *db.EntityNotification, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		entitynotificationDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}

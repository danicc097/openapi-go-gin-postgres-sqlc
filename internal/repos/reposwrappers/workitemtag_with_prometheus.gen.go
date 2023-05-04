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

// WorkItemTagWithPrometheus implements repos.WorkItemTag interface with all methods wrapped
// with Prometheus metrics
type WorkItemTagWithPrometheus struct {
	base         repos.WorkItemTag
	instanceName string
}

var workitemtagDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "workitemtag_duration_seconds",
		Help:       "workitemtag runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewWorkItemTagWithPrometheus returns an instance of the repos.WorkItemTag decorated with prometheus summary metric
func NewWorkItemTagWithPrometheus(base repos.WorkItemTag, instanceName string) WorkItemTagWithPrometheus {
	return WorkItemTagWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.WorkItemTag
func (_d WorkItemTagWithPrometheus) ByID(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtagDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id)
}

// ByName implements repos.WorkItemTag
func (_d WorkItemTagWithPrometheus) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (wp1 *db.WorkItemTag, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtagDurationSummaryVec.WithLabelValues(_d.instanceName, "ByName", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByName(ctx, d, name, projectID)
}

// Create implements repos.WorkItemTag
func (_d WorkItemTagWithPrometheus) Create(ctx context.Context, d db.DBTX, params *db.WorkItemTagCreateParams) (wp1 *db.WorkItemTag, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtagDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Delete implements repos.WorkItemTag
func (_d WorkItemTagWithPrometheus) Delete(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtagDurationSummaryVec.WithLabelValues(_d.instanceName, "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Delete(ctx, d, id)
}

// Update implements repos.WorkItemTag
func (_d WorkItemTagWithPrometheus) Update(ctx context.Context, d db.DBTX, id int, params *db.WorkItemTagUpdateParams) (wp1 *db.WorkItemTag, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtagDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}

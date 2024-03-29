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

// DemoWorkItemWithPrometheus implements repos.DemoWorkItem interface with all methods wrapped
// with Prometheus metrics
type DemoWorkItemWithPrometheus struct {
	base         repos.DemoWorkItem
	instanceName string
}

var demoworkitemDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "demoworkitem_duration_seconds",
		Help:       "demoworkitem runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewDemoWorkItemWithPrometheus returns an instance of the repos.DemoWorkItem decorated with prometheus summary metric
func NewDemoWorkItemWithPrometheus(base repos.DemoWorkItem, instanceName string) DemoWorkItemWithPrometheus {
	return DemoWorkItemWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.DemoWorkItem
func (_d DemoWorkItemWithPrometheus) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		demoworkitemDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id, opts...)
}

// Create implements repos.DemoWorkItem
func (_d DemoWorkItemWithPrometheus) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		demoworkitemDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Paginated implements repos.DemoWorkItem
func (_d DemoWorkItemWithPrometheus) Paginated(ctx context.Context, d db.DBTX, cursor db.WorkItemID, opts ...db.CacheDemoWorkItemSelectConfigOption) (ca1 []db.CacheDemoWorkItem, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		demoworkitemDurationSummaryVec.WithLabelValues(_d.instanceName, "Paginated", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Paginated(ctx, d, cursor, opts...)
}

// Update implements repos.DemoWorkItem
func (_d DemoWorkItemWithPrometheus) Update(ctx context.Context, d db.DBTX, id db.WorkItemID, params repos.DemoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		demoworkitemDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}

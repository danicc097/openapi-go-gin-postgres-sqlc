// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// KanbanStepWithPrometheus implements repos.KanbanStep interface with all methods wrapped
// with Prometheus metrics
type KanbanStepWithPrometheus struct {
	base         repos.KanbanStep
	instanceName string
}

var kanbanstepDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "kanbanstep_duration_seconds",
		Help:       "kanbanstep runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewKanbanStepWithPrometheus returns an instance of the repos.KanbanStep decorated with prometheus summary metric
func NewKanbanStepWithPrometheus(base repos.KanbanStep, instanceName string) KanbanStepWithPrometheus {
	return KanbanStepWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.KanbanStep
func (_d KanbanStepWithPrometheus) ByID(ctx context.Context, d models.DBTX, id models.KanbanStepID, opts ...models.KanbanStepSelectConfigOption) (kp1 *models.KanbanStep, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		kanbanstepDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id, opts...)
}

// ByProject implements repos.KanbanStep
func (_d KanbanStepWithPrometheus) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.KanbanStepSelectConfigOption) (ka1 []models.KanbanStep, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		kanbanstepDurationSummaryVec.WithLabelValues(_d.instanceName, "ByProject", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByProject(ctx, d, projectID, opts...)
}

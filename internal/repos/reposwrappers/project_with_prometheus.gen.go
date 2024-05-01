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

// ProjectWithPrometheus implements repos.Project interface with all methods wrapped
// with Prometheus metrics
type ProjectWithPrometheus struct {
	base         repos.Project
	instanceName string
}

var projectDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "project_duration_seconds",
		Help:       "project runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewProjectWithPrometheus returns an instance of the repos.Project decorated with prometheus summary metric
func NewProjectWithPrometheus(base repos.Project, instanceName string) ProjectWithPrometheus {
	return ProjectWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.Project
func (_d ProjectWithPrometheus) ByID(ctx context.Context, d models.DBTX, id models.ProjectID, opts ...models.ProjectSelectConfigOption) (pp1 *models.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id, opts...)
}

// ByName implements repos.Project
func (_d ProjectWithPrometheus) ByName(ctx context.Context, d models.DBTX, name models.ProjectName, opts ...models.ProjectSelectConfigOption) (pp1 *models.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectDurationSummaryVec.WithLabelValues(_d.instanceName, "ByName", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByName(ctx, d, name, opts...)
}

// IsTeamInProject implements repos.Project
func (_d ProjectWithPrometheus) IsTeamInProject(ctx context.Context, d models.DBTX, arg models.IsTeamInProjectParams) (b1 bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectDurationSummaryVec.WithLabelValues(_d.instanceName, "IsTeamInProject", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.IsTeamInProject(ctx, d, arg)
}

// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TeamWithPrometheus implements _sourceRepos.Team interface with all methods wrapped
// with Prometheus metrics
type TeamWithPrometheus struct {
	base         _sourceRepos.Team
	instanceName string
}

var teamDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "team_duration_seconds",
		Help:       "team runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewTeamWithPrometheus returns an instance of the _sourceRepos.Team decorated with prometheus summary metric
func NewTeamWithPrometheus(base _sourceRepos.Team, instanceName string) TeamWithPrometheus {
	return TeamWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements _sourceRepos.Team
func (_d TeamWithPrometheus) ByID(ctx context.Context, d models.DBTX, id models.TeamID, opts ...models.TeamSelectConfigOption) (tp1 *models.Team, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		teamDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id, opts...)
}

// ByName implements _sourceRepos.Team
func (_d TeamWithPrometheus) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.TeamSelectConfigOption) (tp1 *models.Team, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		teamDurationSummaryVec.WithLabelValues(_d.instanceName, "ByName", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByName(ctx, d, name, projectID, opts...)
}

// Create implements _sourceRepos.Team
func (_d TeamWithPrometheus) Create(ctx context.Context, d models.DBTX, params *models.TeamCreateParams) (tp1 *models.Team, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		teamDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Delete implements _sourceRepos.Team
func (_d TeamWithPrometheus) Delete(ctx context.Context, d models.DBTX, id models.TeamID) (tp1 *models.Team, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		teamDurationSummaryVec.WithLabelValues(_d.instanceName, "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Delete(ctx, d, id)
}

// Update implements _sourceRepos.Team
func (_d TeamWithPrometheus) Update(ctx context.Context, d models.DBTX, id models.TeamID, params *models.TeamUpdateParams) (tp1 *models.Team, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		teamDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}

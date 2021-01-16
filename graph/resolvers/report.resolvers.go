package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *mutationResolver) SubmitReport(ctx context.Context, userID int64, description string) (*models.Report, error) {
	reporter := middlewares.UseAuth(ctx)
	if reporter == nil {
		return nil, errors.New("not authenticated")
	}

	reported, err := new(repositories.UserRepository).GetByID(userID)
	if err != nil {
		return nil, err
	}

	report := &models.Report{
		ReporterID:  reporter.ID,
		Reporter:    *reporter,
		ReportedID:  reported.ID,
		Reported:    *reported,
		Description: description,
	}
	if err := facades.UseDB().Create(report).Error; err != nil {
		return nil, err
	}

	return report, nil
}

func (r *queryResolver) GetReportsByUser(ctx context.Context, id int64) ([]*models.Report, error) {
	var reports []*models.Report
	if err := facades.UseDB().
		Preload("Reporter").
		Preload("Reported").
		Find(&reports, "reported_id = ?", id).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

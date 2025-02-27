package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
)

type JobApplicationService interface {
	Create(ctx context.Context, req request.CreateJobApplication, userId string) (*model.JobApplication, error)
	FindById(ctx context.Context, id int) (*model.JobApplication, error)
	FindAllByUserId(ctx context.Context, userId string) ([]model.JobApplication, error)
	Update(ctx context.Context, jobApplicationId int, req request.UpdateJobApplication) (*model.JobApplication, error)
	Delete(ctx context.Context, jobApplicationId int) error
}

type jobApplicationService struct {
	jobApplicationRepo repo.JobApplicationRepo
}

func NewJobApplicationService(jobApplicationRepo repo.JobApplicationRepo) JobApplicationService {
	return &jobApplicationService{jobApplicationRepo: jobApplicationRepo}
}

func (s *jobApplicationService) Create(ctx context.Context, req request.CreateJobApplication, userId string) (*model.JobApplication, error) {
	var jobApplication model.JobApplication
	jobApplication.Position = req.Position
	jobApplication.Company = req.Company
	jobApplication.Location = req.Location
	jobApplication.Source = req.Source
	jobApplication.Progress = req.Progress
	jobApplication.Notes = req.Notes
	jobApplication.AppliedDate = req.AppliedDate

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	jobApplication.UserID = parsedUUID

	return s.jobApplicationRepo.Save(ctx, &jobApplication)
}
func (s *jobApplicationService) FindById(ctx context.Context, id int) (*model.JobApplication, error) {
	return s.jobApplicationRepo.FindById(ctx, id)
}

func (s *jobApplicationService) FindAllByUserId(ctx context.Context, userId string) ([]model.JobApplication, error) {
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	return s.jobApplicationRepo.FindAllByUserId(ctx, parsedUUID)
}

func (s *jobApplicationService) Update(ctx context.Context, jobApplicationId int, req request.UpdateJobApplication) (*model.JobApplication, error) {
	updates := map[string]interface{}{}
	addIfNotNil(updates, "position", req.Position)
	addIfNotNil(updates, "company", req.Company)
	addIfNotNil(updates, "location", req.Location)
	addIfNotNil(updates, "source", req.Source)
	addIfNotNil(updates, "progress", req.Progress)
	addIfNotNil(updates, "notes", req.Notes)
	if req.AppliedDate != nil {
		updates["applied_date"] = *req.AppliedDate
	}

	return s.jobApplicationRepo.Update(ctx, jobApplicationId, updates)
}

func (s *jobApplicationService) Delete(ctx context.Context, jobApplicationId int) error {
	return s.jobApplicationRepo.Delete(ctx, jobApplicationId)
}

func addIfNotNil(m map[string]interface{}, key string, value *string) {
	if value != nil {
		m[key] = *value
	}
}

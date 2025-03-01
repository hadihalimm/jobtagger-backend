package service

import (
	"context"

	"github.com/hadihalimm/jobtagger-backend/internal/model"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
)

type InterviewService interface {
	Create(ctx context.Context, req request.CreateInterview, applicationId int) (*model.Interview, error)
	FindById(ctx context.Context, interviewId int) (*model.Interview, error)
	FindAllByApplicationId(ctx context.Context, applicationId int) ([]model.Interview, error)
	Update(ctx context.Context, interviewId int, req request.UpdateInterview) (*model.Interview, error)
	Delete(ctx context.Context, interviewId int) error
}

type interviewService struct {
	interviewRepo repo.InterviewRepo
}

func NewInterviewService(interviewRepo repo.InterviewRepo) InterviewService {
	return &interviewService{interviewRepo: interviewRepo}
}

func (s *interviewService) Create(ctx context.Context, req request.CreateInterview, applicationId int) (*model.Interview, error) {
	var interview model.Interview
	interview.ApplicationID = applicationId
	interview.Title = req.Title
	interview.Date = req.Date
	interview.Position = req.Position
	interview.Company = req.Company
	interview.Notes = req.Notes

	return s.interviewRepo.Save(ctx, interview)
}

func (s *interviewService) FindById(ctx context.Context, interviewId int) (*model.Interview, error) {
	return s.interviewRepo.FindById(ctx, interviewId)
}

func (s *interviewService) FindAllByApplicationId(ctx context.Context, applicationId int) ([]model.Interview, error) {
	return s.interviewRepo.FindAllByApplicationId(ctx, applicationId)
}

func (s *interviewService) Update(ctx context.Context, interviewId int, req request.UpdateInterview) (*model.Interview, error) {
	updates := map[string]interface{}{}
	addIfNotNil(updates, "title", req.Title)
	addIfNotNil(updates, "position", req.Position)
	addIfNotNil(updates, "company", req.Company)
	addIfNotNil(updates, "notes", req.Notes)

	if req.Date != nil {
		updates["interview_date"] = *req.Date
	}

	return s.interviewRepo.Update(ctx, interviewId, updates)
}

func (s *interviewService) Delete(ctx context.Context, interviewId int) error {
	return s.interviewRepo.Delete(ctx, interviewId)
}

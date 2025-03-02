package service

import (
	"context"

	"github.com/hadihalimm/jobtagger-backend/internal/model"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
)

type TimelineService interface {
	Create(ctx context.Context, req request.CreateTimeline, jobApplicationId int) (*model.Timeline, error)
	FindAllByApplicationId(ctx context.Context, jobApplicationId int) ([]model.Timeline, error)
	FindById(ctx context.Context, id int) (*model.Timeline, error)
	Update(ctx context.Context, timelineId int, req request.UpdateTimeline) (*model.Timeline, error)
	Delete(ctx context.Context, timelineId int) error
}

type timelineService struct {
	timelineRepo repo.TimelineRepo
}

func NewTimelineService(timelineRepo repo.TimelineRepo) TimelineService {
	return &timelineService{timelineRepo: timelineRepo}
}

func (s *timelineService) Create(ctx context.Context, req request.CreateTimeline, jobApplicationId int) (*model.Timeline, error) {
	var timeline model.Timeline

	timeline.ApplicationID = jobApplicationId
	timeline.Content = req.Content
	timeline.TimelineDate = req.TimelineDate

	return s.timelineRepo.Save(ctx, timeline)
}

func (s *timelineService) FindAllByApplicationId(ctx context.Context, jobApplicationId int) ([]model.Timeline, error) {
	return s.timelineRepo.FindAllByApplicationId(ctx, jobApplicationId)
}

func (s *timelineService) FindById(ctx context.Context, id int) (*model.Timeline, error) {
	return s.timelineRepo.FindById(ctx, id)
}

func (s *timelineService) Update(ctx context.Context, timelineId int, req request.UpdateTimeline) (*model.Timeline, error) {
	updates := map[string]interface{}{}
	addIfNotNil(updates, "content", req.Content)
	if req.TimelineDate != nil {
		updates["timeline_date"] = *req.TimelineDate
	}

	return s.timelineRepo.Update(ctx, timelineId, updates)
}

func (s *timelineService) Delete(ctx context.Context, timelineId int) error {
	return s.timelineRepo.Delete(ctx, timelineId)
}

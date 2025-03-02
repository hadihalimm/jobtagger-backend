package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type TimelineRepo interface {
	Save(ctx context.Context, timeline model.Timeline) (*model.Timeline, error)
	FindAllByApplicationId(ctx context.Context, jobApplicationId int) ([]model.Timeline, error)
	FindById(ctx context.Context, id int) (*model.Timeline, error)
	Update(ctx context.Context, timelineId int, updates map[string]interface{}) (*model.Timeline, error)
	Delete(ctx context.Context, timelineId int) error
}

type timelineRepo struct {
	db *config.Database
}

func NewTimelineRepo(db *config.Database) TimelineRepo {
	return &timelineRepo{db: db}
}

func (r *timelineRepo) Save(ctx context.Context, timeline model.Timeline) (*model.Timeline, error) {
	query := `INSERT INTO timelines 
	(application_id, content, timeline_date) 
	VALUES ($1, $2, $3) 
	RETURNING *`

	var savedTimeline model.Timeline
	err := r.db.Pgx.QueryRow(ctx, query,
		timeline.ApplicationID, timeline.Content, timeline.TimelineDate).Scan(
		&savedTimeline.ID,
		&savedTimeline.ApplicationID,
		&savedTimeline.Content,
		&savedTimeline.TimelineDate,
		&savedTimeline.CreatedAt,
		&savedTimeline.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &savedTimeline, nil
}

func (r *timelineRepo) FindAllByApplicationId(ctx context.Context, jobApplicationId int) ([]model.Timeline, error) {
	query := `SELECT * FROM timelines WHERE application_id=$1`

	rows, err := r.db.Pgx.Query(ctx, query, jobApplicationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timelines []model.Timeline
	for rows.Next() {
		var timeline model.Timeline
		err := rows.Scan(
			&timeline.ID,
			&timeline.ApplicationID,
			&timeline.Content,
			&timeline.TimelineDate,
			&timeline.CreatedAt,
			&timeline.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		timelines = append(timelines, timeline)
	}

	return timelines, nil
}

func (r *timelineRepo) FindById(ctx context.Context, id int) (*model.Timeline, error) {
	query := `SELECT * FROM timelines WHERE id=$1`

	var timeline model.Timeline
	err := r.db.Pgx.QueryRow(ctx, query, id).Scan(
		&timeline.ID,
		&timeline.ApplicationID,
		&timeline.Content,
		&timeline.TimelineDate,
		&timeline.CreatedAt,
		&timeline.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (r *timelineRepo) Update(ctx context.Context, timelineId int, updates map[string]interface{}) (*model.Timeline, error) {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argIndex))
		args = append(args, value)
		argIndex++
	}

	query := fmt.Sprintf(`
		UPDATE timelines 
		SET %s 
		WHERE id = $%d RETURNING *`, strings.Join(setClauses, ", "), argIndex)

	args = append(args, timelineId)

	var updatedTimeline model.Timeline
	err := r.db.Pgx.QueryRow(ctx, query, args...).Scan(
		&updatedTimeline.ID,
		&updatedTimeline.ApplicationID,
		&updatedTimeline.Content,
		&updatedTimeline.TimelineDate,
		&updatedTimeline.CreatedAt,
		&updatedTimeline.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updatedTimeline, nil
}

func (r *timelineRepo) Delete(ctx context.Context, timelineId int) error {
	query := `DELETE FROM timelines WHERE id=$1`

	result, err := r.db.Pgx.Exec(ctx, query, timelineId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("timeline not found")
	}

	return nil
}

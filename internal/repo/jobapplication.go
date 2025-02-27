package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type JobApplicationRepo interface {
	Save(ctx context.Context, jobApplication *model.JobApplication) (*model.JobApplication, error)
	FindById(ctx context.Context, id int) (*model.JobApplication, error)
	FindAllByUserId(ctx context.Context, uuid uuid.UUID) ([]model.JobApplication, error)
	Update(ctx context.Context, jobApplicationId int, updates map[string]interface{}) (*model.JobApplication, error)
	Delete(ctx context.Context, jobApplicationId int) error
}

type jobApplicationRepo struct {
	db *config.Database
}

func NewJobApplicationRepo(db *config.Database) JobApplicationRepo {
	return &jobApplicationRepo{db: db}
}

func (r *jobApplicationRepo) Save(ctx context.Context, jobApplication *model.JobApplication) (*model.JobApplication, error) {
	query := `INSERT INTO job_applications 
	(user_id, position, company, location, source, progress, applied_date, notes) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING *`

	var savedApplication model.JobApplication
	err := r.db.Pgx.QueryRow(ctx, query,
		jobApplication.UserID, jobApplication.Position, jobApplication.Company, jobApplication.Location,
		jobApplication.Source, jobApplication.Progress, jobApplication.AppliedDate, jobApplication.Notes,
	).Scan(
		&savedApplication.ID,
		&savedApplication.UserID,
		&savedApplication.Position,
		&savedApplication.Company,
		&savedApplication.Location,
		&savedApplication.Source,
		&savedApplication.Progress,
		&savedApplication.AppliedDate,
		&savedApplication.Notes,
		&savedApplication.CreatedAt,
		&savedApplication.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &savedApplication, nil

}

func (r *jobApplicationRepo) FindById(ctx context.Context, id int) (*model.JobApplication, error) {
	query := `SELECT * FROM job_applications WHERE id=$1`

	var jobApplication model.JobApplication
	err := r.db.Pgx.QueryRow(ctx, query, id).Scan(
		&jobApplication.ID,
		&jobApplication.UserID,
		&jobApplication.Position,
		&jobApplication.Company,
		&jobApplication.Location,
		&jobApplication.Source,
		&jobApplication.Progress,
		&jobApplication.AppliedDate,
		&jobApplication.Notes,
		&jobApplication.CreatedAt,
		&jobApplication.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &jobApplication, nil
}

func (r *jobApplicationRepo) FindAllByUserId(ctx context.Context, uuid uuid.UUID) ([]model.JobApplication, error) {
	query := `SELECT * FROM job_applications WHERE user_id=$1 ORDER BY updated_at DESC`

	rows, err := r.db.Pgx.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobApplications []model.JobApplication
	for rows.Next() {
		var jobApplication model.JobApplication
		err := rows.Scan(
			&jobApplication.ID,
			&jobApplication.UserID,
			&jobApplication.Position,
			&jobApplication.Company,
			&jobApplication.Location,
			&jobApplication.Source,
			&jobApplication.Progress,
			&jobApplication.AppliedDate,
			&jobApplication.Notes,
			&jobApplication.CreatedAt,
			&jobApplication.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		jobApplications = append(jobApplications, jobApplication)
	}

	return jobApplications, nil
}

func (r *jobApplicationRepo) Update(ctx context.Context, jobApplicationId int, updates map[string]interface{}) (*model.JobApplication, error) {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argIndex))
		args = append(args, value)
		argIndex++
	}

	query := fmt.Sprintf(`
		UPDATE job_applications 
		SET %s 
		WHERE id = $%d RETURNING *`, strings.Join(setClauses, ", "), argIndex)

	args = append(args, jobApplicationId)

	var updatedJob model.JobApplication
	err := r.db.Pgx.QueryRow(ctx, query, args...).Scan(
		&updatedJob.ID,
		&updatedJob.UserID,
		&updatedJob.Position,
		&updatedJob.Company,
		&updatedJob.Location,
		&updatedJob.Source,
		&updatedJob.Progress,
		&updatedJob.AppliedDate,
		&updatedJob.Notes,
		&updatedJob.CreatedAt,
		&updatedJob.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updatedJob, nil

}

func (r *jobApplicationRepo) Delete(ctx context.Context, jobApplicationId int) error {
	query := `DELETE FROM job_applications WHERE id=$1`

	result, err := r.db.Pgx.Exec(ctx, query, jobApplicationId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("job application not found")
	}

	return nil
}

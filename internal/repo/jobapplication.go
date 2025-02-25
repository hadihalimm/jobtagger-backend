package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type JobApplicationRepo interface {
}

type jobApplicationRepo struct {
	db *config.Database
}

func newJobApplicationRepo(db *config.Database) JobApplicationRepo {
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
		jobApplication.Source, jobApplication.Progress, jobApplication.AppliedDate,
	).Scan(
		&savedApplication.ID,
		&savedApplication.UserID,
		&savedApplication.Position,
		&savedApplication.Company,
		&savedApplication.Location,
		&savedApplication.Source,
		&savedApplication.Progress,
		&savedApplication.AppliedDate,
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
		&jobApplication.CreatedAt,
		&jobApplication.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &jobApplication, nil
}

func (r *jobApplicationRepo) FindAllByUserId(ctx context.Context, uuid uuid.UUID) ([]model.JobApplication, error) {
	query := `SELECT * FROM job_applications WHERE user_id=$1`

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

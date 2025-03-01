package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type InterviewRepo interface {
	Save(ctx context.Context, interview model.Interview) (*model.Interview, error)
	FindById(ctx context.Context, id int) (*model.Interview, error)
	FindAllByApplicationId(ctx context.Context, applicationId int) ([]model.Interview, error)
	Update(ctx context.Context, interviewId int, updates map[string]interface{}) (*model.Interview, error)
	Delete(ctx context.Context, interviewId int) error
}

type interviewRepo struct {
	db *config.Database
}

func NewInterviewRepo(db *config.Database) InterviewRepo {
	return &interviewRepo{db: db}
}

func (r *interviewRepo) Save(ctx context.Context, interview model.Interview) (*model.Interview, error) {
	query := `INSERT INTO interviews 
	(application_id, title, interview_date, position, company, notes) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING *`

	var savedInterview model.Interview
	err := r.db.Pgx.QueryRow(ctx, query,
		interview.ApplicationID, interview.Title, interview.Date,
		interview.Position, interview.Company, interview.Notes).Scan(
		&savedInterview.ID,
		&savedInterview.ApplicationID,
		&savedInterview.Title,
		&savedInterview.Date,
		&savedInterview.Position,
		&savedInterview.Company,
		&savedInterview.Notes,
		&savedInterview.CreatedAt,
		&savedInterview.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &savedInterview, nil
}

func (r *interviewRepo) FindById(ctx context.Context, id int) (*model.Interview, error) {
	query := `SELECT * FROM interviews WHERE id=$1`

	var interview model.Interview
	err := r.db.Pgx.QueryRow(ctx, query, id).Scan(
		&interview.ID,
		&interview.ApplicationID,
		&interview.Title,
		&interview.Date,
		&interview.Position,
		&interview.Company,
		&interview.Notes,
		&interview.CreatedAt,
		&interview.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &interview, nil
}

func (r *interviewRepo) FindAllByApplicationId(ctx context.Context, applicationId int) ([]model.Interview, error) {
	query := `SELECT * FROM interviews WHERE application_id=$1`

	rows, err := r.db.Pgx.Query(ctx, query, applicationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interviews []model.Interview
	for rows.Next() {
		var interview model.Interview
		err := rows.Scan(
			&interview.ID,
			&interview.ApplicationID,
			&interview.Title,
			&interview.Date,
			&interview.Position,
			&interview.Company,
			&interview.Notes,
			&interview.CreatedAt,
			&interview.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, interview)
	}
	return interviews, nil
}

func (r *interviewRepo) Update(ctx context.Context, interviewId int, updates map[string]interface{}) (*model.Interview, error) {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argIndex))
		args = append(args, value)
		argIndex++
	}

	query := fmt.Sprintf(`UPDATE interviews 
		SET %s 
		WHERE id = $%d 
		RETURNING *`, strings.Join(setClauses, ", "), argIndex)

	args = append(args, interviewId)

	var updatedInterview model.Interview
	err := r.db.Pgx.QueryRow(ctx, query, args...).Scan(
		&updatedInterview.ID,
		&updatedInterview.ApplicationID,
		&updatedInterview.Title,
		&updatedInterview.Date,
		&updatedInterview.Position,
		&updatedInterview.Company,
		&updatedInterview.Notes,
		&updatedInterview.CreatedAt,
		&updatedInterview.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updatedInterview, nil
}

func (r *interviewRepo) Delete(ctx context.Context, interviewId int) error {
	query := `DELETE FROM interviews WHERE id=$1`
	result, err := r.db.Pgx.Exec(ctx, query, interviewId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("intterview not found")
	}
	return nil
}

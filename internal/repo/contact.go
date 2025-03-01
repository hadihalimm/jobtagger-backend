package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type ContactRepo interface {
	Save(ctx context.Context, contact model.Contact) (*model.Contact, error)
	FindById(ctx context.Context, contactId int) (*model.Contact, error)
	FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]model.Contact, error)
	Update(ctx context.Context, contactId int, updates map[string]interface{}) (*model.Contact, error)
	Delete(ctx context.Context, contactId int) error
}

type contactRepo struct {
	db *config.Database
}

func NewContactRepo(db *config.Database) ContactRepo {
	return &contactRepo{db: db}
}

func (r *contactRepo) Save(ctx context.Context, contact model.Contact) (*model.Contact, error) {
	query := `INSERT INTO contacts 
	(user_id, name, email, phone, notes) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING *`

	var savedContact model.Contact
	err := r.db.Pgx.QueryRow(ctx, query,
		contact.UserID, contact.Name, contact.Email, contact.Phone, contact.Notes).Scan(
		&savedContact.ID,
		&savedContact.UserID,
		&savedContact.Name,
		&savedContact.Email,
		&savedContact.Phone,
		&savedContact.Notes,
		&savedContact.CreatedAt,
		&savedContact.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &savedContact, nil
}

func (r *contactRepo) FindById(ctx context.Context, contactId int) (*model.Contact, error) {
	query := `SELECT * FROM contacts WHERE id=$1`

	var contact model.Contact
	err := r.db.Pgx.QueryRow(ctx, query, contactId).Scan(
		&contact.ID,
		&contact.UserID,
		&contact.Name,
		&contact.Email,
		&contact.Phone,
		&contact.Notes,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *contactRepo) FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]model.Contact, error) {
	query := `SELECT * FROM contacts WHERE user_id=$1`

	rows, err := r.db.Pgx.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact
	for rows.Next() {
		var contact model.Contact
		err := rows.Scan(
			&contact.ID,
			&contact.UserID,
			&contact.Name,
			&contact.Email,
			&contact.Phone,
			&contact.Notes,
			&contact.CreatedAt,
			&contact.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *contactRepo) Update(ctx context.Context, contactId int, updates map[string]interface{}) (*model.Contact, error) {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argIndex))
		args = append(args, value)
		argIndex++
	}

	query := fmt.Sprintf(`UPDATE contacts 
		SET %s
		WHERE id = $%d
		RETURNING *`, strings.Join(setClauses, ", "), argIndex)

	args = append(args, contactId)

	var updatedContact model.Contact
	err := r.db.Pgx.QueryRow(ctx, query, args...).Scan(
		&updatedContact.ID,
		&updatedContact.UserID,
		&updatedContact.Name,
		&updatedContact.Email,
		&updatedContact.Phone,
		&updatedContact.Notes,
		&updatedContact.CreatedAt,
		&updatedContact.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updatedContact, nil
}

func (r *contactRepo) Delete(ctx context.Context, contactId int) error {
	query := `DELETE FROM contacts WHERE id=$1`

	result, err := r.db.Pgx.Exec(ctx, query, contactId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("cannot delete, contact not found")
	}
	return nil
}

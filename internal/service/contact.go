package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
)

type ContactService interface {
	Create(ctx context.Context, req request.CreateContact, userId string) (*model.Contact, error)
	FindById(ctx context.Context, contactId int) (*model.Contact, error)
	FindAllByUserId(ctx context.Context, userId string) ([]model.Contact, error)
	Update(ctx context.Context, contactId int, req request.UpdateContact) (*model.Contact, error)
	Delete(ctx context.Context, contactId int) error
}

type contactService struct {
	contactRepo repo.ContactRepo
}

func NewContactService(contactRepo repo.ContactRepo) ContactService {
	return &contactService{contactRepo: contactRepo}
}

func (s *contactService) Create(ctx context.Context, req request.CreateContact, userId string) (*model.Contact, error) {
	var contact model.Contact

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	contact.UserID = parsedUUID
	contact.Name = req.Name
	contact.Email = req.Email
	contact.Phone = req.Phone
	contact.Notes = req.Notes

	return s.contactRepo.Save(ctx, contact)
}

func (s *contactService) FindById(ctx context.Context, contactId int) (*model.Contact, error) {
	return s.contactRepo.FindById(ctx, contactId)
}
func (s *contactService) FindAllByUserId(ctx context.Context, userId string) ([]model.Contact, error) {
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return s.contactRepo.FindAllByUserId(ctx, parsedUUID)
}

func (s *contactService) Update(ctx context.Context, contactId int, req request.UpdateContact) (*model.Contact, error) {
	updates := map[string]interface{}{}
	addIfNotNil(updates, "name", req.Name)
	addIfNotNil(updates, "email", req.Email)
	addIfNotNil(updates, "phone", req.Phone)
	addIfNotNil(updates, "notes", req.Notes)

	return s.contactRepo.Update(ctx, contactId, updates)
}

func (s *contactService) Delete(ctx context.Context, contactId int) error {
	return s.contactRepo.Delete(ctx, contactId)
}

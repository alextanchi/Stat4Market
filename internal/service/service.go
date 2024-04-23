package service

import (
	"Stat4Market/internal/convert"
	"Stat4Market/internal/models"
	"Stat4Market/internal/repository"
	"context"
)

type Service interface {
	CreateEvent(ctx context.Context, event models.Request) error
}
type EventService struct {
	storage repository.Repository
}

func NewService(repos repository.Repository) Service {
	return EventService{
		storage: repos}
}

func (s EventService) CreateEvent(ctx context.Context, event models.Request) error {
	data := convert.RequestToDomain(event)
	return s.storage.CreateEvent(ctx, data)
}

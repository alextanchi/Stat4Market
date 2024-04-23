package convert

import (
	"Stat4Market/internal/domain"
	"Stat4Market/internal/models"
)

func RequestToDomain(request models.Request) domain.Event {
	return domain.Event{
		EventType: request.EventType,
		UserId:    request.UserID,
		EventTime: request.EventTime,
		Payload:   request.Payload,
	}
}

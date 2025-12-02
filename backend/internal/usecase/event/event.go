package event

import (
	"context"
	"time"

	"github.com/goawwer/devclash/internal/domain"
	eventmodel "github.com/goawwer/devclash/internal/domain/event_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/pkg/helpers"
	"github.com/google/uuid"
)

func (e *EventUsecase) Create(ctx context.Context, id uuid.UUID, role string, input *dto.EventCreationRequest) error {
	if role == "user" {
		return domain.ErrNotForUserRole
	}

	eventID := uuid.New()

	orgID, err := e.r.GetOrganizerIDByAccountID(ctx, id)
	if err != nil {
		return err
	}

	eventTypeID, err := e.r.GetEventTypeIDByName(ctx, input.Type)
	if err != nil {
		return err
	}

	techIDs := make([]uuid.UUID, 0, len(input.TechStack))
	for _, name := range input.TechStack {
		techID, err := e.r.GetTechnologyIDByName(ctx, name)
		if err != nil {
			return err
		}
		techIDs = append(techIDs, techID)
	}

	techIDStrings := make([]string, len(techIDs))
	for i, id := range techIDs {
		techIDStrings[i] = id.String()
	}

	return e.r.CreateEvent(ctx, &eventmodel.Event{
		ID:          eventID,
		OrganizerID: orgID,
		TypeID:      eventTypeID,
		Title:       input.Title,
		CreatedAt:   time.Now(),

		Properties: eventmodel.Properties{
			EventID:       eventID,
			IsOnline:      input.IsOnline,
			IsFree:        input.IsFree,
			NumberOfTeams: input.NumberOfTeams,
			TeamSize:      input.TeamSize,
		},

		Details: eventmodel.Details{
			EventID:         eventID,
			EventPictureURL: input.EventPictureURL,
			StartTime:       input.StartTime,
			EndTime:         input.EndTime,
			Description:     input.Description,
			Prize:           input.Prize,
		},

		Technologies: techIDStrings,
	})
}

func (e *EventUsecase) UpdatePictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	return e.r.UpdateEventPictureByCreatorID(ctx, newURL, accountID)
}

func (e *EventUsecase) GetEventPageByID(ctx context.Context, id uuid.UUID) (*dto.EventPage, error) {
	event, err := e.r.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	eventType, err := e.r.GetEventTypeNameByID(ctx, event.TypeID)
	if err != nil {
		return nil, err
	}

	techIDs := make([]uuid.UUID, 0, len(event.Technologies))
	for _, s := range event.Technologies {
		if s != "" {
			if parsedID, parseErr := uuid.Parse(s); parseErr == nil {
				techIDs = append(techIDs, parsedID)
			}
		}
	}

	techIDToNameMap, err := e.r.GetTechnologyNamesByIDs(ctx, techIDs)
	if err != nil {
		return nil, err
	}

	finalTechNames := make([]string, 0, len(techIDs))
	for _, id := range techIDs {
		if name, ok := techIDToNameMap[id]; ok {
			finalTechNames = append(finalTechNames, name)
		}
	}

	teamIDs := make([]uuid.UUID, 0, len(event.TeamsIDs))
	for _, s := range event.TeamsIDs {
		if s != "" {
			teamIDs = append(teamIDs, uuid.MustParse(s))
		}
	}

	teamsSlice, err := e.r.GetTeamsByIDs(ctx, teamIDs)
	if err != nil {
		return nil, err
	}

	teamName := ""
	teamPictureURL := ""
	teamStatus := ""

	if len(teamsSlice) > 0 {
		team := teamsSlice[0]

		teamName = team.Name

		if team.TeamPictureURL != nil {
			teamPictureURL = *team.TeamPictureURL
		}
		if team.TeamStatus != nil {
			teamStatus = *team.TeamStatus
		}
	}

	return &dto.EventPage{
		Title:           event.Title,
		Description:     event.Details.Description,
		EventPictureURL: event.Details.EventPictureURL,
		Type:            eventType,
		IsOnline:        event.Properties.IsOnline,
		IsFree:          event.Properties.IsFree,
		NumberOfTeams:   event.Properties.NumberOfTeams,
		TeamSize:        event.Properties.TeamSize,
		TechStack:       finalTechNames,
		TeamName:        teamName,
		TeamPictureURL:  teamPictureURL,
		TeamStatus:      teamStatus,
		StartTime:       event.Details.StartTime,
		EndTime:         event.Details.EndTime,
	}, nil
}

func (e *EventUsecase) GetAllEvents(ctx context.Context, filterParams helpers.FilterParameters) ([]*dto.EventListResponse, error) {
	events, err := e.r.GetAllEvents(ctx, filterParams)
	if err != nil {
		return nil, err
	}

	allUniqueTechIDs := make(map[uuid.UUID]struct{})
	for _, event := range events {
		for _, s := range event.Technologies {
			if s != "" {
				if id, err := uuid.Parse(s); err == nil {
					allUniqueTechIDs[id] = struct{}{}
				}
			}
		}
	}

	techIDsToFetch := make([]uuid.UUID, 0, len(allUniqueTechIDs))
	for id := range allUniqueTechIDs {
		techIDsToFetch = append(techIDsToFetch, id)
	}

	idToNameMap, err := e.r.GetTechnologyNamesByIDs(ctx, techIDsToFetch)
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.EventListResponse, 0, len(events))

	for _, event := range events {
		eventTechNames := make([]string, 0, len(event.Technologies))
		for _, techIDStr := range event.Technologies {
			if id, err := uuid.Parse(techIDStr); err == nil {
				if name, ok := idToNameMap[id]; ok {
					eventTechNames = append(eventTechNames, name)
				}
			}
		}

		dto := &dto.EventListResponse{
			ID:            event.ID.String(),
			Title:         event.Title,
			OrganizerName: event.OrganizerName,
			EventTypeName: event.EventTypeName,
			IsOnline:      event.Properties.IsOnline,
			IsFree:        event.Properties.IsFree,
			StartTime:     event.Details.StartTime,
			EndTime:       event.Details.EndTime,
			Description:   event.Details.Description,
			TechStack:     eventTechNames,
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}

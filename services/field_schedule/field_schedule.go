package services

import (
	"context"
	"field-service/common/util"
	"field-service/constants"
	errFieldSchedule "field-service/constants/error/fieldSchedule"
	"field-service/domain/dto"
	"field-service/domain/models"
	"field-service/repositories"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type FieldScheduleService struct {
	repository repositories.IRepositoryRegistry
}

type IFieldScheduleService interface {
	GetAllWithPagination(context.Context, *dto.FieldScheduleRequestParam) (*util.PaginationResult, error)
	GetAllByFieldIDAndDate(context.Context, string, string) ([]dto.FieldScheduleForBookingResponse, error)
	GetByUUID(context.Context, string) (*dto.FieldScheduleResponse, error)
	GenerateScheduleForOneMonth(context.Context, *dto.GenerateFieldScheduleForOneMonthRequest) error
	Create(context.Context, *dto.FieldScheduleRequest) error
	Update(context.Context, string, *dto.FieldScheduleRequest) (*dto.FieldScheduleResponse, error)
	UpdateStatus(context.Context, *dto.FieldScheduleRequest) error
	Delete(context.Context, string) error
}

func NewFieldScheduleService(repository repositories.IRepositoryRegistry) IFieldScheduleService {
	return &FieldScheduleService{repository: repository}
}

func (f *FieldScheduleService) GetAllWithPagination(ctx context.Context, param *dto.FieldScheduleRequestParam) (*util.PaginationResult, error) {
	fieldSchedules, total, err := f.repository.GetFieldSchedule().FindAllWithPagination(ctx, param)
	if err != nil {
		return nil, err
	}

	fieldScheduleResults := make([]dto.FieldScheduleResponse, 0, len(fieldSchedules))
	for _, schedule := range fieldSchedules {
		fieldScheduleResults = append(fieldScheduleResults, dto.FieldScheduleResponse{
			UUID:         schedule.UUID,
			FieldName:    schedule.Field.Name,
			Date:         schedule.Date.Format("2006-01-02"),
			PricePerHour: schedule.Field.PricePerHour,
			Status:       schedule.Status.GetStatusString(),
			Time:         fmt.Sprintf("%s - %s", schedule.Time.StartTime, schedule.Time.EndTime),
			CreatedAt:    schedule.CreatedAt,
			UpdatedAt:    schedule.UpdatedAt,
		})
	}

	pagination := &util.PaginationParam{
		Count: total,
		Limit: param.Limit,
		Page:  param.Page,
		Data:  fieldScheduleResults,
	}

	response := util.GeneratePagination(*pagination)
	return &response, nil
}

func (f *FieldScheduleService) convertMonthName(inputDate string) string {
	date, err := time.Parse(time.DateOnly, inputDate)
	if err != nil {
		return ""
	}

	indonesianMonth := map[string]string{
		"Jan": "Jan",
		"Feb": "Feb",
		"Mar": "Mar",
		"Apr": "Apr",
		"May": "Mei",
		"Jun": "Jun",
		"Jul": "Jul",
		"Aug": "Agu",
		"Sep": "Sep",
		"Oct": "Okt",
		"Nov": "Nov",
		"Dec": "Des",
	}

	formattedDate := date.Format("02 Jan")
	day := formattedDate[:3]
	month := formattedDate[3:]
	formattedDate = fmt.Sprintf("%s %s", day, indonesianMonth[month])
	return formattedDate
}

func (f *FieldScheduleService) GetAllByFieldIDAndDate(ctx context.Context, uuid, date string) ([]dto.FieldScheduleForBookingResponse, error) {
	field, err := f.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	fieldSchedules, err := f.repository.GetFieldSchedule().FindAllByFieldIDAndDate(ctx, int(field.ID), date)
	if err != nil {
		return nil, err
	}

	fieldScheduleResults := make([]dto.FieldScheduleForBookingResponse, 0, len(fieldSchedules))
	for _, schedule := range fieldSchedules {
		pricePerHour := float64(schedule.Field.PricePerHour)
		fieldScheduleResults = append(fieldScheduleResults, dto.FieldScheduleForBookingResponse{
			UUID:         schedule.UUID,
			Date:         f.convertMonthName(schedule.Date.Format(time.DateOnly)),
			Time:         schedule.Time.StartTime,
			Status:       schedule.Status.GetStatusString(),
			PricePerHour: util.RupiahFormat(&pricePerHour),
		})
	}

	return fieldScheduleResults, nil
}

func (f *FieldScheduleService) GetByUUID(ctx context.Context, uuid string) (*dto.FieldScheduleResponse, error) {
	fieldSchedule, err := f.repository.GetFieldSchedule().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	response := dto.FieldScheduleResponse{
		UUID:         fieldSchedule.UUID,
		FieldName:    fieldSchedule.Field.Name,
		PricePerHour: fieldSchedule.Field.PricePerHour,
		Date:         fieldSchedule.Date.Format(time.DateOnly),
		Status:       fieldSchedule.Status.GetStatusString(),
		CreatedAt:    fieldSchedule.CreatedAt,
		UpdatedAt:    fieldSchedule.UpdatedAt,
	}

	return &response, nil
}

func (f *FieldScheduleService) GenerateScheduleForOneMonth(ctx context.Context, request *dto.GenerateFieldScheduleForOneMonthRequest) error {
	field, err := f.repository.GetField().FindByUUID(ctx, request.FieldID)
	if err != nil {
		return err
	}

	times, err := f.repository.GetTime().FindAll(ctx)
	if err != nil {
		return err
	}

	numberOfDays := 30
	fieldSchedules := make([]models.FieldSchedule, 0, numberOfDays)
	now := time.Now().Add(time.Duration(1) * 24 * time.Hour)
	for i := 0; i < numberOfDays; i++ {
		currentDate := now.AddDate(0, 0, i)
		for _, timeItem := range times {
			schedule, err := f.repository.GetFieldSchedule().FindByDateAndTimeID(
				ctx,
				currentDate.Format(time.DateOnly),
				int(timeItem.ID),
				int(field.ID),
			)
			if err != nil {
				return err
			}
			if schedule != nil {
				return errFieldSchedule.ErrFieldScheduleIsExist
			}

			fieldSchedules = append(fieldSchedules, models.FieldSchedule{
				UUID:    uuid.New(),
				FieldID: field.ID,
				TimeID:  timeItem.ID,
				Date:    currentDate,
				Status:  constants.Available,
			})
		}
	}
	err = f.repository.GetFieldSchedule().Create(ctx, fieldSchedules)
	if err != nil {
		return err
	}

	return nil
}

func (f *FieldScheduleService) Create(ctx context.Context, request *dto.FieldScheduleRequest) error {
	field, err := f.repository.GetField().FindByUUID(ctx, request.FieldID)
	if err != nil {
		return err
	}

	fieldSchedules := make([]models.FieldSchedule, 0, len(request.TimeIDs))
	dateParsed, _ := time.Parse(time.DateOnly, request.Date)
	for _, timeID := range request.TimeIDs {
		scheduleTime, err := f.repository.GetTime().FindByUUID(ctx, timeID)
		if err != nil {
			return err
		}

		schedule, err := f.repository.GetFieldSchedule().FindByDateAndTimeID(ctx, request.Date, int(scheduleTime.ID), int(field.ID))
		if err != nil {
			return err
		}

		if schedule != nil {
			return errFieldSchedule.ErrFieldScheduleIsExist
		}

		fieldSchedules = append(fieldSchedules, models.FieldSchedule{
			UUID:    uuid.New(),
			FieldID: field.ID,
			TimeID:  scheduleTime.ID,
			Date:    dateParsed,
			Status:  constants.Available,
		})
	}

	err = f.repository.GetFieldSchedule().Create(ctx, fieldSchedules)
	if err != nil {
		return err
	}
	return nil
}

func (f *FieldScheduleService) Update(ctx context.Context, s string, request *dto.FieldScheduleRequest) (*dto.FieldScheduleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FieldScheduleService) UpdateStatus(ctx context.Context, request *dto.FieldScheduleRequest) error {
	//TODO implement me
	panic("implement me")
}

func (f *FieldScheduleService) Delete(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

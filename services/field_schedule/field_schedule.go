package services

import (
	"context"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
	"fmt"
	"time"
)

type FieldScheduleService struct {
	repository repositories.IRepositoryRegistry
}

type IFieldScheduleService interface {
	GetAllWithPagination(context.Context, *dto.FieldRequestParam) (*util.PaginationResult, error)
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

func (f *FieldScheduleService) GetAllByFieldIDAndDate(ctx context.Context, s string, s2 string) ([]dto.FieldScheduleForBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FieldScheduleService) GetByUUID(ctx context.Context, s string) (*dto.FieldScheduleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FieldScheduleService) GenerateScheduleForOneMonth(ctx context.Context, request *dto.GenerateFieldScheduleForOneMonthRequest) error {
	//TODO implement me
	panic("implement me")
}

func (f *FieldScheduleService) Create(ctx context.Context, request *dto.FieldScheduleRequest) error {
	//TODO implement me
	panic("implement me")
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

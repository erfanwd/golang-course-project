package service

import (
	"context"

	"github.com/erfanwd/golang-course-project/common"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/repository"
)


type BaseService[TModel any, TCreate any, TUpdate any, TResponse any] struct {
	logger logging.Logger
	repository repository.BaseRepo[TModel]
}

func NewBaseService[TModel any, TCreate any, TUpdate any, TResponse any](cfg *config.Config) *BaseService[TModel, TCreate, TUpdate, TResponse] {
	baseService := &BaseService[TModel,TCreate,TUpdate,TResponse]{
		logger: logging.NewLogger(cfg),
		repository: *repository.NewBaseRepository[TModel](db.GetDb()),
	}
	return baseService
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Create(ctx context.Context, req TCreate) (TResponse, error) {
	var response TResponse

	model, _ := common.StructToStructMapper[TModel](req)

	model, err := s.repository.Create(ctx, &model)

	if err != nil {
		return response, err
	}

	response, _ = common.StructToStructMapper[TResponse](model)

	return response, nil
}
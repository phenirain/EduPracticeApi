package service

import (
	"api/internal/infrastructure"
	"context"
	"errors"
	"fmt"
)

func ErrNotFound(tableName string, id int32) error {
	return errors.New(fmt.Sprintf("Object of %s with id: %d not found", tableName, id))
}

type Repository[TDB infrastructure.DbModel[T], T infrastructure.Model] interface {
	Create(ctx context.Context, model T) (T, error)
	ExistsById(ctx context.Context, id int32) (bool, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, model T) error
	Delete(ctx context.Context, id int32) error
}

type ModelRequest[T infrastructure.Model] interface {
	ToModel() (T, error)
}

type Service[CreateRequest ModelRequest[model], UpdateRequest ModelRequest[model], model infrastructure.Model,
	TDB infrastructure.DbModel[model]] struct {
	Repository Repository[TDB, model]
}

func NewService[CreateRequest ModelRequest[model], UpdateRequest ModelRequest[model],
	model infrastructure.Model, TDB infrastructure.DbModel[model]](repo Repository[TDB, model]) *Service[CreateRequest, UpdateRequest,
	model, TDB] {
	return &Service[CreateRequest, UpdateRequest, model, TDB]{Repository: repo}
}

func (s *Service[CreateRequest, UpdateRequest, model, TDB]) Create(ctx context.Context, request CreateRequest) (*model, error) {
	requestModel, err := request.ToModel()
	if err != nil {
		return nil, fmt.Errorf("invalid model: %v", err)
	}
	item, err := s.Repository.Create(ctx, requestModel)
	var dbModel TDB
	if err != nil {
		return nil, fmt.Errorf("errored creating new item of %s: %v", dbModel.TableName(), err)
	}
	return &item, nil
}

func (s *Service[CreateRequest, UpdateRequest, model, TDB]) GetAll(ctx context.Context) ([]model, error) {
	allItems, err := s.Repository.GetAll(ctx)
	var dbModel TDB
	if err != nil {
		return nil, fmt.Errorf("errored getting all %s: %v", dbModel.TableName(), err)
	}

	return allItems, nil
}

func (s *Service[CreateRequest, UpdateRequest, model, TDB]) Update(ctx context.Context,
	id int32, request UpdateRequest) error {
	exists, err := s.Repository.ExistsById(ctx, id)
	if err != nil {
		return err
	}
	var dbModel TDB
	if !exists {
		return ErrNotFound(dbModel.TableName(), id)
	}

	requestModel, err := request.ToModel()
	if err != nil {
		return fmt.Errorf("invalid model: %v", err)
	}
	err = s.Repository.Update(ctx, requestModel)
	if err != nil {
		return fmt.Errorf("errored updating %s with id: %d: %v", dbModel.TableName(), id, err)
	}
	return nil
}

func (s *Service[CreateRequest, UpdateRequest, model, TDB]) Delete(ctx context.Context, id int32) error {
	exists, err := s.Repository.ExistsById(ctx, id)
	if err != nil {
		return err
	}
	var dbModel TDB
	if !exists {
		return ErrNotFound(dbModel.TableName(), id)
	}

	err = s.Repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("errored deleting %s with id: %d: %v", dbModel.TableName(), id, err)
	}
	return nil
}

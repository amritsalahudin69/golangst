package repository

import (
	"context"
	"salt-academy_learn_week2/domain/campus/model"
)

type CampusRepository interface {
	Get(ctx context.Context, campusID int) (*model.Campus, error)
	Create(ctx context.Context, campus model.Campus) error
	Update(ctx context.Context, campus model.Campus) error
	Delete(ctx context.Context, campusID int) error
	GetList(ctx context.Context) ([]*model.Campus, error)
}

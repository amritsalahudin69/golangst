package repository

import (
	"context"
	"salt-academy_learn_week2/model"
)

type CampusTemplate interface {
	Get(ctx context.Context) (*model.Campus, error)
	Create(ctx context.Context, Campus model.Campus) error
	Update(ctx context.Context, CampusId int, Campus model.Campus) error
	Delete(ctx context.Context, CampusId int) error
	GetList(ctx context.Context) ([]*model.Campus, error)
}

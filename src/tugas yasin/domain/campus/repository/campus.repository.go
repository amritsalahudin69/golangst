package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"salt-academy_learn_week2/domain/campus/model"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type campusRepository struct {
	conn *sql.DB
}

func NewCampusRepositoryImpl(conn *sql.DB) CampusRepository {
	return &campusRepository{conn: conn}
}

func (r *campusRepository) Get(ctx context.Context, campusID int) (*model.Campus, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE campus_id = ? ORDER BY created_at DESC LIMIT 1 `, model.Campus{}.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, r.conn, stmt, opts, campusID)

	if result != nil {
		return result.(*model.Campus), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (r *campusRepository) Create(ctx context.Context, campus model.Campus) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("INSERT INTO %s (name, address, rektor, akreditasi, email, created_at) VALUES (?, ?, ?, ?, ?)", campus.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: campus, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		r.conn,
		stmt,
		opts,
		campus.Name,
		campus.Address,
		campus.Rektor,
		campus.Akreditasi,
		campus.Email,
		campus.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *campusRepository) Update(ctx context.Context, campus model.Campus) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("INSERT INTO %s (name, address, rektor, akreditasi, email, updated_at) VALUES (?, ?, ?, ?, ?)", campus.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: campus, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		r.conn,
		stmt,
		opts,
		campus.Name,
		campus.Address,
		campus.Rektor,
		campus.Akreditasi,
		campus.Email,
		campus.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *campusRepository) Delete(ctx context.Context, campusID int) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var campus model.Campus
	deleted := time.Now()
	stmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", campus.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: campus, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(ctx, r.conn, stmt, opts, deleted, campusID)
	if err != nil {
		return err
	}

	return nil
}

func (r *campusRepository) GetList(ctx context.Context) ([]*model.Campus, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s ORDER BY created_at DESC`, model.Campus{}.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: []model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, r.conn, stmt, opts)

	if result != nil {
		return result.([]*model.Campus), nil
	}

	return nil, errors.New("DATA EMPTY")
}

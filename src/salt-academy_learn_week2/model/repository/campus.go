package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"reflect"
	_ "reflect"
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/model"
	"time"
)

type CampusInteractor struct {
	conn *sql.DB
}

func NewCampusRepositoryMysql(connMysql *sql.DB) repository.CampusTemplate {
	return &CampusInteractor{conn: connMysql}
}

func (c *CampusInteractor) Get(ctx context.Context) (*model.Campus, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s LIMIT 1 ORDER BY id DESC`, model.Campus{}.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, c.conn, stmt, opts)

	if result != nil {
		return result.(*model.Campus), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (c *CampusInteractor) GetList(ctx context.Context) ([]*model.Campus, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.Campus{}.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, c.conn, stmt, opts)

	if result != nil || len(result.([]*model.Campus)) == 0 {
		return result.([]*model.Campus), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (c *CampusInteractor) Create(ctx context.Context, Campus model.Campus) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`INSERT INTO %s (nama, email, Address) VALUES (?, ?, ?)`, model.Campus{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	_, err := dbq.E(
		ctx,
		c.conn,
		stmt,
		ops,
		Campus.Name,
		Campus.Email,
		Campus.Address,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func (c *CampusInteractor) Update(ctx context.Context, CampusId int, Campus model.Campus) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET nama = ?, email= ?,  address = ?WHERE nim = ?`, model.Campus{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		c.conn,
		stmt,
		ops,
		CampusId,
		Campus.Name,
		Campus.Email,
		Campus.Address,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (c *CampusInteractor) Delete(ctx context.Context, CampusId int) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, model.Campus{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Campus{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		c.conn,
		stmt,
		ops,
		CampusId,
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
func isZero(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsZero()
}

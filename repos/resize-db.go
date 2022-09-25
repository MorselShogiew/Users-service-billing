package repos

import (
	"encoding/json"

	"github.com/MorselShogiew/Users-service-billing/errs"
	"github.com/MorselShogiew/Users-service-billing/logger"
	"github.com/MorselShogiew/Users-service-billing/models"
	"github.com/MorselShogiew/Users-service-billing/provider"
	"github.com/jmoiron/sqlx"
)

type ResizeDBRepo interface {
	GetUserBalance(id int) (float64, error)
	СreditingFunds(id int, value float64) error
	DebitingFunds(id int, value float64) error
	TransferFunds(idFrom int, idTo int, value float64) error
}

type resizeDB struct {
	db *sqlx.DB
	logger.Logger
}

func NewResizeDBRepo(p provider.Provider, l logger.Logger) ResizeDBRepo {
	return &resizeDB{p.GetResizeDBConn(), l}
}

func (r resizeDB) GetUserBalance(id int) (float64, error) {
	var data []byte
	query := "select * from users where user_id=$1"
	if err := r.db.QueryRow(query, id).Scan(&data); err != nil {
		return 0, errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}
	var res models.User
	if err := json.Unmarshal(data, &res); err != nil {
		return 0, errs.New(err, errs.ErrJSONDecode, true, 500)
	}
	return res.Balance, nil
}

func (r resizeDB) СreditingFunds(id int, value float64) error {

	query := "INSERT INTO public.users (user_id, balance) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET balance=users.balance+$2;"

	r.db.QueryRow(query, id, value)
	return nil
}

func (r resizeDB) DebitingFunds(id int, value float64) error {
	var data []byte

	query := "select * from users where user_id=$1"
	if err := r.db.QueryRow(query, id).Scan(&data); err != nil {
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}
	var res models.User
	if err := json.Unmarshal(data, &res); err != nil {
		return errs.New(err, errs.ErrJSONDecode, true, 500)
	}

	newBalance := res.Balance - value
	if newBalance < 0 {
		return errs.New(nil, errs.ErrNotEnoughFunds, true, 500)
	}

	querySet := "UPDATE users SET balance=balance-$1 WHERE user_id=$2;"
	r.db.QueryRow(querySet, newBalance, id)
	return nil
}

func (r resizeDB) TransferFunds(idFrom int, idTo int, value float64) error {
	var data1 []byte
	query := "select * from users where user_id=$1"
	if err := r.db.QueryRow(query, idFrom).Scan(&data1); err != nil {
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}
	var res1 models.User
	if err := json.Unmarshal(data1, &res1); err != nil {
		return errs.New(err, errs.ErrJSONDecode, true, 500)
	}

	newBalance1 := res1.Balance - value
	if newBalance1 < 0 {
		return errs.New(nil, errs.ErrNotEnoughFunds, true, 500)
	}
	var data2 []byte
	query = "select * from users where user_id=$1"
	if err := r.db.QueryRow(query, idTo).Scan(&data2); err != nil {
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}
	var res2 models.User
	if err := json.Unmarshal(data2, &res2); err != nil {
		return errs.New(err, errs.ErrJSONDecode, true, 500)
	}

	newBalance2 := res2.Balance + value

	tx, err := r.db.Begin()

	queryWithdraw := "UPDATE users SET balance=$1 WHERE user_id=$2;"
	queryAdd := "UPDATE users SET balance=$1 WHERE user_id=$2;"

	if _, err := tx.Exec(queryWithdraw, newBalance1, idFrom); err != nil {
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}

	if _, err := tx.Exec(queryAdd, newBalance2, idTo); err != nil {
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errs.New(err, errs.ErrDatabaseRequest, true, 500)
	}
	return nil
}

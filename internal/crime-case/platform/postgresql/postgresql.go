package postgresql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ellipizle/crime-report/pkg/id"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

type SQLStore struct {
	conn *sql.DB
}

func New(db *sql.DB) *SQLStore {
	return &SQLStore{db}
}

// CreateUser create a new account
func (store *SQLStore) Create(ctx echo.Context, acct *model.Case) (*model.Case, error) {
	acct.Id = id.GenerateNewUniqueCode()
	acct.CreatedAt = time.Now()
	acct.UpdatedAt = time.Now()
	statements := `INSERT INTO user VALUES($1, $2, $3, $4,$5, $6)`
	result, err := store.conn.Exec(statements, acct.Id, acct.Crime, acct.Reporter, acct.Handler, acct.CreatedAt, acct.UpdatedAt)
	if err != nil {
		return acct, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return acct, errors.New("failed to created account")
	}

	return acct, nil
}

// UpdateUser update an account
func (store *SQLStore) Update(ctx echo.Context, acct *model.Case) (*model.Case, error) {
	acct.UpdatedAt = time.Now()
	statements := `UPDATE  user SET account_balance = $1, updated_at = $2,  WHERE account_number = $3)`
	result, err := store.conn.Exec(statements, acct.Crime, acct.UpdatedAt, acct.Reporter)
	if err != nil {
		return acct, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return acct, errors.New("failed to created account")
	}
	return acct, nil
}

// GetUser retreive a single account
func (store *SQLStore) View(ctx echo.Context, id string) (*model.CaseResponse, error) {
	var acct = new(model.CaseResponse)
	row := store.conn.QueryRow("SELECT * FROM user WHERE account_number = $1", id)
	err := row.Scan(&acct.Id, &acct.Crime, &acct.Reporter, &acct.Handler, &acct.CreatedAt)
	if err == sql.ErrNoRows {
		return acct, errors.New("no record found")
	} else if err != nil {
		return acct, err
	}
	return acct, nil
}

// GetUser retreive a single account
func (store *SQLStore) Delete(ctx echo.Context, id string) error {
	// var acct = new(model.CaseResponse)
	// row := store.conn.QueryRow("SELECT * FROM user WHERE account_number = $1", id)
	// err := row.Scan(&acct.Id, &acct.Crime, &acct.Reporter, &acct.Handler, &acct.CreatedAt)
	// if err == sql.ErrNoRows {
	// 	return acct, errors.New("no record found")
	// } else if err != nil {
	// 	return acct, err
	// }
	return nil
}

// GetUsers get filtered user
func (store *SQLStore) List(ctx echo.Context, filter *model.CaseFilterParams) ([]*model.CaseResponse, error) {
	var user []*model.CaseResponse
	statements := `SELECT * FROM user WHERE transaction_type  LIKE = $1% ORDER BY created_at LIMIT =$2, OFFSET =$3`
	rows, err := store.conn.Query(statements, filter.Query, filter.Page, filter.Limit)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		tran := new(model.CaseResponse)
		err := rows.Scan(&tran.Id, &tran.Crime, &tran.Reporter, &tran.Handler, &tran.CreatedAt)
		if err != nil {
			return user, err
		}
		user = append(user, tran)
	}
	if err = rows.Err(); err != nil {
		return user, err
	}

	return user, nil
}

func (store *SQLStore) Count(ctx echo.Context, p *model.CaseFilterParams) int {

	return 0
}

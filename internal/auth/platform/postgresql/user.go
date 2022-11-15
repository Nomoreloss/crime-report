package postgresql

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"database/sql"

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

// View returns single user by ID
func (store *SQLStore) View(ctx echo.Context, id string) (*model.User, error) {
	var user = new(model.User)
	row := store.conn.QueryRow("SELECT * FROM user WHERE id = $1", id)
	err := row.Scan(&user.Id,
		&user.FirstName, &user.LastName, &user.UserType,
		&user.Mobile, &user.About, &user.Address, &user.Status,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return user, errors.New("no record found")
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// Register queries for single user by username
func (store *SQLStore) Register(ctx echo.Context, user *model.User) (*model.User, error) {
	user.Id = id.GenerateNewUniqueCode()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	statements := `INSERT INTO user VALUES($1, $2, $3, $4,$5, $6,$7, $8, $9, $10,$11, $12)`
	result, err := store.conn.Exec(statements, user.Id, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Active, user.Status, user.UserType, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return user, errors.New("failed to created account")
	}

	return user, nil
}

// Activate queries for single user by username
func (store *SQLStore) Activate(ctx echo.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	statements := `UPDATE user SET active = $1, updated_at = $2 WHERE id = $3`
	result, err := store.conn.Exec(statements, user.Active, user.UpdatedAt, user.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("failed to created account")
	}
	return nil
}

// FindByUsername queries for single user by username
func (store *SQLStore) FindByUsername(ctx echo.Context, uname string) (*model.User, error) {
	var user = new(model.User)
	row := store.conn.QueryRow("SELECT * FROM user WHERE username = $1", uname)
	err := row.Scan(&user.Id,
		&user.FirstName, &user.LastName, &user.UserType,
		&user.Mobile, &user.About, &user.Address, &user.Status,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return user, errors.New("no record found")
	} else if err != nil {
		return user, err
	}
	return user, nil
}

// FindByEmail queries for single user by email
func (store *SQLStore) FindByEmail(ctx echo.Context, email string) (*model.User, error) {
	var user = new(model.User)
	row := store.conn.QueryRow("SELECT * FROM user WHERE email = $1", email)
	err := row.Scan(&user.Id,
		&user.FirstName, &user.LastName, &user.UserType,
		&user.Mobile, &user.About, &user.Address, &user.Status,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return user, errors.New("no record found")
	} else if err != nil {
		return user, err
	}
	return user, nil
}

// FindByMobile queries for single user by mobie
func (store *SQLStore) FindByMobile(ctx echo.Context, mobile string) (*model.User, error) {
	var user = new(model.User)
	row := store.conn.QueryRow("SELECT * FROM user WHERE mobile = $1", mobile)
	err := row.Scan(&user.Id,
		&user.FirstName, &user.LastName, &user.UserType,
		&user.Mobile, &user.About, &user.Address, &user.Status,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return user, errors.New("no record found")
	} else if err != nil {
		return user, err
	}
	return user, nil
}

// Update updates user's info
func (store *SQLStore) Update(ctx echo.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	statements := `UPDATE user SET first_name =$2, last_name=$3, username=$4, email=$5, status=$6, type=$7, role=$8, updated_at$9
								WHERE id=$1`
	result, err := store.conn.Exec(statements, user.Id, user.FirstName, user.LastName, user.Username, user.Email, user.Status, user.UserType, user.Role, user.UpdatedAt)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("failed to created account")
	}
	return nil
}

// GenerateCode activatio code
func (store *SQLStore) GenerateCode() string {
	rand.Seed(time.Now().UnixNano())

	var s string
	for i := 0; i < 4; i++ {
		s += fmt.Sprintf("%d", rand.Intn(5))
	}
	return s
}

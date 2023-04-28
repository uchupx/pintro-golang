package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const insertUserQuery = `INSERT INTO user(name, username, password) VALUES(?,?,?)`

const deleteUserQuery = `DELETE FROM user WHERE id= ?`

const updateUserQuery = `UPDATE user SET name=?, username=?, password=? WHERE id=?`

const findUserByUsernameQuery = `
	SELECT
		*
	FROM
		user
		WHERE
		 username= ?
`

type userMySQLRepo struct {
	db *sqlx.DB
}

func (m userMySQLRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	stmt, err := m.db.PrepareContext(ctx, findUserByUsernameQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, username)

	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m userMySQLRepo) Insert(ctx context.Context, user model.User) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, insertUserQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, user.Name, user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &rowsAffected, nil
}

func (m userMySQLRepo) Update(ctx context.Context, user model.User) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, updateUserQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, user.Name, user.Username, user.Password, user.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &rowsAffected, nil
}

func (m userMySQLRepo) Delete(ctx context.Context, user model.User) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, deleteUserQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &rowsAffected, nil
}

func NewUserMySQlRepo(db *sqlx.DB) userMySQLRepo {
	return userMySQLRepo{
		db: db,
	}
}

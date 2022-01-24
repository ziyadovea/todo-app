package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user *models.User) (int, error) {
	err := ap.db.QueryRow(
		`insert into users (Name, Username, Password_hash) values ($1, $2, $3) returning ID`,
		user.Name,
		user.Username,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (*models.User, error) {
	user := &models.User{}

	err := ap.db.QueryRowx(
		`select * from users where Username=$1`,
		username,
	).StructScan(user)

	if err != nil {
		return nil, ErrorIncorrectUsernameOrPassword
	}

	if !isPasswordCorrect(user.Password, password) {
		return nil, ErrorIncorrectUsernameOrPassword
	}

	return user, nil
}

func isPasswordCorrect(hashPassword string, password string) bool {
	pw := []byte(password)
	hw := []byte(hashPassword)
	return bcrypt.CompareHashAndPassword(hw, pw) == nil
}

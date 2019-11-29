package users

import (
	"github.com/cheapeone/goland/api/models"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	DB *sqlx.DB
}

func (us *UserStore) FindByEmail(email string) (*models.User, error) {

}

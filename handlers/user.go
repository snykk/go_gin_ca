package handlers

import (
	"errors"
	"fmt"

	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/models"
)

func (m *module) RetrieveUser(id uint) (user models.User, err error) {
	if user, err = m.db.userRepository.GetUserByID(id); err != nil {
		return models.User{}, fmt.Errorf("cannot find user with id %d", id)
	}
	return
}

func (m *module) UpdateUser(id uint, user datatransfers.UserUpdate) (err error) {
	if err = m.db.userRepository.UpdateUser(models.User{
		ID:    id,
		Email: user.Email,
		Bio:   user.Bio,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}

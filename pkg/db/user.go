package db

import (
	"github.com/choisangh/board-crud-backend/pkg/model"
	"github.com/pkg/errors"
)

func (h *DBHandler) CreateUser(user *model.User) (*model.User, error) {
	result := h.gDB.Create(user)

	return user, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetUserList() ([]*model.User, error) {
	userList := []*model.User{}
	result := h.gDB.Find(&userList)

	return userList, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetUserByID(id uint64) (*model.User, error) {
	user := &model.User{}
	result := h.gDB.First(user, id)
	return user, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetUserByEmail(Email string) (*model.User, error) {
	user := &model.User{}
	if err := h.gDB.Where("email = ?", Email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (h *DBHandler) UpdateUser(id uint, newUser *model.User) (*model.User, error) {
	oldUser := &model.User{}
	result := h.gDB.Model(oldUser).Where(id).Updates(newUser)

	return oldUser, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) DeleteUserByID(id string) error {
	result := h.gDB.Delete(&model.User{}, id)

	return errors.Wrap(result.Error, "db handler error")
}

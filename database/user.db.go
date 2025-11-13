package database

import (
	"errors"

	"github.com/eakarach/AuthJWT/models"
	"gorm.io/gorm"
)

func GetUserByUsername(u string) (*models.User, error) {
	db := gdb
	var user models.User
	if err := db.Where(&models.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

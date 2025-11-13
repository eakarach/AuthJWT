package database

import (
	"errors"

	"github.com/eakarach/AuthJWT/models"
	"gorm.io/gorm"
)

func GetAllUser() (*[]models.User, error) {
	db := gdb
	var user []models.User

	//db.Select("name", "age").Find(&users)
	if err := db.Select("id, name, email, username").Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

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

func GetUserByUserId(id int) (*models.User, error) {
	db := gdb
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	db := gdb
	if err := db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func DeleteUserByUserId(id int) (error) {
	db := gdb
	var user models.User
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return  nil
		}
		return err
	}
	return nil
}

func UpdateUserByUserId(user *models.User) (*models.User, error) {
	db := gdb
	if err := db.Save(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}


package repositories

import (
	"errors"
	"mvc-be20/config"
	"mvc-be20/entities"
	"mvc-be20/models"
)

func InsertUser(newUser entities.UserCore) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Role:        newUser.Role,
	}
	// simpan ke DB
	tx := config.DB.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affexted =0")
	}
	return nil
}

func GetUserById(userId int) (entities.UserCore, error) {
	var userDataGorm models.User
	tx := config.DB.First(&userDataGorm, userId)
	if tx.Error != nil {
		return entities.UserCore{}, tx.Error
	}

	// Proses mapping dari struct GORM model ke struct core
	userDataCore := entities.UserCore{
		ID:          userDataGorm.ID,
		Name:        userDataGorm.Name,
		Email:       userDataGorm.Email,
		Password:    userDataGorm.Password,
		Address:     userDataGorm.Address,
		PhoneNumber: userDataGorm.PhoneNumber,
		Role:        userDataGorm.Role,
		CreatedAt:   userDataGorm.CreatedAt,
		UpdatedAt:   userDataGorm.UpdatedAt,
	}

	return userDataCore, nil
}

func SelectAllUsers() ([]entities.UserCore, error) {
	var usersDataGorm []models.User
	tx := config.DB.Find(&usersDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var usersDataCore []entities.UserCore
	for _, value := range usersDataGorm {
		var userCore = entities.UserCore{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Password:    value.Password,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersDataCore = append(usersDataCore, userCore)
	}

	return usersDataCore, nil
}

func UpdateUserById(id int, userUpdate models.User) error {
	tx := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(userUpdate)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

func DeleteUserById(id int) error {
	tx := config.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}

	return nil
}

func GetProductsByUserId(userId int) ([]entities.Product, error) {
	var productsDataGorm []models.Product
	tx := config.DB.Preload("User").Where("user_id = ?", userId).Find(&productsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Proses mapping dari struct GORM model ke struct core
	var productsData []entities.Product
	for _, value := range productsDataGorm {
		var product = entities.Product{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
		}
		productsData = append(productsData, product)
	}

	return productsData, nil
}

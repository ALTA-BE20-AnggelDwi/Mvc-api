package repositories

import (
	"errors"
	"mvc-be20/config"
	"mvc-be20/entities"
	"mvc-be20/models"
)

func InsertProduct(newProduct models.Product) error {
	// Proses mapping dari struct product core ke model gorm
	productInputGorm := models.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		UserID:      newProduct.UserID,
	}

	// Simpan ke DB
	tx := config.DB.Create(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, rows affected = 0")
	}

	return nil
}

func GetAllProducts() ([]entities.ProductCore, error) {
	var productsDataGorm []models.Product
	tx := config.DB.Preload("User").Find(&productsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Proses mapping dari struct GORM model ke struct core
	var productsDataCore []entities.ProductCore
	for _, value := range productsDataGorm {
		var productCore = entities.ProductCore{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
			// menambahkan data user
			User: entities.UserResponse{
				ID:    value.User.ID,
				Name:  value.User.Name,
				Email: value.User.Email,
			},
		}
		productsDataCore = append(productsDataCore, productCore)
	}

	return productsDataCore, nil
}

func GetProductById(idParam int) (entities.ProductCore, error) {
	var productData models.Product
	tx := config.DB.Preload("User").First(&productData, idParam)
	if tx.Error != nil {
		return entities.ProductCore{}, tx.Error
	}

	// Mapping
	var productCore = entities.ProductCore{
		ID:          productData.ID,
		Name:        productData.Name,
		UserID:      productData.UserID,
		Description: productData.Description,
		// menambahkan data user
		User: entities.UserResponse{
			ID:    productData.User.ID,
			Name:  productData.User.Name,
			Email: productData.User.Email,
		},
	}

	return productCore, nil
}

func UpdateProductById(idParam int, updatedProduct models.Product) error {
	tx := config.DB.Model(&models.Product{}).Where("id = ?", idParam).Updates(updatedProduct)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}

	return nil
}

func DeleteProductById(idParam int) error {
	tx := config.DB.Delete(&models.Product{}, idParam)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}

	return nil
}

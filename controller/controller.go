package controller

import (
	"api-marketplace/config"
	"api-marketplace/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePembeli(c echo.Context) error {
	p := new(model.Pembelis)
	db := config.DB()

	// Binding data
	if err := c.Bind(p); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Create(&p).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": p,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdatePembeli(c echo.Context) error {
	id := c.Param("id")
	p := new(model.Pembelis)
	db := config.DB()

	// Binding data
	if err := c.Bind(p); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_pembeli := new(model.Pembelis)

	if err := db.First(&existing_pembeli, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_pembeli.Nama_pembeli = p.Nama_pembeli
	existing_pembeli.No_telp = p.No_telp
	existing_pembeli.Alamat = p.Alamat
	if err := db.Save(&existing_pembeli).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_pembeli,
	}

	return c.JSON(http.StatusOK, response)
}

func GetPembeli(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var pembelis []*model.Pembelis

	if res := db.Find(&pembelis, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": pembelis[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeletePembeli(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	pembeli := new(model.Pembelis)

	err := db.Delete(&pembeli, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a book has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}

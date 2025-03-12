package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Upload a new image
// @Description Uploads a new image and returns its unique UUID
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file to upload"
// @Success 200 {object} map[string]string "{'id': '123e4567-e89b-12d3-a456-426614174000'}"
// @Failure 400 {object} myError "{'message': 'No image file provided'}"
// @Failure 500 {object} myError "{'message': 'Failed to upload image'}"
// @Router /api/v1/images [post]
func (h *Handler) addImage(c *gin.Context) {
	// Получаем файл из запроса
	fmt.Printf("Получаем файл из запроса\n")
	file, err := c.FormFile("image")
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Считываем содержимое файла в []byte
	fmt.Printf("Считываем содержимое файла в []byte\n")
	fileBytes, err := readFileAsBytes(file)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Генерируем UUID для изображения
	fmt.Printf("// Генерируем UUID для изображения\n")
	id, err := h.services.IImageService.AddImage(fileBytes)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("Возвращаем успешный ответ с ID изображения\n")
	// Возвращаем успешный ответ с ID изображения
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id})

}

func readFileAsBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Открываем временный файл
	fmt.Printf("Открываем временный файл\n")
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Считываем содержимое файла в []byte
	fmt.Printf("Считываем содержимое файла в []byte\n")
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

// @Summary Get image by UUID
// @Description Retrieves an image by its unique UUID and sends it as a downloadable file
// @Tags images
// @Accept json
// @Produce octet-stream
// @Param id path string true "Image UUID" format=uuid
// @Success 200 {file} byte "Image file"
// @Failure 400 {object} myError "{'message': 'Invalid image ID'}"
// @Failure 404 {object} myError "{'message': 'Image not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/images/{id} [get]
func (h *Handler) getImageByImageUUID(c *gin.Context) {
	// Получаем image_id из параметров маршрута
	imageIDStr := c.Param("id")
	imageUUID, err := uuid.Parse(imageIDStr)
	if err != nil || imageUUID == uuid.Nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Получаем изображение из базы данных
	imageBytes, err := h.services.IImageService.GetImageByImageUUID(imageUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Устанавливаем заголовки для автоматической загрузки файла
	c.Writer.Header().Set("Content-Type", "application/octet-stream")                                  // Тип данных - octet-stream
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, imageIDStr)) // Автоматическая загрузка
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(imageBytes)))                             // Размер файла

	// Отправляем бинарные данные изображения
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
}

// @Summary Delete image by UUID
// @Description Deletes an image from the database by its unique UUID
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "Image UUID" format=uuid
// @Success 204 "Image deleted successfully"
// @Failure 400 {object} myError "{'message': 'Invalid image ID'}"
// @Failure 404 {object} myError "{'message': 'Image not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/images/{id} [delete]
func (h *Handler) deleteImageByImageUUID(c *gin.Context) {
	imageUUID, err := uuid.Parse(c.Param("id"))
	if err != nil || imageUUID == uuid.Nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IImageService.DeleteImageByImageUUID(imageUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем статус 204 No Content
	c.Status(http.StatusNoContent)
}

// @Summary Update an existing image
// @Description Updates an image by its unique UUID with a new file
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Image UUID" format=uuid
// @Param image formData file true "New image file to replace the existing one"
// @Success 200 {object} map[string]string "{'message': 'Image updated successfully'}"
// @Failure 400 {object} myError "{'message': 'Invalid request parameters'}"
// @Failure 404 {object} myError "{'message': 'Image not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/images/{id} [put]
func (h *Handler) updateImage(c *gin.Context) {
	imageId, err := uuid.Parse(c.Param("id"))
	if err != nil || imageId == uuid.Nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Получаем новый файл из запроса
	file, err := c.FormFile("image")
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Считываем содержимое файла в []byte
	fileBytes, err := readFileAsBytes(file)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Обновляем изображение в базе данных
	err = h.services.IImageService.UpdateImage(imageId, fileBytes)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}

// @Summary Get image by product ID
// @Description Retrieves the image associated with a specific product and sends it as a downloadable file
// @Tags images
// @Accept json
// @Produce octet-stream
// @Param id path integer true "Product ID"
// @Success 200 {file} byte "Image file"
// @Failure 400 {object} myError "{'message': 'Product ID is required'}"
// @Failure 404 {object} myError "{'message': 'Image not found for the given product ID'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/images/product/{id} [get]
func (h *Handler) getImageByProductId(c *gin.Context) {

	log.Println("hui")
	productIDStr := c.Param("id")
	if productIDStr == "" {
		newErrorResponce(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	log.Println("pizda")
	productId, err := strconv.Atoi(productIDStr)

	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	imageBytes, err := h.services.IImageService.GetImageByProductId(productId)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Устанавливаем заголовки для автоматической загрузки файла
	c.Writer.Header().Set("Content-Type", "application/octet-stream")                                    // Тип данных - octet-stream
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, productIDStr)) // Автоматическая загрузка
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(imageBytes)))                               // Размер файла

	// Отправляем бинарные данные изображения
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)

}

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

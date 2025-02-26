package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ImagePostgresImpl struct {
	db *sqlx.DB
}

func NewImagePostgresImpl(db *sqlx.DB) *ImagePostgresImpl {
	return &ImagePostgresImpl{db: db}
}

func (r *ImagePostgresImpl) AddImage(fileBytes []byte) (uuid.UUID, error) {
	imageId := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id, image) VALUES ($1, $2) RETURNING id", imagesTable)

	_, err := r.db.Exec(query, imageId, fileBytes)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to add image: %w", err)
	}

	return imageId, nil
}

func (r *ImagePostgresImpl) GetImageByImageUUID(imageUUID uuid.UUID) ([]byte, error) {

	var imageBytes []byte
	query := fmt.Sprintf("SELECT image FROM %s WHERE id = $1", imagesTable)

	err := r.db.QueryRow(query, imageUUID).Scan(&imageBytes)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}

func (r *ImagePostgresImpl) DeleteImageByImageUUID(imageUUID uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", imagesTable)

	_, err := r.db.Exec(query, imageUUID)
	return err
}

func (r *ImagePostgresImpl) UpdateImage(imageId uuid.UUID, fileBytes []byte) error {
	query := fmt.Sprintf("UPDATE %s SET image=$1 WHERE id = $2", imagesTable)
	_, err := r.db.Exec(query, fileBytes, imageId)
	return err
}

func (r *ImagePostgresImpl) GetImageByProductId(productId int) ([]byte, error) {
	var imageBytes []byte
	query := fmt.Sprintf("SELECT image from %s WHERE id = (SELECT image_id FROM %s WHERE %s.id = $1)", imagesTable, productsTable, productsTable)
	err := r.db.QueryRow(query, productId).Scan(&imageBytes)

	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}

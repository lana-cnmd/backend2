package types

import (
	"github.com/google/uuid"
)

type SiteImage struct {
	ID    uuid.UUID `db:"id"`
	Image []byte    `db:"image"`
}

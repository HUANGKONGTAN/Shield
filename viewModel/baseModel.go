package viewModel

import (
	"time"
)

type BaseModel struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}



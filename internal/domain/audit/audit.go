package audit

import (
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"time"
)

// Audit represents the audit information of a resource
type Audit struct {
	createdAt time.Time
	createdBy id.ID
	updatedAt time.Time
	updatedBy id.ID
}

// NewDefault creates default audit information with the current time,
// the user performing the action and the resource is active
func NewDefault(userId id.ID) Audit {
	now := time.Now().UTC()
	return New(now, now, userId, userId)
}

// New creates audit information with the provided values
func New(createdAt, updatedAt time.Time, createdBy, updatedBy id.ID) Audit {
	return Audit{
		createdAt: createdAt.UTC(),
		createdBy: createdBy,
		updatedAt: updatedAt.UTC(),
		updatedBy: updatedBy,
	}
}

func (audit Audit) CreatedAt() time.Time {
	return audit.createdAt
}

func (audit Audit) CreatedBy() id.ID {
	return audit.createdBy
}

func (audit Audit) UpdatedAt() time.Time {
	return audit.updatedAt
}

func (audit Audit) UpdatedBy() id.ID {
	return audit.updatedBy
}

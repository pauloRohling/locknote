package audit

import (
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"time"
)

// Audit represents the audit information of a resource
type Audit struct {
	createdAt time.Time
	createdBy id.ID
}

// NewDefault creates default audit information with the current time,
// the user performing the action and the resource is active
func NewDefault(userId id.ID) Audit {
	now := time.Now().UTC()
	return New(now, userId)
}

// New creates audit information with the provided values
func New(createdAt time.Time, createdBy id.ID) Audit {
	return Audit{
		createdAt: createdAt,
		createdBy: createdBy,
	}
}

func (audit Audit) CreatedAt() time.Time {
	return audit.createdAt
}

func (audit Audit) CreatedBy() id.ID {
	return audit.createdBy
}

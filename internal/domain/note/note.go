package note

import (
	"encoding/json"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	"time"
)

// Note represents a note that can be created by a user
type Note struct {
	id      id.ID
	title   text.Title
	content string
	audit   audit.Audit
}

func (note *Note) ID() id.ID {
	return note.id
}

func (note *Note) Title() text.Title {
	return note.title
}

func (note *Note) Content() string {
	return note.content
}

func (note *Note) Audit() audit.Audit {
	return note.audit
}

func (note *Note) Update(title string, content string) error {
	newTitle, err := text.NewTitle(title)
	if err != nil {
		return err
	}

	note.title = newTitle
	note.content = content
	return nil
}

func (note *Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        id.ID      `json:"id"`
		Title     text.Title `json:"title"`
		Content   string     `json:"content"`
		CreatedAt time.Time  `json:"createdAt"`
		CreatedBy id.ID      `json:"createdBy"`
		UpdatedAt time.Time  `json:"updatedAt"`
		UpdatedBy id.ID      `json:"updatedBy"`
	}{
		ID:        note.id,
		Title:     note.title,
		Content:   note.content,
		CreatedAt: note.audit.CreatedAt(),
		CreatedBy: note.audit.CreatedBy(),
		UpdatedAt: note.audit.UpdatedAt(),
		UpdatedBy: note.audit.UpdatedBy(),
	})
}

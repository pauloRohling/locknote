package user

import (
	"encoding/json"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/password"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	"time"
)

type User struct {
	id       id.ID
	name     text.PersonName
	email    email.Email
	password password.Password
	audit    audit.Audit
}

func (user *User) ID() id.ID {
	return user.id
}

func (user *User) Name() text.PersonName {
	return user.name
}

func (user *User) Email() email.Email {
	return user.email
}

func (user *User) Password() password.Password {
	return user.password
}

func (user *User) Audit() audit.Audit {
	return user.audit
}

func (user *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        id.ID           `json:"id"`
		Name      text.PersonName `json:"name"`
		Email     email.Email     `json:"email"`
		CreatedAt time.Time       `json:"createdAt"`
		CreatedBy id.ID           `json:"createdBy"`
	}{
		ID:        user.id,
		Name:      user.name,
		Email:     user.email,
		CreatedAt: user.audit.CreatedAt(),
		CreatedBy: user.audit.CreatedBy(),
	})
}

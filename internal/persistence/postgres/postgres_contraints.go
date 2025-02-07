package postgres

var UniqueViolationErrors = map[string]string{
	"users_email_idx": "The provided email address is already in use",
}

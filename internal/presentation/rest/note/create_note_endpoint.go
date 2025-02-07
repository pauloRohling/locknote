package note

import (
	"github.com/labstack/echo/v4"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"net/http"
)

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateNoteResponse struct {
	*note.Note
}

func (controller *RestController) create(c echo.Context) error {
	body := new(CreateNoteRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	input := &noteApplication.CreateNoteInput{
		Title:   body.Title,
		Content: body.Content,
	}

	output, err := controller.service.Create(c.Request().Context(), input)
	if err != nil {
		return err
	}

	response := &CreateNoteResponse{Note: output.Note}
	return c.JSON(http.StatusCreated, response)
}

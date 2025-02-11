package note

import (
	"github.com/labstack/echo/v4"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"net/http"
)

type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteResponse struct {
	*note.Note
}

func (controller *RestController) updateById(c echo.Context) error {
	noteId := c.Param("id")

	body := new(UpdateNoteRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	input := &noteApplication.UpdateNoteInput{
		ID:      noteId,
		Title:   body.Title,
		Content: body.Content,
	}

	output, err := controller.service.UpdateById(c.Request().Context(), input)
	if err != nil {
		return err
	}

	response := &UpdateNoteResponse{Note: output.Note}
	return c.JSON(http.StatusOK, response)
}

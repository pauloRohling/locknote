package note

import (
	"github.com/labstack/echo/v4"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"net/http"
)

type GetNoteRequest struct {
	NoteID string `json:"id"`
}

type GetNoteResponse struct {
	*note.Note
}

func (controller *RestController) getById(c echo.Context) error {
	noteId := c.Param("id")

	input := &noteApplication.GetNoteInput{
		NoteID: noteId,
	}

	output, err := controller.service.GetById(c.Request().Context(), input)
	if err != nil {
		return err
	}

	response := &GetNoteResponse{Note: output.Note}
	return c.JSON(http.StatusOK, response)
}

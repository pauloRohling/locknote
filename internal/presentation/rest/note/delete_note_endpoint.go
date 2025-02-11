package note

import (
	"github.com/labstack/echo/v4"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"net/http"
)

type DeleteNoteRequest struct {
	NoteID string `json:"id"`
}

type DeleteNoteResponse struct {
	NoteID id.ID `json:"id"`
}

func (controller *RestController) deleteById(c echo.Context) error {
	noteId := c.Param("id")

	input := &noteApplication.DeleteNoteInput{
		NoteID: noteId,
	}

	output, err := controller.service.DeleteById(c.Request().Context(), input)
	if err != nil {
		return err
	}

	response := &DeleteNoteResponse{NoteID: output.NoteID}
	return c.JSON(http.StatusOK, response)
}

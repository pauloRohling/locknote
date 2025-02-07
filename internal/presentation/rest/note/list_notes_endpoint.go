package note

import (
	"github.com/labstack/echo/v4"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/presentation/rest/pagination"
	"github.com/pauloRohling/locknote/pkg/array"
	"net/http"
)

type ListNotesRequest struct {
	NoteID string `json:"noteId"`
}

type ListNotesResponse struct {
	*note.Note
}

func (controller *RestController) list(c echo.Context) error {
	ctx := c.Request().Context()
	paginationParams, err := pagination.GetPagination(ctx)
	if err != nil {
		return err
	}

	input := &noteApplication.ListNotesInput{Pagination: paginationParams}
	output, err := controller.service.List(ctx, input)
	if err != nil {
		return err
	}

	response := array.Map(output.Notes, func(note *note.Note) *ListNotesResponse {
		return &ListNotesResponse{Note: note}
	})

	return c.JSON(http.StatusOK, response)
}

func (controller *RestController) getPaginationMiddleware() echo.MiddlewareFunc {
	return pagination.Middleware([]string{"created_at"}, "id")
}

package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) SegmentStats(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	c := ctx.Request().Context()

	seg, err := h.Segments.GetByID(c, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if seg == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	err = h.Segments.Metadata(c, seg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, seg)
}

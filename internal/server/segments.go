package server

import (
	"net/http"
	"strconv"

	"github.com/Angry-Lab/api/pkg/segment"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListSegments(ctx echo.Context) error {
	c := ctx.Request().Context()

	segments, err := h.Segments.List(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, segments)
}

func (h *Handler) GetSegment(ctx echo.Context) error {
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

	return ctx.JSON(http.StatusOK, seg)
}

func (h *Handler) CreateSegment(ctx echo.Context) error {
	seg := &segment.Segment{}

	if err := ctx.Bind(seg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	c := ctx.Request().Context()

	err := h.Segments.Create(c, seg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, seg)
}

func (h *Handler) UpdateSegment(ctx echo.Context) error {
	seg := &segment.Segment{}

	if err := ctx.Bind(seg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	c := ctx.Request().Context()

	err := h.Segments.PutIfExits(c, seg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, seg)
}

package server

import (
	"encoding/json"
	"fmt"
	"github.com/Angry-Lab/api/pkg/helpers"
	"github.com/Angry-Lab/api/pkg/segment"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) GetSegment(ctx echo.Context) error {
	if id := ctx.Request().URL.Query().Get("id"); id != "" {
		c := ctx.Request().Context()
		var (
			wg      = &sync.WaitGroup{}
			err     error
			segment *segment.Segment
		)

		wg.Add(1)

		go func() {
			defer wg.Done()
			segment, err = h.Segments.Repository().GetByID(c, id)
		}()

		wg.Wait()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}

		return ctx.JSON(http.StatusOK, &segmentResponse{
			Segment: segment,
		})
	} else {
		return errors.New("id field required")
	}
}

func (h *Handler) GetSegmentList(ctx echo.Context) error {
	c := ctx.Request().Context()
	var (
		wg       = &sync.WaitGroup{}
		segments []*segment.Segment
	)

	wg.Add(1)

	go func() {
		defer wg.Done()
		segments, _ = h.Segments.Repository().GetAll(c)
	}()

	wg.Wait()

	return ctx.JSON(http.StatusOK, &segmentsResponse{
		Segments: segments,
	})
}

func (h *Handler) PutSegment(ctx echo.Context) error {
	s := segment.Segment{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&s)
	if err != nil {
		return errors.Wrap(err, "invalid segment struct")
	}
	return nil
}

func (h *Handler) CreateSegment(ctx echo.Context) error {
	s := segment.Segment{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&s)
	if err != nil {
		return errors.Wrap(err, "invalid segment struct")
	}
	return nil
}

func (h *Handler) DeleteSegment(ctx echo.Context) error {
	if lim := ctx.Request().URL.Query().Get("id"); lim != "" {
		return nil
	} else {
		return errors.New("id field required")
	}
}

type segmentResponse struct {
	Segment *segment.Segment `json:"segment"`
}

type segmentsResponse struct {
	Segments []*segment.Segment `json:"segments"`
}

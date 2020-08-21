package server

import (
	"github.com/Angry-Lab/api/pkg/helpers"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) ListUsers(ctx echo.Context) error {
	limit, offset, err := getNav(ctx.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	c := ctx.Request().Context()
	filter := &user.Filter{
		Limit:  limit,
		Offset: offset,
	}

	var (
		wg    = &sync.WaitGroup{}
		mu    = &sync.Mutex{}
		errs  []error
		users []*user.User
		count int64
	)

	wg.Add(2)

	go func() {
		defer wg.Done()

		users, err = h.Users.Repository().Search(c, filter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		count, err = h.Users.Repository().Count(c, filter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
		}
	}()

	wg.Wait()

	if err := helpers.JoinErrs(errs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, &userList{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Users:  users,
	})
}

func (h *Handler) CreateUser(ctx echo.Context) error {
	return nil
}

func (h *Handler) UpdateUser(ctx echo.Context) error {
	return nil
}

func (h *Handler) DeleteUser(ctx echo.Context) error {
	return nil
}

func getNav(r *http.Request) (limit, offset int, err error) {
	if lim := r.URL.Query().Get("limit"); lim != "" {
		limit, err = strconv.Atoi(lim)
		if err != nil {
			return limit, offset, errors.Wrap(err, "bad limit")
		}
	}

	if off := r.URL.Query().Get("limit"); off != "" {
		offset, err = strconv.Atoi(off)
		if err != nil {
			return limit, offset, errors.Wrap(err, "bad offset")
		}
	}

	return
}

type userList struct {
	Count  int64        `json:"count"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
	Users  []*user.User `json:"users"`
}

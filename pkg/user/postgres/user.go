package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/Angry-Lab/api/db/postgres/boiler"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/partyzanex/layer"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type users struct {
	ex layer.BoilExecutor
}

func Users(ex layer.BoilExecutor) user.Repository {
	return &users{ex: ex}
}

func (repo *users) Search(ctx context.Context, filter *user.Filter) ([]*user.User, error) {
	mods := repo.applyFilter(filter, []qm.QueryMod{
		qm.OrderBy("id"),
	})

	c, ex := layer.GetExecutor(ctx, repo.ex)

	models, err := boiler.Users(mods...).All(c, ex)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "search users failed")
	}

	users := make([]*user.User, len(models))

	for i, model := range models {
		users[i] = modelToUser(model)
	}

	return users, nil
}

func (repo *users) Count(ctx context.Context, filter *user.Filter) (int64, error) {
	var mods []qm.QueryMod

	if filter != nil {
		f := *filter
		f.Limit = 0
		mods = repo.applyFilter(&f, mods)
	}

	c, ex := layer.GetExecutor(ctx, repo.ex)

	count, err := boiler.Users(mods...).Count(c, ex)
	if err != nil {
		return 0, errors.Wrap(err, "getting count of users failed")
	}

	return count, nil
}

func (*users) applyFilter(filter *user.Filter, mods []qm.QueryMod) []qm.QueryMod {
	if mods == nil {
		mods = []qm.QueryMod{}
	}

	if filter != nil {
		if n := len(filter.IDs); n > 0 {
			ids := make([]interface{}, n)
			for i, id := range filter.IDs {
				ids[i] = id
			}

			mods = append(mods, qm.WhereIn("id in ?", ids...))
		}

		if filter.Name != "" {
			clause := "%" + filter.Name + "%"
			mods = append(mods, qm.Where("name like ?", clause))
		}

		if filter.Status != "" {
			mods = append(mods, qm.Where("status = ?", filter.Status))
		}

		if filter.Login != "" {
			mods = append(mods, qm.Where("login = ?", filter.Login))
		}

		if filter.Limit > 0 {
			mods = append(mods, qm.Limit(filter.Limit))

			if filter.Offset >= 0 {
				mods = append(mods, qm.Offset(filter.Offset))
			}
		}
	}

	return mods
}

func (repo *users) Create(ctx context.Context, entity user.User) (result *user.User, err error) {
	c, tr := layer.GetTransactor(ctx)
	if tr == nil {
		tr, err = repo.ex.BeginTx(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, layer.ErrCreateTransaction.Error())
		}

		defer func() {
			errTr := layer.ExecuteTransaction(tr, err)
			if errTr != nil {
				err = errors.Wrap(errTr, "transaction error")
			}
		}()
	}

	model := userToModel(&entity)
	model.DTCreated = time.Now()

	err = model.Insert(c, tr, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "inserting user failed")
	}

	return modelToUser(model), nil
}

func (repo *users) Update(ctx context.Context, entity user.User) (result *user.User, err error) {
	c, tr := layer.GetTransactor(ctx)
	if tr == nil {
		tr, err = repo.ex.BeginTx(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, layer.ErrCreateTransaction.Error())
		}

		defer func() {
			errTr := layer.ExecuteTransaction(tr, err)
			if errTr != nil {
				err = errors.Wrap(errTr, "transaction error")
			}
		}()
	}

	model := userToModel(&entity)
	model.DTUpdated = time.Now()

	_, err = model.Update(c, tr, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "updating user failed")
	}

	return modelToUser(model), err
}

func (repo *users) Delete(ctx context.Context, entity user.User) (err error) {
	if entity.ID == 0 {
		return user.ErrRequiredUserID
	}

	c, tr := layer.GetTransactor(ctx)
	if tr == nil {
		tr, err = repo.ex.BeginTx(ctx, nil)
		if err != nil {
			return errors.Wrap(err, layer.ErrCreateTransaction.Error())
		}

		defer func() {
			errTr := layer.ExecuteTransaction(tr, err)
			if errTr != nil {
				err = errors.Wrap(errTr, "transaction error")
			}
		}()
	}

	model, err := boiler.Users(qm.Where("id = ?", entity.ID)).One(c, tr)
	if err == sql.ErrNoRows {
		return user.ErrUserNotFound
	}

	if err != nil {
		return errors.Wrap(err, "search user failed")
	}

	_, err = model.Delete(c, tr)
	if err != nil {
		return errors.Wrap(err, "deleting user failed")
	}

	return
}

func modelToUser(model *boiler.User) *user.User {
	entity := &user.User{
		ID:                model.ID,
		Login:             model.Login,
		Password:          model.Password,
		Status:            user.Status(model.Status),
		Name:              model.Name,
		Role:              user.Role(model.Role),
		DTCreated:         model.DTCreated,
		DTUpdated:         model.DTUpdated,
		DTLastLogged:      model.DTLastLogged,
		PasswordIsEncoded: true,
		Current:           false,
	}

	return entity
}

func userToModel(entity *user.User) *boiler.User {
	model := &boiler.User{
		ID:           entity.ID,
		Login:        entity.Login,
		Password:     entity.Password,
		Status:       string(entity.Status),
		Name:         entity.Name,
		Role:         string(entity.Role),
		DTCreated:    entity.DTCreated,
		DTUpdated:    entity.DTUpdated,
		DTLastLogged: entity.DTLastLogged,
	}

	return model
}

package parcel

import (
	"context"
	"github.com/partyzanex/layer"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct {
	ex layer.BoilExecutor
}

func NewRepo(ex layer.BoilExecutor) *Repository {
	return &Repository{ex: ex}
}

func (repo *Repository) Ensure(ctx context.Context, parcels ...*Parcel) (err error) {
	c, tx := layer.GetTransactor(ctx)
	if tx == nil {
		tx, err = repo.ex.BeginTx(ctx, nil)
		if err != nil {
			return errors.Wrap(err, layer.ErrCreateTransaction.Error())
		}

		defer func() {
			errTr := layer.ExecuteTransaction(tx, err)
			if errTr != nil {
				err = errors.Wrap(errTr, "transaction error")
			}
		}()
	}

	for _, parcel := range parcels {
		err := parcel.Upsert(c, tx, true, []string{"uid"}, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}

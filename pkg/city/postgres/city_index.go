package postgres

import (
	"context"
	"database/sql"

	"github.com/Angry-Lab/api/db/postgres/boiler"
	"github.com/Angry-Lab/api/pkg/city"
	"github.com/partyzanex/layer"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type cityIndexes struct {
	ex layer.BoilExecutor
}

func Cities(ex layer.BoilExecutor) city.Repository {
	return &cityIndexes{ex: ex}
}

func (repo *cityIndexes) Cities(ctx context.Context, postcodes ...int) (map[int]*city.City, error) {
	var (
		mods  []qm.QueryMod
		n     = len(postcodes)
		c, ex = layer.GetExecutor(ctx, repo.ex)
	)

	codes := make([]interface{}, n)

	for i, postcode := range postcodes {
		codes[i] = postcode
	}

	mods = append(mods, qm.WhereIn("postcode IN ?", codes...))

	models, err := boiler.CityIndices(mods...).All(c, ex)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "search cities failed")
	}

	cities := make(map[int]*city.City, len(models))

	for _, model := range models {
		lon, _ := model.Longitude.Float64()
		lat, _ := model.Longitude.Float64()

		cities[model.Postcode] = &city.City{
			ID:        model.ID,
			Postcode:  model.Postcode,
			Name:      model.Name,
			Latitude:  lat,
			Longitude: lon,
		}
	}

	return cities, nil
}

package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Angry-Lab/api/db/postgres/boiler"
	"github.com/Angry-Lab/api/pkg/segment"
	"github.com/partyzanex/layer"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type segments struct {
	ex layer.BoilExecutor
}

func Segments(ex layer.BoilExecutor) segment.Repository {
	return &segments{ex: ex}
}

func (repo *segments) List(ctx context.Context) ([]*segment.Segment, error) {
	c, ex := layer.GetExecutor(ctx, repo.ex)

	models, err := boiler.Segments(qm.OrderBy("id")).All(c, ex)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "getting segments failed")
	}

	segments := make([]*segment.Segment, len(models))

	for i, model := range models {
		segments[i] = modelToSegment(model, nil)
	}

	return segments, nil
}

func (repo *segments) GetByID(ctx context.Context, id int) (*segment.Segment, error) {
	c, ex := layer.GetExecutor(ctx, repo.ex)

	model, err := boiler.Segments(qm.Where("id = ?", id)).One(c, ex)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "getting segment failed")
	}

	return modelToSegment(model, nil), nil
}

func (repo *segments) PutIfExits(ctx context.Context, seg *segment.Segment) error {
	model := segmentToModel(seg)
	c, ex := layer.GetExecutor(ctx, repo.ex)

	_, err := model.Update(c, ex, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "updating segment failed")
	}

	modelToSegment(model, seg)

	return nil
}

func (repo *segments) Create(ctx context.Context, seg *segment.Segment) error {
	model := segmentToModel(seg)
	c, ex := layer.GetExecutor(ctx, repo.ex)

	err := model.Insert(c, ex, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "creating segment failed")
	}

	modelToSegment(model, seg)

	return nil
}

type metadata struct {
	TotalCost float64 `boil:"total_cost"`
	TotalNP   float64 `boil:"total_np"`
	Distance  float64 `boil:"distance"`
	Count     int     `boil:"cnt"`
	Users     int     `boil:"users"`
}

func (repo *segments) Metadata(ctx context.Context, seg *segment.Segment) error {
	q := `
SELECT
       sum(parcel.cost) as total_cost,
       sum(parcel.amount_np) as total_np,
       round(avg(parcel.distance)/1000, 2) as distance,
       (SELECT count(*) FROM parcel WHERE %s) as cnt,
       (SELECT count(*) FROM (SELECT count(*) cnt FROM parcel WHERE %s GROUP BY hid) as st) as users
FROM parcel
WHERE %s;
`
	q = fmt.Sprintf(q, seg.Condition, seg.Condition, seg.Condition)
	tg := &metadata{}
	c, ex := layer.GetExecutor(ctx, repo.ex)

	err := queries.Raw(q).Bind(c, ex, tg)
	if err != nil {
		return nil
	}

	seg.Metadata.TotalNP = tg.TotalNP
	seg.Metadata.TotalCost = tg.TotalCost
	seg.Metadata.Distance = tg.Distance
	seg.Metadata.Count = tg.Count
	seg.Metadata.Users = tg.Users

	return nil
}

func modelToSegment(model *boiler.Segment, seg *segment.Segment) *segment.Segment {
	if seg == nil {
		seg = &segment.Segment{}
	}

	seg.ID = model.ID
	seg.Name = model.Name
	seg.Description = model.Description
	seg.Condition = model.Condition

	return seg
}

func segmentToModel(seg *segment.Segment) *boiler.Segment {
	model := &boiler.Segment{}

	model.ID = seg.ID
	model.Name = seg.Name
	model.Description = seg.Description
	model.Condition = seg.Condition

	return model
}

package server

import (
	"encoding/csv"
	"github.com/Angry-Lab/api/pkg/city"
	"github.com/Angry-Lab/api/pkg/helpers"
	"github.com/Angry-Lab/api/pkg/parcel"
	"github.com/labstack/echo/v4"
	"io"
	"math"
	"net/http"
)

func (h *Handler) UploadParcelsCSV(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	defer helpers.Close(src)

	c := ctx.Request().Context()
	r := csv.NewReader(src)
	i := 0

	cityMap := make(map[int]*city.City)
	var prc []*parcel.Parcel

	for {
		if i%1000 == 0 {
			err = h.Parcels.Ensure(c, prc...)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
			}

			prc = []*parcel.Parcel{}
		}

		row, err := r.Read()
		if err == io.EOF {
			err = h.Parcels.Ensure(c, prc...)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
			}

			break
		}

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}

		if i == 0 {
			i++
			continue
		}

		p, err := parcel.FromRow(row)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}

	CityLoop:
		city1, ok1 := cityMap[p.SenderIndex]
		city2, ok2 := cityMap[p.RecipientIndex]

		if !ok1 || !ok2 {
			cities, err := h.Cities.Cities(c, p.SenderIndex, p.RecipientIndex)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
			}

			j := 0
			for idx, ct := range cities {
				cityMap[idx] = ct
				j++
			}

			if j == 2 {
				goto CityLoop
			}
		}

		if city1 != nil && city2 != nil {
			//log.Println(p.SenderIndex, p.RecipientIndex)
			dist := Distance(city1.Latitude, city1.Longitude, city2.Latitude, city2.Longitude)
			p.Distance.SetFloat64(helpers.Round(dist, 2))
		} else {
			p.Distance.SetFloat64(-1)
		}

		prc = append(prc, p)
		//if i == 2000 {
		//	break
		//}

		i++
	}

	//b, err := ioutil.ReadAll(src)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	//}

	return nil
}

const EarthRadiusM = 6371010.0

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	return EarthRadiusM * 2 * math.Asin(math.Sqrt(
		math.Pow(math.Sin((lat1-math.Abs(lat2))*math.Pi/180/2), 2)+
			math.Cos(lat1*math.Pi/180)*
				math.Cos(math.Abs(lat2)*math.Pi/180)*
				math.Pow(math.Sin((lon1-lon2)*math.Pi/180/2), 2),
	))
}

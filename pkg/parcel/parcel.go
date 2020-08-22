package parcel

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Angry-Lab/api/db/postgres/boiler"
	"github.com/Angry-Lab/api/pkg/helpers"
	"github.com/ericlagergren/decimal"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/types"
	"strconv"
	"strings"
	"time"
)

type Parcel struct {
	*boiler.Parcel
}

func (p *Parcel) GetUID() string {
	if p.UID != "" {
		return p.UID
	}

	weight, _ := p.Weight.Float64()
	cost, _ := p.Cost.Float64()
	amountOTS, _ := p.AmountOts.Float64()
	amountNP, _ := p.AmountNP.Float64()

	uniq := []string{
		p.Hid, p.DT.String(),
		strconv.Itoa(p.SenderIndex),
		strconv.Itoa(p.RecipientIndex),
		strconv.FormatFloat(weight, 'f', -1, 64),
		strconv.FormatFloat(cost, 'f', -1, 64),
		strconv.FormatFloat(amountOTS, 'f', -1, 64),
		strconv.FormatFloat(amountNP, 'f', -1, 64),
		strconv.FormatBool(p.BlankDispatch), strconv.FormatBool(p.ParcelPost),
		strconv.FormatBool(p.Accelerated), strconv.FormatBool(p.International),
		strconv.FormatBool(p.WithAdvertValue), strconv.FormatBool(p.WithImpositionPayment),
		strconv.FormatBool(p.WithListOfAttachments), strconv.FormatBool(p.CautionMark),
		strconv.FormatBool(p.SMSForSender), strconv.FormatBool(p.SMSForRecipient),
	}

	hash := md5.Sum([]byte(strings.Join(uniq, ".")))

	p.UID = hex.EncodeToString(hash[:])

	return p.UID
}

func FromRow(row []string) (*Parcel, error) {
	dt, err := time.Parse("2006-01-02 15:04:05.999999", row[2])
	if err != nil {
		return nil, errors.Wrap(err, "parse date failed")
	}

	senderIdx, err := strconv.Atoi(row[3])
	if err != nil {
		return nil, errors.Wrap(err, "parse sender_index failed")
	}

	recipientIdx, err := strconv.Atoi(row[4])
	if err != nil {
		return nil, errors.Wrap(err, "parse recipient_index failed")
	}

	p := Parcel{
		Parcel: &boiler.Parcel{
			Hid:                   row[1],
			DT:                    dt,
			SenderIndex:           senderIdx,
			RecipientIndex:        recipientIdx,
			Weight:                types.NewDecimal(&decimal.Big{}),
			Cost:                  types.NewDecimal(&decimal.Big{}),
			AmountOts:             types.NewDecimal(&decimal.Big{}),
			AmountNP:              types.NewDecimal(&decimal.Big{}),
			BlankDispatch:         false,
			ParcelPost:            false,
			Accelerated:           false,
			International:         false,
			WithAdvertValue:       false,
			WithImpositionPayment: false,
			WithListOfAttachments: false,
			CautionMark:           false,
			SMSForSender:          false,
			SMSForRecipient:       false,
			Distance:              types.NewDecimal(&decimal.Big{}),
			LowestCost:            false,
		},
	}

	weight, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse weight failed")
	}

	cost, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse cost failed")
	}

	amountOTS, err := strconv.ParseFloat(row[7], 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse amount_ots failed")
	}

	amountNP, err := strconv.ParseFloat(row[8], 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse amount_np failed")
	}

	p.Weight.SetFloat64(weight)
	p.Cost.SetFloat64(helpers.Round(cost, 2))
	p.AmountNP.SetFloat64(helpers.Round(amountNP, 2))
	p.AmountOts.SetFloat64(helpers.Round(amountOTS, 2))

	p.BlankDispatch, _ = strconv.ParseBool(row[9])
	p.ParcelPost, _ = strconv.ParseBool(row[10])
	p.Accelerated, _ = strconv.ParseBool(row[11])
	p.International, _ = strconv.ParseBool(row[12])
	p.WithAdvertValue, _ = strconv.ParseBool(row[13])
	p.WithImpositionPayment, _ = strconv.ParseBool(row[14])
	p.WithListOfAttachments, _ = strconv.ParseBool(row[15])
	p.CautionMark, _ = strconv.ParseBool(row[16])
	p.SMSForSender, _ = strconv.ParseBool(row[17])
	p.SMSForRecipient, _ = strconv.ParseBool(row[18])
	p.GetUID()

	return &p, nil
}

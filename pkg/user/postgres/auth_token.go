package postgres

import (
	"context"
	"database/sql"

	"github.com/Angry-Lab/api/db/postgres/boiler"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/partyzanex/layer"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tokens struct {
	ex layer.BoilExecutor
}

func (repo *tokens) Search(ctx context.Context, token string) (*user.Token, error) {
	c, ex := layer.GetExecutor(ctx, repo.ex)

	model, err := boiler.AuthTokens(qm.Where("token = ?", token)).One(c, ex)
	if err == sql.ErrNoRows {
		return nil, user.ErrTokenNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "search token failed")
	}

	return modelToToken(model), nil
}

func (repo *tokens) Create(ctx context.Context, token user.Token) (result *user.Token, err error) {
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

	model := tokenToModel(&token)

	err = model.Insert(c, tr, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "inserting token failed")
	}

	return modelToToken(model), nil
}

func tokenToModel(token *user.Token) *boiler.AuthToken {
	model := &boiler.AuthToken{
		UserID:    token.UserID,
		Token:     token.Token,
		Type:      string(token.Type),
		DTExpired: token.DTExpired,
		DTCreated: token.DTCreated,
	}

	return model
}

func modelToToken(model *boiler.AuthToken) *user.Token {
	token := &user.Token{
		UserID:    model.UserID,
		Token:     model.Token,
		Type:      user.TokenType(model.Type),
		DTExpired: model.DTExpired,
		DTCreated: model.DTCreated,
		User: &user.User{
			ID: model.UserID,
		},
	}

	return token
}

func Tokens(ex layer.BoilExecutor) user.TokenRepository {
	return &tokens{ex: ex}
}

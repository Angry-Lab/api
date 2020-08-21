package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/partyzanex/testutils"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userCase struct {
	users  Repository
	tokens TokenRepository
}

func (uc *userCase) Validate(entity *User, create bool) error {
	if !create && entity.ID == 0 {
		return ErrRequiredUserID
	}

	if entity.Name == "" {
		return ErrRequiredUserName
	}

	if entity.Login == "" {
		return ErrRequiredUserLogin
	}

	if !govalidator.IsEmail(entity.Login) {
		return ErrInvalidUserLogin
	}

	if !entity.PasswordIsEncoded && entity.Password == "" {
		return ErrRequiredUserPassword
	}

	if !entity.Status.IsValid() {
		return ErrInvalidUserStatus
	}

	if !entity.Role.IsValid() {
		return ErrInvalidUserRole
	}

	return nil
}

func (uc *userCase) SearchByLogin(ctx context.Context, login string) (*User, error) {
	users, err := uc.users.Search(ctx, &Filter{
		Limit:  1,
		Offset: 0,
		Login:  login,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users[0], nil
}

func (uc *userCase) SearchByID(ctx context.Context, id int64) (*User, error) {
	users, err := uc.users.Search(ctx, &Filter{
		IDs:    []int64{id},
		Limit:  1,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users[0], nil
}

func (uc *userCase) SetLastLogged(ctx context.Context, entity *User) error {
	entity.DTLastLogged = time.Now()

	_, err := uc.users.Update(ctx, *entity)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userCase) Register(ctx context.Context, entity *User) error {
	if err := uc.Validate(entity, true); err != nil {
		return err
	}

	err := uc.EncodePassword(entity)
	if err != nil {
		return err
	}

	u, err := uc.users.Create(ctx, *entity)
	if err != nil {
		return err
	}

	*entity = *u

	return nil
}

func (uc *userCase) EncodePassword(entity *User) error {
	if entity.PasswordIsEncoded {
		return nil
	}

	p, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.MinCost)
	if err != nil {
		return errors.Wrap(err, "encoding password failed")
	}

	entity.PasswordIsEncoded = true
	entity.Password = string(p)

	return nil
}

func (uc *userCase) ComparePassword(entity *User, password string) (bool, error) {
	err := uc.EncodePassword(entity)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(password))
	if err != nil {
		err = ErrWrongPassword
	}

	return err == nil, err
}

func (uc *userCase) CreateAuthToken(ctx context.Context, entity *User) (*Token, error) {
	uniq := []string{
		strconv.FormatInt(entity.ID, 10),
		strconv.FormatInt(time.Now().Unix(), 10),
		entity.Login, testutils.RandomString(32),
	}

	t := sha256.Sum256([]byte(strings.Join(uniq, "_")))
	token := &Token{
		User:      entity,
		UserID:    entity.ID,
		Type:      AuthToken,
		Token:     hex.EncodeToString(t[:]),
		DTExpired: time.Now().Add(24 * time.Hour),
	}

	token, err := uc.tokens.Create(ctx, *token)
	if err != nil {
		return nil, errors.Wrap(err, "creating token failed")
	}

	return token, nil
}

func (uc *userCase) SearchToken(ctx context.Context, token string) (*Token, error) {
	authToken, err := uc.tokens.Search(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "search token failed")
	}

	authToken.User, err = uc.SearchByID(ctx, authToken.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "search user failed")
	}

	if authToken.DTExpired.Before(time.Now()) {
		return authToken, ErrTokenExpired
	}

	return authToken, nil
}

func (uc *userCase) Repository() Repository {
	return uc.users
}

func NewUseCase(users Repository, tokens TokenRepository) UseCase {
	return &userCase{
		users:  users,
		tokens: tokens,
	}
}

package usecase

import (
	"context"
	"errors"
	"log"
	"net/mail"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/taniko/rin/internal/domain/model"
	"github.com/taniko/rin/internal/domain/repository"
	"github.com/taniko/rin/internal/dto"
	"github.com/taniko/sumire"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinimumPasswordLength    = 16
	MinimumNameLength        = 3
	MinimumDisplayNameLength = 1
)

var (
	ErrNotMatchPassword   = errors.New("not match password")
	ErrPasswordShorter    = errors.New("password is shorter")
	ErrNameShorter        = errors.New("name is shorter")
	ErrInvalidName        = errors.New("name contains invalid character")
	ErrDisplayNameShorter = errors.New("display name is shorter")
	ErrEmailFormat        = errors.New("email is invalid")
	ErrUsedName           = errors.New("name is already used")
	ErrUsedEmail          = errors.New("email is already used")
)

type Account interface {
	Create(ctx context.Context, req dto.CreateAccount) error
}

type accountUsecase struct {
	logger     *sumire.Sumire
	repository repository.Account
}

func NewAccountUsecase(logger *sumire.Sumire, repository repository.Account) Account {
	return accountUsecase{
		logger:     logger,
		repository: repository,
	}
}

func (a accountUsecase) Create(ctx context.Context, req dto.CreateAccount) error {
	if err := a.validate(req); err != nil {
		return err
	}
	if m, err := a.repository.FindByNameOrEmail(ctx, req.Name, req.Email); err != nil {
		return err
	} else if m != nil {
		log.Println(time.Now().Local().Format(time.RFC3339))
		log.Println(*m)
		if m.Name == req.Name {
			return ErrUsedName
		} else if m.Email == req.Email {
			return ErrUsedEmail
		}
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}
	_, err = a.repository.Create(ctx, model.Account{
		ID:          id.String(),
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Email:       req.Email,
		Password:    string(password),
	})
	if err != nil {
		return err
	}
	return nil
}

func (a accountUsecase) validate(req dto.CreateAccount) error {
	if len(req.Name) < MinimumNameLength {
		return ErrNameShorter
	}
	r, err := regexp.Compile(`^[a-z][a-z\d_]+$`)
	if err != nil {
		return err
	} else if !r.MatchString(req.Name) {
		return ErrInvalidName
	}
	if len(req.DisplayName) < MinimumDisplayNameLength {
		return ErrDisplayNameShorter
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return ErrEmailFormat
	}
	if len(req.Password) < MinimumPasswordLength {
		return ErrPasswordShorter
	}
	if req.Password != req.PasswordConfirmation {
		return ErrNotMatchPassword
	}
	return nil
}

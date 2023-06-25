package usecase

import (
	"context"
	"errors"
	"net/mail"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/taniko/nullable"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/repository"
	"github.com/taniko/rin/internal/dto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
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
	Create(ctx context.Context, req dto.CreateAccount) (nullable.Nullable[account.Token], error)
	Login(ctx context.Context, email account.Email, password account.RawPassword) (nullable.Nullable[account.Token], error)
	VerifyToken(ctx context.Context, token account.Token) (string, error)
}

type accountUsecase struct {
	repo       repository.Account
	signingKey []byte
}

func NewAccountUsecase(repository repository.Account, signingKey []byte) Account {
	return accountUsecase{
		repo:       repository,
		signingKey: signingKey,
	}
}

func (a accountUsecase) Create(ctx context.Context, req dto.CreateAccount) (nullable.Nullable[account.Token], error) {
	if err := a.validate(req); err != nil {
		return nullable.Nullable[account.Token]{}, err
	}
	if v, err := a.repo.FindByNameOrEmail(ctx, req.Name, req.Email); err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return nullable.Nullable[account.Token]{}, err
		}
	} else {
		if v.Name == req.Name {
			return nullable.Nullable[account.Token]{}, ErrUsedName
		} else if v.Email == req.Email {
			return nullable.Nullable[account.Token]{}, ErrUsedEmail
		}
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nullable.Nullable[account.Token]{}, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nullable.Nullable[account.Token]{}, err
	}
	m, err := a.repo.Create(ctx, account.Account{
		ID:          id.String(),
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Email:       req.Email,
		Password:    account.Password(password),
	})
	if err != nil {
		return nullable.Nullable[account.Token]{}, xerrors.Errorf("create account record: %w", err)
	}
	token, err := a.issueToken(m).SignedString(a.signingKey)
	if err != nil {
		return nullable.Nullable[account.Token]{}, xerrors.Errorf("sign token: %w", err)
	}
	return nullable.New(account.Token(token)), nil
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
	if _, err := mail.ParseAddress(string(req.Email)); err != nil {
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

func (a accountUsecase) issueToken(account *account.Account) *jwt.Token {
	now := time.Now()
	return jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        account.ID,
	})
}

func (a accountUsecase) Login(ctx context.Context, email account.Email, password account.RawPassword) (nullable.Nullable[account.Token], error) {
	m, err := a.repo.FindByEmail(ctx, email)
	if err != nil {
		return nullable.Nullable[account.Token]{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)); err != nil {
		return nullable.Nullable[account.Token]{}, repository.ErrNotFound
	}
	token, err := a.issueToken(m).SignedString(a.signingKey)
	if err != nil {
		return nullable.Nullable[account.Token]{}, xerrors.Errorf("sign token: %w", err)
	}
	return nullable.New(account.Token(token)), nil
}

func (a accountUsecase) VerifyToken(_ context.Context, token account.Token) (string, error) {
	v, err := jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, xerrors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})
	if err != nil {
		return "", err
	} else if !v.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := v.Claims.(jwt.RegisteredClaims)
	if !ok {
		return "", errors.New("parse registered claims")
	}
	return claims.ID, nil
}

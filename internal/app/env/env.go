package env

import (
	"github.com/Netflix/go-env"
	"github.com/taniko/nullable"
	"golang.org/x/xerrors"
)

type Environment struct {
	AppEnv string `env:"APP_ENV,required=true"`
	Host   string `env:"HOST,required=true"`
	DB     struct {
		Host string `env:"DB_HOST,required=true"`
		NAME string `env:"DB_NAME,required=true"`
		User string `env:"DB_USER,required=true"`
		Key  string `env:"DB_KEY,required=true"`
	}
}

func Load() (nullable.Nullable[Environment], error) {
	var environment Environment
	if _, err := env.UnmarshalFromEnviron(&environment); err != nil {
		return nullable.Nullable[Environment]{}, xerrors.Errorf("load environment %w", err)
	}
	return nullable.New(environment), nil
}

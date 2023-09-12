package repository

import (
	"context"
)

type Auth interface{
	Login(ctx context.Context, main, password string)
	Register(ctx context.Context, firstName, main, password string)
}

type Repositories struct{
	Auth Auth
}
package auth

import "context"

type AuthService interface {
	CreateAccount(ctx context.Context)
}
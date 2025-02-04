package domain

import "context"

type UserStore interface {
	CreateUser(ctx context.Context, params CreateUserParams) (*User, error)
	FindUser(ctx context.Context, params FindUserParams) (*User, error)
	PatchUser(ctx context.Context, params PatchUserParams) error
	DeleteUser(ctx context.Context, params DeleteUserParams) error
}

type UserService interface {
	FindUserByMe(ctx context.Context, request FindUserByMeRequest) (*User, error)
	PatchByMeUser(ctx context.Context, request PatchUserByMeRequest) error
}

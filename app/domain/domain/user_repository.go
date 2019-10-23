package domain

import "context"

type UserRepository interface {
	Get(ctx context.Context, id int64) (user *User, err error)
	Save(ctx context.Context, user *User) (err error)
}

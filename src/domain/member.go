package domain

import "context"

type Member struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type MemberUsecase interface {
	Fetch(ctx context.Context) ([]Member, error)
}

type MemberRepository interface {
	Fetch(ctx context.Context) (res []Member, err error)
}

package usecase

import (
	"context"

	"github.com/sendaljpt/go-clean/src/domain"
)

type memberUsecase struct {
	memberRepo domain.MemberRepository
}

func NewMemberUsecase(m domain.MemberRepository) domain.MemberUsecase {
	return &memberUsecase{
		memberRepo: m,
	}
}

func (m *memberUsecase) Fetch(c context.Context) (res []domain.Member, err error) {
	res, err = m.memberRepo.Fetch(c)

	if err != nil {
		return nil, err
	}

	return
}

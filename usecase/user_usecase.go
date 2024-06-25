package usecase

import (
	"context"
	"y-u-y-a/template-go/domain"

	"go.uber.org/zap"
)

type UserUsecase struct {
	logger *zap.Logger
}

func NewUserUsecase(logger *zap.Logger) *UserUsecase {
	return &UserUsecase{logger}
}

// お問い合わせ情報を全件取得する
func (u *UserUsecase) GetUser(ctx context.Context) (*domain.User, error) {
	user := domain.NewUser(1, "test@gmail.com", "テスト太郎")
	return user, nil
}

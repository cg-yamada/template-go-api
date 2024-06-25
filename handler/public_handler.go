package handler

import (
	"fmt"
	"net/http"
	"y-u-y-a/template-go/usecase"

	chi "github.com/go-chi/chi/v5"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

type PublicHandler struct {
	logger      *zap.Logger
	userUsecase *usecase.UserUsecase
}

func NewPublicHandler(
	logger *zap.Logger,
	UserUsecase *usecase.UserUsecase,
) *PublicHandler {
	return &PublicHandler{
		logger,
		UserUsecase,
	}
}

func (h *PublicHandler) Route() http.Handler {
	r := chi.NewRouter()
	r.Get("/user", h.getUser)
	return r
}

// ユーザー情報を1件返す
func (h *PublicHandler) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.userUsecase.GetUser(ctx)
	if err != nil {
		fmt.Println("ユーザー情報の取得に失敗しました", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	pp.Println("user", user)
	// if err := json.NewEncoder(w).Encode(user); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

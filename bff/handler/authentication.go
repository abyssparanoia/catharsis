package handler

import (
	"net/http"

	"github.com/abyssparanoia/catharsis/pkg/log"
	"github.com/abyssparanoia/catharsis/pkg/renderer"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/abyssparanoia/catharsis/bff/service"
	"github.com/abyssparanoia/catharsis/pkg/parameter"
)

// AuthenticationHandler ... authentication handler struct
type AuthenticationHandler struct {
	authenticationService service.Authentication
}

// SignIn ... sign in handler
func (h *AuthenticationHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var param struct {
		UserID   string `json:"user_id" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := parameter.GetJSON(r, &param)
	if err != nil {
		log.Errorf(ctx, "parameter.GetJSON", zap.Error(err))
		renderer.HandleError(ctx, w, "parameter.GetJSON", err)
		return
	}

	v := validator.New()
	if err := v.Struct(param); err != nil {
		renderer.HandleError(ctx, w, "validation error: ", err)
		return
	}

	accessToken, refreshToken, err := h.authenticationService.SignIn(ctx, param.UserID, param.Password)
	if err != nil {
		log.Errorf(ctx, "h.authenticationService.SignIn", zap.Error(err))
		renderer.HandleError(ctx, w, "h.authenticationService.SignIn: ", err)
		return
	}

	renderer.JSON(ctx, w, http.StatusOK, struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

package service

import (
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/constants"
	"github.com/lakhan-purohit/net-http/internal/pkg/request"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/pkg/utils"
	"github.com/lakhan-purohit/net-http/internal/rest-api/repository"
	"github.com/lakhan-purohit/net-http/internal/rest-api/schema"
)

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schema.LoginRequest true "Login Credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/v1/public/auth/login [post]
func LoginHandler(repo repository.IAuthRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req schema.LoginRequest

		if err := request.Bind(r, &req); err != nil {
			response.BadRequest(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		user, err := repo.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			response.UnauthorizedAccess(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		response.Success(response.SendParams{
			W:    w,
			Data: user,
		})
	}
}

// @Summary Sign up
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Desired username" example("johndoe")
// @Param email formData string true "User email" example("john@example.com")
// @Param password formData string true "User password" example("password123")
// @Param avatar formData file true "Avatar image file"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/v1/public/auth/sign-up [post]
func SignUpHandler(repo repository.IAuthRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req schema.SignUpRequest

		if err := request.Bind(r, &req); err != nil {
			response.BadRequest(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		avatar, _ := utils.SaveSingle(req.Avatar, constants.UserAvatarDir)

		user, err := repo.SignUp(req.Username, req.Email, req.Password, avatar.Name)
		if err != nil {
			response.InternalError(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		jwtService := utils.NewJWT()

		token, refresh, _ := jwtService.Generate(utils.Claims{
			UserID: user.ID,
			Email:  user.Email,
			Role:   "user",
			UUID:   user.UUID,
		})

		user.Token = token
		user.RefreshToken = refresh

		response.Success(response.SendParams{
			W:    w,
			Data: user,
		})
	}
}

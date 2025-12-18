package service

import (
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/request"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/rest-api/model"
	"github.com/lakhan-purohit/net-http/internal/rest-api/repository"
)

// @Summary Get user list
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit for pagination" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} response.UserListResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /api/v1/private/user/get-list [get]
func UserGetListHandler(repo repository.IUserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse pagination
		var pg request.PaginationRequest
		if err := request.BindQuery(r, &pg); err != nil {
			response.BadRequest(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		// Defaults
		if pg.Limit == 0 {
			pg.Limit = 10
		}

		users, err := repo.GetList(r.Context(), pg.Limit, pg.Offset)
		if err != nil {
			response.InternalError(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		response.Success(response.SendParams{
			W:    w,
			Data: users,
		})
	}
}

// UserGetFullListHandler demonstrates a "Complex API" fetch
// It fetches users and their stats, showing how to avoid the N+1 problem.
// @Summary Get user list with statistics (Complex Example)
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} response.UserFullListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/private/user/get-full-list [get]
func UserGetFullListHandler(repo repository.IUserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var pg request.PaginationRequest
		if err := request.BindQuery(r, &pg); err != nil {
			response.BadRequest(response.SendParams{
				W:       w,
				Message: err.Error(),
			})
			return
		}

		// Defaults
		if pg.Limit == 0 {
			pg.Limit = 10
		}

		// 1. Fetch Users
		users, err := repo.GetList(r.Context(), pg.Limit, pg.Offset)
		if err != nil {
			response.InternalError(response.SendParams{W: w, Message: err.Error()})
			return
		}

		// 2. Collect IDs (to fetch related data in ONE query)
		userIDs := make([]int64, len(users))
		for i, u := range users {
			userIDs[i] = u.ID
		}

		// 3. Fetch Stats in Batch (Scalable way)
		statsMap, err := repo.GetStatsForUsers(r.Context(), userIDs)
		if err != nil {
			response.InternalError(response.SendParams{W: w, Message: err.Error()})
			return
		}

		// 4. Merge into a complex structure (User + Stats)
		fullList := make([]model.UserWithStats, len(users))
		for i, u := range users {
			fullList[i] = model.UserWithStats{
				User:  *u,
				Stats: statsMap[u.ID],
			}
		}

		response.Success(response.SendParams{
			W:    w,
			Data: fullList,
		})
	}
}

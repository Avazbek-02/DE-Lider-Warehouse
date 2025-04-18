package handler

import (
	"strconv"

	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/entity"
	"github.com/gin-gonic/gin"
)

// GetSession godoc
// @Router /session/{id} [get]
// @Summary Get a session by ID
// @Description Get a session by ID
// @Security BearerAuth
// @Tags session
// @Accept  json
// @Produce  json
// @Param id path string true "Session ID"
// @Success 200 {object} entity.Session
// @Failure 400 {object} entity.ErrorResponse
func (h *Handler) GetSession(ctx *gin.Context) {
	var (
		req entity.Id
	)

	req.ID = ctx.Param("id")

	session, err := h.UseCase.SessionRepo.GetSingle(ctx, req)
	if h.HandleDbError(ctx, err, "Error getting session") {
		return
	}

	ctx.JSON(200, session)
}

// GetSessions godoc
// @Router /session/list [get]
// @Summary Get a list of users
// @Description Get a list of users
// @Security BearerAuth
// @Tags session
// @Accept  json
// @Produce  json
// @Param page query number true "page"
// @Param limit query number true "limit"
// @Param user_id query string false "user_id"
// @Success 200 {object} entity.SessionList
// @Failure 400 {object} entity.ErrorResponse
func (h *Handler) GetSessions(ctx *gin.Context) {
	var (
		req entity.GetListFilter
	)

	// Retrieve query parameters with defaults
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")
	userId := ctx.DefaultQuery("user_id", "")

	// Override user_id based on user_type if it's "user"
	if ctx.GetHeader("user_type") == "user" {
		userId = ctx.GetHeader("sub")
	}

	// Convert page and limit to integers and handle potential errors
	req.Page, _ = strconv.Atoi(page)
	req.Limit, _ = strconv.Atoi(limit)

	// Add user_id filter
	if userId != "" {
		req.Filters = append(req.Filters, entity.Filter{
			Column: "user_id",
			Type:   "eq",
			Value:  userId,
		})
	}

	// Set default order by created_at descending
	req.OrderBy = append(req.OrderBy, entity.OrderBy{
		Column: "created_at",
		Order:  "desc",
	})

	// Fetch sessions from the repository
	sessions, err := h.UseCase.SessionRepo.GetList(ctx, req)
	if err != nil {
		// Handle the error appropriately and return a proper response
		h.HandleDbError(ctx, err, "Error getting session")
		return
	}

	// Return the session data as a JSON response
	ctx.JSON(200, sessions)
}

// UpdateSession godoc
// @Router /session [put]
// @Summary Update a session
// @Description Update a session
// @Security BearerAuth
// @Tags session
// @Accept  json
// @Produce  json
// @Param session body entity.Session true "Session object"
// @Success 200 {object} entity.Session
// @Failure 400 {object} entity.ErrorResponse
func (h *Handler) UpdateSession(ctx *gin.Context) {
	var (
		body entity.Session
	)

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		h.ReturnError(ctx, config.ErrorBadRequest, "Invalid request body", 400)
		return
	}

	session, err := h.UseCase.SessionRepo.Update(ctx, body)
	if h.HandleDbError(ctx, err, "Error updating session") {
		return
	}

	ctx.JSON(200, session)
}

// DeleteSession godoc
// @Router /session/{id} [delete]
// @Summary Delete a session
// @Description Delete a session
// @Security BearerAuth
// @Tags session
// @Accept  json
// @Produce  json
// @Param id path string true "Session ID"
// @Success 200 {object} entity.SuccessResponse
// @Failure 400 {object} entity.ErrorResponse
func (h *Handler) DeleteSession(ctx *gin.Context) {
	var (
		req entity.Id
	)

	req.ID = ctx.Param("id")

	err := h.UseCase.SessionRepo.Delete(ctx, req)
	if h.HandleDbError(ctx, err, "Error deleting session") {
		return
	}

	ctx.JSON(200, entity.SuccessResponse{
		Message: "Session deleted successfully",
	})
}

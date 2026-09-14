package create

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/London57/profiles/internal/data/datagen"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	create "github.com/London57/profiles/internal/uc/create"
	get_by_email "github.com/London57/profiles/internal/uc/get_by_email"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ProfileCreateHandler struct {
	create create.ProfileCreate
	getByEmail get_by_email.GetProfileByEmail
}

func (ProfileCreateHandler) New(create create.ProfileCreate, gbe get_by_email.GetProfileByEmail) ProfileCreateHandler {
	return ProfileCreateHandler{
		create: create,
		getByEmail: gbe,
	}
}

// @Summary CreateProfile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param request body datagen.CreateProfileParams true "Record to create" 
// @Success 201 {object} any "Record created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 409 {object} map[string]interface{} "user with this email already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /profiles/create [patch]
func (handler ProfileCreateHandler) CreateProfile(r *gin.Context) {
	req := datagen.CreateProfileParams{}
	err := r.Bind(&req)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("failed to parse JSON: %s", err.Error()),
		})
		return
	}
	_, err = handler.getByEmail.Exec(r.Request.Context(), req.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		r.JSON(http.StatusInternalServerError, err)
		return
	}
	if err == nil {
		r.JSON(http.StatusConflict, gin.H{
			"error": "user with this email already exists",
		})
		return
	}

	resp, err := handler.create.Exec(r.Request.Context(), req)
	if err != nil {
		r.JSON(http.StatusInternalServerError, err)
		return
	}
	r.JSON(http.StatusCreated, resp)
}




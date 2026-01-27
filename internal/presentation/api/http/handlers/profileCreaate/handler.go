package handlers

import (
	"errors"
	"net/http"

	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	interactors "github.com/London57/profiles/internal/uc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ProfileCreateHandler struct {
	create interactors.ProfileCreate
	getByEmail interactors.GetProfileByEmail
}

func (handler ProfileCreateHandler) CreateProfile(r *gin.Context) {
	req := request.ProfileCreateRequest{}
	err := r.Bind(&req)
	if err != nil {
		r.JSON(http.StatusBadRequest, err)
		return
	}
	_, err = handler.getByEmail.Exec(r.Request.Context(), req.Email)
	if !errors.Is(err, pgx.ErrNoRows) {
		r.JSON(http.StatusInternalServerError, err)
		return
	}
	if err == nil {
		r.JSON(http.StatusConflict, "user with this email already exists")
		return
	}

	resp, err := handler.create.Exec(r.Request.Context(), req)
	if err != nil {
		r.JSON(http.StatusInternalServerError, err)
		return
	}
	r.JSON(http.StatusCreated, resp)
}

func (handler ProfileUpdateHandler) UpdateProfile(r *gin.Context) {
	req := request.ProfileUpdateRequest{}
	err := r.Bind(&req)
	if err != nil {
		r.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := handler.update.Exec(r.Request.Context(), req)
	if err != nil {
		r.JSON(http.StatusInternalServerError, err)
		return
	}
	r.JSON(http.StatusCreated, resp)
}


package handlers

import (
	"errors"
	"net/http"

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




package handlers

import (
	"net/http"

	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	interactors "github.com/London57/profiles/internal/uc"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateHandler struct {
	update interactors.ProfileUpdate
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


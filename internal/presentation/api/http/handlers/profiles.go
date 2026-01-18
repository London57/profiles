package handlers

import (
	"net/http"

	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	interactors "github.com/London57/profiles/internal/uc"
	"github.com/gin-gonic/gin"
)

type ProfilesHandler struct {
	create interactors.ProfileCreate
}

func (handler ProfilesHandler) CreateProfile(r *gin.Context) {
	req := request.ProfileCreateRequest{}
	err := r.Bind(&req)
	if err != nil {
		r.JSON(http.StatusBadRequest, err)
	}
	resp, err := handler.create.Exec(r.Request.Context(), req)
	if err != nil {
		r.JSON(http.StatusInternalServerError, err)
	}
	r.JSON(http.StatusCreated, resp)
}

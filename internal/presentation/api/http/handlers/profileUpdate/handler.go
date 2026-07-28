package update

import (
	"net/http"

	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	update "github.com/London57/profiles/internal/uc/update"
	"github.com/gin-gonic/gin"
)

func (ProfileUpdateHandler) New(update update.ProfileUpdate) ProfileUpdateHandler {
	return ProfileUpdateHandler{
		update: update,
	}
}

type ProfileUpdateHandler struct {
	update update.ProfileUpdate
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
	r.JSON(http.StatusOK, resp)
}


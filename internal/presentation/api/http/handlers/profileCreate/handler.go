package create

import (
	"errors"
	"fmt"
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

func  (ProfileCreateHandler) NewProfleCreateHandler(create create.ProfileCreate, gbe get_by_email.GetProfileByEmail) ProfileCreateHandler {
	return ProfileCreateHandler{
		create: create,
		getByEmail: gbe,
	}
}


func (handler ProfileCreateHandler) CreateProfile(r *gin.Context) {
	req := request.ProfileCreateRequest{}
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




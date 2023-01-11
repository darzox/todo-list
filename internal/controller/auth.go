package controller

import (
	"fmt"
	"net/http"

	"todo-app/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body model.User true "registration info"
// @Success 200 {integer} int
// @Failure 400,404 {object} errorHandler
// @Failure 500 {object} errorHandler
// @Failure default {object} errorHandler
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {
	type singInInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

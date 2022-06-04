package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin_learn/models/usermodel"
	"github.com/gin_learn/util"
)

type userHandler struct {
	userService usermodel.Service
}

func NewUserHandler(userService usermodel.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var user usermodel.UserInput
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	createdUser, err := h.userService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": createdUser,
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var user usermodel.UserInput
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	u := h.userService.Login(user)

	if u.Id == 0 {
		c.JSON(404, gin.H{
			"message": "User Not Found",
		})
		return
	}

	if err := u.ComparePassword(user.Password); err != nil {
		c.JSON(404, gin.H{
			"message": "Incorrect password",
		})
		return
	}

	// jwt token
	token, err := util.GenerateJWT(strconv.Itoa(int(u.Id)))
	if err != nil {
		c.JSON(500, gin.H{"message": err})
		return
	}

	c.SetCookie("jwt", token, int(time.Now().Add(time.Hour*24).Unix()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"jwt":     token,
		"message": "Succes",
	})
}

func (h *userHandler) User(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	id, _ := util.ParseJwt(cookie)
	// if err != nil {
	// 	return
	// }

	convert, _ := strconv.Atoi(id)
	user := h.userService.GetUser(convert)
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// TODO: answer here

	// Membuat model userLogin
	var user model.UserLogin
	// Melakukan decode dan melakukan error hendling
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}
	// error hendling ketika data login kosong
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("email or password is empty"))
		return
	}
	// membuat modelUser
	modelUser := model.User{
		Email:    user.Email,
		Password: user.Password,
	}
	// error hendling ketika mengalami error ketika mengakses serviceLogin
	token, err := u.userService.Login(&modelUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	// mengambalikan data berupa token (JWT) yang akan disimpan ke cookie
	cookie, err := c.Cookie("session_token")
	// cek apakan sudah terdapat cookie
	if err != nil {
		cookie = *token
		c.SetCookie("session_token", cookie, 3600, "/", "", false, true)
	}
	// mengembalikan status succes
	c.JSON(http.StatusOK, gin.H{
		// "user_id": *token,
		"message": "login success",
	})
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// TODO: answer here
	task, err := u.userService.GetUserTaskCategory()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

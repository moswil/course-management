package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/user-service/internal/core/dto"
	"github.com/moswil/course-management/user-service/internal/core/port/service"
)

// UserController handles HTTP requests related to courses.
type UserController struct {
	userService service.UserService
}

// NewCourseController creates a new UserController instance.
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateCourse handles the creation of a new course.
func (c *UserController) CreateUser(ctx *gin.Context) {
	var createRequest dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := c.userService.CreateUser(ctx, &createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": err.Error()})
		return
	}

	res := &dto.UserDTO{
		UserId: user.UserId,
	}

	ctx.JSON(http.StatusCreated, res)
}

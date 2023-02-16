package router

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/AlbatozK/go_backend_boilerplate/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	service *service.UserService
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		service: service.NewUserService(),
	}
}

func (ur *UserRouter) Init(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id", ur.GetById)
	}
}

func (ur *UserRouter) GetById(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	user, err := ur.service.GetUser(intId)
	switch err {
	case nil:
		c.JSON(http.StatusOK, user)
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

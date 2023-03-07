package router

import (
	"io"
	"os"

	"github.com/choisangh/board-crud-backend/pkg/api"
	"github.com/choisangh/board-crud-backend/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"authorization", "Content-Type"}

	r.Use(cors.New(config))

	boardRouter := r.Group("/board")

	boardRouter.GET("", apis.GetBoardList)
	boardRouter.POST("", apis.CreateBoard)
	boardRouter.GET("/:id", apis.GetBoardByID)
	boardRouter.PATCH("/:id", apis.UpdateBoard)
	boardRouter.DELETE("/:id", apis.DeleteBoardByID)

	userRouter := r.Group("/user")
	userRouter.POST("register", apis.CreateUser)
	userRouter.POST("login", apis.LoginUser)

	usersRouter := r.Group("/users")
	usersRouter.GET("/:id", apis.GetUserByID)

	userAuthRouter := r.Group("/user")
	userAuthRouter.Use(utils.JWTAuthMiddleware())
	userAuthRouter.GET("current", apis.GetUserCurrent)

	userlistRouter := r.Group("/userlist")
	userlistRouter.Use(utils.JWTAuthMiddleware())
	userlistRouter.GET("", apis.GetUserList)

	return r
}

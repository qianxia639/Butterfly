package handler

import (
	"Butterfly/config"
	db "Butterfly/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
	conf   config.Config
	store  *db.Store
}

func NewHandler(conf config.Config,store *db.Store) *Handler {

	handler := &Handler{
		conf: conf,
		store: store,
	}

	handler.setupRouter()

	return handler
}

func (handler *Handler) setupRouter() {
	router := gin.Default()

	router.GET("/example", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successfully..."})
	})

	// Router
	{
		router.POST("/refresh")
		router.PUT("/reset/pwd")
		router.POST("/email/code")
	}

	// User Router
	{
		router.POST("/login")
		router.POST("/user")

	}

	// Friend Request Router
	{
	}

	// Friendship Router
	{
	}

	// Group Router
	{
	}

	handler.Router = router
}

func (h *Handler) Start(address string) error {
	return h.Router.Run(address)
}

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	_ "github.com/mustafayilmazdev/musarchive/doc/statik"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/rakyll/statik/fs"
)

// * Note [codermuss]: Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	Router     *gin.Engine
	lm         *localization.LocalizationManager
}

// * Note [codermuss]: NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker, lm: localization.GetInstance()}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.New()

	router.Use(ZerologMiddleware())
	router.Use(gin.Recovery())

	// Serve the API endpoints
	api := router.Group("/v1")
	{
		api.GET("/onboardings", server.GetOnboardings)
		api.POST("/register", server.RegisterUser)
		api.POST("/login", server.LoginUser)

	}
	authRoutes := router.Group("/v1").Use(authMiddleware(server.tokenMaker))
	{
		authRoutes.GET("/posts", server.GetPosts)
		authRoutes.GET("/followed_posts", server.GetFollowedPosts)
		authRoutes.POST("/create_post", server.CreatePost)
	}

	// Serve the bundled static files
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	router.StaticFS("/swagger", statikFS)
	server.Router = router
}

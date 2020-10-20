package main

import (
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

type login struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"text":     "Hello World.",
	})
}

func main() {

	userRepository := infra.NewUserPersistence()

	// the jwt middleware
	authMiddleware, err := authMiddlewareByJWT()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST","PUT", "PATCH","DELETE"},
		AllowHeaders:     []string{"Origin","Access-Control-Allow-Origin","Content-type","Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/login", authMiddleware.LoginHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
	router.POST("/new", handler.TemporaryRegisterUser)
	router.POST("/register", handler.RegisterUser)
	router.GET("/evaluated",handler.CheckUserAlreadyLikedMovie)

	api := router.Group("/api/v1")
	api.GET("/movie-and-comments",handler.GetMovieAndComments)
	api.GET("/index-movies",handler.IndexMovie)

	auth := router.Group("/auth/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/user", handler.GetLoggedInUserData)

		changeUserNameUsecase := usecase.NewChangeUserName(userRepository)
		changeUserNameHandler := handler.NewChangeUserName(userRepository,changeUserNameUsecase)
		auth.PUT("/user-name", changeUserNameHandler.ChangeUserName)

		changePasswordUsecase := usecase.NewChangePassword(userRepository)
		changePasswordHandler := handler.NewChangePassword(userRepository,*changePasswordUsecase)
		auth.PUT("/password", changePasswordHandler.ChangePassword)

		auth.PUT("/profile-image", handler.ChangeUserProfileImage)
		auth.POST("/comments", handler.PostComment)
		auth.GET("/hello", helloHandler)
		auth.POST("/movie", handler.UploadMovieFile)
		auth.PUT("/movie",handler.UpdateMovie)
		auth.PUT("/thumbnail",handler.ChangeThumbnail)
		auth.POST("/evaluates", handler.EvaluateMovie)
		auth.POST("/play-lists",handler.CreatePlayList)
		auth.POST("/play-list-items",handler.AddPlayListMovie)
		auth.PUT("/play-list-items",handler.ChangeOrderOfPlayListMovies)
		auth.POST("/follows",handler.FollowUser)

		auth.GET("/movies",handler.GetUploadedMovies)
		auth.GET("/play-lists",handler.IndexPlayListsInMyPage)
		auth.GET("/play-list-items/:play_list_id",handler.IndexPlaylistMovies)

		auth.GET("play-lists/:movie_id",handler.IndexPlayListInMovieListPage)

		auth.DELETE("/play-list-items",handler.DeletePlayListMovie)
		auth.DELETE("/play-lists",handler.DeletePlayList)
	}

	router.GET("/health",func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Healthy."})
	})

	router.Run()
}

func authMiddlewareByJWT() (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       os.Getenv("Realm"),
		Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(model.UserID); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			//claims := jwt.ExtractClaims(c)
			//return &User{
			//	UserName: claims[identityKey].(string),
			//}
			return true
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email, _ := model.NewUserEmail(loginVals.Email)
			password := loginVals.Password

			userRepository := infra.NewUserPersistence()
			userExistsUsecaes := usecase.NewUserExists(userRepository)

			userExists, err := userExistsUsecaes.CheckUserExistsForAuth(email, password)

			if userExists != nil && err == nil {
				return userExists.ID, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		//TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

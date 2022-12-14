package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/config"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/endpoint"
	adminEndpoint "github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/endpoint/admin"
	userEndpoint "github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/endpoint/user"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/middleware"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load app.env file data to struct
	config, err := config.LoadConfig(".")
	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn: "",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	r := gin.Default()

	allowedOrigins := "*"
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWildcard = true
	corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",") // contain whitelist domain
	corsConfig.AllowHeaders = []string{"*", "Content-Type", "Accept"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	r.GET("/", func(c *gin.Context) {
		var success string = fmt.Sprintf("Server listening with version %s", config.Version)
		c.JSON(http.StatusOK, &model.Response{
			Success: true,
			Error:   nil,
			Msg:     success,
			Data:    nil,
		})
	})
	r.POST("/login", endpoint.Login)

	// Admin Endpoint
	admin := r.Group("/admin")
	admin.PUT("/info/update", adminEndpoint.AdminChangeInfo)

	adminSubmission := admin.Group("/submission")
	adminSubmission.Use(middleware.ValidateRoleAccess)
	adminSubmission.GET("/list", adminEndpoint.SubmissionList)
	adminSubmission.POST("/approve", adminEndpoint.SubmissionApprove)
	adminSubmission.POST("/reject", adminEndpoint.SubmissionReject)

	adminRoom := admin.Group("/room")
	adminRoom.Use(middleware.ValidateRoleAccess)
	adminRoom.POST("/register", adminEndpoint.RegisterRoom)
	adminRoom.POST("/list", adminEndpoint.ListRoom)

	// User Endpoint
	user := r.Group("/user")
	user.POST("/register", userEndpoint.Register)
	user.PUT("/info/update", userEndpoint.UserChangeInfo)
	user.GET("/token/validate", userEndpoint.UserValidateToken)

	userSubmission := user.Group("/submission")
	userSubmission.Use(middleware.MiddlewareValidateToken)
	userSubmission.POST("/create", userEndpoint.SubmissionCreate)
	userSubmission.PUT("/update", userEndpoint.SubmissionUpdate)
	userSubmission.DELETE("/delete", userEndpoint.SubmissionDelete)
	userSubmission.GET("/list", userEndpoint.SubmissionList)

	userRoom := user.Group("/room")
	userRoom.Use(middleware.MiddlewareValidateToken)
	userRoom.GET("/list", userEndpoint.UserRoomList)

	port, _ := strconv.Atoi(config.Port)
	log.Infof("Service version: %s", config.Version)
	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error(err)
	}
}

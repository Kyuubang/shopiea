package main

import (
	"flag"
	"fmt"
	"github.com/Kyuubang/shopiea/db"
	"github.com/Kyuubang/shopiea/handlers"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var (
	migrate bool
)

func init() {
	// set flag migrate to run database migration
	flag.BoolVar(&migrate, "migrate", false, "run database migration")
	// set gin mode
	if os.Getenv("SHOPIEA_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	// init db
	var conn = db.Config{
		Host:     os.Getenv("SHOPIEA_DB_HOST"),
		Port:     os.Getenv("SHOPIEA_DB_PORT"),
		User:     os.Getenv("SHOPIEA_DB_USER"),
		Password: os.Getenv("SHOPIEA_DB_PASSWORD"),
		DBName:   os.Getenv("SHOPIEA_DB_NAME"),
		TZ:       os.Getenv("SHOPIEA_TZ"),
	}

	err := conn.InitDB(migrate)
	if err != nil {
		panic(err)
	}

	// init router
	var router *gin.Engine

	// check if environment is production
	if os.Getenv("SHOPIEA_ENV") == "production" {
		// create a new logger
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		// create a new Gin engine
		router = gin.New()

		// add the JSON logging middleware to the router
		router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	} else {
		router = gin.Default()
	}

	// authentication endpoints
	// handlers for login user
	router.POST("/auth/login", handlers.Login)

	// handlers versioning - version 1
	apiV1 := router.Group("/v1", handlers.AuthMiddleware())
	{
		// handlers for get all course
		apiV1.GET("/course", handlers.GetCourses)
		// handlers for get all labs
		apiV1.GET("/labs", handlers.GetLabs)
		// handlers for push score
		apiV1.POST("/score", handlers.PushScore)
		// handlers for get score
		apiV1.GET("/score", handlers.GetScore)
		// handlers for check token with middleware
		apiV1.POST("/auth/check", handlers.CheckToken)

		// admin handlers
		admin := apiV1.Group("/admin", handlers.AdminOnly())
		{
			// handlers for check admin
			admin.POST("/check", handlers.CheckToken)
			// handlers for get all user
			admin.GET("/user", handlers.GetUsers)
			// handlers for get all class
			admin.GET("/class", handlers.GetClasses)
			// handlers for create class
			admin.POST("/class", handlers.CreateClass)
			// handlers for create course
			admin.POST("/course", handlers.CreateCourse)
			// handlers for create labs
			admin.POST("/labs", handlers.CreateLabs)
			// handlers for create user
			admin.POST("/user", handlers.CreateUser)
			// handlers for export score
			admin.GET("/export", handlers.ExportScore)
		}
	}

	var addr = os.Getenv("SHOPIEA_HOST") + ":" + os.Getenv("SHOPIEA_PORT")
	// create http server on port 8080
	server := &http.Server{
		Addr:         addr, //+ os.Getenv("SHOPIEA_PORT"),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = server.ListenAndServe()
	fmt.Println("Server running on port", addr)

	if err != nil {
		panic(err)
	}

}

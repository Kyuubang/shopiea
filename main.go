package main

import (
	"encoding/json"
	"flag"
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
	seed    bool
)

func init() {
	// set flag migrate to run database migration
	flag.BoolVar(&migrate, "migrate", false, "run database migration")
	// set flag seed to run database seeding
	flag.BoolVar(&seed, "seed", false, "run database seeding")
	// set gin mode
	if os.Getenv("SHOPIEA_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	// parse flags
	flag.Parse()

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

	// run database seeding if seed flag is set
	if seed {
		err = db.Seed()
		if err != nil {
			panic(err)
		}
		os.Exit(0)
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

	// handlers for anonymous function config
	router.GET("/info", func(c *gin.Context) {
		// json object that include date, config
		data, err := json.Marshal(map[string]interface{}{
			"date":    time.Now().Format(time.RFC822),
			"version": "1.0.0",
			"config": map[string]string{
				"case_repo":    os.Getenv("SHOPIEA_CASE_REPO"),
				"case_branch":  os.Getenv("SHOPIEA_CASE_BRANCH"),
				"infra_repo":   os.Getenv("SHOPIEA_INFRA_REPO"),
				"infra_branch": os.Getenv("SHOPIEA_INFRA_BRANCH"),
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when marshal json",
			})
			return
		}

		// return data json object to client with status code 200
		c.Data(http.StatusOK, "application/json", data)
		return
	})

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
			// handlers for create user
			admin.POST("/user", handlers.CreateUser)
			// handlers for delete user
			admin.DELETE("/user", handlers.DeleteUser)
			// handlers for update user
			admin.PUT("/user", handlers.UpdateUser)

			// handlers for get all class
			admin.GET("/class", handlers.GetClasses)
			// handlers for create class
			admin.POST("/class", handlers.CreateClass)
			// handlers for delete class
			admin.DELETE("/class", handlers.DeleteClass)
			// handlers for update class
			admin.PUT("/class", handlers.UpdateClass)

			// handlers for create course
			admin.POST("/course", handlers.CreateCourse)
			// handlers for update course
			admin.PUT("/course", handlers.UpdateCourse)
			// handlers for delete course
			admin.DELETE("/course", handlers.DeleteCourse)

			// handlers for create labs
			admin.POST("/labs", handlers.CreateLabs)
			// handlers for update labs
			admin.PUT("/labs", handlers.UpdateLabs)
			// handlers for delete labs
			admin.DELETE("/labs", handlers.DeleteLabs)

			// handlers for export score
			admin.GET("/export", handlers.ExportScore)

		}
	}

	// create http server on port 8080
	server := &http.Server{
		Addr:         ":9898",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

package main

import (
	"awx-consul-inventory/handlers/healthcheck"
	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "awx-consul-inventory"))

	// Start routes
	router.GET("/health", healthcheck.HealthCheck)

	// RUN rabit run
	router.Run() // listen and serve on 0.0.0.0:8080
}
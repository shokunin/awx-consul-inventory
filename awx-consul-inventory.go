package main

import (
	"awx-consul-inventory/handlers/awx"
	"awx-consul-inventory/handlers/consul"
	"awx-consul-inventory/handlers/healthcheck"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "awx-consul-inventory"))

	// Start routes
	router.GET("/health", healthcheck.HealthCheck)
	router.GET("/fail", awx.GetFailedHosts)
	router.GET("/consul/:server/:inventoryname", consul.GenInventory)
	router.GET("/nodes/:server/:inventoryname", consul.GenNodes)

	// RUN rabit run
	router.Run() // listen and serve on 0.0.0.0:8080
}

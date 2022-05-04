package api

import (
	"github.com/gin-gonic/gin"
	"vsign/service"
)

func HandleAppStoreApi(s *gin.Engine) {
	// 证书
	cert := s.Group("/asca/certificates")
	{
		cert.POST("/", service.CreateCert)
		cert.GET("/", service.ListCert)
		cert.GET("/:id", service.DownloadCert)
		cert.DELETE("/:id", service.RevokeCert)
	}

	// bundle ids

	// bundle id capabilities

	// devices

	// profiles
}

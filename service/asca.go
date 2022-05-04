package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vsign/applestore"
	"vsign/logger"
	"vsign/model"
	"vsign/result"
)

const api = "https://api.appstoreconnect.apple.com"

// ListCert List and Download Certificates
// https://developer.apple.com/documentation/appstoreconnectapi/list_and_download_certificates
func ListCert(c *gin.Context) {
	res, err := applestore.HttpRequest(c.GetHeader("token"), http.MethodGet, api+"/v1/certificates", nil)
	if err != nil {
		logger.Error("%s", err)
		c.JSON(http.StatusInternalServerError, result.R{}.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	if res.StatusCode() != http.StatusOK && res.StatusCode() != http.StatusCreated {
		var d model.ErrorResponse
		json.Unmarshal(res.Body(), &d)
		logger.Error("证书获取异常: %+v", d)
		c.JSON(res.StatusCode(), result.R{}.Error(res.StatusCode(), fmt.Sprintf("%+v", d)))
		return
	}

	var data model.Certificates
	json.Unmarshal(res.Body(), &data)
	c.JSON(http.StatusOK, result.R{}.Success(data))
}

// CreateCert Create a Certificate
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_certificate
func CreateCert(c *gin.Context) {

}

// DownloadCert Read and Download Certificate Information
// https://developer.apple.com/documentation/appstoreconnectapi/read_and_download_certificate_information
func DownloadCert(c *gin.Context) {

}

// RevokeCert Revoke a Certificate
// https://developer.apple.com/documentation/appstoreconnectapi/revoke_a_certificate
func RevokeCert(c *gin.Context) {

}

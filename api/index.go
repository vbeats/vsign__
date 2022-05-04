package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"vsign/applestore"
	"vsign/config"
	"vsign/logger"
)

func init() {

	go func() {
		// 加载配置
		cfg, err := os.OpenFile("vsign.yml", os.O_RDONLY, fs.ModePerm)
		if err != nil {
			logger.Error("配置文件加载失败：", err.Error())
			return
		}
		content, _ := ioutil.ReadAll(cfg)
		config := &config.Config{}
		yaml.Unmarshal(content, config)

		gin.SetMode(gin.ReleaseMode)
		s := gin.New()

		s.Use(gin.Recovery())
		s.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			Output:    os.Stdout,
			SkipPaths: []string{"/favicon.ico"},
		}))

		// 中间件

		// 生成jwt token
		applestore.HandleJwt(s)

		// *************************************appstore connect api***************************************

		HandleAppStoreApi(s)

		// *************************************fastlane***************************************

		logger.Info("vsign api server started and listening on port %d", config.Port)
		s.Run(":" + strconv.Itoa(config.Port))
	}()

}

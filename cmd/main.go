package main

// @title 项目标题
// @version 0.0.1
// @description 这是一个gin web开发脚手架
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath /api/
func main() {
	webApp, err := initWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Start()
}

func Init() {
	// "xs.bbs/pkg/log"
	// "xs.bbs/pkg/tool/snowflake"
	// err = log.Init(config)
	// if err != nil {
	// 	return nil, err
	// }
	// err = snowflake.Init(config)
	// if err != nil {
	// 	return nil, err
	// }
}

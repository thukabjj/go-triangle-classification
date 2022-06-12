package main

import (
	_ "github.com/thukabjj/go-triangle-classification/docs"
	"github.com/thukabjj/go-triangle-classification/infrastructure"
)

// @title     		Triangle API
// @version         1.0
// @description    	A triangle management service API in Go using Gin framework.
// @termsOfService  https://www.linkedin.com/in/arthur-alves-da-costa/
// @contact.name   	Arthur Alves
// @contact.url    	https://twitter.com/prayformercy_tv
// @contact.email  	arthur.alvesdeveloper@gmail.com
// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html
// @host      		localhost:8080
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
// @BasePath  /
func main() {
	infrastructure.Run()
}

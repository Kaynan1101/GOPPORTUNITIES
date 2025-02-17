package router

import "github.com/gin-gonic/gin"

func Initalize() {

	//Inicializa a Router utilizando as configurações do Default da Gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//rodando a API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

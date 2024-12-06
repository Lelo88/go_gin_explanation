package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Bienvenido a la aplicación de ejemplo!")
	// 01- creamos una nueva instancia de gin
	router := gin.Default()

	// 02- definimos una ruta
	// Se le envía un mensaje en formato JSON, con un endpoint "/" y una función anónima que recibe un contexto
	// y devuelve un mensaje en formato JSON
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 200 es el código de estado HTTP que indica que la solicitud se ha completado con éxito
			"message": "Hola Mundo!",
		})
	})

	// 03- arrancamos el servidor en el puerto 8080
	router.Run(":8080")
}
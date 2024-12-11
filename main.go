package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"` // `json:"id"` es una etiqueta que se utiliza para serializar y deserializar objetos JSON
	Title  string  `json:"title"` // `json:"title"` es una etiqueta que se utiliza para serializar y deserializar objetos JSON
	Artist string  `json:"artist"` // `json:"artist"` es una etiqueta que se utiliza para serializar y deserializar objetos JSON
	Price  float64 `json:"price"` // `json:"price"` es una etiqueta que se utiliza para serializar y deserializar objetos JSON
}

// Creamos una lista de álbumes
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//Creamos un controlador
func getAlbums(c *gin.Context) {
	// la recomendación de golang es utilizar el paquete http para devolver respuestas HTTP
	c.IndentedJSON(http.StatusOK, albums) // 200 es el código de estado HTTP que indica que la solicitud se ha completado con éxito
}

// vamos a crear una función que nos permita agregar un elemento al listado en memoria
func postAlbum(c *gin.Context) {
	var newAlbum album
	
	// BindJSON es un método que se utiliza para enlazar la solicitud JSON con la estructura de datos
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum) // 201 es el código de estado HTTP que indica que la solicitud se ha completado con éxito y se ha creado un nuevo recurso
}

// vamos a crear una función que nos permita obtener un álbum por su ID a través de una URL
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {
	fmt.Println("Bienvenido a la aplicación de ejemplo!")
	// 01- creamos una nueva instancia de gin
	router := gin.Default()

	/* // 02- definimos una ruta
	// Se le envía un mensaje en formato JSON, con un endpoint "/" y una función anónima que recibe un contexto
	// y devuelve un mensaje en formato JSON
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 200 es el código de estado HTTP que indica que la solicitud se ha completado con éxito
			"message": "Hola Mundo!",
		})
	})

	// 03- arrancamos el servidor en el puerto 8080
	router.Run(":8080") */

	// 04 - Definimos una ruta para obtener todos los álbumes
	router.GET("/albums", getAlbums)
	// 05 - Vamos a crear un nuevo album
	router.POST("/albums", postAlbum)

	// 06 - Vamos a crear una ruta para obtener un álbum por su ID
	router.GET("/albums/:id", getAlbumsByID)

	// 07 - arrancamos el servidor en el puerto 8080
	router.Run(":3000")

}
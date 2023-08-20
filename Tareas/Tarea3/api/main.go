package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type Album struct {
	AlbumID         int    `json:"album_id"`
	TituloAlbum     string `json:"titulo_album"`
	Artista         string `json:"artista"`
	AnioLanzamiento string `json:"anio_lanzamiento"`
	GeneroMusical   string `json:"genero_musical"`
}

func main() {
	// Configurar las credenciales de la base de datos desde variables de entorno
	dbUser := "root"
	dbPass := ""
	dbName := "tarea3"
	dbHost := "localhost:3006"

	/*dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")*/

	// Crear la cadena de conexión a la base de datos MySQL
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Conectar a la base de datos
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicializar el enrutador Gin
	r := gin.Default()

	// Configura CORS para permitir solicitudes desde tu dominio de React
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Reemplaza con tu URL de frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	r.Use(corsMiddleware(c))

	// Ruta para insertar datos en la base de datos
	r.POST("/insert", func(c *gin.Context) {
		// Parsear los datos del cuerpo de la solicitud
		var album struct {
			TituloAlbum     string `json:"titulo_album"`
			Artista         string `json:"artista"`
			AnioLanzamiento string `json:"anio_lanzamiento"`
			GeneroMusical   string `json:"genero_musical"`
		}
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insertar datos en la base de datos
		_, err := db.Exec("INSERT INTO Albums (TituloAlbum, Artista, AnioLanzamiento, GeneroMusical) VALUES (?, ?, ?, ?)", album.TituloAlbum, album.Artista, album.AnioLanzamiento, album.GeneroMusical)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Datos insertados correctamente"})
	})

	// Ruta para obtener todos los registros
	r.GET("/albums", func(c *gin.Context) {
		rows, err := db.Query("SELECT AlbumID, TituloAlbum, Artista, AnioLanzamiento, GeneroMusical FROM Albums")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var albums []Album
		for rows.Next() {
			var album Album
			err := rows.Scan(&album.AlbumID, &album.TituloAlbum, &album.Artista, &album.AnioLanzamiento, &album.GeneroMusical)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			albums = append(albums, album)
		}

		c.JSON(http.StatusOK, albums)
	})
	// Iniciar el servidor
	r.Run(":8080")
}

func corsMiddleware(c *cors.Cors) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	}
}

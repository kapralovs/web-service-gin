package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//структура album обозначающая способ хранения данных об альбоме
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//слайс структур album для того чтобы заполнить данные об альбомах
var albums = []album{
	{ID: "1", Title: "Load", Artist: "Metallica", Price: 36.99},
	{ID: "2", Title: "Siren Charms", Artist: "In Flames", Price: 28.59},
	{ID: "3", Title: "Fiction", Artist: "Dark Tranqullity", Price: 44.98},
}

//функция-обработчик выдающая в качестве ответа список альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//Вызываем BindJSON для того чтобы заполнить поля структуры newAlbum данными из полученного JSON-файла.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Добавляем newAlbum в слайс альбомов
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//функция getAlbumByID находит альбом,чье значение ID совпадает с указанным в запросе и возвращает этот альбом в качестве ответа.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Проходим циклом по списку альбомов и, если id из запроса совпадает с id альбома,то выводим этот альбом в формате JSON в качестве ответа.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

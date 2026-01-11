package main

//import the packages you’ll need to support the code you’ve just written.
import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

/*
func main() is the entry point of the program
when running a Go Application, execution always starts from the main() function
it must be inside a package named main
*/
func main() {
    // Initialize a Gin router using Default.
    router := gin.Default()
    // Use the GET function to associate the GET HTTP method and /albums path with a handler function.
    router.GET("/albums", getAlbums)
    // Associate the /albums/:id path with the getAlbumByID function
    router.GET("/albums/:id", getAlbumByID)
    // Associate the POST method at the /albums path with the postAlbums function.
    router.POST("/albums", postAlbums)
    // Use the Run function to attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

/* postAlbums adds an album from JSON received in the request body.
Use Context.BindJSON to bind the request body to newAlbum.
Append the album struct initialized from the JSON to the albums slice.
Add a 201 status code to the response, along with JSON representing the album you added.
*/
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

/* getAlbumByID locates the album whose ID value matches the id
parameter sent by the client, then returns that album as a response.
Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
/Loop over the album structs in the slice, looking for one whose ID field value matches the id parameter value. If it’s found, you serialize that album struct to JSON and return it as a response with a 200 OK HTTP code.
As mentioned above, a real-world service would likely use a database query to perform this lookup.
*/

//Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

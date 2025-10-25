package main

import (
	"fmt"
	"os"
	"strconv"
	"uma/echo-test/file"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	storage_go "github.com/supabase-community/storage-go"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

var (
	bucketId  string = "images"
	projectId string = "jaurcadmqgccmhzllhxr"
	apiKey    string = "secret"
	rawUrl    string = fmt.Sprintf("https://%s.supabase.co/storage/v1", projectId)
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.BodyLimit("10M"))

	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})
	e.POST("/upload", uploadFile)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})
	e.GET("/test", test)
	e.GET("/test-upload", testUpload)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.POST("/upload", uploadFile)
func uploadFile(c echo.Context) error {
	return file.UploadHandler(c)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.GET("/test", test)
func test(c echo.Context) error {
	storageClient := storage_go.NewClient(fmt.Sprintf("https://%s.supabase.co/storage/v1", projectId), apiKey, nil)
	result, err := storageClient.GetBucket(bucketId)
	if err != nil {
		fmt.Println("error")
		errJson := map[string]string{
			"statusCode":   strconv.Itoa(http.StatusInternalServerError),
			"errorMessage": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errJson)
	}
	fmt.Println(result)

	files, err := storageClient.ListFiles(bucketId, "", storage_go.FileSearchOptions{
		Limit:  10,
		Offset: 0,
		SortByOptions: storage_go.SortBy{
			Column: "",
			Order:  "",
		},
	})
	fmt.Println(files, err)

	return c.String(http.StatusOK, "test")
}

// e.GET("/test-upload", testUpload)
func testUpload(c echo.Context) error {
	storageClient := storage_go.NewClient(rawUrl, apiKey, nil)
	filename := "test.jpeg"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error")
		errJson := map[string]string{
			"statusCode":   strconv.Itoa(http.StatusInternalServerError),
			"errorMessage": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errJson)
	}

	fmt.Println(file.Name())
	fileType := "image/jpeg"
	fileOptions := storage_go.FileOptions{ContentType: &fileType}
	resp, err := storageClient.UploadFile(bucketId, filename, file, fileOptions)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("error")
		errJson := map[string]string{
			"statusCode":   strconv.Itoa(http.StatusInternalServerError),
			"errorMessage": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errJson)
	}
	return c.String(http.StatusOK, "test-upload")
}

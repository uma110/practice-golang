package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"

	storage_go "github.com/supabase-community/storage-go"
)

var (
	bucketId  string = "images"
	projectId string = "jaurcadmqgccmhzllhxr"
	apiKey    string = "secret"
	rawUrl    string = fmt.Sprintf("https://%s.supabase.co/storage/v1", projectId)
)

const MaxUploadSize = 1024 * 1024

func UploadHandler(c echo.Context) error {
	fmt.Printf("Content Length : %d\n", c.Request().ContentLength)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("error")
		errJson := map[string]string{
			"statusCode":   strconv.Itoa(http.StatusInternalServerError),
			"errorMessage": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errJson)
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fmt.Println("tmp", file.Filename)
	fmt.Println(src)

	// 存在していなければ、保存用のディレクトリを作成します。
	err = os.MkdirAll("./uploadimages", os.ModePerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// ファイルのmimetypeの取得
	// ファイルを512バイト分のみ読み込み
	buff := make([]byte, 512)
	_, err = src.Read(buff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 読み込んだバッファからmimetypeを推定
	filetype := http.DetectContentType(buff)
	fmt.Println(filetype)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return c.JSON(http.StatusInternalServerError, "許可されていないファイルタイプです。JPEGかPNGをアップロードしてください")
	}
	// 読み込んだバッファ分を戻す
	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// // 保存用ディレクトリ内に新しいファイルを作成します。
	// dst, err := os.Create(fmt.Sprintf("./uploadimages/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename)))
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// defer dst.Close()

	// // アップロードされたファイルを先程作ったファイルにコピーします。
	// _, err = io.Copy(dst, src)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	storageClient := storage_go.NewClient(rawUrl, apiKey, nil)

	resp, err := storageClient.UploadFile(bucketId, file.Filename, src)
	fmt.Println(resp)
	if err != nil {
		errJson := map[string]string{
			"statusCode":   strconv.Itoa(http.StatusInternalServerError),
			"errorMessage": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errJson)
	}

	fmt.Println("Upload successful")
	// return c.String(http.StatusOK, "uploaded")
	return c.Redirect(http.StatusFound, "/")
}

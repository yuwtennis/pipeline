package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"pipelines/internal/clients"
	"pipelines/internal/helpers"
)

// Run is the orchestrator of this application
func Run() {
	var decodedSlice []string

	downloadedFile := "/tmp/file.zip"
	from := "20204"
	to := "20214"
	prefecture := "13"
	city := "13999"

	requestUrl := fmt.Sprintf(
		"https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?DLF=true&TTC-From=%s&TTC-To=%s&TDK=%s&SKC=%s",
		from,
		to,
		prefecture,
		city)

	requestHeader := make(map[string]string)
	requestHeader["referer"] = fmt.Sprintf(
		"https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?TDK=%s&SKC=%s&TDIDFrom=%s&TDIDTo=%s",
		prefecture,
		city,
		from,
		to)

	// TODO Use channel
	contents := clients.Download(requestUrl, requestHeader)

	os.WriteFile(downloadedFile, contents, 0644)
	read, err := zip.OpenReader(downloadedFile)
	helpers.Check(err)

	for _, file := range read.File {
		rawFile, err := file.Open()
		lines, err := io.ReadAll(rawFile)
		helpers.Check(err)

		decodedString := string(helpers.ShiftJisToUTF8(lines))
		decodedSlice = append(decodedSlice, decodedString)
		log.Println(decodedString)
	}

	err = os.Remove(downloadedFile)
	helpers.Check(err)
}

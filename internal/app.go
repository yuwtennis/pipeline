package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"pipelines/internal/clients"
)

// Run is the orchestrator of this application
func Run() {
	downloadedFile := "/tmp/file.zip"
	var decodedSlice []string

	requestUrl := fmt.Sprintf(
		"https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?DLF=true&TTC-From=%s&TTC-To=%s&TDK=13&SKC=13115", "20204", "20214")
	requestHeader := make(map[string]string)
	requestHeader["referer"] = fmt.Sprintf("https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?TDK=13&SKC=13115&TDIDFrom=%s&TDIDTo=%s",
		"20204", "20214")

	// TODO Use channel
	contents := clients.Download(requestUrl, requestHeader)

	os.WriteFile(downloadedFile, contents, 0644)
	read, err := zip.OpenReader(downloadedFile)
	Check(err)

	for _, file := range read.File {
		rawFile, err := file.Open()
		lines, err := io.ReadAll(rawFile)
		Check(err)

		decodedString := string(ShiftJisToUTF8(lines))
		decodedSlice = append(decodedSlice, decodedString)
		log.Println(decodedString)
	}

	err = os.Remove(downloadedFile)
	Check(err)
}

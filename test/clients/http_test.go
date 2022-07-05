package clients

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pipelines/internal/clients"
	"testing"
)

func TestDownloadWithHeader(t *testing.T) {
	from := "20204"
	to := "20214"
	prefecture := "13"
	city := "13115"

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

	contents := clients.Download(requestUrl, requestHeader)

	assert.NotNil(t, contents, "Invalid contents.")
}

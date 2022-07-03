package clients

import (
	"io"
	"net/http"
	"pipelines/internal"
)

// Download Download contents from specified url
func Download(requestUrl string, header map[string]string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)
	internal.Check(err)

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	internal.Check(err)
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	internal.Check(err)

	return contents
}

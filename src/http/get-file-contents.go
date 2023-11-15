package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetFileContents(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return body
}

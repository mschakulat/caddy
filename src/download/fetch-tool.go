package download

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path"
)

func FetchTool(downloadLink string, target string, description string) string {
	req, err := http.NewRequest("GET", downloadLink, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		fmt.Println(aurora.Yellow("Could not download binary. Please check your version constraint."))
		fmt.Println("This usually happens when an incorrect version is specified in your package.json file.")
		os.Exit(0)
	}
	defer check(resp.Body.Close)

	filePath := fmt.Sprintf("%s/%s", target, path.Base(downloadLink))
	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			},
		),
	)

	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	println()

	return filePath
}

func Description(tool string, version string) string {
	return fmt.Sprintf("[cyan][bold]Fetching[reset] %s@%s", tool, version)
}

func check(f func() error) {
	if err := f(); err != nil {
		fmt.Fprintf(os.Stderr, "received error: %v\n", err)
	}
}

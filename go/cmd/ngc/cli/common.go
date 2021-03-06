package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

var coreAssets = []string{
	"starmap.json",
	"details.json",
	"starmap.jpg",
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func downloadProgress(filepath string, url string, quiet bool) error {
	parts := strings.Split(url, "/")
	name := url
	if len(parts) > 0 {
		name = parts[len(parts)-1]
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	if !quiet {
		s.FinalMSG = " "
		s.Suffix = fmt.Sprintf("\t%s  ", name)
		s.Color("green")
		s.Start()
	}
	err := downloadFile(filepath, url)
	if !quiet {
		s.Stop()
		fmt.Println()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
	return err
}

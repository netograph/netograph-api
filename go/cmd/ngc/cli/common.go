package cli

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var coreAssets = []string{
	"starmap.json",
	"contents.json",
	"details.json",
	"starmap.jpg",
}

func resolveIP(ip string) (string, error) {
	ips, err := net.LookupIP(ip)
	if err != nil {
		return "", err
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String(), err
		}
	}
	return "", fmt.Errorf("Could not resolve address: %s", ip)
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
	defer s.Stop()
	if !quiet {
		s.Suffix = fmt.Sprintf("\t%s  ", name)
		s.Color("green")
		s.Start()
	}
	return downloadFile(filepath, url)
}

var timeFormats = []string{
	time.RFC3339,
	"2006-01-02",
	"2006/01/02",
}

func parseTime(s string) (*timestamp.Timestamp, error) {
	if s == "" {
		return &timestamp.Timestamp{}, nil
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	var t time.Time
	for _, f := range timeFormats {
		t, err = time.ParseInLocation(f, s, loc)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("could not parse date specification: %s", s)
	}
	return ptypes.TimestampProto(t)
}

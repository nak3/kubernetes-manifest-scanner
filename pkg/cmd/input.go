package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

// Forked from: https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/util/helpers.go
// Add insecure option
func ReadConfigDataFromLocation(location string, insecure bool) ([]byte, error) {
	// we look for http:// or https:// to determine if valid URL, otherwise do normal file IO
	if strings.Index(location, "http://") == 0 || strings.Index(location, "https://") == 0 {
		var resp *http.Response
		var err error
		if insecure == true {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			resp, err = client.Get(location)
		} else {
			resp, err = http.Get(location)
		}
		if err != nil {
			return nil, fmt.Errorf("unable to access URL %s: %v\n", location, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("unable to read URL, server reported %d %s", resp.StatusCode, resp.Status)
		}
		return cmdutil.ReadConfigDataFromReader(resp.Body, location)
	} else {
		file, err := os.Open(location)
		if err != nil {
			return nil, fmt.Errorf("unable to read %s: %v\n", location, err)
		}
		return cmdutil.ReadConfigDataFromReader(file, location)
	}
}

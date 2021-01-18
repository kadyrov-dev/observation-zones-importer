package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/kadyrov-dev/observation-zones-importer/internal/authorizer"
	"github.com/kadyrov-dev/observation-zones-importer/internal/exporter"
	"github.com/kadyrov-dev/observation-zones-importer/internal/importer"

	"github.com/labstack/gommon/log"
)

func main() {
	login := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")
	outputDir := os.Getenv("OUTPUT_DIR")

	if err := checkEnvironmentVariables(login, password, outputDir); err != nil {
		log.Fatalf("CheckEnvironmentVariables: %v", err)
	}

	httpClient := createHTTPClient()

	// Authorizer
	auth := authorizer.NewAuthorizer(httpClient)
	accessToken, err := auth.GetAccessToken(login, password)

	if err != nil {
		log.Fatalf("Authorizer: %s", err.Error())
	}

	// Importer
	importer := importer.NewImporter(accessToken, httpClient)
	zones, err := importer.Import()

	if err != nil {
		log.Fatalf("Importer: %v", err)
	}

	// Exporter
	exporter := exporter.NewExporter(outputDir)

	if err := exporter.Export(zones); err != nil {
		log.Fatalf("Exporter: %v", err)
	}
}

func checkEnvironmentVariables(login, password, outputDir string) error {
	if login == "" || password == "" {
		return errors.New("you must set LOGIN and PASSWORD environment variables")
	}

	if outputDir == "" {
		return errors.New("you must set OUTPUT_DIR environment variable")
	}

	if err := checkDirectoryPermissions(outputDir); err != nil {
		return fmt.Errorf("checkDirectoryPermissions(): %v", err)
	}

	return nil
}

func checkDirectoryPermissions(outputDir string) error {
	stat, err := os.Stat(outputDir)

	if err != nil {
		return fmt.Errorf("directory %s does not exist", outputDir)
	}

	if !stat.IsDir() {
		return fmt.Errorf("path %s is not a directory", outputDir)
	}

	if stat.Mode().Perm()&(1<<(uint(7))) == 0 {
		return fmt.Errorf("write permission bit is not set on %s directory for user", outputDir)
	}

	return nil
}

func createHTTPClient() *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

package bot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func downloadImage(fileURL string) (bytes.Buffer, error) {
	res, err := http.Get(fileURL) //nolint:gosec
	if err != nil {
		log.Fatalf("http.Get: %v", err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Fatalf("res.Body.Close: %v", err)
		}
	}()

	imageBuffer := bytes.Buffer{}
	_, err = io.Copy(&imageBuffer, res.Body)
	if err != nil {
		log.Fatalf("io.Copy: %v", err)
	}

	return imageBuffer, nil
}

func zbarRun(input bytes.Buffer) (string, error) {
	zbarimgPath, err := exec.LookPath("zbarimg")
	if err != nil {
		log.Printf("exec.LookPath: %v", err)
		return "", err
	}

	args := []string{
		"-S*.disable",
		"-Sqrcode.enable",
		"-",
	}

	zbarimgCmd := exec.Command(zbarimgPath, args...)

	zbarimgCmd.Stdin = &input

	output, err := zbarimgCmd.CombinedOutput()
	if err != nil {
		return "barcode data was not detected", err
	}

	outputString := fmt.Sprint(string(output), err)

	outputSlice := strings.Split(outputString, "\n")
	result := ""
	for i, line := range outputSlice {
		if strings.Contains(line, "QR-Code") {
			line = strings.Replace(line, "QR-Code:", "", 1)
			line = strings.TrimSpace(line)
			result += fmt.Sprintf("%v: %v\n", i+1, line)
		} else if strings.Contains(line, "scanned") {
			line = strings.TrimSpace(line)
			log.Printf("zbarReport: %v", line)
		}
	}

	return result, nil
}

func processImage(fileURL string) (string, error) {
	imageBuffer, err := downloadImage(fileURL)
	if err != nil {
		log.Printf("downloadImage: %v", err)
		return "", err
	}

	result, err := zbarRun(imageBuffer)
	if err != nil {
		return result, err
	}

	return result, nil
}

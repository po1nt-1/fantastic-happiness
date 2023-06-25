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
		log.Println(err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	imageBuffer := bytes.Buffer{}
	_, err = io.Copy(&imageBuffer, res.Body)
	if err != nil {
		log.Println(err)
	}

	return imageBuffer, nil
}

func zbarRun(input bytes.Buffer) (string, error) {
	zbarimgPath, err := exec.LookPath("zbarimg")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	outputString := fmt.Sprint(string(output), err)
	outputString = strings.Replace(outputString, "QR-Code:", "", 1)
	outputSlice := strings.Split(outputString, "\n")

	result := ""
	for _, line := range outputSlice {
		if strings.Contains(line, "\n") {
			result += line + "\n"
		}
	}

	return result, nil
}

func processImage(fileURL string) (string, error) {
	imageBuffer, err := downloadImage(fileURL)
	if err != nil {
		return "", err
	}

	result, err := zbarRun(imageBuffer)
	if err != nil {
		log.Println(err)
	}

	return result, nil
}

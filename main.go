package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const version = "0.1.0"
const baseURL = "http://control.d-imaging.sony.co.jp/GPS/"
const datFile = "assistme.dat"
const subDir = "PRIVATE/SONY/GPS"

func main() {
	versionFlag := flag.Bool("version", false, "Print version information.")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" <path to SD card>")
		fmt.Fprintln(os.Stderr, "example: "+os.Args[0]+" /Volumes/Untitled")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Fprintf(os.Stdout, "assistme %s\n", version)
		os.Exit(0)
	}

	if len(os.Args) != 2 {
		exitErrorf("Usage: " + os.Args[0] + " <path to SD card>")
	}
	mountPoint := os.Args[1]
	if !exists(mountPoint) {
		exitErrorf("SD card not found")
	}

	createSubDir(getGPSFolderPath(mountPoint))

	// Create the file
	out, err := os.Create(getGPSFolderPath(mountPoint) + "/" + datFile)
	if err != nil {
		exitErrorf(err.Error())
	}
	defer out.Close()

	// Get the GPS data
	resp, err := http.Get(baseURL + datFile)
	if err != nil {
		exitErrorf(err.Error())
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		exitErrorf(err.Error())
	}

	fmt.Fprintln(os.Stdout, "GPS datafile updated")

}

func createSubDir(path string) {
	if !exists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getGPSFolderPath(mountPoint string) string {
	return mountPoint + "/" + subDir
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

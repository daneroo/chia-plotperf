package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"regexp"
	"time"
)

func main() {
	SetupLoggingFormat()
	log.Printf("Extracting Chia Plotting Performance From `.plot`s\n")
	directory := "/Volumes/ChiaPlots/"
	err := getPlots(directory)
	if err != nil {
		log.Fatal(err)
	}
}

func getPlots(dir string) error {
	//Important Note: ReadDir guarantees order by filename
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	count := 0
	// map each os.FileInfo to a full path
	for _, file := range files {
		// log.Printf("Considering %s\n", file.Name())
		count += onePlot(file)
	}
	log.Printf("Examined %d `.plot`s\n", count)

	return nil
}

func onePlot(file fs.FileInfo) int {
	re := regexp.MustCompile(`plot-k32-(?P<YMDHM>\d{4}-\d{2}-\d{2}-\d{2}-\d{2})-[[:xdigit:]]{64}\.plot`)

	matches := re.FindStringSubmatch(file.Name())
	if matches == nil {
		return 0
	}
	YMDHM := matches[re.SubexpIndex("YMDHM")]
	start, err := time.ParseInLocation("2006-01-02-15-04", YMDHM, time.Local)
	if err != nil {
		fmt.Println(err)
		return 0

	}
	end := file.ModTime()
	elapsed := end.Sub(start)

	fmt.Printf("[ %s - %s ]: %v\n", start.Format(fmtRFC3339Local), end.Format(fmtRFC3339Local), elapsed)
	return 1
}

// logging setup
const (
	fmtRFC3339Local  = "2006-01-02T15:04:05"
	fmtRFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(fmtRFC3339Millis) + " - " + string(bytes))
}

// SetupFormat initializes logging
func SetupLoggingFormat() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

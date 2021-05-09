package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"time"
)

func main() {
	SetupLoggingFormat()
	log.Printf("Starting Chia Plotting Performance\n")
	directory := "/Volumes/ChiaPlots/"
	plots, err := getPlots(directory)
	if err != nil {
		log.Fatal(err)
	}
	if len(plots) == 0 {
		log.Printf("%v\n", plots)

	}
}

func getPlots(dir string) ([]string, error) {
	//Important Note: ReadDir guarantees order by filename
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// map each os.FileInfo to a full path
	var filenames []string // == nil
	for _, file := range files {
		filename := path.Join(dir, file.Name())
		// log.Printf("Considering %s\n", file.Name())

		// re := regexp.MustCompile(`plot-k32-(?P<Y>\d{4})-(?P<M>\d{2})-(?P<D>\d{2})-(?P<h>\d{2})-(?P<m>\d{2})-(?P<Digest>[[:xdigit:]]{64})\.plot`)
		re := regexp.MustCompile(`plot-k32-(?P<YMDHM>\d{4}-\d{2}-\d{2}-\d{2}-\d{2})-[[:xdigit:]]{64}\.plot`)

		matches := re.FindStringSubmatch(file.Name())
		if matches == nil {
			continue
		}
		// fmt.Printf("%#v\n", matches)
		// fmt.Printf("%#v\n", re.SubexpNames())
		// Y, _ := strconv.Atoi(matches[re.SubexpIndex("Y")])
		// M, _ := strconv.Atoi(matches[re.SubexpIndex("M")])
		// D, _ := strconv.Atoi(matches[re.SubexpIndex("D")])
		// h, _ := strconv.Atoi(matches[re.SubexpIndex("h")])
		// m, _ := strconv.Atoi(matches[re.SubexpIndex("m")])
		// start := time.Date(Y, M, D, h, m, 0, 0, time.Local)
		YMDHM := matches[re.SubexpIndex("YMDHM")]
		start, err := time.ParseInLocation("2006-01-02-15-04", YMDHM, time.Local)
		if err != nil {
			fmt.Println(err)
			continue
			// return nil, err
		}
		end := file.ModTime()
		elapsed := end.Sub(start)
		fmt.Printf("Start: %v End: %v Elapsed: %v\n", start, end, elapsed)

		filenames = append(filenames, filename)
	}

	return filenames, nil
}

// logging setup
const (
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

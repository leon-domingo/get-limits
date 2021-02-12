package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const (
	optionTop    string = "top"
	optionBottom string = "bottom"
	optionBoth   string = "both"
)

type stringsSlice []string
type batteryLimits [][2]int

var defaultLimits = [][2]int{
	{20, 95},
	{25, 90},
}

func (s stringsSlice) Exists(item string) bool {
	found := false
	for i := range s {
		if strings.ToLower(s[i]) == strings.ToLower(item) {
			found = true
			break
		}
	}
	return found
}

var options = stringsSlice{
	optionTop,
	optionBottom,
	optionBoth,
}

func main() {
	argsLength := len(os.Args)
	if argsLength < 2 || argsLength > 3 {
		usageAndExit()
	}

	option := strings.ToLower(os.Args[1])
	if !options.Exists(option) {
		usageAndExit()
	}

	var t time.Time
	var err error
	if argsLength > 2 {
		theDate := os.Args[2]
		t, err = time.Parse("20060102", theDate)
		if err != nil {
			usageAndExit()
		}
	} else {
		today := time.Now()
		t, err = time.Parse("20060102",
			fmt.Sprintf("%d%02d%02d",
				today.Year(),
				today.Month(),
				today.Day()))
		if err != nil {
			usageAndExit()
		}
	}

	var limits batteryLimits

	limitsPath := path.Join(os.Getenv("HOME"), ".config/get-limits/limits.json")
	limitsFile, err := os.Open(limitsPath)
	if os.IsNotExist(err) {
		limits = defaultLimits[:]
	} else if err != nil {
		log.Fatal(err)
	}
	defer limitsFile.Close()

	var readerValues = bufio.NewReader(limitsFile)
	var contenido []byte

	readerValues.Read(contenido)
	limitsJSONDecoder := json.NewDecoder(readerValues)
	limitsJSONDecoder.Decode(&limits)

	daysFromEpoch := t.Unix() / (24 * 3600)
	indexAccordingToDate := daysFromEpoch % int64(len(limits))
	bottom, top := limits[indexAccordingToDate][0], limits[indexAccordingToDate][1]

	switch option {
	case optionTop:
		fmt.Printf("%d\n", top)
	case optionBottom:
		fmt.Printf("%d\n", bottom)
	case optionBoth:
		fmt.Printf("%d %d\n", top, bottom)
	}
}

func usageAndExit() {
	log.Fatal("usage: get-limits <top|bottom|both> [<YYYYMMDD>]")
}

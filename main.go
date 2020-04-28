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
	"top",
	"bottom",
	"both",
}

func main() {

	// only executable and options is valid (2 or 3 args)
	argsLength := len(os.Args)
	if argsLength < 2 || argsLength > 3 {
		usageAndExit()
	}

	// get the option from the command-line
	option := strings.ToLower(os.Args[1])

	// is it a valid option?
	if !options.Exists(option) {
		usageAndExit()
	}

	// get the date from the command-line
	var t time.Time
	var err error
	if argsLength > 2 {
		theDate := os.Args[2]
		t, err = time.Parse("20060102", theDate)
		if err != nil {
			usageAndExit()
		}
	} else {
		// today
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

	// read the JSON file containing the limits:
	// arbitrary values ranging from 10,20,25,30 (bottom) to 90,95,98 (top)
	limitsPath := path.Join(os.Getenv("HOME"), ".config/get-limits/limitsx.json")
	limitsFile, err := os.Open(limitsPath)

	if os.IsNotExist(err) {
		// log.Printf("WARN: File \"%s\" does not exist\n", limitsPath)
		limits = defaultLimits[:]
	} else if err != nil {
		log.Fatal(err)
	}
	defer limitsFile.Close()

	var readerValues = bufio.NewReader(limitsFile)
	var contenido []byte

	readerValues.Read(contenido)
	dec := json.NewDecoder(readerValues)

	dec.Decode(&limits)

	// days from epoch
	daysFromEpoch := t.Unix() / (24 * 3600)

	i := daysFromEpoch % int64(len(limits))
	bottom, top := limits[i][0], limits[i][1]

	switch option {
	case "top":
		fmt.Printf("%d\n", top)
	case "bottom":
		fmt.Printf("%d\n", bottom)
	case "both":
		fmt.Printf("%d %d\n", top, bottom)
	}
}

// usageAndExit shows the usage of the command and exit the program with exit status = 1
func usageAndExit() {
	log.Fatal("usage: get-limits <top|bottom|both> [<YYYYMMDD>]")
}

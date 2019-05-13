package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type stringsSlice []string

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

var options = stringsSlice{"top", "bottom", "both"}

// arbitrary values ranging from 10,20,25,30 (bottom) to 90,95,98 (top)
var limits = [][2]int{
	{20, 95},
	{25, 90},
	{25, 95},
	{20, 98},
	{30, 95},
	{25, 90},
	{25, 95},
	{10, 98},
	{25, 95},
	{25, 90},
	{30, 95},
	{20, 90},
	{25, 95},
	{20, 90},
}

func main() {

	// only executable and option is valid (2 args)
	if len(os.Args) != 2 {
		usageAndExit()
	}

	// get the option from the command-line
	option := strings.ToLower(os.Args[1])

	// is it a valid option?
	if !options.Exists(option) {
		usageAndExit()
	}

	// days from epoch
	daysFromEpoch := time.Now().Unix() / 24 * 3600
	i := int(daysFromEpoch) % len(limits)

	switch option {
	case "top":
		fmt.Printf("%d\n", limits[i][1])
	case "bottom":
		fmt.Printf("%d\n", limits[i][0])
	case "both":
		fmt.Printf("%d %d\n", limits[i][0], limits[i][1])
	}
}

// usageAndExit shows the usage of the command and exit the program with exit status = 1
func usageAndExit() {
	log.Fatal("usage: get-limits [top|bottom|both]")
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type stringsSlice []string

func (s stringsSlice) exists(item string) bool {
	found := false
	for i := range s {
		if strings.ToLower(s[i]) == strings.ToLower(item) {
			found = true
			break
		}
	}

	return found
}

var opciones = stringsSlice{"top", "bottom", "both"}

var limites = [][2]int{
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

	if len(os.Args) != 2 {
		usageAndExit()
	}

	opcion := strings.ToLower(os.Args[1])

	if !opciones.exists(opcion) {
		usageAndExit()
	}

	// días desde EPOCH
	dias := time.Now().Unix() / 24 * 3600
	i := int(dias) % len(limites)

	switch opcion {
	case "top":
		fmt.Printf("%d\n", limites[i][1])
	case "bottom":
		fmt.Printf("%d\n", limites[i][0])
	case "both":
		fmt.Printf("%d %d\n", limites[i][0], limites[i][1])
	}
}

// usageAndExit shows the usage of the command and exit the program with exit status = 1
func usageAndExit() {
	log.Fatal("usage: get-limits [top|bottom|both]")
}

package main

import (
	"gifs/may"
	"log"
	"os"
	"strconv"
)

var funcs []func()

func main() {
	funcs = append(funcs,
		may.May01,
		may.May02,
		may.May03,
		may.May04,
		may.May05,
		may.May06,
		may.May07,
		may.May08,
		may.May09,
		may.May10,
		may.May11,
		may.May12,
		may.May13,
		may.May14,
		may.May15,
		may.May16,
		may.May17,
		may.May18,
		may.May19,
		may.May20,
		may.May21,
		may.May22,
		may.May23,
		may.May24,
		may.May25,
		may.May26,
		may.May27,
		may.May28,
		may.May29,
		may.May30,
	)
	arg, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	funcs[arg-1]()

}

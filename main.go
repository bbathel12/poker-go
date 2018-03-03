package main

import (
	//_ box "github.com/nsf/termbox-go"
	"log"
	"os"
	"runtime/pprof"
	_ "time"
)

func main() {

	f, err := os.Create("./cpu")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	Run()
	mf, err := os.Create("./mem")
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(mf)
	mf.Close()
}

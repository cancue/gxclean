package main

import (
	"flag"
	"log"
	"os"

	"github.com/cancue/gxclean/cleaner"
)

func main() {
	name := flag.String("n", "", "name to delete")
	d := flag.Bool("d", false, "only directory")
	f := flag.Bool("f", false, "only file")
	flag.Parse()

	cfg, err := cleaner.NewConfig(*name, *d, *f)
	if err != nil {
		log.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	err = cleaner.FindAndDeleteAll(cfg)
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

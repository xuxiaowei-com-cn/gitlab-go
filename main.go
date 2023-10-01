package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v0.0.1-snapshot" // 版本号

func main() {
	app := &cli.App{
		Name:    "boom",
		Version: version,
		Usage:   "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

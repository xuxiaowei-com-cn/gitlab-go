package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	BuildDate         string  // 2023-07-19T12:20:54Z
	GitCommitSha      string  // fa3d7990104d7c1f16943a67f11b154b71f6a132
	GitCommitShortSha string  // fa3d7990
	GitVersion        = "dev" // v1.27.4
)

func main() {
	fmt.Printf("BuildDate: %s\n", BuildDate)
	fmt.Printf("GitCommitSha: %s\n", GitCommitSha)
	fmt.Printf("GitCommitShortSha: %s\n", GitCommitShortSha)
	fmt.Printf("GitVersion: %s\n", GitVersion)
	app := &cli.App{
		Name:    "boom",
		Version: GitVersion,
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

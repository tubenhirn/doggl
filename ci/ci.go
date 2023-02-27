package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tubenhirn/dagger-ci-modules/v2/goreleaser"
)

func main() {
	dir, _ := os.Getwd()

	version, err := ioutil.ReadFile(dir + "/version")
	if err != nil {
		panic(err)
	}

	options := goreleaser.GoReleaserOpts{
		Source:     dir,
		Snapshot:   true,
		RemoveDist: true,
		Env: map[string]string{
			"APP_VERSION": string(version),
		},
	}

	if err := goreleaser.Release(context.Background(), options); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/tubenhirn/dagger-ci-modules/v2/goreleaser"
)

func main() {
	snapshot := flag.Bool("snapshot", true, "the string of the platform to run renovate on.")
	flag.Parse()

	fmt.Println("snapshot build " + strconv.FormatBool(*snapshot))


	dir, _ := os.Getwd()

	version, err := ioutil.ReadFile(dir + "/version")
	if err != nil {
		panic(err)
	}

	options := goreleaser.GoReleaserOpts{
		Source:     dir,
		Snapshot:   *snapshot,
		RemoveDist: true,
		Env: map[string]string{
			"APP_VERSION": string(version),
		},
		Secret: []string{"GITHUB_TOKEN"},
	}

	if err := goreleaser.Release(context.Background(), options); err != nil {
		fmt.Println(err)
	}
}

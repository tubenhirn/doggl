package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tubenhirn/dagger-ci-modules/v2/goreleaser"
)

func main() {
	snapshot := flag.Bool("snapshot", true, "the string of the platform to run renovate on.")
	clean := flag.Bool("clean", true, "the string of the platform to run renovate on.")
	flag.Parse()

	fmt.Println("running with flags:", "\nsnapshot", *snapshot, "\nremovedist", *clean)

	dir, _ := os.Getwd()

	version, err := ioutil.ReadFile(dir + "/version")
	if err != nil {
		panic(err)
	}

	options := goreleaser.GoReleaserOpts{
		Source:     dir,
		Snapshot:   *snapshot,
		RemoveDist: *clean,
		Env: map[string]string{
			"APP_VERSION": string(version),
		},
		Secret: []string{"GITHUB_TOKEN"},
	}

	if err := goreleaser.Release(context.Background(), options); err != nil {
		fmt.Println(err)
	}
}

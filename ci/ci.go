package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"dagger.io/dagger"
	"github.com/tubenhirn/dagger-ci-modules/v2/goreleaser"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}

	defer client.Close()

	task := flag.String("task", "", "the task to execute.")
	if *task == "release" {
		snapshot := flag.Bool("snapshot", true, "the string of the platform to run renovate on.")
		clean := flag.Bool("clean", true, "the string of the platform to run renovate on.")
		flag.Parse()

		fmt.Println("running with flags:", "\nsnapshot", *snapshot, "\nremovedist", *clean)

		var secrets = make(map[string]dagger.SecretID)
		githubTokenId, err := client.Host().EnvVariable("GITHUB_TOKEN").Secret().ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId

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
			Secret: secrets,
		}

		if err := goreleaser.Release(context.Background(), *client, options); err != nil {
			fmt.Println(err)
		}
	} else if *task == "tag" {
		var secrets = make(map[string]dagger.SecretID)
		githubTokenId, err := client.Host().EnvVariable("GITHUB_TOKEN").Secret().ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId

		// dir, _ := os.Getwd()

		options :=
	}
}

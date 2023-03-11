package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"dagger.io/dagger"
	"github.com/tubenhirn/dagger-ci-modules/goreleaser"
	"github.com/tubenhirn/dagger-ci-modules/semanticrelease"
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
	snapshot := flag.Bool("snapshot", true, "do a snapshot build.")
	clean := flag.Bool("clean", true, "clean dist directory.")

	flag.Parse()
	if *task == "release" {
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

		dir, _ := os.Getwd()

		options := semanticrelease.SemanticOpts{
			Source:   dir,
			// use "git" for tag only
			// release is done with goreleaser
			Platform: "git",
			Env:      map[string]string{},
			Secret:   secrets,
		}

		if err := semanticrelease.Semanticrelease(ctx, *client, options); err != nil {
			fmt.Println(err)
		}
	}
}

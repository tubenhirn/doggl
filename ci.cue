package ci

import (
	"dagger.io/dagger"

	"universe.dagger.io/alpha/go/golangci"
	"github.com/tubenhirn/dagger-ci-modules/goreleaser"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents:       dagger.#FS
	client: filesystem: "./dist": write: contents: actions.build.export.directories."/src/dist"

	client: env: {
		GITHUB_TOKEN: dagger.#Secret
	}

	actions: {
		_source:  client.filesystem["."].read.contents
		build: goreleaser.#Release & {
			source:     _source
			snapshot:   true
			removeDist: true
		}

		lint: {
			go: golangci.#Lint & {
				source:  _source
				version: "1.48"
			}
		}

		release: goreleaser.#Release & {
			source:     _source
			removeDist: true
			env: {
				"GITHUB_TOKEN": client.env.GITHUB_TOKEN
			}
		}

	}
}

# Copyright © 2023 Johannes Stang <tubenhirn@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
#  *******     *******     ********    ********  **
# /**////**   **/////**   **//////**  **//////**/**
# /**    /** **     //** **      //  **      // /**
# /**    /**/**      /**/**         /**         /**
# /**    /**/**      /**/**    *****/**    *****/**
# /**    ** //**     ** //**  ////**//**  ////**/**
# /*******   //*******   //********  //******** /********
# ///////     ///////     ////////    ////////  ////////

test: ## test doggl
	@echo "testing..."
	go test ./...

build: test ## build doggl
	@echo "building..."
	go run .dagger/ci.go --task=release --snapshot=true

tag: ## tag doggl
	@echo "taggin..."
	go run .dagger/ci.go --task=tag

release: test ## release doggl
	@echo "release..."
	go run .dagger/ci.go --task=release --snapshot=false

.PHONY: help
help:  ## 🤔 Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

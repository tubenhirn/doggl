# doggl

A simple toggl cli.


## Configuration

Copy the `.doggl` file from the example folder to your homedirectory.\
Fill the missing values `TOGGL_TOKEN`, `PROJECT_ID` and `WORKSPACE_ID`.

## Usage

Simple:

``` shell
doggl book
```

Custom duration:

``` shell
doggl book 1800
```

## Build

With `goreleaser`:

``` shell
export APP_VERSION="v1.0.0"
goreleaser release --snapshot --rm-dist
```


# doggl

A simple toggl cli.\
In upper bavaria we "sometimes" pronounce the `t` like a `d`.

## installation

### homebrew

``` shell
brew tap tubenhirn/homebrew-formulae
brew install doggl
```

## configuration

Copy the `.doggl` file from the example folder to your homedirectory.\
Fill the missing values `TOGGL_TOKEN`, `PROJECT_ID` and `WORKSPACE_ID`.

## usage

Simple:

``` shell
doggl book
```

Custom duration:

``` shell
doggl book 1800
```

## build

``` shell
make build
```

## tag

``` shell
make tag
```

## release

``` shell
make release
```

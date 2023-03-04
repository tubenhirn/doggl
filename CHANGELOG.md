## [1.6.3](https://github.com/tubenhirn/doggl/compare/v1.6.2...v1.6.3) (2023-03-04)


### Bug Fixes

* fix ci parameters ([1f2a81c](https://github.com/tubenhirn/doggl/commit/1f2a81c152a73410168448d11447322d4a1afce1))

## [1.6.2](https://github.com/tubenhirn/doggl/compare/v1.6.1...v1.6.2) (2023-03-04)


### Bug Fixes

* set platform "git" ([342c706](https://github.com/tubenhirn/doggl/commit/342c706807c6c1aaf27a5161c289150e4cc4e159))

## [1.6.1](https://github.com/tubenhirn/doggl/compare/v1.6.0...v1.6.1) (2023-03-04)


### Bug Fixes

* **deps:** update github.com/tubenhirn/dagger-ci-modules/v2 digest to 2956bae ([22ceaee](https://github.com/tubenhirn/doggl/commit/22ceaee841c8a3a705daa885cc3667d110e8282d))
* **deps:** update github.com/tubenhirn/dagger-ci-modules/v2 digest to 7d6f6b7 ([b6ac15f](https://github.com/tubenhirn/doggl/commit/b6ac15f19d6c363a58dee4a27248740b34d14d21))


### Continuous Integration

* add flag for dagger go sdk release ([771ddc0](https://github.com/tubenhirn/doggl/commit/771ddc0ce7d939f59d1c2f84c7dcde7295f21acd))
* add makefile ([702fc7f](https://github.com/tubenhirn/doggl/commit/702fc7fd24716896bcfcee5345057f35a5dac95c))
* add more flags to configure goreleaser from command line ([6c19b4b](https://github.com/tubenhirn/doggl/commit/6c19b4bcaaf83f2e5c995088316f36dbf6c52718))
* rework ci steps ([8b9170d](https://github.com/tubenhirn/doggl/commit/8b9170de168fcf16b6a99bcde836c9e59f0d4eec))
* rework ci with makefile ([02bce7a](https://github.com/tubenhirn/doggl/commit/02bce7a91bd166f91d25d4a403471ad328e64991))
* update ci pipeline ([3d18266](https://github.com/tubenhirn/doggl/commit/3d182665278cc167d062faaa52a61c548c09846b))

# [1.6.0](https://github.com/tubenhirn/doggl/compare/v1.5.0...v1.6.0) (2023-03-02)


### Bug Fixes

* **deps:** update github.com/tubenhirn/dagger-ci-modules/v2 digest to 262a10f ([d40f337](https://github.com/tubenhirn/doggl/commit/d40f33713ba08356cb127771cd80ff52a4bdc557))
* **deps:** update github.com/tubenhirn/dagger-ci-modules/v2 digest to 5470090 ([7474179](https://github.com/tubenhirn/doggl/commit/747417910a2f76850bfed86c9ac7bf0bf8625574))
* **deps:** update github.com/tubenhirn/dagger-ci-modules/v2 digest to c660d48 ([b6f3a36](https://github.com/tubenhirn/doggl/commit/b6f3a367695dbd48e37d73d70cb2c62ff5c65387))


### Code Refactoring

* remove dagger dps from doggl ([08fb06a](https://github.com/tubenhirn/doggl/commit/08fb06a06cfdf7b6905d09b11cebb9534fb0651a))


### Continuous Integration

* rebuild dagger-cue ci with go sdk ([5bff365](https://github.com/tubenhirn/doggl/commit/5bff3655e259bf169377c8bc23e3929f1b05df27))


### Features

* add some more parameters for book command ([4ec4f17](https://github.com/tubenhirn/doggl/commit/4ec4f17a6eb2ba9000bbd2913fe08ee1dfa71a40))

# [1.5.0](https://github.com/tubenhirn/doggl/compare/v1.4.0...v1.5.0) (2023-02-25)


### Continuous Integration

* update dagger-ci-modules version ([da2cf1a](https://github.com/tubenhirn/doggl/commit/da2cf1a650f946241003766ea3ca193a702aefb2))


### Features

* add api token cli flag ([efc05a6](https://github.com/tubenhirn/doggl/commit/efc05a6a5d898c99caeb951f759cf07bb2b6d5cb))

# [1.4.0](https://github.com/tubenhirn/doggl/compare/v1.3.2...v1.4.0) (2023-02-23)


### Features

* add better error handling for failed requests ([8683022](https://github.com/tubenhirn/doggl/commit/86830220e8e695c0dc2d979ef4f42fcd229a0515))


### Miscellaneous Chores

* **deps:** update module go to 1.20 ([a995a34](https://github.com/tubenhirn/doggl/commit/a995a347509c18526f56da7447fc7a111fd4f693))

## [1.3.2](https://github.com/tubenhirn/doggl/compare/v1.3.1...v1.3.2) (2023-01-29)


### Miscellaneous Chores

* update semantic-release version to v2.10.1 ([632805e](https://github.com/tubenhirn/doggl/commit/632805e2b70c28655cf2cfc7dc1b3eb101838fe0))

## [1.3.1](https://github.com/tubenhirn/doggl/compare/v1.3.0...v1.3.1) (2023-01-29)


### Bug Fixes

* **deps:** update module github.com/spf13/viper to v1.15.0 ([b830eb4](https://github.com/tubenhirn/doggl/commit/b830eb40f738e8d1fb7b6ad28c3f7a859faa8049))
* make keys const ([e99c001](https://github.com/tubenhirn/doggl/commit/e99c001e974024b08fd009c89a49bd6023ac86b8))


### Code Refactoring

* add some docs and error handling ([8c50b6a](https://github.com/tubenhirn/doggl/commit/8c50b6a02fb35a1b7c6f5a917455266f3c0b7b08))


### Continuous Integration

* add proper brew description ([08ad29a](https://github.com/tubenhirn/doggl/commit/08ad29a74ad1bef5149d1fe23ce3bb2417779a90))
* add renovate config ([62ed1a1](https://github.com/tubenhirn/doggl/commit/62ed1a1ac230119d06862ffc59bb3b213080d4de))
* add semantic release job ([5b5d8e5](https://github.com/tubenhirn/doggl/commit/5b5d8e544a4fd4f1071a0fd10e78ab08210cd01a))


### Miscellaneous Chores

* add release config ([50663d0](https://github.com/tubenhirn/doggl/commit/50663d090f896a2c66720cb579358f0cb0408b52))
* update dagger-ci modules ([4e6b5cf](https://github.com/tubenhirn/doggl/commit/4e6b5cf6ebdbf7fada08298e7b2dbfc79c7ccf56))
* update dagger-ci-modules to v1.4.0 ([760661a](https://github.com/tubenhirn/doggl/commit/760661a8b3cf6cc2d7fd2bafa12bb69bb6a65d22))
* update semanticRelease version ([c6990dd](https://github.com/tubenhirn/doggl/commit/c6990dd2c16073fc939ebb661990840c7c5be726))

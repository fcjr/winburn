<p align="center">
<img src="img/logo.svg" width="125" alt="winburn logo">
</p>

# winburn

winburn is a cross-platform cli for creating bootable windows usbs.

[![GitHub release (latest by date)][release-img]][release]
[![GolangCI][golangci-lint-img]][golangci-lint]
[![Go Report Card][report-card-img]][report-card]

## Development

### Building the current commit

This project uses [goreleaser](https://github.com/goreleaser/goreleaser/).

 1) Install [go](https://golang.org/doc/install).
 2) Install [snapcraft](https://snapcraft.io/docs/installing-snapcraft).
 3) Install goreleaser via the steps [here](https://goreleaser.com/install/).
 4) Build current commit via `goreleaser release --snapshot --skip-publish --rm-dist`.

[release-img]: https://img.shields.io/github/v/release/fcjr/winburn
[release]: https://github.com/fcjr/winburn/releases
[golangci-lint-img]: https://github.com/fcjr/winburn/workflows/lint/badge.svg
[golangci-lint]: https://github.com/fcjr/winburn/actions?query=workflow%3Alint
[report-card-img]: https://goreportcard.com/badge/github.com/fcjr/winburn
[report-card]: https://goreportcard.com/report/github.com/fcjr/winburn
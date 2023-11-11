<p align="center">
  <a href="https://github.com/kevincobain2000/go-hexe">
    <img alt="gobrew" src="https://imgur.com/pCMsDyq.png" width="360">
  </a>
</p>
<p align="center">
  Static sites to executable binaries converter.
  <br>
  Portable binaries for any systems.
  <br>
  <br>
  Convert sites exported from Next.js, Gatsby, Hugo, Jekyll, Sveltekit
  <br>
  to executable binary.
</p>

**Quick Setup:** Easy setup.

**Hassle Free:** Portable binaries to run anywhere on any system.

**Platform:** Supports (arm64, arch64, Mac, Mac M1, Ubuntu and Windows).

**Graceful Reloads:** Reloads server gracefully without dropping any connections.

**Supports:** Sites exported using Next.js, Gatsby, Hugo, Jekyll, Hexo, Sveltekit and etc.

**Supports:** HTML, CSS, JS, Images, Fonts, SVGs, JSON, XML, Markdown and etc.

**Supports:** Base Path, Custom Port.

# Build Status

| Branch          | Status                                                                                                                                                    |
| :-------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------- |
| master          | ![CI](https://github.com/kevincobain2000/go-hexe/actions/workflows/build.yml/badge.svg)                                                                   |
| Coverage        | [![coveritup](https://coveritup.app/embed/kevincobain2000/go-hexe?branch=master&type=coverage)](https://coveritup.app/kevincobain2000/go-hexe)            |
| Binary Size     | [![coveritup](https://coveritup.app/embed/kevincobain2000/go-hexe?branch=master&type=go-binary-size)](https://coveritup.app/kevincobain2000/go-hexe)      |
| Mod Dependecies | [![coveritup](https://coveritup.app/embed/kevincobain2000/go-hexe?branch=master&type=go-mod-dependencies)](https://coveritup.app/kevincobain2000/go-hexe) |
| Unit Test Time  | [![coveritup](https://coveritup.app/embed/kevincobain2000/go-hexe?branch=master&type=unit-test-run-time)](https://coveritup.app/kevincobain2000/go-hexe)  |
| Average PR Days | [![coveritup](https://coveritup.app/embed/kevincobain2000/go-hexe?branch=master&type=average-pr-days)](https://coveritup.app/kevincobain2000/go-hexe)     |



## Requirements

Install `go` using https://github.com/kevincobain2000/gobrew


## Convert static site to executable binary

```sh
git clone https://github.com/kevincobain2000/go-hexe.git

# make sure you copy the folder to dist in current dir
cp -r ../path/to/your/website dist

go build main.go
./main
```

## Development

```sh
go run main.go
# or
air
```

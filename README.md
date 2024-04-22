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

![CI](https://github.com/kevincobain2000/go-hexe/actions/workflows/build.yml/badge.svg)
![unit-test-run-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=unit-test-run-time&branch=master)
![build-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=build-time&branch=master)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=coverage&branch=master)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=go-binary-size&branch=master)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=go-mod-dependencies&branch=master)
![go-sec-issues](https://coveritup.app/badge?org=kevincobain2000&repo=go-hexe&type=go-sec-issues&branch=master)


![unit-test-run-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=unit-test-run-time&output=svg&width=160&height=160&branch=master)
![build-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=build-time&output=svg&width=160&height=160&branch=master)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=coverage&output=svg&width=160&height=160&branch=master)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=go-binary-size&output=svg&width=160&height=160&branch=master)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=go-mod-dependencies&output=svg&width=160&height=160&branch=master)
![go-sec-issues](https://coveritup.app/chart?org=kevincobain2000&repo=go-hexe&type=go-sec-issues&output=svg&width=160&height=160&branch=master)


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

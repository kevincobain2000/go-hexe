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

# Build Status


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

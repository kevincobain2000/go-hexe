<p align="center">
  <a href="https://github.com/kevincobain2000/aketemite">
    <img alt="gobrew" src="https://imgur.com/4tkG6Zh.png" width="360">
  </a>
</p>
<p align="center">
  Simple CLI tool written in Go, to ping a url and get the status code.
  <br>
  Monitor up status for websites, api and URLs, with automatic crawling capability.
</p>

**Quick Setup:** One command to ping multiple urls and get the result.

**Crawling capability:** Crawls entire page, finds the links and obtains the status code.

**Colorful:** Colorful output.


# Build Status


## Usage

```sh
go run main.go sample.yml
```

## Screenshots

```yml
# sample.yml
urls:
  #success case with crawling
  - name: https://kevincobain2000.github.io
    timeout: 2000 #in ms
    enabled: true
    crawl: true
  #error case
  - name: https://kevincobain2001.github.io
    timeout: 2000 #in ms
    enabled: true
    crawl: true
  #success case without crawling
  - name: https://github.com
    timeout: 2000
    enabled: true
    crawl: false

```

![Screenshot](https://imgur.com/q6w2IEJ.png)
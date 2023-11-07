package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/kevincobain2000/aketemite/pkg"
	"github.com/olekukonko/tablewriter"
)

const (
	TRUNCATE_URL_LENGTH = 100
)

func main() {
	finalExitCode := 0

	if len(os.Args) != 2 {
		fmt.Println("Usage: aketemite <config.yaml>")
		os.Exit(1)
	}

	configPath := os.Args[1]

	config := pkg.NewConfig(configPath)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Is Alive", "Response Code", "Url", "Response Time", "Response Size"})
	tableData := [][]string{}

	alreadyPingUrls := []string{}

	// non-crawling urls
	for _, url := range config.URLs {
		if !url.Enabled {
			continue
		}
		if url.Crawl {
			continue
		}

		hc := pkg.NewHttpChallenge(time.Duration(url.Timeout), url.Crawl)
		if pkg.StringInSlice(url.Name, alreadyPingUrls) {
			continue
		}
		alreadyPingUrls = append(alreadyPingUrls, url.Name)
		color.Notice.Println("Pinging: ", url.Name)
		hc.Ping(url.Name)

		alive := ""
		respCode := ""
		url := ""
		if hc.Result.IsAlive {
			alive = color.Success.Render("Yes")
			respCode = color.Success.Render(hc.Result.ResponseCode)
			url = color.Success.Render(pkg.TruncateString(hc.Result.Url, TRUNCATE_URL_LENGTH))
			color.Success.Printf("OK: %d %s\n", hc.Result.ResponseCode, hc.Result.Url)
		} else {
			alive = color.Error.Render("No")
			respCode = color.Error.Render(hc.Result.ResponseCode)
			url = color.Error.Render(pkg.TruncateString(hc.Result.Url, TRUNCATE_URL_LENGTH))
			color.Danger.Printf("FAIL: %d %s\n", hc.Result.ResponseCode, hc.Result.Url)
			finalExitCode = 1
		}
		dataItem := []string{
			alive,
			respCode,
			url,
			hc.Result.ResponseTime,
			fmt.Sprintf("%d KB", hc.Result.ResponseSize),
		}
		tableData = append(tableData, dataItem)
		fmt.Println()
	}

	// crawling urls
	for _, url := range config.URLs {
		if !url.Enabled {
			continue
		}
		if !url.Crawl {
			continue
		}
		hc := pkg.NewHttpChallenge(time.Duration(url.Timeout), url.Crawl)
		color.Notice.Println("Crawling: ", url.Name)
		urls := hc.CrawlHrefs(url.Name)
		color.Warn.Println("Located: ", len(urls), " urls")
		fmt.Println()
		for _, url := range urls {
			if pkg.StringInSlice(url, alreadyPingUrls) {
				continue
			}
			alreadyPingUrls = append(alreadyPingUrls, url)
			color.Notice.Println("Pinging: ", url)
			hc.Ping(url)

			alive := ""
			respCode := ""
			url := ""
			if hc.Result.IsAlive {
				alive = color.Success.Render("Yes")
				respCode = color.Success.Render(hc.Result.ResponseCode)
				url = color.Success.Render(pkg.TruncateString(hc.Result.Url, TRUNCATE_URL_LENGTH))
				color.Success.Printf("OK: %d %s\n", hc.Result.ResponseCode, hc.Result.Url)
			} else {
				alive = color.Error.Render("No")
				respCode = color.Error.Render(hc.Result.ResponseCode)
				url = color.Error.Render(pkg.TruncateString(hc.Result.Url, TRUNCATE_URL_LENGTH))
				color.Danger.Printf("FAIL: %d %s\n", hc.Result.ResponseCode, hc.Result.Url)
				finalExitCode = 1
			}
			dataItem := []string{
				alive,
				respCode,
				url,
				hc.Result.ResponseTime,
				fmt.Sprintf("%d KB", hc.Result.ResponseSize),
			}
			tableData = append(tableData, dataItem)
			fmt.Println()
		}
	}
	table.AppendBulk(tableData)
	table.Render()
	os.Exit(finalExitCode)
}

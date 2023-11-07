package pkg

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type HttpResult struct {
	IsAlive      bool
	ResponseCode int
	ResponseTime string
	ResponseSize int
	Title        string
	Url          string
}

type HttpChallenge struct {
	browse *browser.Browser
	crawl  bool
	Result HttpResult
}

func NewHttpChallenge(timeout time.Duration, crawl bool) *HttpChallenge {
	b := surf.NewBrowser()
	b.SetUserAgent("Go/aketemite")
	b.SetTimeout(timeout * time.Millisecond)

	return &HttpChallenge{
		browse: b,
		crawl:  crawl,
		Result: HttpResult{},
	}
}

func (hc *HttpChallenge) Ping(url string) {
	// response timer
	start := time.Now()
	err := hc.browse.Open(url)
	elapsed := time.Since(start).Round(time.Millisecond)

	var result HttpResult

	if err != nil {
		result = HttpResult{
			IsAlive:      false,
			ResponseCode: 0,
			Title:        "",
			Url:          url,
			ResponseTime: elapsed.String(),
			ResponseSize: 0,
		}
	} else {
		result = HttpResult{
			IsAlive:      hc.isStatusSuccess(hc.browse.StatusCode()),
			ResponseCode: hc.browse.StatusCode(),
			Title:        hc.browse.Title(),
			Url:          url,
			ResponseTime: elapsed.String(),
			ResponseSize: hc.responseSize(hc.browse.Body()),
		}
	}

	hc.Result = result
}

func (hc *HttpChallenge) responseSize(body string) int {
	bytes := len(body)
	kb := bytes / 1024
	return kb
}
func (hc *HttpChallenge) isStatusSuccess(code int) bool {
	return code >= 200 && code < 400
}

func (hc *HttpChallenge) CrawlHrefs(url string) []string {
	urls := []string{}
	urls = append(urls, url)
	err := hc.browse.Open(url)
	if err != nil {
		color.Error.Println("Error opening URL: ", err)
		return urls
	}

	// crawl the page and print all links
	hc.browse.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		href = hc.relativeToAbsoluteURL(href)

		isSubset := hc.isURL2SubsetOfURL1(url, href)
		if isSubset {
			urls = append(urls, href)
		}
	})
	urls = UniqueStrings(urls)
	return urls
}

func (hc *HttpChallenge) relativeToAbsoluteURL(href string) string {
	if !strings.HasPrefix(href, "http") && !strings.HasPrefix(href, "//") {
		href = fmt.Sprintf("%s://%s%s", hc.browse.Url().Scheme, hc.browse.Url().Host, href)
	}
	return href
}

func (hc *HttpChallenge) isURL2SubsetOfURL1(url1 string, url2 string) bool {
	// Parse both URLs
	parsedURL1, err := url.Parse(url1)
	if err != nil {
		return false
	}

	parsedURL2, err := url.Parse(url2)
	if err != nil {
		return false
	}

	// Check the scheme and host
	if parsedURL1.Scheme != parsedURL2.Scheme || parsedURL1.Host != parsedURL2.Host {
		return false
	}

	if !strings.HasPrefix(parsedURL2.Path, parsedURL1.Path) {
		return false
	}

	// Check query parameters
	params1 := parsedURL1.Query()
	params2 := parsedURL2.Query()

	for key, values := range params1 {
		if val2, ok := params2[key]; !ok || !IsEqualSlice(values, val2) {
			return false
		}
	}

	return true
}

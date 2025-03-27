package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

const (
	URL_CONTI = "http://continewsnv5otx5kaoje7krkto2qbu3gtqef22mnr7eaxw3y6ncz3ad.onion"

	TOR1 = "socks5://127.0.0.1:9050"
	TOR2 = "socks5://127.0.0.1:9051"
	TOR3 = "socks5://127.0.0.1:9052"

	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent(USER_AGENT),
	)
	rp, err := proxy.RoundRobinProxySwitcher(TOR1, TOR2, TOR3)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// On every a element script
	c.OnHTML("script", func(r *colly.HTMLElement) {

		fmt.Println("Probando la de conti y su script, visitando", r.Request.URL.String(), "texto...", r.Text)
	})

	c.OnHTML("a[href]", func(r *colly.HTMLElement) {
		r.Request.Visit(r.Attr("href"))
	})

	/* c.OnRequest(func(a *colly.Request) {
		fmt.Println("Visiting:", a.URL.String())

	})
	*/

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.Visit(URL_CONTI)

	c.Wait()
}

/* func createFile() {

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("conti_.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//if _, err := f.Write([]byte("appended some data\n")); err != nil {
	if _, err := f.Write([]byte(link)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
*/

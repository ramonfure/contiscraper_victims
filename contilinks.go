package main

import (
	//"encoding/json"
	//"flag"

	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

//var path = "/home/kanjid/projects/conticrw/dbconti/conti_1.txt
//
const (
      
      //Conti URL
      CONTI_URL_ONION = "http://continewsnv5otx5kaoje7krkto2qbu3gtqef22mnr7eaxw3y6ncz3ad.onion"

      //User Agents
      USER_AGENT_TORBROWSER = "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0"
      
      //Proxy Tor
      TOR1 = "socks5://127.0.0.1:9050"
    
  )
func main() {

	c := colly.NewCollector()

	c.SetProxy(TOR1)

	c.OnHTML("script", func(e *colly.HTMLElement) {

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(Abba)
	fmt.Println("Done!")
    }
}

func crawl(urlSearch string, user_agent string) {

  urlConti, _ := url.Parse(CONTI_URL_ONION)
  
  //Instantitae new NewCollector
  c := colly.NewCollector(
    colly.UserAgent(USER_AGENT_TORBROWSER),
    colly.AllowedDomains(urlConti.Host),
    colly.Async(true),
  )
  
  //No revisit
  c.AllowURLRevisit = false
  c.SetRequestTimeout(20 * time.Second)
  
  //Setting proxy
  c.SetProxy(TOR1)


  c.OnHTML(("a[href]", func(e *colly.HTMLElement))) {
	link := e.Attr("href")
	absoluteUrl := e.Request.AbsoluteURL(link)

  }
  c.OnHTML(("script", func(e *colly.HTMLElement))) {
        text := e.Text()

  }
  c.OnRes


}

func createFile(){
 link := e.Text
  // If the file doesn't exist, create it, or append to the file
 f, err := os.OpenFile("conti_bbdd.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

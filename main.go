package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {

	var u string
	var j string
	flag.StringVar(&j, "j", "", "javascript to be executed")
	flag.StringVar(&u, "u", "", "the URL to be scanned for Prototype Pollution")
	flag.Parse()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		chromedp.Evaluate(j, &res),
	)
	if err != nil {
		fmt.Printf("[!] %s does not appear to be vulnerable", u)
	} else {
		fmt.Printf("[+] %s is VULNERABLE!", u)
	}
}

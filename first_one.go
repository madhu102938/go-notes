// Exercise : to err is human

package main

import (
	"fmt"
	"net/url"
)

func main() {
	google, err := url.Parse("https://a b.com/")
	if err != nil {
		if errs, ok := err.(*url.Error); ok {
			fmt.Printf("%#v\n", errs)
			// &url.Error{Op:"parse", URL:"https://a b.com/", Err:" "}
		}
		return
	}
	fmt.Println(google)
}
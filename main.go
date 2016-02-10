package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// KagomeURL is the address for the Japanese morphological analyser Kagome service
var KagomeURL = "localhost:8080"

func main() {
	anaconda.SetConsumerKey(Secrets.ConsumerKey)
	anaconda.SetConsumerSecret(Secrets.ConsumerSecret)
	api := anaconda.NewTwitterApi(Secrets.AccessToken, Secrets.AccessTokenSecret)

	// searchResult, _ := api.GetSearch("日本語", nil)
	v := url.Values{}
	v.Add("track", "トレクル")
	stream := api.PublicStreamFilter(v)
	for {
		select {
		case t := <-stream.C:
			tweet := t.(anaconda.Tweet)
			fmt.Println(tweet)
			// putRequestKagome(tweet.Text)
			break
		default:
			break
		}
	}
}

func putRequestKagome(tweet string) {
	b := bytes.NewBufferString(tweet)
	// client, err := http.NewRequest("PUT", KagomeURL, b)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(client.Body)
	// fmt.Println(strings.NewReader("ひらがな"))
	// resp, _ := http.Post("localhost", "json", strings.NewReader("ひらがな"))
	resp, _ := http.Post("localhost", "application/string", b)
	fmt.Println(resp)
}

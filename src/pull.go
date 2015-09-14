package main

import "github.com/nii236/gank/vendor/github.com/ChimeraCoder/anaconda"

func main() {
	anaconda.SetConsumerKey("your-consumer-key")
	anaconda.SetConsumerSecret("your-consumer-secret")
	api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")
}

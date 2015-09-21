package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey("m03RgEpgBDLsODQZ7ew6yxstf")
	anaconda.SetConsumerSecret("fhsKFBvYHs0d097F09ijpv3WRxg2Fi8BnHOrMJxvW8QstBFOXl")
	api := anaconda.NewTwitterApi("16979832-LdYmktTeGmAj7IK7sTg9YrU2N5RebroA8deJ7FO9L", "Az1ABD07L8BvtFk0Pq7fjaSNNwLIOlCESrXQLjXfSArWn")
	searchResult, _ := api.GetSearch("日本語", nil)

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

	for {
	}
}

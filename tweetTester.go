package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey       = "p15iIog2URiyoMNK6bXPoxnbX"
	consumerSecret    = "PrahT9mvsUChhUTKZoKNDBSFOGvNJBTnYcFPLc8SP16GBz3Kuc"
	accessToken       = "916906728549052416-6FerwFdLDGNBcPbcQTnSEwZub1hb1zq"
	accessTokenSecret = "kWvf3TJuSGnzOWPU6lwv6ZQfz2MByut8YMkGn9AjJAOMV"
)

func main() {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	// AccessTokenとAccessTokenSecretのセット
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	_, err := api.PostTweet("hello GoLang!!", nil)
	if err != nil {
		fmt.Println("tweet faild.")
	}
}

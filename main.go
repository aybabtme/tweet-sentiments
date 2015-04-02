package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	log.Printf("parsing tweets from stdin...")
	tweets, err := ParseTweets(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d tweets found...", len(tweets))

	log.Printf("extracting %d features from tweets...", len(featureExtractors))
	for i, tweet := range tweets {
		// extract features
		features := []Feature{}
		for _, extractor := range featureExtractors {
			features = append(features, extractor(tweet))
		}
		tweets[i].Features = features
	}

	log.Printf("writing ARFF to stdout...")
	if err := MakeArff(tweets, os.Stdout); err != nil {
		log.Fatal(err)
	}
	log.Printf("done!")
}

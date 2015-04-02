package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Sentiment string

const (
	Positive  Sentiment = "positive"
	Negative  Sentiment = "negative"
	Objective Sentiment = "objective"
	Neutral   Sentiment = "neutral"
)

func ParseSentiment(str string) (Sentiment, bool) {
	str, _ = strconv.Unquote(str)
	switch Sentiment(str) {
	case Positive:
		return Positive, true
	case Negative:
		return Negative, true
	case Objective:
		return Objective, true
	case Neutral:
		return Neutral, true
	}
	return Sentiment(""), false
}

type FeatureType int

const (
	Numeric FeatureType = iota
	NumericFloat
	String
)

type Feature struct {
	Name  string
	Type  FeatureType
	Value interface{}
}

type Tweet struct {
	SID       int
	UID       int
	Sentiment Sentiment
	Corpus    string

	Features []Feature
}

func ParseTweets(rd io.Reader) ([]Tweet, error) {
	r := bufio.NewReader(rd)

	tweets := []Tweet{}
	for i := 0; ; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		values := strings.Split(line, "\t")

		for i, v := range values {
			values[i] = strings.TrimSpace(strings.ToLower(v))
		}

		if len(values) != 4 {
			continue
		}

		sid, err := strconv.Atoi(values[0])
		if err != nil {
			return tweets, fmt.Errorf("%d: sid: %v", i, err)
		}
		uid, err := strconv.Atoi(values[1])
		if err != nil {
			return tweets, fmt.Errorf("%d: uid: %v", i, err)
		}
		sentiment, ok := ParseSentiment(values[2])
		if !ok {
			return tweets, fmt.Errorf("%d: not a sentiment: %v", i, values[2])
		}

		tweets = append(tweets, Tweet{
			SID:       sid,
			UID:       uid,
			Sentiment: sentiment,
			Corpus:    values[3],
		})
	}
	return tweets, nil
}

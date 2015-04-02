package main

import (
	"bufio"
	"io"
	"log"
	"os"
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

func main() {
	r := bufio.NewReader(os.Stdin)

	for i := 0; ; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		values := strings.Split(line, "\t")

		for i, v := range values {
			values[i] = strings.TrimSpace(v)
		}

		if len(values) != 4 {
			continue
		}

		sid, err := strconv.Atoi(values[0])
		if err != nil {
			log.Printf("%d: sid: %v", i, err)
			continue
		}
		uid, err := strconv.Atoi(values[1])
		if err != nil {
			log.Printf("%d: uid: %v", i, err)
			continue
		}
		sentiment, ok := ParseSentiment(values[2])
		if !ok {
			log.Printf("%d: not a sentiment: %v", i, values[2])
			continue
		}

		t := Tweet{
			SID:       sid,
			UID:       uid,
			Sentiment: sentiment,
			Corpus:    values[3],
		}

		log.Printf("%+v", t)

	}
}

type Tweet struct {
	SID       int
	UID       int
	Sentiment Sentiment
	Corpus    string
}

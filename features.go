package main

import (
	"path"
	"regexp"
	"strings"
)

var featureExtractors = []func(Tweet) Feature{
	ExclamationMarks,
	QuestionMarks,
	HappyEmoticon,
	AngryEmoticon,
	DCSList,
	PositiveListCount,
	NegativeListCount,
	Posemo,
	Negemo,
}

func ExclamationMarks(t Tweet) Feature {
	return Feature{
		Name:  "exclamation_marks",
		Type:  Numeric,
		Value: strings.Count(t.Corpus, "!"),
	}
}

func QuestionMarks(t Tweet) Feature {
	return Feature{
		Name:  "question_marks",
		Type:  Numeric,
		Value: strings.Count(t.Corpus, "?"),
	}
}

func HappyEmoticon(t Tweet) Feature {

	count := 0
	for _, emoji := range happyEmojies {
		count += strings.Count(t.Corpus, emoji)
	}

	return Feature{
		Name:  "happy_emoji",
		Type:  Numeric,
		Value: count,
	}
}

func AngryEmoticon(t Tweet) Feature {
	count := 0
	for _, emoji := range angryEmojies {
		count += strings.Count(t.Corpus, emoji)
	}

	return Feature{
		Name:  "angry_emoji",
		Type:  Numeric,
		Value: count,
	}
}

func DCSList(t Tweet) Feature {
	score := 0.0
	wordRE := regexp.MustCompile(`\w+`)
	matchs := wordRE.FindAllString(t.Corpus, -1)
	for _, word := range matchs {
		if v, ok := dsclist[strings.ToLower(strings.TrimSpace(word))]; ok {
			score += v
		}
	}
	return Feature{
		Name:  "dcs_list_score",
		Type:  NumericFloat,
		Value: score,
	}
}

func PositiveListCount(t Tweet) Feature {
	count := 0
	wordRE := regexp.MustCompile(`\w+`)
	matchs := wordRE.FindAllString(t.Corpus, -1)
	for _, word := range matchs {
		if _, ok := positive[strings.ToLower(strings.TrimSpace(word))]; ok {
			count++
		}
	}
	return Feature{
		Name:  "positive_list",
		Type:  Numeric,
		Value: count,
	}
}

func NegativeListCount(t Tweet) Feature {
	count := 0
	wordRE := regexp.MustCompile(`\w+`)
	matchs := wordRE.FindAllString(t.Corpus, -1)
	for _, word := range matchs {
		if _, ok := negative[strings.ToLower(strings.TrimSpace(word))]; ok {
			count++
		}
	}
	return Feature{
		Name:  "negative_list",
		Type:  Numeric,
		Value: count,
	}
}

func Posemo(t Tweet) Feature {
	count := 0
	wordRE := regexp.MustCompile(`\w+`)
	matchs := wordRE.FindAllString(t.Corpus, -1)
	for _, word := range matchs {
		word := strings.ToLower(strings.TrimSpace(word))
		for _, wildcard := range posemoWildcard {
			if ok, _ := path.Match(wildcard, word); ok {
				count++
			}
		}
	}
	return Feature{
		Name:  "posemo_list",
		Type:  Numeric,
		Value: count,
	}
}

func Negemo(t Tweet) Feature {
	count := 0
	wordRE := regexp.MustCompile(`\w+`)
	matchs := wordRE.FindAllString(t.Corpus, -1)
	for _, word := range matchs {
		word := strings.ToLower(strings.TrimSpace(word))
		for _, wildcard := range posemoWildcard {
			if ok, _ := path.Match(wildcard, word); ok {
				count++
			}
		}
	}
	return Feature{
		Name:  "negemo_list",
		Type:  Numeric,
		Value: count,
	}
}

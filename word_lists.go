package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
Here we do the dirty work of importing word lists or building
them on the fly.
*/

var happyEmojies = func() []string {
	var eyes = []string{
		":", ";", "=", "X", ":-", ";-", "=-", "X-",
	}
	var mouths = []string{
		"]", ")", "}", "D", ">", "P", "p",
	}

	emojies := []string{
		"😍",
		"<3",
		"♥",
	}
	for _, mouth := range mouths {
		for _, eye := range eyes {
			emojies = append(emojies, eye+mouth)
		}
	}

	return emojies
}()

var angryEmojies = func() []string {
	var eyes = []string{
		":", ";", "=", "X", ":-", ";-", "=-", "X-",
	}
	var faces = []string{
		"[", "(", "{", "<", "X", "p",
	}

	emojies := []string{
		"😉",
		"</3",
	}
	for _, face := range faces {
		for _, eye := range eyes {
			emojies = append(emojies, eye+face+"\n")
		}
	}

	return emojies
}()

var q = struct{}{}

var badWordList = []string{
	"fuck",
	"suck",
	"asshole",
}

var goodWordList = []string{
	"happy",
	"joy",
}

var dsclist = func() map[string]float64 {

	list, err := wordlists_dsclist_2_txt()
	if err != nil {
		log.Panic(err)
	}

	rd := bufio.NewReader(bytes.NewReader(list.bytes))
	out := map[string]float64{}

	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		values := strings.Split(string(line), " ")
		if len(values) != 2 {
			// invalid line
			continue
		}
		word := strings.ToLower(strings.TrimSpace(values[0]))
		score, err := strconv.ParseFloat(strings.TrimSpace(values[1]), 64)
		if err != nil {
			log.Panic(err)
			continue
		}
		out[word] = score
	}

	return out
}()

var negative = func() map[string]struct{} {

	list, err := wordlists_negative_txt()
	if err != nil {
		log.Panic(err)
	}

	out := map[string]struct{}{}

	wordRE := regexp.MustCompile(`\w+`)
	words := wordRE.FindAllString(string(list.bytes), -1)

	for _, word := range words {
		out[strings.TrimSpace(strings.ToLower(word))] = struct{}{}
	}

	return out
}()

var positive = func() map[string]struct{} {

	list, err := wordlists_positive_txt()
	if err != nil {
		log.Panic(err)
	}

	out := map[string]struct{}{}

	wordRE := regexp.MustCompile(`\w+`)
	words := wordRE.FindAllString(string(list.bytes), -1)

	for _, word := range words {
		out[strings.TrimSpace(strings.ToLower(word))] = struct{}{}
	}

	return out
}()

var negemoWildcard = func() []string {

	list, err := wordlists_negemo_txt()
	if err != nil {
		log.Panic(err)
	}

	out := []string{}
	uniq := map[string]struct{}{}

	wordRE := regexp.MustCompile(`\w+\*?`)
	words := wordRE.FindAllString(string(list.bytes), -1)

	for _, word := range words {
		word = strings.TrimSpace(strings.ToLower(word))
		if _, ok := uniq[word]; !ok {
			uniq[word] = struct{}{}
			out = append(out, word)
		}
	}

	return out
}()

var posemoWildcard = func() []string {

	list, err := wordlists_posemo_txt()
	if err != nil {
		log.Panic(err)
	}

	out := []string{}
	uniq := map[string]struct{}{}

	wordRE := regexp.MustCompile(`\w+\*?`)
	words := wordRE.FindAllString(string(list.bytes), -1)

	for _, word := range words {
		word = strings.TrimSpace(strings.ToLower(word))
		if _, ok := uniq[word]; !ok {
			uniq[word] = struct{}{}
			out = append(out, word)
		}
	}

	return out
}()

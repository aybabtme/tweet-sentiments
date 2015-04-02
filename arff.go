package main

import (
	"fmt"
	"io"
)

type stickyWriter struct {
	err error
	w   io.Writer
}

func (s *stickyWriter) Write(p []byte) (int, error) {
	if s.err != nil {
		return 0, s.err
	}
	n, err := s.w.Write(p)
	if err != nil {
		s.err = err
	}
	return n, err
}

func MakeArff(tweets []Tweet, out io.WriteCloser) error {
	defer out.Close()
	w := &stickyWriter{w: out}

	featureNames := uniqueFeatures(tweets)

	fmt.Fprintf(w, "@relation opinion\n")
	// fmt.Fprintf(w, "@attribute sentence string\n")
	fmt.Fprintf(w, "@attribute category {positive,negative,neutral,objective}\n")

	for _, ft := range featureNames {
		switch ft.Type {
		case Numeric, NumericFloat:
			fmt.Fprintf(w, "@attribute %s numeric\n", ft.Name)
		case String:
			fmt.Fprintf(w, "@attribute %s string\n", ft.Name)
		}

	}

	fmt.Fprintf(w, "@data\n")
	if w.err != nil {
		return w.err
	}

	for i, tweet := range tweets {

		// fmt.Fprintf(w, "%q,", tweet.Corpus)
		fmt.Fprintf(w, "%s", tweet.Sentiment)

		for _, ft := range tweet.Features {

			fmt.Fprint(w, ",")

			switch ft.Type {
			case Numeric:
				fmt.Fprintf(w, "%d", ft.Value)
			case NumericFloat:
				fmt.Fprintf(w, "%.2f", ft.Value)
			case String:
				fmt.Fprintf(w, "' %s '", ft.Value)
			}

		}
		fmt.Fprint(w, "\n")
		if w.err != nil {
			return fmt.Errorf("writing tweet %d: %v", i, w.err)
		}
	}

	return w.err

}

func uniqueFeatures(tweets []Tweet) []Feature {
	features := []Feature{}
	uniq := map[string]struct{}{}
	for _, tweet := range tweets {
		for _, feature := range tweet.Features {
			if _, ok := uniq[feature.Name]; !ok {
				features = append(features, feature)
			}
			uniq[feature.Name] = struct{}{}

		}
	}
	return features
}

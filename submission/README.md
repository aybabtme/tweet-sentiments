To run the feature extractor and generate a ARFF file:

* Install Go 1.4.2, see https://golang.org/doc/install
* In this folder, run:

    $ go run *.go < semeval_twitter_data.txt > own.arff

You should see the following output:

    parsing tweets from stdin...
    7230 tweets found...
    extracting 14 features from tweets...
    writing ARFF to stdout...
    done!

And a new file names `own.arff` should have been created. You can use
this file in Weka to train a classifier.

- Antoine

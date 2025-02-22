package main

import (
	processing "./processing"
)

func main() {
	// Here we enter a list of sentences
	// corresponding to text in documents
	DocList := []string{
		"new home sales top forecasts",
		"home sales rise in July",
		"increase in home sales in July",
		"July new home sales rise"}

	// Generate the inverted index
	index := processing.GenerateInvertedIndex(DocList)

	// Search for "Sales" and output the
	// list of documents containing it
	processing.Find(index, "Sales")

	// Ability to show that search
	// outputs values that dont
	// exist in the documents
	processing.Find(index, "June")
}

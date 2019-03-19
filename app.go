package main

import "fmt"

// Example implementation

func main() {

	// Supply the FastText model file location
	model, err := New("basic-model.bin")
	if err != nil {
		panic(err)
	}

	// Label the sentence with that FastText model
	sentence := "Sentence to predict"
	result, err := model.Predict(sentence)
	if err != nil {
		panic(err)
	}

	// Switch based on the custom logic mapping
	if result == LabelA {
		fmt.Printf("`%s` -> `%s`\n", sentence, "label A")
	} else if result == LabelB {
		fmt.Printf("`%s` -> `%s`\n", sentence, "label B")
	} else {
		fmt.Printf("`%s` -> `%s`\n", sentence, "No Label")
	}
}

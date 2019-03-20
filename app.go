package main

// Example implementation

func main() {

	// Supply the FastText model file location
	model, err := New("basic-model.bin")
	if err != nil {
		panic(err)
	}

	// Label the sentence with that FastText model
	sentence := "Sentence to predict"
	err = model.Predict(sentence)
	if err != nil {
		panic(err)
	}
}

package fasttext

import (
	"testing"
)

func TestGetDimension(t *testing.T) {
	model, err := New("test_data/model.bin")
	if err != nil {
		t.Errorf("error loading model: %v", err)
	}
	d, err := model.GetDimension()
	if err != nil {
		t.Errorf("error getting dimension: %v", err)
	}
	if d != 50 {
		t.Errorf("wrong dimension")
	}

}

func TestSaveModel(t *testing.T) {
	var newFileName = "test_data/model_.bin"
	model, err := New("test_data/model.bin")
	if err != nil {
		t.Errorf("error loading model: %v", err)
	}
	err = model.SaveModel(newFileName)
	if err != nil {
		t.Errorf("error writing to a file: %v: %v", newFileName, err)
	}

}

func TestTrain(t *testing.T) {
	var (
		modelType      = "supervised"
		inputFileName  = "test_data/train"
		outputFileName = "test_data/clf"
		epoch          = 10
		wordNGrams     = 2
		thread         = 10
		lr             = 0.1
	)

	modelNew := &Model{isInitialized: false}

	err := modelNew.Train(modelType, inputFileName, outputFileName, epoch, wordNGrams, thread, lr)
	if err != nil {
		t.Errorf("error training model: %v", err)
	}

	err = modelNew.SaveModel(outputFileName)
	if err != nil {
		t.Errorf("error writing to a file: %v: %v", outputFileName, err)
	}
}

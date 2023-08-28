package fasttext

import (
	"testing"
)

func TestGetDimension(t *testing.T) {
	model, err := New("model.bin")
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
	var newFileName = "model_.bin"
	model, err := New("model.bin")
	if err != nil {
		t.Errorf("error loading model: %v", err)
	}
	err = model.SaveModel(newFileName)
	if err != nil {
		t.Errorf("error writing to a file: %v", newFileName)
	}

}

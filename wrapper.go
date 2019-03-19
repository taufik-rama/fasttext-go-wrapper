package main

// #cgo LDFLAGS: -L${SRCDIR}/lib -lfasttext-wrapper -lstdc++ -lm -pthread
// #include <stdlib.h>
// int load_model(char *path);
// int predict(char *query, float *prob, char *buf, int buf_size);
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	_ = iota

	// LabelA is an example prediction value label
	LabelA

	// LabelB is an example prediction value label
	LabelB

	// NoLabel is an example prediction value label
	NoLabel
)

// Model uses FastText for it's prediction
type Model struct {
	isInitialized bool
}

// New should be used to instantiate the model.
// FastTest needs some initialization for the model binary located on `file`.
func New(file string) (*Model, error) {

	status := C.load_model(C.CString(file))

	if status != 0 {
		return nil, fmt.Errorf("Cannot initialize model on `%s`", file)
	}

	return &Model{
		isInitialized: true,
	}, nil
}

// Predict the `keyword`
func (m *Model) Predict(keyword string) (int, error) {

	if !m.isInitialized {
		return NoLabel,
			errors.New("The FastText model needs to be initialized first. It's should be done inside the `New()` function")
	}

	var (
		label string
		prob  float64
		err   error
	)

	labelbufsize := 32

	var cprob C.float
	cbuflabel := (*C.char)(C.malloc(C.ulong(labelbufsize)))

	status := C.predict(
		C.CString(keyword),
		&cprob,
		cbuflabel,
		C.int(labelbufsize),
	)

	label = C.GoString(cbuflabel)
	prob = float64(cprob)
	if status != 0 {
		err = fmt.Errorf("Exception when predicting `%s`", keyword)
	}

	C.free(unsafe.Pointer(cbuflabel))

	if err != nil {
		return NoLabel, err
	}

	// Probability if needed
	fmt.Println("Probability is", prob)

	// Simple mapping from string label to constant
	if label == "fasttext_label_a" {
		return LabelA, nil
	} else if label == "fasttext_label_b" {
		return LabelB, nil
	}

	return NoLabel, err
}

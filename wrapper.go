package main

// #cgo LDFLAGS: -L${SRCDIR}/fastText/lib -lfasttext-wrapper -lstdc++ -lm -pthread
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
func (m *Model) Predict(keyword string) error {

	if !m.isInitialized {
		return errors.New("The FastText model needs to be initialized first. It's should be done inside the `New()` function")
	}

	resultSize := 32
	result := (*C.char)(C.malloc(C.ulong(resultSize)))

	var cprob C.float

	status := C.predict(
		C.CString(keyword),
		&cprob,
		result,
		C.int(resultSize),
	)
	if status != 0 {
		return fmt.Errorf("Exception when predicting `%s`", keyword)
	}

	// Here's the result from C
	label := C.GoString(result)
	prob := float64(cprob)
	fmt.Println(label, prob)

	C.free(unsafe.Pointer(result))

	return nil
}

package main

/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// #cgo LDFLAGS: -L${SRCDIR} -lfasttext -lstdc++ -lm
// #include <stdlib.h>
// void load_model(char *path);
// int predict(char *query, float *prob, char *buf, int buf_sz);
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	// Load the model
	// LoadModel("fastText/models/basic-model.bin")

	// Get the predicted value
	// fmt.Println(Predict("Chell"))
}

// LoadModel based on the `path`
func LoadModel(path string) {
	C.load_model(C.CString(path))
}

// Predict the `sentence`
func Predict(sentence string) (label string, prob float32, err error) {

	var cprob C.float
	cbuf := (*C.char)(C.malloc(C.ulong(len(sentence))))

	status := C.predict(
		C.CString(sentence),
		&cprob,
		cbuf,
		C.int(len(sentence)),
	)

	label = C.GoString(cbuf)
	prob = float32(cprob)
	if status != 0 {
		err = fmt.Errorf("Exception when predicting `%s`", sentence)
	}

	C.free(unsafe.Pointer(cbuf))
	return
}

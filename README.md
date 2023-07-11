# FastText Golang Wrapper

## Overview

Here's my attempt at wrapping FastText C++ library with Golang CGO.

## Requirements

- `git`
- `make`
- And [other requirements](https://github.com/facebookresearch/fastText/#requirements) for the FastText library.

## Compiling

- Clone the `FastText` git repository & compile it.

    ```Bash
    $ wget https://github.com/facebookresearch/fastText/archive/v0.9.2.zip
    $ unzip v0.9.2.zip
    $ cd fastText-0.9.2
    $ make
    ```

- Clone this repository & copy all the `.o` from `fastText-0.9.2` into directory inside `fasttext-go-wrapper/fastText/obj`.

    ```Bash
    $ git clone https://github.com/fkurushin/fasttext-go-wrapper
    $ mkdir fasttext-go-wrapper/fastText/obj
    $ cp fastText-0.9.2/*.o fasttext-go-wrapper/fastText/obj/
    ```

- Compile the C project

    ```Bash
    $ cd fasttext-go-wrapper/fastText && make
    ```

- Build the Go package normally (in the `fasttext-go-wrapper/` dir)

    ```Bash
    $ go build
    ```

    ```Bash
    $ sudo cp libfasttext-wrapper.a /usr/local/lib/
    ```

## Basic usage
- Initialization
    ```
    model, err = fasttext.New(modelName)
    if err != nil {
        panic(err)
    }
    ```
    
- Predict vector
    ```
    vec, err := model.GetSentenceVector(sentence)
    if err != nil {
        panic(err)
    }
    ```
- Get model dimension
    ```
    d, err := model.GetDimension()
	if err != nil {
		panic(err)
	}
    ```
## Example of Dockerfile 
    WORKDIR /src
    RUN wget https://github.com/facebookresearch/fastText/archive/v0.9.2.zip && \
        unzip v0.9.2.zip

    WORKDIR /src/fastText-0.9.2
    RUN make

    WORKDIR /src
    RUN git clone https://github.com/fkurushin/fasttext-go-wrapper && \
        rm v0.9.2.zip

    WORKDIR /src/fasttext-go-wrapper/fastText/obj
    RUN cp ../../../fastText-0.9.2/*.o .

    WORKDIR /src/fasttext-go-wrapper/fastText
    RUN make

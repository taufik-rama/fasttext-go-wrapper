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
    $ git clone https://github.com/facebookresearch/fastText
    # Cloning...

    $ cd fastText && make
    # Compiling...
    ```

- Clone this repository & copy all the `.o` from previous compile result into directory inside `fastText/obj`.

    ```Bash
    $ git clone https://github.com/taufik-rama/fasttext-go-binding
    # Cloning...

    $ mkdir fastText/obj

    $ cp /path/to/previous/repo/*.o fastText/obj/
    ```

- Compile the C project

    ```Bash
    $ cd fastText && make
    # Compiling...
    ```

- Build the Go package normally

    ```Bash
    $ go build
    # Compiling...
    ```
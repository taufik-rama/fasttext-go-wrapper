# FastText Golang Wrapper

## Overview

Here's my attempt at wrapping FastText C++ library with Golang CGO

## Requirements

- `git`
- `make`
- And [other requirements](https://github.com/facebookresearch/fastText/#requirements) for the FastText library

## Compiling

- Clone the `FastText` git repository & compile it

    ```Bash
    $ git clone https://github.com/facebookresearch/fastText
    # Cloning...

    $ cd fastText && make
    # Compiling...
    ```

- Copy all the `.o` file into `fastText/obj` directory

    ```Bash
    $ cp *.o path/to/fastText/obj/
    $ ls path/to/fastText/obj/
    ... # All the object file
    ```

- Compile the project and copy the resulting static archived object (`.a`) into `lib` directory. Then compile the golang package

    ```Bash
    # On `fastText` directory
    $ make
    # ...
    $ cp build/*.a ../lib/
    $ cd ..

    # Just compile it normally
    $ go build
    ```
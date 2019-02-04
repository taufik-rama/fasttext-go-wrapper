# How To's

- Copy the fasttext `.bin` model into `fastText/models`

- Copy the static fasttext library into main package directory:

```bash
$ cd fastText
$ mkdir build
$ cd build
$ cmake ..
... # Output omitted

$ make
... # Output omitted

$ cp libfasttext.a ../../
```

- Then, build it normally

```bash
# On main package directory
$ go build
$ ./fasttext-go-binding
```
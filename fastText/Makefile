CPPFLAGS += -g3 -std=c++11 -Iinclude

libfasttext-wrapper.a: libfasttext-wrapper.o
	ar rcs build/$@ obj/*.o

libfasttext-wrapper.o: fasttext-wrapper.cc
	if [ ! -d build ]; then mkdir build; fi
	g++ -c $(CPPFLAGS) $^ -o obj/$@

.PHONY: clean

clean:
	rm -rf build
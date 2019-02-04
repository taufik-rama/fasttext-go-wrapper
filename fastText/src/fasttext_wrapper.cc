#include <iostream>
#include <istream>
#include "fasttext.h"
#include "real.h"
#include <streambuf>
#include <cstring>

extern "C" {

    struct input_buffer : std::streambuf {
        input_buffer(char* begin, char* end) {
            this->setg(begin, begin, end);
        }
    };

    fasttext::FastText ft_model;
    bool ft_initialized = false;

    void load_model(char *path) {
        if (!ft_initialized) {
            ft_model.loadModel(std::string(path));
            ft_initialized = true;
        }
    }

    int predict(char *query, float *prob, char *out, int out_size) {

        int32_t k = 1;
        fasttext::real threshold = 0.0;

        input_buffer buffer(query, query + strlen(query));
        std::istream in(&buffer);

        std::vector<std::pair<fasttext::real, std::string>> predictions;

        if(!ft_model.predictLine(in, predictions, k, threshold)) {
            *prob = -1;
            strncpy(out, "", out_size);
            return 1;
        }

        for(const auto &prediction : predictions) {
            *prob = prediction.first;
            strncpy(out, prediction.second.c_str(), out_size);
        }

        return 0;
    }
}

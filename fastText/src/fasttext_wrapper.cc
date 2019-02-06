/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

#include <iostream>
#include <istream>
#include <sstream>
#include <cstring>
#include "real.h"
#include "fasttext.h"
#include "fasttext_wrapper.h"

extern "C" {

    fasttext::FastText ft_model;
    bool ft_initialized = false;

    bool has_newline(std::string str) {
        return (0 == str.compare(str.length() - 1, 1, "\n"));
    };

    void load_model(char *path) {
        if (!ft_initialized) {
            ft_model.loadModel(std::string(path));
            ft_initialized = true;
        }
    }

    int predict(char *q, float *prob, char *out, int out_size) {

        std::cout << "Running predict" << "\n";

        int32_t k = 1;
        fasttext::real threshold = 0.0;

        std::string query(q);

        if(!has_newline(query)) {
            query.append("\n");
        }

        std::istringstream inquery(query);
        std::istream &in = inquery;

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

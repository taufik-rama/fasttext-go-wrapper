/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

#include <unistd.h>
#include <iostream>
#include <istream>
#include <sstream>
#include <cstring>
#include <fastText/fasttext.h>
#include <fastText/args.h>
#include <fasttext-wrapper.hpp>

extern "C" {

    fasttext::FastText ft_model;
    bool ft_initialized = false;

    bool ft_has_newline(std::string str) {
        return (0 == str.compare(str.length() - 1, 1, "\n"));
    };

    int ft_load_model(const char *path) {
        if (!ft_initialized) {
            if(access(path, F_OK) != 0) {
                return -1;
            }
            ft_model.loadModel(std::string(path));
            ft_initialized = true;
        }
        return 0;
    }

    int ft_predict(const char *query_in, float *prob, char *out, int out_size) {

        int32_t k = 1;
        fasttext::real threshold = 0.0;

        std::string query(query_in);

        if(!ft_has_newline(query)) {
            query.append("\n");
        }

        std::istringstream inquery(query);
        std::istream &in = inquery;

        std::vector<std::pair<fasttext::real, std::string>> predictions;

        if(!ft_model.predictLine(in, predictions, k, threshold)) {
            *prob = -1;
            strncpy(out, "", out_size);
            return -1;
        }

        for(const auto &prediction : predictions) {
            *prob = prediction.first;
            strncpy(out, prediction.second.c_str(), out_size);
        }

        return 0;
    }

    int ft_get_vector_dimension()
    {
        if(!ft_initialized) {
            return -1;
        }
        return ft_model.getDimension();
    }

    int ft_get_sentence_vector(const char* query_in, float* vector_out, int vector_size)
    {
        std::string query(query_in);

        if(!ft_has_newline(query)) {
            query.append("\n");
        }

        std::istringstream inquery(query);
        std::istream &in = inquery;
        fasttext::Vector svec(ft_model.getDimension());

        ft_model.getSentenceVector(in, svec);
        if(svec.size() != vector_size) {
            return -1;
        }
        memcpy(vector_out, svec.data(), vector_size*sizeof(float));
        return 0;
    }
}

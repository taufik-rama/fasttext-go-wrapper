/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
#pragma once

extern "C" {
    
    /**
     * Initialize the fasttext model located on `path`
     * returns 0 on success
     */
    int ft_load_model(const char* path);

    /**
     * Predict a given keyword
     * `query_in`: The actual keyword to predict
     * `prob`: floating value to determine the probability of the result
     * `out`: Predicted value
     * `out_size`: How much characted to be copied into `out`
     * returns 0 on success
     */
    int ft_predict(const char* query_in, float* prob, char* out, int out_size);

    /**
     * get dimension of vector from loaded model
     * returns positive dimension on success
     */
    int ft_get_vector_dimension();
    
    /**
     * get vector representation from given sentence
     * `query_in`: The actual keyword to predict
     * `vector`: PRE-ALLOCATED buffer for sentence vector
     * `vector_size`: dimmension of allocated vector
     * returns 0 on success
     */
    int ft_get_sentence_vector(const char* query_in, float* vector, int vector_size);

}
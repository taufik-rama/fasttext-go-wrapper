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
    int load_model(char* path);

    /**
     * Predict a given keyword
     * `q`: The actual keyword to predict
     * `prob`: floating value to determine the probability of the result
     * `out`: Predicted value
     * `out_size`: How much characted to be copied into `out`
     * returns 0 on success
     */
    int predict(char* q, float* prob, char* out, int out_size);
    
}
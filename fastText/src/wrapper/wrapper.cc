/**
 * Copyright (c) 2019-present, Tokopedia, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

#include "../fasttext.h"

// class Wrapper {

//     bool initialized;
//     fasttext::FastText ft();

//   public:
//     Wrapper() : initialized(false) {}
//     ~Wrapper() {};

//     bool Initialize(char *path) {
//         if(!initialized) {
//             ft.loadModel(std::string(path));
//             initialized = true;
//         }
//     }

//     bool Predict(char *keyword) {
//         return false;
//     }
// };

bool initialized;
fasttext::FastText ft;

bool Initialize(char *path) {
    if(!initialized) {
        // ft.loadModel(std::string(path));
        initialized = true;
    }
}

bool Predict(char *keyword) {
    return false;
}

// Wrapper::Wrapper() 
// Wrapper::~Wrapper() {}

// bool Wrapper::Initialize

// bool Wrapper::Predict
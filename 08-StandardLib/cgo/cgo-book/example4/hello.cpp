// hello.cpp

#include <iosstream>

extern "C" {
    #include "hello.h"
}

void SayHello(const * s){
    std::count<< s;
}
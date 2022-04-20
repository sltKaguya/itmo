#pragma once

#include <cstring>
#include <iostream>

template<typename T>
class CircularBuffer {
public:
    /**
     * @brief CircularBuffer constructor
     * 
     * @param capacity capacity of object, zero as default
     */
    CircularBuffer(size_t capacity=0)
    : capacity(capacity)
    , data(new T[capacity]) 
    {
        std::cout << "ctor" << std::endl;
    }

    /**
     * @brief CircularBuffer copy constructor
     * 
     * @param other adress of object to copy
     */
    CircularBuffer(const CircularBuffer &other)
    : capacity(other.capacity)
    , data(new T[other.capacity]) 
    {
        std::memcpy(data, other.data, capacity);
        std::cout << "copyctor" << std::endl;
    }

    /**
     * @brief CirularBuffer copy assignment constructor
     * 
     * @param other adress of object to copy
     */
    CircularBuffer& operator=(const CircularBuffer &other) {
        if (&other == this) return *this;

        capacity = other.capacity;
        std::memcpy(data, other.data, capacity);
        std::cout << "copy assgnment ctor" << std::endl;

        return *this;
    }

    /**
     * @brief CircularBuffer destructor
     */
    ~CircularBuffer() {
        std::cout << "~ctor" << std::endl;
        delete[] data;
    };
private:
    /**
     * @brief Capacity of CircularBuffer
     */
    size_t capacity;

    /**
     * @brief Pointer to array of data of CircularBuffer
     */
    T* data;

};
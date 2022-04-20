#include "../include/circular_buffer.hpp"

int main() {
    CircularBuffer<int> intBuffer {10};
    CircularBuffer<int> copyIntBuffer = intBuffer;
    CircularBuffer<int> copyAssgnmntTest {23};
    copyAssgnmntTest = intBuffer;
}
#include "../include/point.h"

int main() {
    Point a {3, 4};
    Point b {2, 3};
    
    Point c = a;
    PrintPoint(c);
    
    c = b;
    PrintPoint(c);

    return 0;
}
#include <iostream>
#include <string>
#include <vector>

using std::cout;

struct Dot {
    int x, y; //position on X and Y axes
    
    Dot(int x_pos, int y_pos) {
        x = x_pos;
        y = y_pos;
    }
};

class BrokenLine {
    int breaks; //numbers of breaks i.e. numbers of dots + 2 (first and last)
    class LinkedList {
    };
};

class ClosedBroken {
};

class Polygon {
};

class Triangle {
};

class Trapeze {
};

class RegularPolygon {
};

int main() {
    Dot dot(3, 5);
    cout << "dot:" << dot.x << " " << dot.y; 
    return 0;
}
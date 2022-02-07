#include <iostream>
#include <string>
#include <vector>

class Dot {
    int x, y; //position on X and Y axes
    Dot *prevDot, *nextDot; //pointers for lines

    void createDot(int x, int y);
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
    std::cout << "Please, select an object to create. To see list of available objects type \"list\"" << std::endl;
    std::string comand = "";
    std::cin >> comand;
    
}
#include <iostream>
#include <string>
#include <vector>

class Dot {
    public:
        int x_pos, y_pos; //position on X and Y axes
        Dot* prevDot, * nextDot; //pointers for lines
        Dot(int x, int y) {
            x_pos = x;
            y_pos = y;
        };
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
    std::string objects = "List of availables objects:\n1) dot";
    std::string commands = "List of availables commands:";
    std::cout << "Please, select a command and an object to interact.\n\
        To see the list of available objects type \"objects\"\n\
        To see the list of commands type \"commands\"\n\
        For help type \"help\"" << std::endl;
    std::string command = "";
    getline(std::cin, command);

    if (command == "objects") {
        std::cout << objects << std::endl;
    } else if (command == "commands") {
        std::cout << commands << std::endl;
    } else if (command == "help") {
        std::cout << "there is no help..." << std::endl;
    }

    return 0;
}
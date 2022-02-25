#include "../include/point.h"
#include "../include/polychain.h"
#include "../include/closedchain.h"
#include "../include/polygon.h"
#include "../include/triangle.h"
#include "../include/trapeze.h"

int main() {
    Point a {3, 4};
    Point b {2, 3};
    Point d {5, 6};
    
    Point c = a;
    PrintPoint(c);
    
    c = b;
    PrintPoint(c);

    PolygonalChain myChain;

    myChain.PushBack(&a);
    myChain.PushBack(&b);
    myChain.PushBack(&b);
    myChain.PushBack(&a);

    PolygonalChain secondChain;

    secondChain.PushBack(&b);
    secondChain.PushBack(&a);
    secondChain.PushBack(&d);

    PolygonalChain thirdChain = secondChain;
    thirdChain = myChain;

    PrintChain(secondChain);

    std::cout << "Lenght of second Chain is " 
            << secondChain.Perimeter()
            << std::endl;

    ClosedChain closed;
    PrintChain(closed);
    std::cout << "Lenght of closed Chain is "
            << closed.Perimeter()
            << std::endl;
    closed.PushBack(&a);
    PrintChain(closed);
    std::cout << "Lenght of closed Chain is "
            << closed.Perimeter()
            << std::endl;
    closed.PushBack(&b);
    PrintChain(closed);
    std::cout << "Lenght of closed Chain is "
            << closed.Perimeter()
            << std::endl;
    closed.PushBack(&d);
    PrintChain(closed);
    std::cout << "Lenght of closed Chain is "
            << closed.Perimeter()
            << std::endl;
    closed.PushBack(&d);
    PrintChain(closed);
    std::cout << "Lenght of closed Chain is "
            << closed.Perimeter()
            << std::endl;

    Point one {0, 0};
    Point two {2, 0};
    Point three {2, 2};
    Point four {0, 2};

    Polygon square;
    square.PushBack(&one);
    square.PushBack(&two);
    square.PushBack(&three);
    square.PushBack(&four);
    PrintChain(square);
    std::cout << "Area of this square is "
            << square.Area()
            << std::endl;
    
    Point t1 {0, 0};
    Point t2 {2, 0};
    Point t3 {2, 2};
    Point t0 {0, 2};
    Triangle triangle;
    triangle.PushBack(&t1);
    triangle.PushBack(&t2);
    triangle.PushBack(&t3);
    triangle.PushBack(&t0);
    PrintChain(triangle);
    std::cout << "Area of this triangle is "
            << triangle.Area()
            << std::endl;

    Point tr1 {0, 0};
    Point tr2 {2, 0};
    Point tr3 {1.5, 1};
    Point tr0 {0.5, 1};
    Trapeze trapeze;
    trapeze.PushBack(&tr1);
    trapeze.PushBack(&tr2);
    trapeze.PushBack(&tr3);
    trapeze.PushBack(&tr0);
    std::cout << "Area of this trapeze is "
            << trapeze.Area()
            << std::endl;
    return 0;
}
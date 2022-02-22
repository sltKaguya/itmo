#include "../include/point.h"

Point::Point(int x_val, int y_val): x(x_val), y(y_val) {}

Point::Point(const Point &other): x(other.x), y(other.y) {}

Point::~Point() {}

int Point::x_val() const {return x;}

int Point::y_val() const {return y;}

Point& Point::operator=(const Point &other) {
    if (&other == this) {
        return *this;
    }

    x = other.x;
    y = other.y;
    return *this;
}

void PrintPoint(Point &point) {
    std::cout << "("
            << point.x_val()
            << "; "
            << point.y_val()
            << ")"
            << std::endl;
}
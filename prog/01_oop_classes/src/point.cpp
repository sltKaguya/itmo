#include "../include/point.h"

Point::Point() {}

Point::Point(int a, int b): x(a), y(b) {}

Point::Point(const Point &other): x(other.x), y(other.y) {}

Point::~Point() {}
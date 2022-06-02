#include "../include/point.hpp"

Point::Point(double xVal, double yVal) : x(xVal), y(yVal) {}

Point::~Point() {}

Point::Point(const Point &other) : x(other.x), y(other.y) {}

Point &Point::operator=(const Point &other) {
  if (&other == this) {
    return *this;
  }

  x = other.x, y = other.y;
  return *this;
}

double Point::getPos(const double *xRead, const double *yRead) const {
  xRead = &x;
  yRead = &y;
}
double Point::getPosX() const { return x; }
double Point::getPosY() const { return y; }

void Point::setPos(double newX, double newY) { x = newX, y = newY; }
void Point::setPosX(double newX) { x = newX; }
void Point::setPosY(double newY) { y = newY; }
// std::cout << "" << std::endl;
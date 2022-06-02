#include "../include/point.hpp"

int main() {
  Point a;
  Point b = a;
  Point c(3, 2);
  c = a;
  double ax, ay = a.getPos();
  double bx = b.getPosX();
  double by = b.getPosY();
  std::cout << ax << " " << ay << " " << bx << " " << by << std::endl;
  return 0;
}
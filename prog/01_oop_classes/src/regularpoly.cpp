#include "../include/regularpoly.h"

RegularPolygon::RegularPolygon(double xCenter, double yCenter, double rad, \
    int nCor, int aCor)
    : xC(xCenter)
    , yC(yCenter)
    , size(nCor+1)
    , points({})
{
    double radius = rad;
    int n = nCor;
    int Phi = aCor;

    if(n == 0) {
        std::cout << "Incorrect Polygon!" << std::endl;
        return;
    }

    for (int i = 0; i < n; i++) {
        double x_temp = xC + radius * cos((double) Phi + (2*M_PI*i)/n);
        double y_temp = yC + radius * cos((double) Phi + (2*M_PI*i)/n);
        Point tempPoint = {x_temp, y_temp};
        points.push_back(&tempPoint);
    }

    points.push_back(points[0]);
}

RegularPolygon::RegularPolygon(RegularPolygon const &other)
    : xC(other.xC)
    , yC(other.yC)
    , points(other.points)
    , size(other.size)
{}
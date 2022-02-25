#include "../include/regularpoly.h"

RegularPolygon::RegularPolygon(): Polygon(), regcoords({}) {}

void RegularPolygon::SetPolygon(double xCenter, double yCenter, double rad, \
    int nity, int Phi) {
    size = nity + 1;
    if(nity < 3) {
        std::cout << "Incorrect Polygon!" << std::endl;
        return;
    }

    for (int i = 0; i < nity; i++) {
        double x_temp = xCenter + (rad * cos(Phi + ((2*PI*i)/nity)));
        double y_temp = yCenter + (rad * sin(Phi + ((2*PI*i)/nity)));
        Point tempPoint {x_temp, y_temp};
        regcoords.push_back(tempPoint);
    }
    regcoords.push_back(regcoords[0]);
}

RegularPolygon::RegularPolygon(RegularPolygon const &other) {
    regcoords = other.regcoords;
    size = other.size;
}

double RegularPolygon::RPArea() {
    double area = 0;
    if (size < 4) {
        std::cout << "Can't count area, not a polygon" <<std::endl;
        
        return area;
    }
    for (int i = 0; i < size - 2; i++) {
        area += regcoords[i].x_val()*regcoords[i+1].y_val() - \
            regcoords[i+1].x_val()*regcoords[i].y_val();
    }
    area += regcoords[size - 2].x_val()*regcoords[0].y_val() - \
        regcoords[0].x_val()*regcoords[size - 2].y_val();
    area = 0.5 * std::abs(area);

    return area;
}
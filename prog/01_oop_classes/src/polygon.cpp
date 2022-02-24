#include "../include/polygon.h"

Polygon::Polygon(): ClosedChain() {}

double Polygon::Area() {
    double area = 0;
    if (size < 4) {
        std::cout << "Can't count area, not a polygon" <<std::endl;
        
        return area;
    }
    for (int i = 0; i < size - 2; i++) {
        area += points[i]->x_val()*points[i+1]->y_val() - \
            points[i+1]->x_val()*points[i]->y_val();
    }
    area += points[size - 2]->x_val()*points[0]->y_val() - \
        points[0]->x_val()*points[size - 2]->y_val();
    area = 0.5 * std::abs(area);

    return area;
}
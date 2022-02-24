#include "../include/triangle.h"

Triangle::Triangle(): Polygon(), state(true) {}

void Triangle::PushBack(Point *point) {
    if (!state) {
        std::cout << "Triangle already set" <<std::endl;
        return;
    }
    if (size == 0) {
        this->points.push_back(point);
    } else if (size == 1) {
        this->points.push_back(point);
        this->points.push_back(points[0]);
    } else{
        if (DistancePoints(this->points[size-3], this->points[size-2]) + \
            DistancePoints(this->points[size-2], point) == \
            DistancePoints(this->points[size-3], point)) {
            
            this->points[size-2] = point;
            return;
        }
        this->points[size - 1] = point;
        this->points.push_back(points[0]);
    }
    size = points.size();
    if (size == 4) {
        state = false;
    }
}
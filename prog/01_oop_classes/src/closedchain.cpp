#include "../include/closedchain.h"

ClosedChain::ClosedChain(): PolygonalChain() {}

void ClosedChain::PushBack(Point *point) {
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
}
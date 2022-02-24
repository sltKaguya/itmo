#include "../include/polychain.h"

PolygonalChain::PolygonalChain(): size(0), points({}) {}

void PolygonalChain::PushBack(Point *point) {
    if (size > 1) {
        if (DistancePoints(this->points[size-2], this->points[size-1]) + \
            DistancePoints(this->points[size-1], point) == \
            DistancePoints(this->points[size-2], point)) {
            
            this->points[size-1] = point;
            return;
        }
    }
    this->points.push_back(point);
    size++;
}
PolygonalChain::PolygonalChain(PolygonalChain const &other)
    : size(other.size)
    , points(other.points)
{}

PolygonalChain::~PolygonalChain() {}

int PolygonalChain::size_val() const {return size;}

PolygonalChain& PolygonalChain::operator=(const PolygonalChain &other) {
    if (&other == this) {
        return *this;
    }

    size = other.size;
    points = other.points;
    return *this;
}

double PolygonalChain::Perimeter() {
    double ans = 0;
    for (int i = 0; i < size - 1; i++) {
        ans += DistancePoints(this->points[i], this->points[i+1]);
    }
    return ans;
}

double DistancePoints(Point *first, Point *second) {
    double dist = pow(pow(second->x_val() - first->x_val(), 2) + \
        pow(second->y_val() - first->y_val(), 2), 0.5);
    
    return dist;
}

void PrintChain(PolygonalChain &chain) {
    std::cout << "Chain: [" << std::endl;
    for (int i = 0; i < chain.size; i++) {
        PrintPoint(*chain.points[i]);
    }
    std::cout << "]" << std::endl;
}
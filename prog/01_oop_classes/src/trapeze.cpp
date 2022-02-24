#include "../include/trapeze.h"

Trapeze::Trapeze(): Polygon(), k({}), x_offset({}), y_offset({}) {}

void Trapeze::PushBack(Point *point) {
    if (size == 0) {
        this->points.push_back(point);
    } else if (size == 1 && point != this->points[0]) {
        this->points.push_back(point);
        k.push_back((this->points[0]->y_val() - this->points[1]->y_val()) / \
            (this->points[0]->x_val() - this->points[1]->x_val()));
        x_offset.push_back(this->points[1]->x_val() - this->points[0]->x_val());
        y_offset.push_back(this->points[1]->y_val() - this->points[0]->y_val());       
    } else if (size == 2) {
        if (DistancePoints(this->points[size-2], this->points[size-1]) + \
            DistancePoints(this->points[size-1], point) == \
            DistancePoints(this->points[size-2], point)) {
            
            this->points[1] = point;
            
            return;
        } else {
            this->points.push_back(point);
            k.push_back((this->points[1]->y_val() - this->points[2]->y_val()) / \
                (this->points[2]->x_val() - this->points[2]->x_val()));
        }
    } else if (size == 3) {
        k.push_back((this->points[2]->y_val() - point->y_val()) / \
            (this->points[2]->x_val() - point->x_val()));
        k.push_back((point->y_val() - this->points[0]->y_val()) / \
            (point->x_val() - this->points[0]->x_val()));
        x_offset.push_back(point->x_val() - this->points[2]->x_val());
        y_offset.push_back(point->y_val() - this->points[2]->y_val());

        if ((k[0] != k[2]) || (k[1] == k[3]) || (x_offset[0]*x_offset[1] > 0) \
            || (y_offset[0]*y_offset[1] > 0)) {
                std::cout << "no)" << std::endl;
                return;
        } else {
            this->points.push_back(point);
            size++;
        }
    } else {
        std::cout << "cringe" << std::endl;
        return;
    }
    size++;
}
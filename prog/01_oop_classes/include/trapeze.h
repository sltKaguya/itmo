#pragma once

#include "polygon.h"

class Trapeze: public Polygon {
    public:
    /**
     * Create a new empty Trapeze object
     */
    Trapeze();

    /**
     * Add a Point in the Trapeze if possible
     * No more than 4 Points
     * base and top should be parallel
     * left and right shouldn't
     */
    void PushBack(Point *point);

    private:
    /**
     * Coefficient k in y = kx + b, base, right side, top, left side
     */
    std::vector<double> k;

    /**
     * Offset of x and y from Point 1 to Point 2 in first case
     * The offset from Point 3 to Point 4 in second case must be of opposite sign
     */
    std::vector<double> x_offset, y_offset;
};
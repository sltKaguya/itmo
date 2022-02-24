#pragma once

#include "../include/point.h"

#include <vector>

class BrokenLine {
    public:
    /**
     * Construct a new empty BrokenLine
     */
    BrokenLine();

    bool push_back(Point *point);
    private:
    /**
     * Number of breaks
     */
    int breaks;
    std::vector<Point*> points;
};
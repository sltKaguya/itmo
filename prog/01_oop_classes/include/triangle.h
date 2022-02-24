#pragma once

#include "../include/polygon.h"

class Triangle: public Polygon {
    public:
    /**
     * Construct a new empty Polygon object
     */
    Triangle();

    /**
     * Add a new Point object to the chain
     */
    void PushBack(Point *point);

    private:
    /**
     * State of Triangle. If it's done (three points), the state is set to
     * false and no more access to this figure
     * Quite useless but I like it :)
     */
    bool state;
};
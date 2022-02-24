#pragma once

#include "polychain.h"

class ClosedChain: public PolygonalChain {
    public:
    /**
     * Construct a new empty ClosedChain object
     */
    ClosedChain();

    /**
     * Add a new Point object to the chain and close it to the first Point
     */
    void PushBack(Point *point);
};
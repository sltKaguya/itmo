#pragma once

#include "closedchain.h"

class Polygon: public ClosedChain {
    public:
    /**
     * Construct a new empty Polygon object
     */
    Polygon();

    /**
     * Find area of Polygon object via Gauss
     */
    double Area();
};
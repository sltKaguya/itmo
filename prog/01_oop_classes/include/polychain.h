#pragma once

#include "point.h"

#include <cmath>
#include <vector>

class PolygonalChain {
    public:
    /**
     * Construct a new empty PolygonalChain object
     */
    PolygonalChain();

    /**
     * Add a new Point object to the chain
     */
    void PushBack(Point *point);

    /**
     * Construct a copy of a given PolygonalChain object
     * 
     * other - PolygonalChain object to copy 
     */
     PolygonalChain(PolygonalChain const &other);

    /**
     * Destroy object
     */
    ~PolygonalChain();

    /**
     * Get the size of the Chain
     */
    int size_val() const;

    /**
     * Assign one Chain object values to another existing
     * 
     * other - Chain object to assign
     */
    PolygonalChain& operator=(const PolygonalChain &other);

    /**
     * Print Points in Chain
     * 
     * chain - given Chain to print
     */
    friend void PrintChain(PolygonalChain &chain);


    /**
     * Count perimeter of Chain
     */
    double Perimeter();

    protected:
    /**
     * Number of breaks
     */
    int size;

    /**
     * Vector of pointers to Point objects 
     */
    std::vector<Point*> points;
};

/**
 * Count distance between 2 Point objects: sqrt of sum of x and y offset squares
 * 
 * first, second - given Point objects
 */
double DistancePoints(Point *first, Point *second);

/**
 * Print Points in Chain
 * 
 * chain - given Chain to print
 */
void PrintChain(PolygonalChain &chain);
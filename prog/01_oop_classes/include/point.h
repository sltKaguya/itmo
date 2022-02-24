#pragma once

#include <iostream>

class Point {
    public:
    /**
     * Construct a new Point object with given values or zeros for default
     * 
     * x_val - position on X-axis
     * y_val - position on Y-axis
     */
    Point(double x_val=0, double y_val=0);

    /**
     * Construct a copy of a given Point object
     * 
     * other - Point object to copy
     */
    Point(const Point &other);

    /**
     * Destroy object
     */
    ~Point();

    /**
     * Get the value on X-axis
     */
    double x_val() const;

    /**
     * Get the value on Y-axis
     */
    double y_val() const;

    /**
     * Assign one Point object values to another existing
     * 
     * other - Point object to assign
     */
    Point& operator=(const Point &other);

    private:
    /**
     * Positions on X and Y axes
     */
    double x, y;
};

/**
 * Print given Point object values on X and Y axes in (x; y) format
 * 
 * point - Point object to print
 */
void PrintPoint(Point &point);
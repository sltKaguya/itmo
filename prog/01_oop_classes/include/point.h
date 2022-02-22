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
    Point(int x_val=0, int y_val=0);

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
    int x_val() const;

    /**
     * Get the value on Y-axis
     */
    int y_val() const;

    /**
     * Assign one Point object values to another existing
     * 
     * other - Point object to assign
     */
    Point& operator =(const Point &other);

    private:
    /**
     * Positions on X and Y axes
     */
    int x, y;
};

/**
 * Print given Point objects values on X and Y axes
 * 
 * point - Point object to print
 */
void PrintPoint(Point &point);
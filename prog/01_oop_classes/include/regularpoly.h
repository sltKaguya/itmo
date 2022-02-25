#pragma once

#include "polygon.h"

class RegularPolygon: public Polygon {
    public:
    /**
     * Create a RegularPolygon object
     * With given xCenter and yCenter values make centerPoint
     * rad - radius of circumcircle
     * nCor - amount of corners
     * aCor - angle coordinate Phi
     */
    RegularPolygon(double xCenter=0, double yCenter=0, double rad=0, \
        int nCor=1, int aCor=0);
    
    /**
     * Construct a copy of a given RegularPolygon object
     * 
     * other - RegularPolygon object to copy
     */
    RegularPolygon(RegularPolygon const &other);

    private:

    /**
     * Size because it's easier RN
     */
    int size;
    
    /**
     * Points because it's easier RN
     */
    std::vector<Point*> points;
    /**
     * Center of circumcircle
     */
    double xC, yC;

    /**
     * Radius of inscribed circumcircle
     */
    //double radius;

    /**
     * Amount of corners and the angle coordinate
     */
    //int n, Phi;
};
#pragma once

#include "polygon.h"

const double PI = acos(-1);

class RegularPolygon: public Polygon {
    public:
    /**
     * Create a RegularPolygon object
     * With given xCenter and yCenter values make centerPoint
     * rad - radius of circumcircle
     * nCor - amount of corners
     * aCor - angle coordinate Phi
     */
    RegularPolygon();

    void SetPolygon(double xCenter=0, double yCenter=0, double rad=0, \
    int nity=0, int Phi=0);
    
    /**
     * Construct a copy of a given RegularPolygon object
     * 
     * other - RegularPolygon object to copy
     */
    RegularPolygon(RegularPolygon const &other);

    double RPArea();

    private:

    /**
     * Size because it's easier RN
     */
    //int size;
    
    /**
     * Points because it's easier RN
     */
    //std::vector<Point*> points;
    /**
     * Center of circumcircle
     */
    //double xC, yC;

    /**
     * Radius of inscribed circumcircle
     */
    //double radius;

    /**
     * Amount of corners and the angle coordinate
     */
    //int n, Phi;

    std::vector<Point> regcoords;
};
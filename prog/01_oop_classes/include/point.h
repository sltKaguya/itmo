#pragma once

/**
 * I. Define Point class with 2 position on X and Y axes:
 ** 1) void constructor;
 ** 2) consructor with given arguments;
 ** 3) copy constructor;
 ** 4) destructor 
**/
class Point {
    public:
        Point();

        Point(int a, int b);

        Point(const Point &other);

        ~Point();

    private:
        int x, y;
};
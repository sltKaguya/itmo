#pragma once

#include <iostream>

class Point {
public:
  /**
   * @brief Construct a new Point object with given values or zeros for default
   *
   * @param x_val position on X-axis
   * @param y_val position on Y-axis
   */
  Point(double x_val = 0, double y_val = 0);

  /**
   * @brief Construct a copy of a given Point object
   *
   * @param other Point object to copy
   */
  Point(const Point &other);

  /**
   * @brief Default destructor
   */
  ~Point();

  /**
   * @brief Getter
   *
   * @return Value on X-axis
   */
  double getXVal() const;

private:
};
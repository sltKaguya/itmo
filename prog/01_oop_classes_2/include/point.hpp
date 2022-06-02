#pragma once

#include <iostream>

class Point {
public:
  /**
   * @brief Constructor
   *
   * @param xVal position on X-axis (0 for default)
   * @param yVal position on Y-axis (0 for default)
   */
  Point(double xVal = 0, double yVal = 0);

  /**
   * @brief Destructor
   */
  ~Point();

  /**
   * @brief Copy constructor
   *
   * @param other object to copy
   */
  Point(const Point &other);

  /**
   * @brief Copy assignment constructor
   *
   * @param other object to assign
   */
  Point &operator=(const Point &other);

  /**
   * @brief Getter
   *
   * @return Positions on X and Y axes
   */
  double getPos(const double *xRead, const double *yRead) const;

  /**
   * @brief Getter
   *
   * @return Position on X-axis
   */
  double getPosX() const;

  /**
   * @brief Getter
   *
   * @return Position on Y-axis
   */
  double getPosY() const;

  /**
   * @brief Setter
   *
   * @param newX new value on X-axis
   * @param newY new value on Y-axis
   */
  void setPos(double newX, double newY);

  /**
   * @brief Setter
   *
   * @param newX new value on X-axis
   */
  void setPosX(double newX);

  /**
   * @brief Setter
   *
   * @param newY new value on Y-axis
   */
  void setPosY(double newY);

private:
  /**
   * @brief Position on X/Y-axis
   */
  double x, y;
};
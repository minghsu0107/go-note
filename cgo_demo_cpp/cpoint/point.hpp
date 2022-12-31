#ifndef POINT_HPP
#define POINT_HPP

#include <cmath>

class Point {
public:
    Point(double x, double y);
    double x();
    double y();
    static double distance(Point *p, Point *q);
private:
    double _x;
    double _y;
};

Point::Point(double x, double y) {
    this->_x = x;
    this->_y = y;
}

double Point::x() {
    return this->_x;
}

double Point::y() {
    return this->_y;
}

double Point::distance(Point *p, Point *q) {
    double dx = p->x() - q->x();
    double dy = p->y() - q->y();

    return sqrt(dx * dx + dy * dy);
}

#endif  /* POINT_HPP */
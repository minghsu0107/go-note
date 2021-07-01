#include <assert.h>
#include <math.h>
#include <stdlib.h>
#include "point.h"

struct point_t {
    double x;
    double y;
};

point_t * point_new(double x, double y) {
    point_t *pt = (point_t *) malloc(sizeof(point_t));
    if (!pt)
        return pt;

    pt->x = x;
    pt->y = y;

    return pt;
}

void point_delete(void *self) {
    assert(self);

    free(self);
}

double point_x(point_t *self) {
    assert(self);

    return self->x;
}

double point_y(point_t *self) {
    assert(self);

    return self->y;
}

double point_distance(point_t *p, point_t *q) {
    assert(p);
    assert(q);

    double dx = p->x - q->x;
    double dy = p->y - q->y;

    return sqrt(dx * dx + dy * dy);
}
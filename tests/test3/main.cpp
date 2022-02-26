#include <iostream>
#include "out/test1.hpp"
#include "out/test2.hpp"
using std::cout;
using std::endl;

int main() {
    Test1 t1;
    t1.t = new Test2();
    t1.t->t = &t1;
    t1.t->t->value = 10;
    cout << t1.t->t->value << " " << t1.value << endl;
}
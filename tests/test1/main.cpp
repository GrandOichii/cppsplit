#include <iostream>
#include "out/test.hpp"

using std::cout;

int main() {
    simpleFunc("ere", 4);
    anotherFunc();
    Human h("Igor", 19);
    h.printInfo();
    MyStruct s = createMS();
    s.value += 10;
    cout << s.value << endl;
}
#include <iostream>
using std::cout;
using std::endl;

class A {
protected:
    int value;
public:
    A(int value) {
        this->value = value;
    }
};

class B : A {
public:
    B() : A(10) {

    }
    void print() {
        cout << this->value << endl;
    }
};
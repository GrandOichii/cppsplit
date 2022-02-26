#include <iostream>
#include <string>
using std::cout;
using std::endl;
using std::string;

void simpleFunc(string line, int count) {
    for (int i = 0; i < count; i++)
        cout << line << endl;
}

void anotherFunc() {
    cout << "Hello" << endl;
}

struct MyStruct {
    int value;
};

MyStruct createMS() {
    struct MyStruct s;
    s.value = 1;
    return s;
}

class Human {
private:
    string name;
    int age;
public:
    Human(string name, int age) {
        this->name = name;
        this->age = age;
    }
    void printInfo() {
        cout << name << "\t" << age << endl;
    }
};
#include "test.hpp"

void simpleFunc(string line, int count) {
    for (int i = 0; i < count; i++)
        cout << line << endl;
}

void anotherFunc() {
    cout << "Hello" << endl;
}

Human::Human(string name, int age) {
	this->name = name;
	this->age = age;
}

void Human::printInfo() {
	cout << name << "\t" << age << endl;
}


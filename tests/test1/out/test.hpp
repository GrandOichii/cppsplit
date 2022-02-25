#pragma once

#include <iostream>
#include <string>
using std::cout;
using std::endl;
using std::string;

void simpleFunc(string line, int count);

void anotherFunc();

class Human {
private:
    string name;
    int age;
public:
    Human(string name, int age) ;
    void printInfo() ;
};

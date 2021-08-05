#include <iostream>

class Sounder {
public:
    virtual void foo() = 0;
    virtual void sound() = 0;
};

class Animal: public Sounder {
public:
    void foo() {
        std::cout << "foo from animal" << std::endl;
    }

    void sound() {
        this->foo();
        std::cout << "Pruttel" << std::endl;
    }
};


class Duck: public Animal {
public:
    void foo() {
        std::cout << "foo from animal" << std::endl;
    }
};


int main() {
    Animal a{};
    Duck d{};

    Sounder & ref_a = a;
    Sounder & ref_d = d;

    ref_a.sound();
    ref_d.sound();

    return 0;
}

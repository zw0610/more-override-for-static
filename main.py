class Animal(object):
    def __init__(self):
        pass

    def foo(self):
        print("foo from animal")

    def sound(self):
        self.foo()
        print("Pruttel")


class Duck(Animal):
    def __init__(self):
        super(Duck, self).__init__()

    def foo(self):
        print("foo from duck")


if __name__ == '__main__':
    a = Animal()
    a.sound()

    d = Duck()
    d.sound()

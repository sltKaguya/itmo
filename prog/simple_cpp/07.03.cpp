#include <iostream>


class FileHandle {
    public:
    FileHandle(const char *path, const char *mode) {
        file = fopen(path, mode);

        if (file == nullptr) {
            //some code
        }
    }

    ~FileHandle() {
        if (file) {
            fclose(file);
        }
    }

    //Правило трёх: запретить или убрать в приватные поля копирование и присваивание
    printf() {}//... придётся писать слишком много функций

    FILE *getFile() {

    }//Везде, где нужно будет обратиться к файлу, нужно будет писать file.getFile()

    FILE *operator()(){
        return file;
    }

    operator FILE*() {// Перегузка оператора приведения типов
        return file;
    }

    private:
    FILE *file;
};

class Foo {
    public:
    Foo(int value): value(value) {}

    int value() const {
        return value;
    }

    private:
    int value;
};

class FooPtr {// Оболочка класса
    public:
    FooPtr(int value): ptr(new Foo(value)) {}

    ~FooPtr() {
        if (ptr) {
            delete ptr;
            ptr = nullptr;
        }
    }

    Foo* operator->() {//каждый раз при обращении к указателю вызываем функцию
        return ptr;
    }

    Foo& operator*() {
        return *ptr;
    }
    private:
    Foo *ptr;
};

//Наследование

class Person {
    //constructor does name and age
    Person(const std::string& name, unsigned age)
        : name(name)
        , age(age)
    {}

    protected:
    std::string name;
    unsigned age;
};

class Student: public Person {
    public:
    //constructor does Person + university
    Student(const std::string& name, unsigned age, const std::string& university)
        : Person(name, age)
        , university(university) 
    {}

    private:
    std::string university;
}

virtual void Hello(const Person& p) {}//virtual 

int main() {
    FileHandle File("smth.txt", "w");
    fprintf(file, "test\n");

    Student ivan;
    Hello(ivan)// Можно передать класс наследник в функцию
}

/**
 * Обычное открытие файла и функция fprintf
 * Может не закрыться, если вдруг программа словит ошибку.
 * Соотвественно, надо создать класс и засунуть fclose в деструктор
 * 
 * Идиома RAII
 */
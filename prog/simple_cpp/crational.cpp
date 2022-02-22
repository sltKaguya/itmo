#include <iostream>

class CRational {
    public:

        CRational(int num=0, unsigned denom=1) //=default if none arg given
            : numerator(num)
            , denominator(denom)
        {
            std::cout << "ctor" << std::endl;
        };

        ~CRational() {
            std::cout << "~ctor" << std::endl;
        };

        CRational(const CRational &other) //copy constructor
            : numerator(other.numerator)
            , denominator(other.denominator)
        {
            std::cout << "copy ctor" << std::endl;
        };

        CRational(CRational &&other) //move constructor
            : numerator(std::__exchange(other.numerator, 0))
            , denominator(std::__exchange(other.denominator, 0))
        {}

        CRational &operator =(const CRational &other) {
            if (&other == this) {
                std::cout << "same nums" << std::endl;
                return *this;
            }

            numerator = other.numerator;
            denominator = other.denominator;
            std::cout << "copy assignement" << std::endl;
            return *this;
        };

        int Numerator() const {
            return numerator;
        };
    
        unsigned Denominator() const {
            return denominator;
        }

        void PrintRational(const CRational &number) {
            std::cout << number.Numerator()
                    << "/"
                    << number.Denominator()
                    << std::endl;
        };

    private:
        int numerator;
        unsigned denominator;
};

int main() {
    CRational damn;
    CRational fuck {1, 4};
    fuck.PrintRational(fuck);
    return 0;
}

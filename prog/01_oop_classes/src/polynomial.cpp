#include "../include/polynomial.h"

Node::Node(double coeff, unsigned pow): coefficient(coeff), pow(pow) {}

Node::~Node() {}

double Node::get_coeff() const {
    return coefficient;
}

unsigned Node::get_pow() const {
    return pow;
}

Polynomial::Polynomial(): coeffs({}) {}

Polynomial::~Polynomial() {}

void Polynomial::CreatePolynomial(std::vector<Node> vector_in) {
    for (int i=0; i < (int)vector_in.size(); i++) {
        PushNode(vector_in[i]);
    }
}

void Polynomial::PushNode(Node node) {
    if (coeffs.size() == 0) {
        coeffs.push_back(node);
        return;
    }
    for (unsigned i=0; i < coeffs.size(); i++) {
        if (coeffs[i].get_pow() < node.get_pow()) {
            coeffs.insert(coeffs.begin() + i, node);
            return;
        } else if (coeffs[i].get_pow() == node.get_pow()) {
            std::cout << "Incorrect input of Node" << std::endl;
            return;
        } else {
            coeffs.push_back(node);
        }
    }
}

void Polynomial::PrintPolynomial() {

    for (unsigned i=0; i < coeffs.size(); i++) {
        if (coeffs[i].get_pow() != 0) {
            std::cout << coeffs[i].get_coeff()
                      << "x^"
                      << coeffs[i].get_pow()
                      << " + "
                      << std::endl;
        } else {
            std::cout << coeffs[i].get_coeff()
                      << std::endl;
        }
    }

}
#pragma once

#include <iostream>
#include <vector>

class Node {
    public:

    /**
     * Construct a new Node object with given values or zeros if no values
     */
    Node(double coeff=0, unsigned pow=0);

    /**
     * Destroy object
     */
    ~Node();

    /**
     * Getter for the coefficient of Node object
     */
    double get_coeff() const;
    /**
     * Getter for power of Node object
     */
    unsigned get_pow() const;

    private:
    /**
     * Coefficient value
     */
    double coefficient;

    /**
     * Power of the coefficient's argument
     */
    unsigned pow;
};

class Polynomial {
    public:
    /**
     * Construct a new empty Polynomial object
     */
    Polynomial();

    /**
     * Destroy object
     */
    ~Polynomial();

    /**
     * Create a Polynomial with given Nodes
     */
    void CreatePolynomial(std::vector<Node> vector_in);

    /**
     * Pushes a Node object into Polynomial in the right place
     * If already exists a Node with same pow return error
     */
    void PushNode(Node node);

    /**
     * Print the Polynomial coefficients in "coeff x^pow +..." format
     */
    void PrintPolynomial();

    private:
    /**
     * Vector of coefficients
     */
    std::vector<Node> coeffs;
};
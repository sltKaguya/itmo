CC = g++
FLAGS = -Wextra -Wall -g -c
INC = include
SRC = src
MODULES = polynomial.o regularpoly.o trapeze.o triangle.o polygon.o closedchain.o polychain.o point.o main.o

main.exe: $(MODULES)
	$(CC) $(MODULES) -o main.exe

main.o: $(SRC)/main.cpp
	$(CC) $(FLAGS) $(SRC)/main.cpp

point.o: $(SRC)/point.cpp $(INC)/point.h
	$(CC) $(FLAGS) $(SRC)/point.cpp

polychain.o: $(SRC)/polychain.cpp $(INC)/polychain.h
	$(CC) $(FLAGS) $(SRC)/polychain.cpp

closedchain.o: $(SRC)/closedchain.cpp $(INC)/closedchain.h
	$(CC) $(FLAGS) $(SRC)/closedchain.cpp

polygon.o: $(SRC)/polygon.cpp $(INC)/polygon.h
	$(CC) $(FLAGS) $(SRC)/polygon.cpp

triangle.o: $(SRC)/triangle.cpp $(INC)/triangle.h
	$(CC) $(FLAGS) $(SRC)/triangle.cpp

trapeze.o: $(SRC)/trapeze.cpp $(INC)/trapeze.h
	$(CC) $(FLAGS) $(SRC)/trapeze.cpp

regularpoly.o: $(SRC)/regularpoly.cpp $(INC)/regularpoly.h
	$(CC) $(FLAGS) $(SRC)/regularpoly.cpp

polynomial.o: $(SRC)/polynomial.cpp $(INC)/polynomial.h
	$(CC) $(FLAGS) $(SRC)/polynomial.cpp

clean:
	rm -f $(MODULES)
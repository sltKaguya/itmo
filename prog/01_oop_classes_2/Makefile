CC = g++
FLAGS = -Wextra -Wall -g -c
INC = include
SRC = src
MODULES = point.o main.o

main.exe: $(MODULES)
	$(CC) $(MODULES) -o main.exe

main.o: $(SRC)/main.cpp
	$(CC) $(FLAGS) $(SRC)/main.cpp

point.o: $(SRC)/point.cpp $(INC)/point.hpp
	$(CC) $(FLAGS) $(SRC)/point.cpp

clean:
	rm -f $(MODULES)
CC = g++
FLAGS = -Wextra -Wall -g -c -std=c++17
INC = include
SRC = src

main.exe: main.o
	$(CC) main.o -o main.exe

main.o: $(SRC)/main.cpp $(SRC)/circular_buffer.hpp
	$(CC) $(FLAGS) $(SRC)/main.cpp

$(SRC)/circular_buffer.hpp:
	@echo changes in circular_buffer.hpp

clean:
	rm -rf main.o
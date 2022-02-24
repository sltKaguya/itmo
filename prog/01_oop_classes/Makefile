CC = g++
FLAGS = -Wextra -Wall -g -c
INC = include
SRC = src
MODULES = polygon.o closedchain.o polychain.o point.o main.o

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

clean:
	rm -f $(MODULES)
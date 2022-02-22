CC = g++
FLAGS = -Wextra -Wall -g -c
INC = include
SRC = src
MODULES = point.o main.o

classes.exe: $(MODULES)
	$(CC) $(MODULES) -o logs.exe

main.o: $(SRC)/main.c
	$(CC) $(FLAGS) $(SRC)/main.c

point.o: $(SRC)/point.c $(INC)/point.h
	$(CC) $(FLAGS) $(SRC)/point.c

clean:
	rm -f $(MODULES)
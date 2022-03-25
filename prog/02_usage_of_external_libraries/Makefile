CC = g++
FLAGS = -Wextra -Wall -g -c -std=c++17
INC = include
SRC = src
MODULES = main.o curlSetUp.o jsonSetUp.o

main.exe: $(MODULES)
	$(CC) $(MODULES) -o main.exe -lcurl

main.o: $(SRC)/main.cpp
	$(CC) $(FLAGS) $(SRC)/main.cpp

curlSetUp.o: $(SRC)/curlSetUp.cpp
	$(CC) $(FLAGS) $(SRC)/curlSetUp.cpp

jsonSetUp.o: $(SRC)/jsonSetUp.cpp
	$(CC) $(FLAGS) $(SRC)/jsonSetUp.cpp

clean:
	rm -f $(MODULES)
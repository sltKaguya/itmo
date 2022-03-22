CC = clang
FLAGS = -Wextra -Wall -g -c
BIN = first_try

all: $(BIN)

%: %.c
	$(CC) $(FLAGS) $< -o $@ -lcurl

clean:
	$(RM) -rf $(BIN) *.dSYM
SRC_FILES = main.go game.go minimax.go
OUT_EXT = out
OUT_FILE = connectfour.$(OUT_EXT)

ifeq ($(OS),Windows_NT)
	OUT_EXT = exe
endif

all: build

build:
	go build -o $(OUT_FILE) $(SRC_FILES)

run: build
	@$(OUT_FILE)

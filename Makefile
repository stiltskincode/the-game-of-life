.PHONY: build run_glider run_random

build: 
	go build -o the-game-of-life .

run_glider:
	./the-game-of-life -map 1

run_random:
	./the-game-of-life -map 2

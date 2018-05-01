default:
	@tmux send-keys -t right 'make show' C-m

build:
	@go build gifs.go
	@./gifs

ifeq ($(shell uname -s), Darwin)
show: build
	@qlmanage -p out/latest.gif 2>/dev/null
else
show: build
	@eog 2>/dev/null out/latest.gif
endif

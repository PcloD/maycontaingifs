default:
	@tmux send-keys -t right 'make build' C-m

build:
	@go build gifs.go
	@./gifs
	@eog 2>/dev/null out/latest.gif

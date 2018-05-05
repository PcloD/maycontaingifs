date := 07

default:
	tmux send-keys -t right 'make display' C-m

ifeq ($(shell uname -s), Darwin)
display: out/may$(date).gif
	open out/may$(date).gif
else
display: out/may$(date).gif
	eog out/may$(date).gif
endif

out/may$(date).gif: gifs gifs.go may/may$(date).go
	go build gifs.go
	./gifs $(date)

gifs:
	go build gifs.go

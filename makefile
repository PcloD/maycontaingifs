<<<<<<< HEAD
date := 12
||||||| merged common ancestors
date := 13

default:
	tmux send-keys -t right 'make display' C-m

ifeq ($(shell uname -s), Darwin)
display: out/may$(date).gif
	open out/may$(date).gif
else
display: out/may$(date).gif
	eog 2>/dev/null out/may$(date).gif
endif

out/may$(date).gif: gifs.go may/may$(date).go
	go build gifs.go
	./gifs $(date)

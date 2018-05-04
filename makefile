date := 12

default:
	tmux send-keys -t right 'make display' C-m

out/may$(date).gif: gifs may/may$(date).go
	go build gifs.go
	./gifs $(date)

ifeq ($(shell uname -s), Darwin)
display: may/may$(date).go
	# @qlmanage -p out/latest.gif 2>/dev/null
	open out/may$(date).gif
else
display: may/may$(date).go out/may$(date).gif
	eog 2>/dev/null out/may$(date).gif
endif

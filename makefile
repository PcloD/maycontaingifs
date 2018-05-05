date := 14

default:
	@tmux_send 'make display'

display: out/may$(date).gif
	@img_view $<

launch:
	@img_view out/may$(date).gif

clean:
	@tmux_send "# NO CLEAN!"

out/may$(date).gif: gifs gifs.go may/may$(date).go
	go build gifs.go
	./gifs $(date)

gifs:
	go build gifs.go

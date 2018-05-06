date := 16

vpath %.gif out
vpath %.go may

default:
	@tmux_send 'make display'

display: may$(date).gif
	@img_view $<

launch:
	@img_view may$(date).gif

clean:
	@tmux_send "# NO CLEAN!"

out/may$(date).gif: gifs gifs.go may$(date).go
	go build gifs.go
	./gifs $(date)

gifs:
	go build gifs.go

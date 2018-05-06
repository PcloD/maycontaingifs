date := 16

# imagemagick calculates delay as 100ths of seconds.
# thus, 100/30 is 30fps.
# 100/30 = 3.3333...
delay := 3.3333

vpath %.gif out
vpath %.go may

default:
	@tmux_send 'make display'

display: may$(date).gif
	@img_view $<

launch:
	@img_view may$(date).gif

out/may$(date).gif: clean frames
	@convert -delay $(delay) frames/*.png $@
	@echo "\nComplete: $@"

clean:
	@rm -rf frames

frames: gifs gifs.go may$(date).go
	@echo may/may$(date).gif
	@echo Creating frames...
	@mkdir frames
	@go build gifs.go
	@./gifs $(date)

gifs:
	@go build gifs.go

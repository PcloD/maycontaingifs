date := 24

gif := out/may$(date).gif
gofile := may/may$(date).go

# imagemagick calculates delay as 100ths of seconds.
# thus, 100/30 is 30fps.
# 100/30 = 3.3333...
delay := 3.3333


default:
	@tmux_send 'make display'

clean:
	@rm -rf frames

display: $(gif)
	@img_view $<

launch:
	@img_view $(gif)

$(gif): gifs.go $(gofile)
	@echo target: $(date)
	@echo Creating frames...

	@# clean/create frames dir
	@rm -rf frames
	@mkdir frames

	@# build executable and run it
	@go build gifs.go
	@./gifs $(date)

	@# turn the frames into an animation
	@echo "\nMaking the animation..."
	@convert -delay $(delay) frames/*.png $@
	@echo "Complete: $@"


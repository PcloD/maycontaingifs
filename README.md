# maycontaingifs
Source code for my entries to https://twitter.com/hashtag/MayContainGIFs?src=hash

To compile any of this you'll need my `blgo` library, as well as `go-cairo` and the Cairo graphics libs installed. Instructions here:

https://github.com/bit101/blgo

Then check the makefile. It determines which gif will be built. Each day's code generates a folder full of pngs. These are turned into an animated gif using ImageMagick.

`tmux_send` is a tmux script that allows me to run the makefile's `build` target in a separate tmux pane from Vim.

https://github.com/bit101/dotfiles/blob/master/tmux/tmux_send

# maycontaingifs
Source code for my entries to https://twitter.com/hashtag/MayContainGIFs?src=hash

To compile any of this you'll need my `blgo` library, as well as `go-cairo` and the Cairo graphics libs installed. Instructions here:

https://github.com/bit101/blgo

Then check the makefile. It determines which gif will be built. Each day's code generates a folder full of pngs. These are turned into an animated gif using ImageMagick.

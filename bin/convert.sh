#!/bin/bash

if [[ $1 == *.pdf ]]; then
	pdftoppm -png -scale-to-x 1920 -scale-to-y 1080 "$1" "$2"
	rm -f "$1"
else
	convert "$1" -resize 1920x1080 -background black -gravity center -extent 1920x1080 "$1"
fi

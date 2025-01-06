#!/usr/bin/bash

# Run a local jekyll
# website with one command!

main () {
    jekyll build
    jekyll serve
}

main

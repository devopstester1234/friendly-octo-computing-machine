SHELL := /usr/bin/env bash
add-drone: sign
	git add .
sign:
	source .env && /usr/local/bin/drone sign devopstester1234/friendly-octo-computing-machine --save
install: setup
	go get -u github.com/mtib/mtiblib/probability
	go get -u github.com/mtib/mtiblib/numbers
	go get -u github.com/mtib/mtiblib/simplehttp

setup:
	git pull --ff-only
	git submodule init
	git submodule update

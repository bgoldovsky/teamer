.PHONY:compose
compose:
	 docker-compose build && docker-compose up

.PHONY:build
build:
	 docker-compose build

.PHONY:up
up:
	 docker-compose up

.DEFAULT_GOAL := compose
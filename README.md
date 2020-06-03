# [Quiet Hacker News](https://quiethackernews.com)

[![License](https://img.shields.io/github/license/fuzzingbits/quiet-hacker-news)](https://github.com/fuzzingbits/quiet-hacker-news/blob/master/LICENSE)
[![Docker Image](https://img.shields.io/badge/container-Docker-blue)](https://hub.docker.com/r/fuzzingbits/quiet-hacker-news)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuzzingbits/quiet-hacker-news)](https://goreportcard.com/report/github.com/fuzzingbits/quiet-hacker-news)
[![Go](https://github.com/fuzzingbits/quiet-hacker-news/workflows/Go/badge.svg)](https://github.com/fuzzingbits/quiet-hacker-news/actions)
[![Coverage Status](https://coveralls.io/repos/github/fuzzingbits/quiet-hacker-news/badge.svg?branch=master)](https://coveralls.io/github/fuzzingbits/quiet-hacker-news?branch=master)

A quieter approach to Hacker News. Links and nothing else.

![screenshot](/ops/screenshot.png)

## What is it?
- Just a list of the top 30 links from Hacker News
- Updates about once an hour

## Extra Features
- Uses `prefers-color-scheme` for light and dark mode
- Works well on small screens

## Usage
- From Source:
    - Build binary from source: `make`
    - Run for local development: `make dev`
    - See [Makefile](/Makefile) for additional targets
- From Docker Image:
    - `docker run -p 9090:9090 fuzzingbits/quiet-hacker-news`

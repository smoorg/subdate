# Subtime

A simple utility to subtract two dates in "HH:MM:SS" format. It returns absolute difference between the two timestamps and returns it to the console. It has no dependencies other than built in packages and does the job done.

## Installation

To install it on your machine simply build it

```sh
$ go build .
```

To make it available to your system you might just copy it over to `/usr/local/bin` so that your `$PATH` recognizes it.

## Usage

In order to use simply use `d1` and `d2` flags like so:

```sh
$ subtime -d1 12:43:22 -d2 15:04:01
3:39:21
```

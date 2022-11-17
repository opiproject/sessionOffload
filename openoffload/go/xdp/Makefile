SHELL=/bin/bash

KERN_HEADERS=kern/xdp_maps.h \
             kern/xdp_parser.h

GOEBPF_SHA=398a646c8fc1acc65194af3888e88718f83f01f2

MACHINE := $(shell uname -m)

ifeq ($(MACHINE),x86_64)
	MACHINE_INCLUDE := -I/usr/include/x86_64-linux-gnu
endif

ifeq ($(MACHINE),aarch64)
	MACHINE_INCLUDE := -I/usr/include/aarch64-linux-gnu/asm -I/usr/include/aarch64-linux-gnu
endif

all: loader/xdploader kern/xdp.elf

goebpf:
	git clone https://github.com/dropbox/goebpf.git
	cd goebpf && git checkout 398a646c8fc1acc65194af3888e88718f83f01f2

kern/xdp.elf: kern/xdp.c $(KERN_HEADERS) goebpf
	clang-9 -g -O2 -target bpf -I/opt/xdp/goebpf $(MACHINE_INCLUDE) -c kern/xdp.c -o xdp.elf

loader/xdploader: loader/main.go goebpf
	cd loader && go build -ldflags="-s -w" -o xdploader


#!/bin/bash
# Copyright (C) 2022 Intel Corporation, or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

# Turn on debug
set -x

# Create dummy device
ip link add dummy0 type dummy

# Set MTUs
ip link set dummy0 mtu 9000

# Turn off TSO
ethtool -K dummy0 tso off

# Mount bpffs
mount bpffs /sys/fs/bpf -t bpf

# Load the XDP programs
/opt/xdp/xdploader -file /opt/xdp/xdp.elf -i dummy0

# Start the bridge
/opi-sessionoffload-bridge -port=50151 -simulate 7

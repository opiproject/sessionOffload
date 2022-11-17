/* SPDX-License-Identifier: (GPL-2.0-or-later OR BSD-3-clause) */
/*
 * Copyright (c) 2022, Intel Corporation, or its subsidiaries.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 * * Neither the name of the copyright holder nor the names of its
 *   contributors may be used to endorse or promote products derived from
 *   this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

#ifndef __XDP_SESSION_H
#define __XDP_SESSION_H

#include "bpf_helpers.h"

/* IP Version */
enum ip_version {
	SESSION_IPV4 = 0,
	SESSION_IPV6,
};

/* Action */
enum action {
    _DROP = 0,
    _FORWARD,
    _MIRROR,
    _SNOOP,
};

struct action_s {
    enum action action_type;
    __be32      action_next_hop;
    __be32      action_next_hop_v6[4];
};

enum session_state {
    _ESTABLISHED = 0,
    _CLOSING_1,
    _CLOSING_2,
    _CLOSED,
    _UNKNOWN_STATE,
};

enum session_close_code {
    _NOT_CLOSED = 0,
    _FINACK,
    _RST,
    _TIMEOUT,
    _UNKNOWN_CLOSE_CODE,
};

struct session {
    __u32           in_lif;
    __u32           out_lif;
    enum ip_version ip_ver;
    struct {
        __be32 saddr;
        __be32 daddr;
        __be16 sport;
        __be16 dport;
    } ipv4;
    struct {
        __be32 saddr[4];
        __be32 daddr[4];
        __be16 sport;
        __be16 dport;
    } ipv6;
    __u32                   protocol;
    struct action_s         action;
    /* Cache Timeout: cache_timeout      uint32 */
    __u64                   in_packets;
    __u64                   out_packets;
    __u64                   in_bytes;
    __u64                   out_bytes;
    enum session_state      state;
    enum session_close_code close_code;
    /* Start Time: start_time         time.Time */
    /* End Time: end_time           time.Time */
};

#endif /* __XDP_SESSION_H */

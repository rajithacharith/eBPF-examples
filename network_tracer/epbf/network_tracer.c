#ifndef __u16
typedef unsigned short u16;
#endif
#ifndef __u32
typedef unsigned int u32;
#endif
#ifndef PT_REGS_PARM1
# if defined(__x86_64__)
#  define PT_REGS_PARM1(x) ((x)->di)
# elif defined(__aarch64__)
#  define PT_REGS_PARM1(x) ((x)->regs[0])
# else
#  define PT_REGS_PARM1(x) 0
# endif
#endif
#ifndef __bpf_ntohs
#define __bpf_ntohs(x) __builtin_bswap16(x)
#endif

struct sock_common {
    unsigned short skc_family;
    unsigned short skc_dport;
    unsigned int skc_daddr;
};
struct sock {
    struct sock_common __sk_common;
};
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/tcp.h>
#include <linux/ptrace.h>
#include <linux/in.h>
#include <linux/socket.h>

SEC("kprobe/tcp_v4_connect")
int trace_connect(struct pt_regs *ctx) {
    struct sock *sk = (struct sock *)PT_REGS_PARM1(ctx);
    u16 dport = 0;
    u32 daddr = 0;
    bpf_probe_read_kernel(&dport, sizeof(dport), &sk->__sk_common.skc_dport);
    bpf_probe_read_kernel(&daddr, sizeof(daddr), &sk->__sk_common.skc_daddr);
    dport = __bpf_ntohs(dport);
    u32 pid = bpf_get_current_pid_tgid() >> 32;
    bpf_printk("tcp connect: pid=%d daddr=%u dport=%u\n", pid, daddr, dport);
    return 0;
}

char LICENSE[] SEC("license") = "GPL";

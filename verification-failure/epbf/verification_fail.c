#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>


SEC("tracepoint/syscalls/sys_enter_execve")
int verifier_fail(void *ctx) {
    char buf[8];
    // Out-of-bounds stack access: always rejected by verifier
    buf[100] = 1;
}

char LICENSE[] SEC("license") = "GPL";

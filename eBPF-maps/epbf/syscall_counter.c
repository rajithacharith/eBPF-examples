#ifndef u32
typedef unsigned int u32;
#endif
#ifndef u64
typedef unsigned long long u64;
#endif
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1);
    __type(key, u32);
    __type(value, u64);
} execve_count SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_execve")
int count_execve(struct trace_event_raw_sys_enter *ctx) {
    u32 key = 0;
    u64 *val = bpf_map_lookup_elem(&execve_count, &key);
    u64 one = 1;
    if (val) {
        __sync_fetch_and_add(val, 1);
    } else {
        bpf_map_update_elem(&execve_count, &key, &one, BPF_ANY);
    }
    return 0;
}

char LICENSE[] SEC("license") = "GPL";

#ifndef __TRACE_EVENT_RAW_SCHED_SWITCH_DEFINED
#define __TRACE_EVENT_RAW_SCHED_SWITCH_DEFINED
struct trace_event_raw_sched_switch {
    unsigned long long pad;
    int prev_pid;
    int prev_prio;
    long prev_state;
    int next_pid;
};
#endif
#ifndef u32
#define u32 unsigned int
#endif
#ifndef u64
#define u64 unsigned long long
#endif
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/sched.h>

struct key_t {
    u32 pid;
    char comm[16];
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, struct key_t);
    __type(value, u64);
} cpu_time SEC(".maps");

SEC("tracepoint/sched/sched_switch")
int on_cpu_switch(struct trace_event_raw_sched_switch *ctx) {
    u32 prev_pid = ctx->prev_pid;
    u32 next_pid = ctx->next_pid;
    u64 ts = bpf_ktime_get_ns();
    struct key_t key = {};
    key.pid = prev_pid;
    bpf_get_current_comm(&key.comm, sizeof(key.comm));
    u64 *val = bpf_map_lookup_elem(&cpu_time, &key);
    if (val) {
        *val += ts;
    } else {
        bpf_map_update_elem(&cpu_time, &key, &ts, BPF_ANY);
    }
    return 0;
}

char LICENSE[] SEC("license") = "GPL";

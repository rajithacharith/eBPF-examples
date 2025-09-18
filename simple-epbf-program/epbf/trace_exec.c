#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

SEC("tracepoint/syscalls/sys_enter_execve")
int trace_exec(void *ctx) {
	unsigned long pid_tgid = bpf_get_current_pid_tgid();
	unsigned int pid = pid_tgid >> 32;
	char comm[16];
	bpf_get_current_comm(&comm, sizeof(comm));
	bpf_printk("execve: pid=%d comm=%s\n", pid, comm);
	return 0;
}

char LICENSE[] SEC("license") = "GPL";

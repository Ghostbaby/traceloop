#ifndef __STRACEBACK_TAILCALL_BPF_H
#define __STRACEBACK_TAILCALL_BPF_H

#include <linux/types.h>

#ifndef TASK_COMM_LEN
#define TASK_COMM_LEN 16
#endif

#ifndef __NR_openat
#define __NR_openat 257
#endif

#ifndef PIN_GLOBAL_NS
#define PIN_GLOBAL_NS 2
#endif

#define PARAM_LEN 128
/* reexport for Cgo */
const __u64 PARAM_LENGTH = PARAM_LEN;
/* special values used to refer to dynamic length */
const __u64 USE_NULL_BYTE_LENGTH = 0xffffffffffffffffULL;
const __u64 USE_RET_AS_PARAM_LENGTH = 0xfffffffffffffffeULL;
/* INDEX(x) is not defined (Cgo cannot access macros),
 * use bit arithmetic with 0xf to get value and addition to generate */
const __u64 USE_ARG_INDEX_AS_PARAM_LENGTH = 0xfffffffffffffff0ULL;

#define SYSCALL_EVENT_TYPE_ENTER 0
#define SYSCALL_EVENT_TYPE_EXIT  1
#define SYSCALL_EVENT_TYPE_CONT  2

struct syscall_event_t {
	__u64 timestamp;
	__u64 typ;

	/* how many syscall_event_cont_t messages to expect after */
	__u16 cont_nr;
	__u16 cpu;
	__u32 id;
	__u64 pid;
	char comm[TASK_COMM_LEN];
	__u64 args[6];
	/* __u64 ret stored in args[0] */
};

struct syscall_event_cont_t {
	__u64 timestamp;
	__u64 typ;
	__u64 index;
	__u64 length;

	char param[PARAM_LEN];
	__u8 failed;
};

struct syscall_def_t {
	__u64 args_len[6];
};

#endif

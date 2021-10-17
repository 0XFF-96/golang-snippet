# hacking document code 

### 目的

This is a living document and at times it will be out of date. 
It is intended to articulate how programming in the Go runtime differs from
writing normal Go. 

It focuses on 【pervasive concepts】 rather than
details of particular interfaces.


### 概念

1. Gs, Ms, and Ps. It's important to understand these 
even if you're not working on the scheduler.

2. A "G" is simply a goroutine. It's represented by type `g`.

3. An "M" is an OS thread that can be executing user Go code.

4. Finally, a "P" represents the resources required to execute user Go
   code, such as scheduler and memory allocator state.

5. User stacks and system stacks

6. `getg()` and `getg().m.curg`

### scheduler's purpose 编译器的目的

The scheduler's job is to match up a G (the code to execute), an M
(where to execute it), and a P (the rights and resources to execute
it). When an M stops executing user Go code,


### 为什么  G,M,P 都分配在堆里面？

All `g`, `m`, and `p` objects are heap allocated, but are never freed,
so their memory remains type stable. As a result, the runtime can
avoid write barriers in the depths of the scheduler.


### 栈空间的分配（什么是用户栈和什么是系统栈？）

1. Every non-dead G has a *user stack* associated with it, which is what
   user Go code executes on. User stacks start small (e.g., 2K) and grow
   or shrink dynamically.
   
2. Every M has a *system stack* associated with it (also known as the M's
   "g0" stack because it's implemented as a stub G) and, on Unix
   platforms, a *signal stack* (also known as the M's "gsignal" stack).


3. Runtime code often temporarily switches to the system stack using
   `systemstack`, `mcall`, or `asmcgocall` to perform tasks that must not
   be preempted, that must not grow the user stack, or that switch user
   goroutines.
(不懂，为什么系统栈是不可抢占和不能 grow 的呢？ 切换过程是如何进行？)


### 代码层面上

1. To get the current user `g`, use `getg().m.curg`.


2. To determine if you're running on the user stack or the system stack,
   use `getg() == getg().m.curg`.
   
   
### Error 相关 (Error handling and reporting)

1. Errors that can reasonably be recovered from in user code should use
   `panic` like usual.
   
2. panic != fatal 

3. For runtime error debugging, it's useful to run with
   `GOTRACEBACK=system` or `GOTRACEBACK=crash`.

(什么情况会出现 runtime crash ?)

### runtime 的同步机制

> The runtime has multiple synchronization mechanisms.

2. The simplest is `mutex`, which is manipulated using `lock` and
   `unlock`.
   
3. For one-shot notifications, use `note`, which provides `notesleep` and
   `notewakeup`. Unlike traditional UNIX `sleep`/`wakeup`, `note`s are
   race-free, so `notesleep` returns immediately if the `notewakeup` has
   already happened.
(不懂， 什么是 note 机制？)

### 什么是 gopark 和 goready ?

1. `gopark` parks the current goroutine—putting it in the
   "waiting" state and removing it from the scheduler's run queue—and
   schedules another goroutine on the current M/P. 
   
2. `goready` puts a
   parked goroutine back in the "runnable" state and adds it to the run
   queue.

### runtime 中的原子操作, 回避使用 atmoic 的理由？

The runtime uses its own atomics package at `runtime/internal/atomic`.
This corresponds to `sync/atomic`,

In general, we think hard about the uses of atomics in the runtime and
try to avoid unnecessary atomic operations.

### Unmanaged memory (下次继续)
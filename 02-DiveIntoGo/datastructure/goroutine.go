package datastructure

type g struct {

	// goroutine 使用的栈
	stack       stack   // offset known to runtime/cgo
	// 用于栈的扩张和收缩检查，抢占标志
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic         *_panic // innermost panic - offset known to liblink
	_defer         *_defer // innermost defer
	// 当前与 g 绑定的 m
	m              *m      // current m; offset known to arm liblink
	// goroutine 的运行现场
	sched          gobuf
	syscallsp      uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc      uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp       uintptr        // expected sp at top of stack, to check in traceback
	// wakeup 时传入的参数
	param          unsafe.Pointer // passed parameter on wakeup
	atomicstatus   uint32
	stackLock      uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid           int64
	// g 被阻塞之后的近似时间
	waitsince      int64  // approx time when the g become blocked
	// g 被阻塞的原因
	waitreason     string // if status==Gwaiting
	// 指向全局队列里下一个 g
	schedlink      guintptr
	// 抢占调度标志。这个为 true 时，stackguard0 等于 stackpreempt
	preempt        bool     // preemption signal, duplicates stackguard0 = stackpreempt
	paniconfault   bool     // panic (instead of crash) on unexpected fault address
	preemptscan    bool     // preempted g does scan for gc
	gcscandone     bool     // g has scanned stack; protected by _Gscan bit in status
	gcscanvalid    bool     // false at start of gc cycle, true if G has not run since last scan; TODO: remove?
	throwsplit     bool     // must not split stack
	raceignore     int8     // ignore race detection events
	sysblocktraced bool     // StartTrace has emitted EvGoInSyscall about this goroutine
	// syscall 返回之后的 cputicks，用来做 tracing
	sysexitticks   int64    // cputicks when syscall has returned (for tracing)
	traceseq       uint64   // trace event sequencer
	tracelastp     puintptr // last P emitted an event for this goroutine
	// 如果调用了 LockOsThread，那么这个 g 会绑定到某个 m 上
	lockedm        *m
	sig            uint32
	writebuf       []byte
	sigcode0       uintptr
	sigcode1       uintptr
	sigpc          uintptr
	// 创建该 goroutine 的语句的指令地址
	gopc           uintptr // pc of go statement that created this goroutine
	// goroutine 函数的指令地址
	startpc        uintptr // pc of goroutine function
	racectx        uintptr

	waiting        *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt        []uintptr      // cgo traceback context
	labels         unsafe.Pointer // profiler labels
	// time.Sleep 缓存的定时器
	timer          *timer         // cached timer for time.Sleep

	gcAssistBytes int64
}

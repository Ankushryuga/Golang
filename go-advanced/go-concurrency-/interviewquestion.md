A. Goroutine Behavior (15)

What happens if you don’t WaitGroup.Wait() after launching goroutines?

The program may terminate before goroutines finish, resulting in partial or no output.

What does go keyword actually do?

It creates a new lightweight thread (goroutine) and schedules it for concurrent execution.

Will a goroutine always run in order? Why not?

No. Goroutine execution order is non-deterministic due to Go’s scheduler.

How to pass a loop variable safely to a goroutine?

Pass it as an argument: go func(i int) { ... }(i) to avoid closure capture bugs.

What happens when main() finishes before goroutines?

The program exits immediately, killing all running goroutines.

Are goroutines OS threads?

No. They are user-space constructs managed by Go runtime.

What is the stack size of a goroutine?

Starts small (2KB) and grows/shrinks dynamically.

How does Go scheduler handle thousands of goroutines?

With M:N scheduling (many goroutines on fewer OS threads), preemptively managed by the runtime.

What’s a goroutine leak and how to detect it?

A goroutine that never finishes. Use pprof, runtime.NumGoroutine(), or code audits to detect.

Is go go(), go fmt.Println(), and go wg.Done() legal?

Yes. You can use go before any function call.

Can you return a value from a goroutine?

Not directly. Use channels to send return values.

What happens when a goroutine panics?

It crashes that goroutine. Use recover() to handle it.

Can a goroutine be cancelled?

Yes, using context.Context cancellation.

What is preemption in Go?

Forced scheduling switch to prevent goroutines from blocking CPU.

How to limit goroutine concurrency using a pool?

Use buffered channels as semaphores.

📂 B. Channels Basics (15)

What happens when you send on a nil channel?

Blocks forever.

What happens when you receive on a nil channel?

Blocks forever.

What happens if you send on a closed channel?

Panics: "send on closed channel"

What’s the difference between buffered and unbuffered channels?

Unbuffered blocks until receiver is ready; buffered does not block until full.

How does a channel block?

Send blocks until a receiver is ready; receive blocks until a sender sends.

How to detect if a channel is closed?

Use v, ok := <-ch. If ok is false, the channel is closed.

How to use range on a channel?

Loops until the channel is closed.

How to implement a timeout with time.After()?

Use select { case <-ch: ...; case <-time.After(time.Second): ... }

How to send multiple values concurrently using a channel?

Use multiple goroutines sending to the same channel.

Can a channel be shared across multiple producers and consumers?

Yes, channels are safe for concurrent use.

What happens when multiple goroutines write to the same channel?

Go handles it; first one that gets scheduled wins.

Can you close a channel from the receiver side?

No. Only the sender should close a channel.

Why shouldn’t you close a channel twice?

It causes a panic: "close of closed channel"

How to simulate a broadcast channel?

Use multiple channels or fan-out pattern.

How to safely close a channel?

Ensure only one sender closes it, typically after a loop ends.

✨ C. Select Statement (10)

What is the select statement used for?

To wait on multiple channel operations.

How to implement a timeout with select?

Use select { case <-ch: ...; case <-time.After(...) }

What does default do in select?

It executes if no other cases are ready (non-blocking select).

What happens if multiple cases in select are ready?

One is picked randomly.

Can you use select with only one case?

Yes, though it's not very useful unless used with timeout or default.

How to use select to avoid blocking?

Use default: case.

How to use select with infinite loop to listen for multiple channels?

Use for { select { ... } } loop.

How to terminate select with a done channel?

Add case <-done: return to the select.

What’s the difference between select and switch?

select is for channels; switch is for any values.

Can select be used for multiplexing multiple input sources?

Yes. That’s its primary use-case.

🔐 D. sync.Mutex & sync.RWMutex (10)

What is the purpose of sync.Mutex?

To ensure mutual exclusion when accessing shared data.

What happens if you unlock a mutex that isn’t locked?

It causes a runtime panic.

Can a goroutine lock a mutex multiple times?

No, it would deadlock itself. Use sync.RWMutex with care if you need multiple locks.

What’s the difference between sync.Mutex and sync.RWMutex?

RWMutex allows multiple readers or one writer, but not both simultaneously.

When should you use RLock vs Lock?

Use RLock when reading shared data; Lock when writing.

How do mutexes prevent race conditions?

By ensuring only one goroutine accesses a critical section at a time.

What happens if you forget to unlock a mutex?

Other goroutines will block forever, causing a deadlock.

Can mutexes cause deadlocks?

Yes. If not carefully ordered or released.

Can you use defer with mutexes?

Yes. It's a best practice to avoid forgetting to unlock.

What’s the difference between sync.Mutex and sync.Once?

sync.Once ensures a function runs only once, even across multiple goroutines.

⏱ E. sync.WaitGroup (10)

What happens if you call wg.Done() more times than Add()?

Panics with "negative WaitGroup counter".

Can wg.Add() be called after wg.Wait()?

No. It causes a race condition.

What if wg.Wait() is never reached?

Main might exit early or hang, depending on context.

What’s the consequence of forgetting wg.Done()?

WaitGroup.Wait() blocks forever.

Can WaitGroup be reused after waiting?

Yes, but not concurrently during wait.

How to use WaitGroup with multiple loops?

Call Add() outside loops; use Done() inside goroutines.

Can you use WaitGroup to synchronize single goroutines?

Yes, even one is fine.

What is a correct pattern to parallelize N jobs?

Loop with wg.Add(1) + go func() { defer wg.Done() ... }()

Can a single WaitGroup be shared across functions?

Yes, pass it as a pointer.

Is WaitGroup thread-safe?

Yes, as long as you follow correct usage.


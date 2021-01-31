### Context 使用的基本原则

1. Incoming requests to a server should create a Context.
2. Outgoing calls to servers should accept a Context.
3. Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
4. The chain of function calls between them must propagate the Context.
5. Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
6. When a Context is canceled, all Contexts derived from it are also canceled.
7. The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
8. Do not pass a nil Context, even if a function permits it. Pass a TODO context if you are unsure about which Context to use.
9. Use context values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.


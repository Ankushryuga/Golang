## How panic, recover, and defer work together?
- **defer** statements are executed in LIFO order.
- A **panic** will execute all deferred functions before the program crashes.
- **recover()** must be called within a deferred function to stop the panic.

## When to use panic?
- Truly unrecoverable situations(e.g. corrupt memory, logic invariant failure)
- programming errors(e.g., nil pointer dereferece)
- internal bugs or contract violated.
- library code

## When not to use:
  - **Input validation** (use errors instead)
  - **Expected failure**(use error return values)
  - in **control flow** (it's not try/catch)


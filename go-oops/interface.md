## Non-concrete Interface:
    => Non-concrete interface: they typically mean an interface that:
      Does not have a specific, concrete implementation at the time of definition.
      It's an abstract definition of behaviour, not tied to any specific struct or type.


## What is interface?
    => An interface in Go defines a set of methods signatures, Any type that implements those methods implicitly satisfies the interface:
      type Shape interface{
        Area() float64
      }

      Here:, Shape is an interface, it is non-concrete because:
      1. It doesn't describe any data or implementation.
      2. It only specifies behavior that some concrete type like (Circle, Rectangle) must provide..


# 🎯 Concrete vs. Non-Concrete Interface

| Type                       | Description                          | Example                                 |
| -------------------------- | ------------------------------------ | --------------------------------------- |
| **Concrete Type**          | Has a full implementation and data   | `struct`, `int`, `string`, etc.         |
| **Non-Concrete Interface** | Abstract behavior, no implementation | `io.Reader`, `error`, custom interfaces |


# Example:
    type Reader interface {
    Read(p []byte) (n int, err error)
}
This Reader interface is non-concrete — it doesn't say how reading happens.

Now, a concrete type that implements it:

type File struct {}

func (f File) Read(p []byte) (int, error) {
    // Implementation here
    return len(p), nil
}



type File struct {}

func (f File) Read(p []byte) (int, error) {
    // Implementation here
    return len(p), nil
}

//Now File is a concrete type that implements the non-concrete interface Reader.



✅ Why Use Non-Concrete Interfaces?
Flexibility: Allows writing functions that work with any type that satisfies the interface.

Abstraction: Decouple function logic from specific types.

Testing: You can easily mock interfaces in tests.


🧠 Quick Analogy
Non-concrete interface: Like a job description ("must be able to write code").

Concrete type: Like a specific employee who knows how to write code in Go.

🔒 Gotchas
Avoid empty interfaces (interface{}) unless truly necessary — they are non-concrete, but too generic (like any).

If an interface has only one or two implementations, consider whether it's worth abstracting.


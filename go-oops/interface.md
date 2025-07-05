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



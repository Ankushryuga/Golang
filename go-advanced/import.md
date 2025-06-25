# Go Import:
    =>
    it tells go compiler that, i want to use this package in my program.


# import multiple package:
    =>
    import (
        "fmt"
        "math"
    )

# Aliasing imports:
    =>
    import (
        f "fmt"
        m "math"
    )

# Dot import:
    => it allows you to use the package's exported identifiers without qualifying them with the package name.
    NOTE: Dot imports are discouraged., it make code harder to read and understand.


# Blank Import:
    =>
    package that is being used for its side effects (e.g., initialization functions). In such cases, you can use blank identifiers: _
    # 
    import _ "image/png"

    # This is often used when the package registers itself with the Go runtime, such as image format handlers.


/**
defer: keyword is used to delay the execution of a function untill the surrounding function completes,
the deferred function calls are executed in LIFO order, that means the most recently deferred function
is executed first, followed by the second most recently deferred function, and so on.

## syntax:
defer functionName(arguments)
*/

package deferExample

import "fmt"

func DeferExample(name string) {
	fmt.Println("Hello ", name)
}

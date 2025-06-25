/**
Errorf(): formatting the error message using Errorf:
This function is present inside the fmt package, so we can directly use this if we have imported the fmt package.
*/

package deferExample

import "fmt"

func divide(num1, num2 int) error {
	if num2 == 0 {
		return fmt.Errorf("%d / %d \n cannot divide a number by 0 ", num1, num2)
	}
	return nil
}

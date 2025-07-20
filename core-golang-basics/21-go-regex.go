/***
| Pattern  | Meaning                                  |        |
| -------- | ---------------------------------------- | ------ |
| `.`      | Any character                            |        |
| `\d`     | Digit                                    |        |
| `\w`     | Word character (letter/digit/underscore) |        |
| `\s`     | Whitespace                               |        |
| `^`      | Start of string                          |        |
| `$`      | End of string                            |        |
| \`a      | b\`                                      | a or b |
| `[abc]`  | a, b, or c                               |        |
| `[^abc]` | Not a, b, or c                           |        |
| `(abc)`  | Capture group                            |        |
| `a*`     | 0 or more a                              |        |
| `a+`     | 1 or more a                              |        |
| `a{2,4}` | Between 2 and 4 a’s                      |        |


## Validate Email
emailRe := regexp.MustCompile(`^[\w._%+\-]+@[\w.\-]+\.[a-zA-Z]{2,}$`)
fmt.Println(emailRe.MatchString("hello@example.com")) // true

## Split by Regex.
re := regexp.MustCompile(`[,\s]+`)
result := re.Split("one, two three\tfour", -1)
fmt.Println(result) // ["one" "two" "three" "four"]



*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "My phone numbers are 123-456-7890 and 987-654-3210."

	// Compile regex
	re := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)

	// Match?
	fmt.Println("MatchString:", re.MatchString(text))

	// First match
	fmt.Println("FindString:", re.FindString(text))

	// All matches
	fmt.Println("FindAllString:", re.FindAllString(text, -1))

	// Replace with placeholder
	masked := re.ReplaceAllString(text, "[REDACTED]")
	fmt.Println("Masked:", masked)

	// Using capturing groups
	emailRe := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
	email := "Contact me at user@example.com"
	matches := emailRe.FindStringSubmatch(email)
	if len(matches) > 0 {
		fmt.Println("Full match:", matches[0])
		fmt.Println("Username:", matches[1])
		fmt.Println("Domain:", matches[2])
	}
}

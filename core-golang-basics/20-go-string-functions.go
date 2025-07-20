/***

| Section             | Functions Covered                                                 |
| ------------------- | ----------------------------------------------------------------- |
| **Search**          | `Contains`, `HasPrefix`, `HasSuffix`, `Index`, `LastIndex`        |
| **Modify**          | `ToLower`, `ToUpper`, `TrimSpace`, `Trim`, `ReplaceAll`, `Repeat` |
| **Split/Join**      | `Split`, `Join`, `Fields`                                         |
| **Compare**         | `Compare`, `EqualFold`                                            |
| **String ⇄ Number** | `Atoi`, `Itoa`, `ParseFloat`, `FormatFloat`                       |
| **Unicode Safe**    | Iteration with `range` (handles emojis and UTF-8)                 |


*/

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("===== STRING FUNCTIONS DEMO =====")

	// -------- SEARCH / CHECK --------
	text := "  GoLang is awesome!  "
	fmt.Println("\n🔍 SEARCH / CHECK")

	fmt.Println("Original text:", text)
	fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))
	fmt.Println("HasPrefix '  Go':", strings.HasPrefix(text, "  Go"))
	fmt.Println("HasSuffix '!  ':", strings.HasSuffix(text, "!  "))
	fmt.Println("Index of 'Lang':", strings.Index(text, "Lang"))
	fmt.Println("LastIndex of 'e':", strings.LastIndex(text, "e"))

	// -------- MODIFY / TRANSFORM --------
	fmt.Println("\n✂️ MODIFY / TRANSFORM")

	fmt.Println("ToLower:", strings.ToLower(text))
	fmt.Println("ToUpper:", strings.ToUpper(text))
	fmt.Println("TrimSpace:", strings.TrimSpace(text))
	fmt.Println("Trim '<> Go <>':", strings.Trim("<> Go <>", "<> "))
	fmt.Println("ReplaceAll 'Go-Go-Go':", strings.ReplaceAll("Go-Go-Go", "Go", "Run"))
	fmt.Println("Repeat 'ha' x 3:", strings.Repeat("ha", 3))

	// -------- SPLIT / JOIN --------
	fmt.Println("\n📏 SPLIT / JOIN")

	data := "a,b,c,d"
	split := strings.Split(data, ",")
	fmt.Println("Split:", split)
	fmt.Println("Join:", strings.Join(split, "-"))
	fmt.Println("Fields (whitespace split):", strings.Fields("a b\tc\n d"))

	// -------- COMPARE --------
	fmt.Println("\n🔠 COMPARE")

	fmt.Println("Compare 'apple' vs 'banana':", strings.Compare("apple", "banana")) // -1
	fmt.Println("EqualFold 'Go' vs 'go':", strings.EqualFold("Go", "go")) // true

	// -------- CONVERT STR <-> NUM --------
	fmt.Println("\n🔢 STRCONV (STRING <-> NUMBER)")

	// String to int
	strInt := "42"
	num, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Println("Atoi failed:", err)
	} else {
		fmt.Println("Atoi:", num)
	}

	// Int to string
	intVal := 99
	strFromInt := strconv.Itoa(intVal)
	fmt.Println("Itoa (int to string):", strFromInt)

	// Parse float
	strFloat := "3.14159"
	floatVal, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("ParseFloat failed:", err)
	} else {
		fmt.Println("Parsed float64:", floatVal)
	}

	// Float to string
	fmt.Println("FormatFloat (2 decimal):", strconv.FormatFloat(floatVal, 'f', 2, 64))

	// -------- CHARACTER ITERATION --------
	fmt.Println("\n🔤 RUNE ITERATION (UTF-8 SAFE)")

	emojiStr := "Go😊Lang"
	for i, r := range emojiStr {
		fmt.Printf("Index %d: Rune %q\n", i, r)
	}

	fmt.Println("\n✅ Demo complete!")
}

# Approach 1: Using bufio.NewScanner
The first approach is to use the bufio package and its NewScanner function to read the file word by word. This function takes an io.Reader interface as input and returns a Scanner object that can be used to scan the file word by word.
    => 
      package main
      import (
         "bufio"
         "fmt"
         "os"
      )
      
      func main() {
         file, err := os.Open("filename.txt")
         if err != nil {
            panic(err)
         }
         defer file.Close()
         
         scanner := bufio.NewScanner(file)
         
         // Set the split function for the scanning operation.
         scanner.Split(bufio.ScanWords)
         
         // Scan all words from the file.
         for scanner.Scan() {
            fmt.Println(scanner.Text())
         }
         
     if err := scanner.Err(); err != nil {
        panic(err)
     }
  }

# Approach 2: Using fmt.Fscanf
      =>
      The second approach is to use the fmt package and its Fscanf function to read the file word by word. This function takes an io.Reader interface as input and returns the number of items successfully scanned.

# Approach 3: Using bufio.NewReader and strings.Fields
      => 
      The third approach is to use the bufio package and its NewReader function along with the strings package and its Fields function to read the file word by word. The NewReader function creates a new buffered reader, and the Fields function splits the string into words.


# Approach 4: Using ioutil.ReadFile and strings.Fields


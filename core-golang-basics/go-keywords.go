/***
  **new** keyword is a built-in function used to allocate memory for a new value of a specified type.
  ## Purpose:
    - new(T) allocates memory for a variable of type T, initialize it to the zero value of 
    returns a pointer to it (*T).

  ## Key Points:
    - The memory allocated by new is zeroed.
    - It is often used when you need a pointer to a value but don't want to manage memory manually
    - For slices, maps, and channels, the make function is preferred as it initializes the underlying data structures.
    
*/

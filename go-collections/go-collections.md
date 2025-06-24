# Collections in Golang:

# Arrays: 
        => 
        fixed-size, homogenous data structure that store elements of the same type.
        =>
        example:
        var myArray [10]int //declares an array of integers with a size of 5.

        # Initialize an array with values:
        myArray := [5]int{1,2,3,4,5}



# Slices: 
        =>
        Slices are dynamic, resizable arrays that can grow or shrink in size. A slice is a reference to an underlying array, which means that any changes made to the slice will also affect the original array.
        #### To create a slice use make function:
        ## Example:
        
        mySlice := make([]int, 5)   //creates a slice of integers with a length of 3.

        ### Create a slice from an existing array:
        myArray := [5]int{1,2,3,4,5}
        nySlice := myArray[1:4] // creates a slice containing elements at index 1 to 3.

        ### adding elements to a slice can be done using the append function:
        mySlice := []int{1,2,3}
        mySlice = append(mySlice, 4)    //myslice now contains [1,2,3,4]


# Maps:
        => 
        maps in go are unordered collections of key-value pairs, similar to java's HashMap. Maps are created using the make functino:

        myMap := make(map[string]int) // creates a map with string keys and integer values

        # you can add or update elements in a map using the following syntax:

        myMap["one"] = 1
        myMap["two"] = 2

        ## Retrieving an element from map is done by specifying its key:
        => 
        value := myMap["one"]   //1

        # you can also delete elements from a map using the delete function:

        delete(myMap, "one")//removes the key "one" and its associated value from the map.


## Custom Collections: 
        => you can create custom collections by defining your own struct types and implementing methods for them.
        ## Example: create a custom collection: a Queue.

        type Queue struct {
            items[] int
        } 
        func (q *Queue) Enqueue(item int){
            q.items=append(q.items, item)
        }

        func (q *Queue) Dequeue() int{
            if len(q.items) <1 {
                return -1       //return -1 if the queue is empty.
            }
            item := q.items[0]
            q.items = q.items[1:]
            return item
        }


## comparing java collection and Golang Collections:

    =>
    1. Arrays: In java, we have fixed size arrays, so do in Golang.

    2. Lists: In java, we have List, and in golang we have Slices
    
    3. Maps: Java provides a variety of Map interfaces and implementations, such as HashMap, TreeMap, and LinkedHashMap, in Golang we have built-in map type, which is similar to Java's HashMap functionality.

    4. Sets: Java has a Set interface with implementations like HashSet, TreeSet, and LinkedHashSet. Golang does not provide a built-in set type, but you can create a custom set collection.
    
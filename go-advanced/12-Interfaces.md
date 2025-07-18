# Interfaces:
    => Its a set of method signatures, the struct data type implements these interfaces to have method definition.

   => interfaces enable polymorphism, allowing code to work with different types as long as they implement the required methods of an interface.
   
# Syntax:
    =>
          /* define an interface */
      type interface_name interface {
         method_name1 [return_type]
         method_name2 [return_type]
         method_name3 [return_type]
         ...
         method_namen [return_type]
      }
      
      /* define a struct */
      type struct_name struct {
         /* variables */
      }
      
      /* implement interface methods*/
      func (struct_name_variable struct_name) method_name1() [return_type] {
         /* method implementation */
      }
      ...
      func (struct_name_variable struct_name) method_namen() [return_type] {
         /* method implementation */
      }


# Example:
      =>
        package main
        
        import ("fmt" "math")
        
        /* define an interface */
        type Shape interface {
           area() float64
        }
        
        /* define a circle */
        type Circle struct {
           x,y,radius float64
        }
        
        /* define a rectangle */
        type Rectangle struct {
           width, height float64
        }
        
        /* define a method for circle (implementation of Shape.area())*/
        func(circle Circle) area() float64 {
           return math.Pi * circle.radius * circle.radius
        }
        
        /* define a method for rectangle (implementation of Shape.area())*/
        func(rect Rectangle) area() float64 {
           return rect.width * rect.height
        }
        
        /* define a method for shape */
        func getArea(shape Shape) float64 {
           return shape.area()
        }
        
        func main() {
           circle := Circle{x:0,y:0,radius:5}
           rectangle := Rectangle {width:10, height:5}
           
           fmt.Printf("Circle area: %f\n",getArea(circle))
           fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
        }

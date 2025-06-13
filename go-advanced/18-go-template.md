# templates :
      => 
      refer to a powerful feature used for generating dynamic content. They are used for formatting and producing output, such as HTML for web applications.

# Go provides two types of templates, They are:
1. Text Templates (from text/template package): Ideal for generating plain text output like emails or configuration files.
2. HTML Templates (from html/template package): It is designed for generating HTML and includes automatic escaping to prevent security vulnerabilities like cross-site scripting (XSS)
# example:
      =>
      package main
      import (
         "os"
         "text/template"
      )
      func main() {
         tmpl, _ := template.New("example").Parse("Hello, {{.Name}}!")
         tmpl.Execute(os.Stdout, map[string]string{"Name": "Revathi"})
      }

# Basic Template Syntax
    => Go templates use double curly braces ({{ }}) as placeholders to inject data, iterate, or apply logic (e.g., conditionals, variables, loops ) into templates.

# Syntax:
    => 
    {{.}}              
    {{.FieldName}} 
    {{"constant text"}}

# Example:
      =>
      package main
      import (
         "os"
         "text/template"
      )
      func main() {
         // Define a template
         const tmpl = "Hello, {{.}}!"
         
         // Parse the template
         t, err := template.New("test").Parse(tmpl)
         if err != nil {
         	 panic(err)
         }  
         // Execute the template with data
         err = t.Execute(os.Stdout, "World")
         if err != nil {
         	 panic(err)
         }
      }

# Using Structs in Templates:
    =>
    You can pass a Go struct to a template and reference its fields (.FieldName) to inject structured data dynamically into the output.
# Syntax:
    =>
    {{.FieldName}}      
    {{.Nested.Field}}   
    Loops and Conditionals

# Example:
    => 
    package main
    import (
       "os"
       "text/template"
    )
    type Person struct {
       Name string
       Age  int
    }
    func main() {
       const tmpl = "Name: {{.Name}}, Age: {{.Age}}"
       
       t, err := template.New("test").Parse(tmpl)
       if err != nil {
          panic(err)
       }
    
       p := Person{Name: "Revathi", Age: 24}
       err = t.Execute(os.Stdout, p)
       if err != nil {
          panic(err)
       }
    }

# Loops and conditional:
      => 
      The Templates support control flow by using "range" to iterate over slices or maps and "if" to add conditional logic, tailoring the output based on data.

# Syntax:
      =>
      {{range .}}         
        {{.}}            
      {{end}}
      
      {{if .Condition}}   
        {{.FieldName}}     
      {{else}}             
        {{.OtherField}}  
      {{end}}
    
# Example:
      =>
      package main
      import (
         "os"
         "text/template"
      )
      func main() {
             const tmpl = `{{range .}}{{if .Active}}Name: {{.Name}}, Age: {{.Age}}
      {{end}}{{end}}`
      
         t, err := template.New("test").Parse(tmpl)
         if err != nil {
            panic(err)
         }
      
         type Person struct {
            Name   string
            Age    int
            Active bool
         }
         people := []Person{
            {Name: "Revathi", Age: 25, Active: true},
            {Name: "Tapas", Age: 35, Active: false},
            {Name: "Vivek", Age: 27, Active: true},
         }
      
         err = t.Execute(os.Stdout, people)
         if err != nil {
            panic(err)
         }
      } 

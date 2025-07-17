🌐 2. Embedding to Implement Interfaces (e.g. HTTP Handlers)
Scenario:
You want to create reusable HTTP handlers that implement http.Handler.

Example: Basic Handler Struct
go
Copy
Edit
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from MyHandler")
}
Now embed this in another handler:
go
Copy
Edit
type LoggingHandler struct {
    http.Handler // Embedded interface
}

func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request to %s", r.URL.Path)
    h.Handler.ServeHTTP(w, r) // Delegate to embedded handler
}
Usage:
go
Copy
Edit
h := LoggingHandler{
    Handler: MyHandler{},
}

http.ListenAndServe(":8080", h)
Now every request is logged before your custom handler executes — perfect for middleware-like behavior using embedding.

🔌 3. Composition in Services
Let’s say you have shared logic for services — like logging or validation — and you want to reuse it across multiple services.

go
Copy
Edit
type ServiceBase struct{}

func (s ServiceBase) LogAction(action string) {
    fmt.Println("Action:", action)
}

type UserService struct {
    ServiceBase // Embed for reuse
}

func (s UserService) CreateUser(name string) {
    s.LogAction("create user: " + name)
    fmt.Println("User created:", name)
}
Usage:
go
Copy
Edit
svc := UserService{}
svc.CreateUser("Eve")  // Logs + Creates
This pattern helps modularize and DRY (Don't Repeat Yourself) your business logic.

🔁 4. Embedding and Interfaces
You can embed a struct that implements an interface, and the outer struct will implement it too (thanks to method promotion).

Interface Example:
go
Copy
Edit
type Notifier interface {
    Notify(msg string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Notify(msg string) {
    fmt.Println("Sending email:", msg)
}

type UserService struct {
    EmailNotifier // Embed it
}
Now UserService also satisfies Notifier:
go
Copy
Edit
var n Notifier = UserService{}
n.Notify("Welcome!")  // Works!
🧠 Recap: Why Use Struct Embedding?
Use Case	Why Use Embedding?
DB Models	Share common fields like ID, timestamps
Middleware/HTTP	Chain handlers using ServeHTTP
Service Logic Reuse	Promote shared methods (e.g., logging)
Interface Composition	Inherit interface behavior


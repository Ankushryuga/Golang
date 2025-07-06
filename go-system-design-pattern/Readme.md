# Go design patterns
    => refers to archiectural approaches and reusable solutions to common large-scale s/w design problems:

# 1. Layerd Architecture
    =>Organize your code into well-separated layers:
    1. handlers (API layer)
    2. services (business logic)
    3. repositories (database access)
    4. models (domain types)

    ✅ Why Use:
    Separation of concerns
    Testable & maintainable
    Clear dependency direction
        ┌────────────┐
        │  Handler   │  ← Web or CLI
        └────┬───────┘
             ↓
        ┌────────────┐
        │  Service   │  ← Business logic
        └────┬───────┘
             ↓
        ┌────────────┐
        │ Repository │  ← DB or external systems
        └────────────┘

# 2. Factory Pattern:
    => Used to create objects (e.g., DB Clients, services) without exposing the creation logic
      func NewUserService(repo UserRepository) *UserService {
        return &UserService{repo: repo}
      }

# 3. Singletoon Pattern:
    => Ensure only 1 instance of a service(e.g., config loader, logger).
    var instance *Logger
    var once sync.Once
    
    func GetLogger() *Logger {
        once.Do(func() {
            instance = &Logger{}
        })
        return instance
    }

# 4. Observer Pattern:
    => Use for event-driven architecture or pub-sub systems(e.g., notify subscribers on state changes).
    1. useful for message queues or logging systems.


# etc
    

    

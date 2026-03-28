# How to approach a system design question.

System designing is an **open-ended conversation**. You are expected to lead it.


## Step 1: Outline use cases, constraints, and assumptions

Gather requirements and scope the problem. Ask questions to clarify use cases and constraints. 

> [!IMPORTANT]
> Discuss assumptions

1. Who is going to use it?
2. How are they going to use it?
3. How many users are there?
4. What does the system do?
5. What are the inputs and outputs of the system?
6. How much data do we expect to handle?
7. How many requests per second do we expect?
8. What is the expected read to write ratio?

## Step 2: Create a high level design.

Outline a high level design with all important components.

1. Sketch the main components and connections
2. Justify your ideas

## Step 3: Design core components

Dive into details for each component. For example, if you were asked to design to [design a url shortening service](), 

> [!IMPORTANT]
> discuss:
> 
> 1. Generating and storing a hash of the full url
>       1. [MD5] and [Base62]
>       2. Hash collisions
>       3. SQL or NoSQL
>       4. Database schema
>
> 2. Translating a hashed url to the full url
>       1. Database lookup
>
> 3. API and object-oriented design

## Step 4: Scale the design

Identify and address bottlenecks, given the constraints. For example, do you need the following to address scalability issues?

1. Load balancer
2. Horizontal scaling
3. Caching
4. Database sharding


> [!IMPORTANT]
> Discuss potential solutions and trade-offs. Everything is a trade-off. Address bottlenecks using [principles of scalable system design]().


## Back-of-the-envelope calculations

You might be asked to do some estimates by hand. Refer to the [Appendix]() for the following resources.

1. [Use back of the envelope calculations]()
2. [Powers of two table]()
3. [Latency numbers every programmer should know]()

## More resource
coming soon

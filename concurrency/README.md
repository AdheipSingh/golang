# Concurrency 

- As mentioned by Rob Pike concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

- Before  jumping to go routines lets just understand synchronous and asynchronous a bit. In simple terms when we execute something synchronously, you wait for it to finish before moving to another task where as in async, we can move on to another task before the first once finishes.

- For computers in their context, this translates into a an executing process. A thread is a series of commands that exists as a unit of work. 

- When working with concurrent code, there are a few options for safe operation. 
1. Synchronization primitives for sharing memory
2. Synchronization via communicating

- In Go, use the go keyword to run a function concurrently. This shall create a GoRoutine. 



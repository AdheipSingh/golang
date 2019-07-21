# Go Routines

- Every go program has at least one go routine ie. the main go routine, which is started automatically when the process begins. 

- A go routine is a function that is running concurrently alongside other code, not necessarily in parallel. 

# Go Routines Examples 

```
git clone https://github.com/AdheipSingh/golang-guide.git
cd concurrency/Go-Routines/
```
- Run without go routines

```
go run without-GR/no-goroutines.go 
```

- Run with go routines

```
go run with-GR/goroutines.go 
```
# Wait Groups

- Run with waitgroup. A waitgroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the go routines runs and calls Done when finshed.

- sync.WaitGroup also provides 3 methods: Add, Done and Wait. Add method is used to identify how many goroutines need to be waited. When a goroutine exits, it must call Done. The main goroutine blocks on Wait, Once the counter becomes 0, the Wait will return, and main goroutine can continue to run.

```
 go run waitgroups/wg.go
```

# Race Conditions

- A race condition occurs when two or more operations must execute in the correct order, but the program somehow does not maintain that guarantee.

- Where one concurrent operation attempts to read a variable while at some undetermined time another concurrent operation is attempting to write to the same vairable. This is called a data race.


- To check if a race condition occurs
```
go run -race Go-Routines/race-conditions/rc.go 
```


# MUTEX 

- Mutex stands for mutually exclusive or mutual exclusion object. It is a program that is created so that multiple program thread can take turns sharing the same resource. 

- Mutex helps solve multi-threading problems.

```
go run -race concurrency/Go-Routines/mutex/mutex.go
```
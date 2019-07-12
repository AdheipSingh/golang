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
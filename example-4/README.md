# Concurrency in Go!

## Not Green threads
```
In computer programming, green threads are threads that are scheduled by a runtime library or virtual machine (VM) instead of natively by the underlying operating system.
```

## Not Coroutines
```
A Coroutine is a component that generalizes a subroutine to allow multiple entry points for suspending and resuming execution at certain locations. Unlike subroutines, coroutines can exit by calling other coroutines, which may later return to the point where they were invoked in the original coroutine.
```

## Go Routines are
```
Lightweight processes (not OS) managed by the go runtime, which are multiplexed onto 1 or more OS level threads 
```


A real world example...

Any other questions around this you can look at:
* https://www.youtube.com/watch?v=f6kdp27TYZs&t=914s
* https://www.youtube.com/watch?v=cN_DpYBzKso 
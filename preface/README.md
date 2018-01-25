# Preface

# What the hell is Go?....

Go is an open source language developed by some old and geeky unix guys at google because they got tired of C/C++.

1. ## Go compiles down to native binaries
    * Multiple Arch backends all by setting a variable at build time.
      * ARM
      * Windows
      * Darwin
      * Linux
      * ...
2. ## Go is garbage collected ( no pointer arithmetic )
    * `tricolor mark-and-sweep algorithm` =
    * Lower latency garbage collection mean reduced throughput
3. ## Go is open source

## Now the stuff you'll probably hear everywhere else...

1. Go is great for concurrency!
    * first class support for concurrency primitives (`go` keyword)
    * concurrency ecosystem influenced by Tony Hoare's paper https://en.wikipedia.org/wiki/Communicating_sequential_processes (ala the `erlang vm` and others)
2. Go is performant!
3. Go is fast!
4. Go is inelegant!
    ```go
      err := doWork()
      if err != nil {
        handleError()
      }
      ```
5. Go is verbose
6. Go doesn't have generics!
6. Go is arcane! and looks like C! ```go var a *int = &1```
7. [Gofmt](https://golang.org/cmd/gofmt/) uses tabs instead of spaces!
8. Go doesn't have ***sets*/*queues*/*stacks*/*dequeues*/*linkedLists***! etc...
9. Go doesn't have any of the cool FP features like pattern matching, Functors/Applicatives/Monads/Monoids
10. (usually) Huge frameworks like Spring, and ORM's are discouraged!
11. Go embraces code generation


## What you may not know, but you probably do know...
1. ## Written in / Heavily Utilizing Go
    * ### [Docker](https://github.com/moby/moby)
    * ### [terraform (and typically all other Hashicorps products)](https://github.com/mitchellh?tab=repositories)
    * ### [Kubernetes](https://github.com/kubernetes/kubernetes)
    * ### [Etcd](https://github.com/coreos/etcd)
    * ### [SoundCloud](https://developers.soundcloud.com/blog/go-at-soundcloud) *outdated link*
    * ### [CloudFlare](https://github.com/cloudflare)
    * ### [CoreOS](https://github.com/coreos)
    * ### and more....



# Anthropomorphic Language Mascot

![golang gopher](./gopher.png)
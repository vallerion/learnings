# [Rust vs GO](https://bitfieldconsulting.com/golang/rust-vs-go)

1. Rust is a low-level statically-typed multi-paradigm programming language that’s focused on safety and performance.
2. both Rust and Go eliminate such issues completely by using a standard formatting tool (gofmt for Go, rustfmt for
   Rust)
3. Go is focused on concurrency as a first-class concept. That is not to say you cannot find aspects of Go’s
   actor-oriented concurrency in Rust, but it is left as an exercise to the programmer.
4. “Fighting with the borrow checker” is a common syndrome for new Rust programmers
5. “Fighting the borrow checker” becomes “The compiler can detect that? Cool!”
6. My take: Go for the code that has to ship tomorrow, Rust for the code that has to keep running untouched for the next
   five years.
7. Rust aims to give the programmer complete control of the underlying hardware, it’s possible to optimise Rust programs
   to be pretty close to the maximum theoretical performance of the machine.
8. Not having GC in Rust makes it really fast (especially if you need guaranteed latency, not just high throughput) and
   enables features and programming patterns that are not possible in Go (or at least not without sacrificing
   performance).
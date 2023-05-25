# [Ultimate Go Notebook](https://www.amazon.com/Ultimate-Go-Notebook-William-Kennedy/dp/1737384426)

## Chapter 2: Language Mechanics

1. Escape analysis. May sound that value escape from stack to heap, but it's not the case. Escape analysis is algorithm
   happens during compile and preform analyse to determine whether variable will tend to overgrow stack or can be safely
   allocated on stack.
2. Function stack size is being calculated on compilation stage. If compiler can determine what size of stack is needed
   the value must be constructed on the heap.
3. Constants in Golang could be typed and untyped:
    1. `const pi = 3.141592 // kind: floating point`
    2. `const pi float64 = 3.141592 // type: float64`

## Chapter 3: Data Structures

1. Each core inside processor has its own memory (L1 and L2) and shared memory L3.
2. The **latency** of accessing each layer changes from least to most: L1 -> L2 -> L3. 
3. The **linked list** is twice **slower** than a **row traversal** dut to cache line misses but has fewer TLB (
   Translation Lookaside Buffer) misses.
4. The OS shares physical memory by breaking it into pages and provide access to physical memory through pages for every
   running program. The TLB is a small buffer inside CPU which helps OS reduce latency on translating virtual memory
   address to physical memory address.
5. Performance in my app is how fast I can get data into processor. 
6. A **string** is immutable **two-word** structure: pointer to backing array of bytes and length. 
7. There are two ways to iterate over a **slice** in Go:
    1. Pointer semantic
   ```go
    for i := range urls {
        fmt.Println(urls[i])
    } 
   ```
    2. Value semantic
   ```go
    for i, v := range urls {
        fmt.Println(v)
    } 
   ```
   When we iterate through value semantics we get a copy of element every iteration, also this copy rewrites every
   iteration. In pointer iteration we use direct access to slice. 
8. In Go slice of array (not slice) has to be known on compile time. 
9. A slice is a **three-word structure**: pointer to backing array, length and capacity.
10. If a size of slice is known at compile time, the backing array could be constructed on stack. 
11. It is recommended to use value semantics to move slice around program since slice just contain a link to actual data.
12. Marshaling and unmarshaling are only recommended cases to use pointer semantics with slice. 
13. If maximum capacity reached `append` creates its own copy of provided slice, mutate it and return.
14. `append` creates a new backing array (doubling or growing by 25%), updates slice pointer, copies values from old to
    new array, add new value and return.
15. Once **backing array** reached limit in 1024 elements growing happens at 25%. 
16. Map is a data structure that support storing and accessing data by the key. It uses a hashmap and bucket system to
   support contiguous block of memory.

## Chapter 4: Decoupling

1. **Methods** are just a functions provide the ability for data to exhibit behavior.
2. When I call a method, the compiler converts it to a function call. 
3. Functions are values in Go and belongs to set of internal types.
4. Assigning a function to variable create a copy of this function.
   ```go
      d := data{
        name: "Bill",
      }
      f1 := d.displayName // f1 is copy
      d.name = "Taras"
      f1()            // print Bill
      d.displayName() // prints Taras
   ```
5. Interfaces in Go mainly used for decoupling between software components and to encourage design by **composition**. 
6. Go is about convention over configuration. 
7. **Polymorphism** means that a piece of code (e.g. function) can change its behaviour without changing internal code.
   ```go
      func read(r Reader) (int, error) { // Reader interface can change behaviour
         ...
      }
   ```
8. Interface in Go is **two-word structure**: type descriptor and a data pointer. Data pointer refer to actual data being
   stored. The type descriptor is also a pointer that points to iTable structure. The iTable (interface table) structure
   contains concrete data type (implementation) and pointer to set of functions of implementation being stored.
   ```
      +----------------------+                     
      |       Interface      |                     iTable:
      +----------------------+                     +----------------+
      |   Type: *iTable      |                     |  Type          |  -->  ConcreteType
      |   Value: *Data       |                     +----------------+
      +----------------------+                     |  NumMethods    |  -->  total number of methods defined in the interface
                                                   +----------------+
                                                   |  Method1       |
                                                   +----------------+
                                                   |  Method2       |
                                                   +----------------+
   ```
9. Constants only live at compile time. The value is never on the stack or heap. 
10. **Embedding** could be two types: value semantics and pointer semantics.
    ```go
       type user struct {}
       type admin struct {
          *user // pointer
       }
       type admin struct {
          user // value
       }
    ```
11. `user` implement interface then embedded to `admin` so `admin` also implements interface.
12. If both embedded and outer type has same method. Promotion of embedded method won't happen and outer type method will
    be used. 
13. Avoid implementing `New` function that returns unexported type.

## Chapter 5: Software Design

1. Types should be grouped by common behaviour (methods) not type or "DNA".
2. Embed types only if I need that behaviour. 
3. Discover interfaces not design them.
4. Convention over configuration 
5. Handling error means decide to logging the error, propagating or not and determining if the goroutine should be
   terminated.
6. Use error type as a context of error.

## Chapter 6: Concurrency

1. Golang scheduler use OS threads under the hood. Usually we use next symbols: M - system thread, p - logical
   processor, G - goroutine.
2. On the begging Go runtime asks machine how much OS thread available. Available through `runtime.GOMAXPROCS()`.
3. Concurrency - undefined execution order. Parallelism - doing many things at once (require at least two physical cores
   or CPU threads).
4. CPU Bound work is work requires CPU calculation and do not cause moving thread to waiting state.
5. I/O Bound work is work that cause thread move to waiting state, e.g. waiting for http-response.
6. Data race is when we have at least two goroutines accessing same memory and at least one of them writing.
7. Atomics provide synchronization at hardware level. sync.WaitGroup use atomics.
8. Mutex is boxing block of code that could be executed only by one goroutine at a time. If forget unlock then all
   goroutines wait indefinite -> deadlock.
9. Keep code block inside mutex as small as possible but do not use it as optimization, e.g. call Lock/Unlock twice in
   the same function.
10. RWMutex allow all goroutines reading at the same time, but once write take place all readings and writings wait.
11. Channels are not data structure, but a mechanic for signaling.
12. Sending on unbuffered channel will lock the sender until receiver is available. Buffered channel won't block sender until the buffer is full.
13. Don't need to close channel to release memory, use that for signaling.
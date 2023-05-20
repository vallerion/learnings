# [Ultimate Go Notebook](https://www.amazon.com/Ultimate-Go-Notebook-William-Kennedy/dp/1737384426)

## Chapter 2: Language Mechanics

1. Escape analysis. May sound that value escape from stack to heap, but it's not the case. Escape analysis is algorithm
   happens during compile and preform analyse to determine whether variable will tend to overgrow stack or can be safely
   allocated on stack.
2. Only consumption of the heap called an allocation in Go.
3. Function stack size is being calculated on compilation stage. If compiler can determine what size of stack is needed
   the value must be constructed on the heap.
4. Constants in Golang could be typed and untyped:
    1. `const pi = 3.141592 // kind: floating point`
    2. `const pi float64 = 3.141592 // type: float64`

## Chapter 3: Data Structures

### 3.1 CPU Caches

1. Each core inside processor has its own memory (L1 and L2) and shared memory L3.
2. The **latency** of accessing each layer changes from least to most: L1 -> L2 -> L3.

### 3.2 Translation Lookaside Buffer (TLB)

2. The **linked list** is twice **slower** than a **row traversal** dut to cache line misses but has fewer TLB (Translation
   Lookaside Buffer) misses.
3. The OS shares physical memory by breaking it into pages and provide access to physical memory through pages for every
   running program. The TLB is a small buffer inside CPU which helps OS reduce latency on translating virtual memory
   address to physical memory address.
4. Performance in my app is how fast I can get data into processor.

### 3.3 Declaring and Initializing Values

1. A **string** is immutable **two-word** structure: pointer to backing array of bytes and length.

### 3.5 Iterating Over Collections

1. There are two ways to iterate over a slice in Go:
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

### 3.9 Different Type Arrays

1. In Go slice of array (not slice) has to be known on compile time.

### 3.11 Constructing Slices

1. A slice is a three-word structure: pointer to backing array, length and capacity.
2. If a size of slice is known at compile time, the backing array could be constructed on stack.
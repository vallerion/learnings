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

2. The **linked list** is twice **slower** than a **row traversal** dut to cache line misses but has fewer TLB (
   Translation Lookaside Buffer) misses.
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

### 3.13 Data Semantic Guideline For Slices

1. It is recommended to use value semantics to move slice around program since slice just contain a link to actual data.
2. Marshaling and unmarshaling are only recommended cases to use pointer semantics with slice.

### 3.15 Appending With Slices

1. If maximum capacity reached `append` creates its own copy of provided slice, mutate it and return.
2. `append` creates a new backing array (doubling or growing by 25%), updates slice pointer, copies values from old to
   new array, add new value and return.
3. Once backing array reached limit in 1024 elements growing happens at 25%.

### 3.22 Declaring And Constructing Maps

1. Map is a data structure that support storing and accessing data by the key. It uses a hashmap and bucket system to
   support contiguous block of memory.

## Chapter 4: Decoupling

### 4.5 Methods Are Just Functions

1. Methods are just a functions provide the ability for data to exhibit behavior.
2. When I call a method, the compiler converts it to a function call.

### 4.6 Know The Behavior of the Code

1. Functions are values in Go and belongs to set of internal types.
2. Assigning a function to variable create a copy of this function.
   ```go
      d := data{
        name: "Bill",
      }
      f1 := d.displayName // f1 is copy
      d.name = "Taras"
      f1()            // print Bill
      d.displayName() // prints Taras
   ```

### 4.7 Interfaces

1. Interfaces in Go mainly used for decoupling between software components and to encourage design by composition.

### 4.9 Implementing Interfaces

1. Go is about convention over configuration.

### 4.10 Polymorphism

1. Polymorphism means that a piece of code (e.g. function) can change its behaviour without changing internal code.
   ```go
      func read(r Reader) (int, error) { // Reader interface can change behaviour
         ...
      }
   ```
2. Interface in Go is two-word structure: type descriptor and a data pointer. Data pointer refer to actual data being
   stored. The type descriptor is also a pointer that points to iTable structure. The iTable (interface table) structure contains concrete
   data type (implementation) and pointer to set of functions of implementation being stored.
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
   

                                     
                                     
                                     
                                     
                                     

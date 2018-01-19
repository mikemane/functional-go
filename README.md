# Functional Go
Here are a list of higher order functions written in golang.

### Notation
- f(a) -> bool: This means function f applied to value a returns a bool.
- f(a) -> type(a): This means function f applied to a will return a value that is of type of a
- A: A List of values a1,....an 
- F: A series of function f1,....fn

### This includes 
- Map: The map function will take an array of elements and apply the same function to all the elements.
> Given a list A and a function f, Map(A, f) -> type(A) = f(A) = f(a1), f(a2),...f(an)

- ParMap: This is a parallel version of the map function that takes in the same arguments as the map function. The difference is that this function takes in an extra argument. This argument specifies the amount of threads to use whilst performing the operation.

- Filter: The filter function takes in a function that returns a boolean and a list and return the values that satisfies the predicate function.
> Filter(A, f(a)) -> type(A) where f(a) -> bool and a is an element of A.

- Reduce: Given a list and a function that takes in two parameters of the same type of the function return the acculumated value based on the function.
> Reduce(A, f(a = type(a), b =type(a)) -> type(a), initialValue = type(a)) -> type(a) where a and b are different elements of A.

- Compose: This applies a series of function on a list of objects from right to left. 
> Compose(F) = where F is in the order f1 ----> fn, applies the function list in the order fn(fn-1(fn-2(.......f1))).

- Pipe: This applies a series of function on a list of objects from left to right.
> Pipe(F) = where F is in the order f1 ----> fn, applies the functions in the order f1(f2(f3(.......fn))).
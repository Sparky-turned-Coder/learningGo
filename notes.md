# Golang

## Pointers

### Basics:

 - Every variable is a memory location
 - Every memory location has its own (defined)address 
 - These memory 'addresses' can be accessed by using the ampersand sign(&) and this denotes an address.

Example:
    
    x := 5

    fmt.Println(x) 
    fmt.Println(&x)

    - The first print statement will print the *value* stored in the variable 'x'
    - The second print statment will print the *memory address* assigned to the variable 'x'

## Loops

Most programming languages use three types of loops: for, while, do-while.

Golang uses only 1 loop. The FOR LOOP.

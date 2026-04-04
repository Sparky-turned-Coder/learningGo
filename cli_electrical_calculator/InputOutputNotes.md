# Notes

## Handling user input and output in Go

I'm working with Gippity as my teacher for this one.

### 1. The % (Format Specifiers)

The % symbol starts a __format specifier__ (or "verb"), which tells Go what type of data to expect from the input. Common 
specifiers include:
* __%d__: Integer (decimal)
* __%s__: String (reads until the first space)
* __%f__: Floating-point number
* __%t__: Boolean (true or false)
* __%c__: A single character (rune)

### 2. The __&__ (Address-of Operator)
The __&__ symbol is the __address-of operator__. 
It is placed before a variable name (e.g., &name) to provide the memory address of that variable to the __Scanf__ function.

* __Why it's needed:__ Go passes arguments by value (making a copy). To allow __Scanf__ to modify the original variable's 
value with the user's input, _you must pass a __pointer__ to that variable's location in memory._

#### Think of it like this:

We have __userInput__...

Imagine a box:
* __userInput__ = what's inside the box
* __&userInput__ = location of the box

#### Example:

    package main

    import "fmt"

    func main() {
	    var input string

	    fmt.Println("Enter your name: ")
	    fmt.Scanln(&input)
	    fmt.Printf("Hello, %s\n", input + "!")
    }

* __Declare variable__
    - Sticking to our box metaphor, _This creates an empty box_
            
            var input string
* __Prompt User__
    - This just displays text. Nothing more.
        
            fmt.Println("Enter your name: ", input)
* __Read Input__
    - Takes what the user types
    - Stores it inside __input__
        
            fmt.Scanln(&input)
* __Print Result__
    - Reads the value from __input__
    - Displays it

            fmt.Printf("Hello, %s\n", input + "!")

#### Key Takeaway:

__Functions can either READ data or WRITE data__
* _Println_ - reads your variables and prints them
* _Scanln_  - writes into your variables

#### Mental Model:
* __input__ - the box (holds a value)
* __&input__ - the address of the box (where to store the new data)
* __Scanln(&input)__ - Go goes to that address and WRITES the user's input into the box

#### Compare:

    // This reads a value (outputs to screen)
    fmt.Println(input)

    // This writes a value (takes user input)
  fmt.Scanln(&input)

Think __Println__ reads, __Scanln__ writes.

### Key Behaviors and Tips
* __Space Seperation:__ By default, __%s__ stops reading at the first whitespace.
* __Newlines:__ _fmt.Scanf_ can be sensitive to newline characters. If you find it skipping inputs in a loop, adding a 
newline or space to the format string _("%d\n")_ can help consume the trailing "Enter" key press.
* __Return Values:__ _Scanf_ returns two values: the number of items successfully scanned and an error (if any). Which I 
would assume means we need to include handling of potential errors when using _Scanf_...
* __Alternatives:__ For simpler input without complex formatting, _fmt.Scanln_ is ofte easier as it automatically stops 
at a newline and doesn't require a format string.




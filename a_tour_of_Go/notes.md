# A Tour of Go

## Basic Types

* __bool__
* __string__

* __int  int8  int16  int32  int64__
* __uint uint8 uint16 uint32 uint64 uintptr__

* __byte__ // alias for uint8
* __rune__ // alias for int32
           // represents a Unicode code point

* __float32 float64__

* __complex64 complex128__

### Regarding 'Runes'

It is important to note the difference between the rune type itself and how the character is represented within a string:

* rune type: An int32 (4 bytes) used to represent a single Unicode code point. This fixed size guarantees it can hold any code point.
* String representation: Strings in Go are slices of bytes encoded with UTF-8, which is a variable-length encoding.
* An ASCII character (like 'A' or ' ') takes only 1 byte in a string.
* Other characters, like the euro symbol ('€') or an emoji ('😊'), can take 2, 3, or up to 4 bytes within the string's byte slice.

When you iterate over a string using a for range loop, Go correctly decodes the variable-length UTF-8 sequences and presents each individual code point as a rune (4-byte integer).

## Zero values

Variables declared __without__ an explicit initial value(initializer) are given their _zero value._

The zero values are:

> * __Numeric types:__ 0, 0.0, 0+0i = int, float, complex
> * __Boolean types:__ false
> * __Strings:__ "" (empty string)
> * __Pointers, functions, interfaces, slices, channels, maps:__ nil

Key Concepts

* __Predictable Behavior:__
    The zero value principle ensures code is safer and more predictable, as variables are never left uninitialized with arbitrary memory contents.

* __Structs:__
    The zero value of a 'struct' is a new struct with all its fields set to their respective zero values recursively.

* __nil vs. Zero Value:__
    'nil' is used to represent the absence of a value for specific complex types, but it's important to note that 'nil' is the _actual zero value_ for those types(pointers, slices, etc.). Basic types like booleans and numeric types can never be 'nil'.

* __Usability of Zero Values:__
    Many standard library types, such as 'sync.Mutex' and 'bytes.Buffer', are designed to be used safely from their zero value without explicit initialization.

* __Distinguishing Unset Values:__
    For use cases like databases or APIs where a value of 0 might be a valid input but needs to be distinct from an unset or missing value, Go developers often use pointers to the type (where 'nil' indicates the value was unset) or custom types, such as those in the 'database/sql' package.

## Recursion (for the uninitiated)

__Explain Recursion to me like I've never seen it before:__

Imagine you have a __Russian Nesting Doll__ (a Matryoshka doll).

1. __The Goal:__ You want to make sure every single doll, from the largest one on the outside to the tiniest one in the very middle, is clean (set to zero).
2. __No Recursion(Manual):__ You open the big one and clean it. Then you open the next one and clean it. Then the next... you have to keep doing this _manually_ until you hit the tiny, solid one.
3. __Recursive Zero Value(Go's Way):__ You take the big, uninitialized doll and tell Go: "Clean this."

    1. Go opens it, cleans the outer shell.
    2. Go sees another doll inside, so it __calls itself__ to clean that inner doll. (notice that it __calls itself__.)
    3. This keeps happening automatically-diving deeper and deeper-until it hits the last, solid doll (the base case).

__"Recursively" just means Go doesn't stop at the surface.__ It dives into structs, then into arrays inside those structs, then into pointers, until everything is zeroed out.

__In Go specifically,__ recursive initialization applies to all composite types that contain other types, specifically __arrays__ and __structs__. The rule is that if an array or struct is declared _without_ an explicit value, Go recursively zeros each element or field.

* __Structs:__ _type A struct { B int }_ - A new 'A' will have 'B' set to 0.
* __Arrays:__ _var arr [5]string_ - Creates an array of 5 empty strings ("").
* __Nested Types:__ If a struct contains an array or another struct, the recursion descends through all levels, initializing all nested fields and elements to their respective zero values.

Non-composite types (integers, strings, booleans) are set to 0, "", or false, respectively. Types that hold references (pointers, slices, maps, channels, functions) are initialized to __nil__.

__NOTE:__ You've used the built-in 'Error' type before. Error is an __interface type__, therefore its zero value is __nil__.

> This is a key part of how error handling works in Go: a 'nil' error value means that there was no error.

Function calls conventionally return a non-__nil__ error value if an abnormal state occurred (non-nil == there was an error), which is why code frequently checks:

> if err != nil    >>  then we have a problem

## Type Converstion

In Go, type conversion is an __explicit__ process used to convert a value from one data type to another. It _does not support_ automatic or implicit conversions, even between compatible types like __int__ and __int64__.

### General Syntax

The basic syntax for converting a value(v) to a type(T) is:

> T(v)

For example, to convert an integer to a float64:

> var i int = 42
>
> var f float64 = float64(i)  
>
> _i is converted to a float64_

Conversely, if we were to convert a float to an integer, the decimal portion would be _truncated_, not rounded.

### Common Conversions

Specific functions or methods are required for conversions between strings and numeric types, primarily found in the standard library's __strconv__ package.

* __Numeric Types:__
    * int to float64: float64(myInt)
    * float64 to int: int(myFloat)  _(truncates decimals)_
    * Between different integer sizes (int32 to int64): int64(smallerInt)
* __Strings and Numbers:__
    * string to int: Use __strconv.Atoi(myString)__. This function returns both the integer and a potential error, which must be handled.
    * int to string: Use __strconv.Itoa(myInt)__
    * string to float64: Use __strconv.ParseFloat(myString, 64)__
* __Strings to Bytes/Runes:__
    * string to []byte: []byte(myString)
    * []byte to string: string(myBytes)
    * string to []rune: []rune(myString)
    * []rune to string: string(myRunes)

### Type Assertions and Interfaces

For variables of an __interface{}__ (or __any__) type, a __type assertion__ is used to access the underlying concrete value. The syntax is __x.(T)__, where __x__ is the interface variable and __T__ is the asserted type.

> var a interface{} = 42
>
> i, ok := a.int(int)  _(Type assertion with the "comma, ok" idiom)_
>
> if ok {
>
> _// use __i__ as an int_
>
>}

Using the "comma, ok" idiom is safer as it prevents a runtime panic if the assertion fails. A __type switch__ is also useful for handling multiple possible types stored in an interface.
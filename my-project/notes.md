# Notes - Learn Go with tests

### Writing tests

Writing a test is just like writing a function, with a few rules:

    - It needs to be in a file with a name like <xxx_test.go>
    - The test function must start with the word <Test>
    - The test function takes one argument only: <t *testing.T>
    - To use the *testing.T type, you need to  <import "testing">, like we did 
    with 'fmt' in the other file.

For now, it's enough to know that your 't' of type <*testing.T> is your 'hook' 
into the testing framework so you can do things like <t.Fail()> when you want to fail.

### Go's Documentation

Go has a built-in tool, doc, which lets you examine any package installed on 
your system, or the module you're currently working on.

To view that documentation, enter this in the terminal:

    - $ go doc fmt

Consider the function below:

    func Hello(name string) string {
        return "Hello, " + name + "!"
    }

- the func Hello now accepts an arguement (type 'string') which we call 'name'
- the second 'string' declared after the paranthesis tells the function to return a value of type 'string'.
- finally, our function returns the string(name) passed to us by concatenating it with 
"Hello, " and "!"  Resulting in:   "Hello, Chris!"



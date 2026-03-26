# Notes - Learn Go with tests

### Writing tests

Writing a test is just like writing a function, with a few rules:

    1. It needs to be in a file with a name like <xxx_test.go>
    2. The test function must start with the word <Test>
    3. The test function takes one argument only: <t *testing.T>
    4. Use the *testing.T type, you need to  <import "testing">, like we did 
    with 'fmt' in the other file.

For now, it's enough to know that your 't' of type <*testing.T> is your 'hook' 
into the testing framework so you can do things like <strong>t.Fail()</strong> when you want to fail.

### Go's Documentation

Go has a built-in tool, <strong>doc</strong>, which lets you examine any package installed on 
your system, or the module you're currently working on.

To view that documentation, enter this in the terminal:

    $ go doc fmt


<strong>EXPLAIN:</strong> Consider the function below:

    func Hello(name string) string {
        return "Hello, " + name + "!"
    }

- this Hello function accepts an arguement (type 'string') which we call 'name'
- the second 'string' declared after the paranthesis tells the function to return a value of type 'string'.
- finally, our function returns the string(name) passed to us by concatenating it with "Hello, " and "!"  Resulting in:   "Hello, Chris!"

### A note on source control

At this point, if you are using source control (which we do. We use Git). I would commit the code as it is. We have working software backed by a test.

(This is referring to our Hello World test, which passed after successfully concatenating our string with the value of 'name')

I wouldn't push to main though, because I plan to refactor next. It is nice to commit at this point in case you somehow get into a mess with refactoring - 
you can always go back to the working version.



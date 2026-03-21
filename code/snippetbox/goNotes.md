# Notes for "Let's Go" book by Alex Edwards

## Creating a module

### Inside your root project directory, run the 'go mod init' command - passing in your chosen module path as a parameter like so:  

### $ go mod init <project>.<modualpathname>.net*

*example:* snippetbox.chrisjones.net

    - Inside a valid 'go.mod' file in the root of your project directory, your project 
    *is a module.*   You have just set your project up as a module.

*Note on standard practice*

    - Go documentation advises using a fully qualified domain name (like .com, 
    .net, or .org) for module paths, even if you do not plan to publish the module.

    - In the Go community, a common convention is to base your module paths on a 
    URL that you own.

*Module paths for downloadable packages*

    - If you're creating a project which can be downloaded and used by other people 
    programs, then it's good practice for your 'module path' to equal the location 
    that the code can be downloaded from.

        - For instance, if your package is hosted at https://github.com/foo/bar 
        then the module path for the project should be: 'github.com/foo/bar'

## Hello World

    - Create a 'main.go' in your project directory and write a simple function 
    that prints 'hello world'
    - In the terminal, run:  $ go run .

## Web Application Basics

### Now that everything is set up correctly, let's make the first iteration of our 
### web application. We'll begin with the three absolute essentials:

    - The first thing we need is a handler.
    - The second component is a router(or servemux in Go terminology)
    - The last thing we need is a web server.


# KEY NOTE ON GO TYPES

## In Go, every type has a zero value.

    - int       >             0
    - string    >             ""
    - bool      >             false
    - pointers/errors   >     nil

## NOTE on this code:

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
### Steps for code above:
- We extract the value of id wildcard from the request using r.PathValue()
- Then we try to convert that value (which HTTP transmits as a "string") to an integer
- If it cannot be converted to an integer, err != nil is True.
- If the value is less than 1 (whether by conversion or the default value of 0), this is True.
- If one OR the other is TRUE, we return the error "NotFound"

Once you put it in context, it’s really just:

HTTP always speaks strings
Your app thinks in numbers (IDs, counts, indexes)
Conversion + validation = safe bridge from the web to Go logic

Think of it as a guardrail:

    - strconv.Atoi says: “I’ll give you a number if I can, otherwise I’ll signal an error.”
    - The err != nil || id < 1 check says: “If anything’s wrong, reject it right away.”

That’s why you see this pattern all over Go web apps—it’s idiomatic, safe, and clear once you understand what the ID represents.

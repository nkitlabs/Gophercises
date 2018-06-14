# Exercise #2: URL Shortener

you can follow the details and instruction more on the [link](https://gophercises.com/exercises/urlshort)

## Exercise details

We create a URL shortener handler in [handler.go](handler.go) that will look at the path of any incoming web request and determine if it should redirect the user to a new page.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

We also create a function to map YAML format into map structure in go using the [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2) package.

We implement the program [main/main.go](main/main.go) to test the correctness of the function.

**Big Credit : (https://gophercises.com/exercises/)**
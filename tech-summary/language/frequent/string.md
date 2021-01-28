## C++

### pass string via function parameters

- Using the string as an id (will not be modified). Passing it in by const reference is probably the best idea here: (`std::string const&`)
- Modifying the string but not wanting the caller to see that change. Passing it in by value is preferable: (`std::string`)
- Modifying the string but wanting the caller to see that change. Passing it in by reference is preferable: (`std::string &`)
- Sending the string into the function and the caller of the function will never use the string again. Using move semantics might be an option (`std::string &&`)

[ref](https://stackoverflow.com/questions/10789740/passing-stdstring-by-value-or-reference)

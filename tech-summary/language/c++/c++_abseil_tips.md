- [Abseil Tips of the week](#abseil-tips-of-the-week)
  - [24: Copies](#24-copies)
  - [101: Return Values, References, and Lifetimes](#101-return-values-references-and-lifetimes)

# Abseil Tips of the week


## [24: Copies](https://abseil.io/tips/24)

Want speed, return value.

```C++
std::string build();

std::string foo(std::string arg) {
  return arg;  // no copying here, only one name for the data “arg”.
}

void bar() {
  std::string local = build();  // only 1 instance -- only 1 name

  // no copying, a reference won’t incur a copy
  std::string& local_ref = local;

  // one copy operation, there are now two named collections of data.
  std::string second = foo(local);
}
```

## [101: Return Values, References, and Lifetimes](https://abseil.io/tips/101)

When you see code like
```C++
const string& name = obj.GetName();
std::unique_ptr<Consumer> consumer(new Consumer(name));
```
You should worry about
- What type is returned?  Is GetName() returning by value or by reference?
- What type are we storing into/initializing? (The type of `name`)
- Does the constructor for Consumer take string, const string& or string_view?
- Does the constructor have any lifetime requirements placed on that parameter? (If it isn’t just string.)

There are several things need to distinguish
- returning `string`, initializing `string`: [RVO](https://en.wikipedia.org/wiki/Copy_elision#Return_value_optimization) or a `move`, no copy
- returning `const string&` or `string&`, initializing `string`: There is a `Copy`, be aware of liveness of object which are returning a reference to
- (*)returning `string`, initializing `string&`: This won’t compile, as you cannot bind a reference to a temporary.
   + 13:33: error: invalid initialization of non-const reference of type 'std::string& {aka std::basic_string<char>&}' from an rvalue of type 'std::string {aka std::basic_string<char>}'
- returning `const string&`, initializing `string&`:This won't compile, `const_cast`
   + 13:33: error: invalid initialization of reference of type 'std::string& {aka std::basic_string<char>&}' from expression of type 'const string {aka const std::basic_string<char>}'
- (#)returning `const string&`, initializing `const string&`: No cost
   + You must aware of the lifetime of that reference, you have inherited a existing lifetime restriction
- returning `string&`, initializing `string&`: No cost, same as (#)
- returning `string&`, initializing `const string&`:No cost, same as (#)
- returning `string`, initializing `const string&`: Similar situation as (*), but this one works: the language has special-case support for this: if you initialize a const T& with a temporary T, that T (in this case string) is not destroyed until the reference goes out of scope (in the common case of automatic or static variables)



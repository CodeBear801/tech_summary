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
- 

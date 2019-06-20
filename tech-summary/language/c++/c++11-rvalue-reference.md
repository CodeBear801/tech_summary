# C++ rvalue reference

```C++
//Lvalue Examples:
int i;
Dog d1;
i = 1;  // i

//Rvalue Examples:
int x = 1;  // 1
int x = x + 1; // x+ 1
d1 = dog();    // dog() is rvalue of user defined type (class)

int sum(int x, int y) { return x+y; }
int i = sum(1, 2);  // sum(1, 2) is rvalue

// lvalue reference
int& r = x;   // r
const int& r = 1;  // assign rvalue to const lvalue reference 

// lvalue can be used to create an rvalue
int i = 1;
int x = i + 2; 

// rvalue can be used to create an lvalue
int v[3];
*(v+2) = 4;

// function or operator always yields rvalues.
// lvalues are modifiable
// rvalue are not modifiable

// rvalue referernce
int&& c = 3;

```

- The most useful place for rvalue reference is overloading copy constructor and assignment operator, to achieve move semantics.
```C++
X& X::operator=(X const & rhs); 
X& X::operator=(X&& rhs);
```

- Move semantics is implemented for all STL containers, which means:
 1. Move to C++ 11, You code will be faster without changing a thing.
 2. You should use passing by value more often.

```C++
vector<int> foo() { std::vector<int> vTmp;  return vTmp; }
```

- The most frequent std::forward showc case
```C++
template <typename ...Args> void f(Args && ...args)
{
  g(std::forward<Args>(args)...);
}

```
That's because of the [reference collapsing rules](http://thbecker.net/articles/rvalue_references/section_08.html): If T = U&, then T&& = U&, but if T = U&&, then T&& = U&&, so you always end up with the correct type inside the function body. Finally, you need forward to turn the lvalue-turned x (because it has a name now!) back into an rvalue reference if it was one initially.  

**Std::forward avoid using pointer to pass rvalue reference.  For example, an object is allocated on stack, use std::forward act like std::move, to pass obj from one place to another.**  


- Reference Collapsing Rules ( C++ 11 ):
    * T& &   ==>  T&
    *  T& &&  ==>  T&
    * T&& &  ==>  T&
    * T&& && ==>  T&&

```C++
// Example 1
template< classs T >
struct remove_reference;    // It removes reference on type T

// T is int&
remove_refence<int&>::type i;  // int i;

// T is int
remove_refence<int>::type i;   // int i;


// Example 2
template< typename T >
void relay(T&& arg ) {
   ...
}
// T&& variable is intialized with rvalue => rvalue reference
  relay(9); =>  T = int&& =>  T&& = int&& && = int&&
// T&& variable is intialized with lvalue => lvalue reference
  relay(x); =>  T = int&  =>  T&& = int& && = int&

// Example 3
template< typename T >
void relay(T&& arg ) {
  foo( std::forward<T>( arg ) );
}




```



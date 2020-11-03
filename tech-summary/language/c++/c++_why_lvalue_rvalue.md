# Why lvalue & rvalue


## Lvalue & Rvalue

Lvalue & rvalue are <span style="color:red">**sematic properties of expressions**</span>


### In C

- Every expression in C either a `lvalue` or a `rvalue`
- An `lvalue` is an expression referring to an object.  An object is a region of storage.
- An `rvalue` is simply an expression that's not an lvalue
- Caveat: although this is true for none-class type in C++, its not true for class types
- Java: build-in type-> pass by value, none build-in -> pass by reference


<img src="resource/pictures/c++_lvalue_rvalue_lvalues.png" alt="c++_lvalue_rvalue_lvalues" width="500"/>

<span style="color:red">Why lvalue is none moveable - lvalue means it could have more than one way to access that variable </span>

<img src="resource/pictures/c++_lvalue_rvalue_rvalues.png" alt="c++_lvalue_rvalue_rvalues" width="500"/>

Function return value: no name, you can't take it address, candidate for moving




### Why lvalue & rvalue



<span style="color:blue">Why differentiate lvalue & rvalue</span>
<br/>
<img src="resource/pictures/c++_lvalue_rvalue_why_1.png" alt="c++_lvalue_rvalue_why_1" width="500"/>
<br/>
<img src="resource/pictures/c++_lvalue_rvalue_why_2.png" alt="c++_lvalue_rvalue_why_2" width="500"/>
<br/>
Number 1 exists in code section.  
<img src="resource/pictures/c++_lvalue_rvalue_why_3.png" alt="c++_lvalue_rvalue_why_3" width="500"/>
<br/>
<img src="resource/pictures/c++_lvalue_rvalue_why_4.png" alt="c++_lvalue_rvalue_why_4" width="500"/>
<br/>
<span style="color:red">rvalue give compiler the permission for optimization and validation.</span>

More info:   


### Data storage

<img src="resource/pictures/c++_lvalue_rvalue_data_storage.png" alt="c++_lvalue_rvalue_data_storage" width="500"/>


### Basic case

```C
int m, n
m=n     // n is lvalue, this assignment uses the lvalue expressions n as an rvalue
        // due to c++ performs an lvalue-to-rvalue conversion
        // assign to lvalue m

m+n     // the result in a compiler-generated temporary object, such as a CPU register -> rvalue
```

More examples


```C++
//Lvalue Examples:
int i;
Dog d1;
i = 1;  // i

// more explain
// i is lvalue
// - named object, address can be taken
// - could be referred for multiple times


//Rvalue Examples:
int x = 1;  // 1
int x = x + 1; // x+ 1
d1 = dog();    // dog() is rvalue of user defined type (class)

// more explain
// 1 is rvalue
// - can't take its address
// - all build-in numeric literals are rvalues

int sum(int x, int y) { return x+y; }
int i = sum(1, 2);  // sum(1, 2) is rvalue

template <typename T1, typename T2>
int sizeDiff(const T1& t1, const T2& t2)
{ return t1.size() - t2.size(); }

// sizeDiff, t1, t2 are lvalue
// Return value is rvalue
// - address can't be taken
// - no way to get pointer/reference to it

int *px;
*px = sizeDiff(v, s);
// sizeDiff(v, s) is rvalue

// lvalue reference
int& r = x;   // r
const int& r = 1;  // assign rvalue to const lvalue reference 

// lvalue can be used to create an rvalue
int i = 1;
int x = i + 2; 

// rvalue can be used to create an lvalue
int v[3];
*(v+2) = 4;

// - function or operator always yields rvalues.
// - lvalues are modifiable
// - rvalue are not modifiable

// rvalue referernce
int&& c = 3;

```


### Enumeration Constants
- When used in expressions, enum constans are also rvalues

```C
enum color {red, green, blue};
color c;

c = green;   // OK, c is an lvalue, green is rvalue
blue = green; // wrong, blue is an rvalue

```


### unary &
&e is valid expression only if e is an lvalue, <span style="color:red">&e is rvalue</span>  
&e = n is not work


### unary *
In contrast to unary &, unary * yields an lvalue  

<span style="color:red">A pointer p can point to an object, so *p is an lvalue</span>  

<span style="color:red">However, its operand can be an rvalue, while result is lvalue</span>

```C
Int *p = a;
*p = 3;

char *s = Null
*s = '\0'.  // undefined behavior

*(p + 1) = 4   // p+1 is an rvalue, but *(p+1) is an lvalue, we store 4 into the object referenced by *(p + 1)

```



### const

#### Const type

An lvalue is non-modifiable if it has a const-qualified type

```C
Char const name[] = "dan"
Name[0] = 'x'.     // wrong, name[0] is lvalue, but can not change
```

<img src="resource/pictures/c++_lvalue_rvalue_rules.png" alt="c++_lvalue_rvalue_rules" width="500"/>


#### Const object

<img src="resource/pictures/c++_lvalue_rvalue_const_obj.png" alt="c++_lvalue_rvalue_const_obj" width="500"/>


### references
```C++
int &ri = a;  // ri is an alias for a
```

<img src="resource/pictures/c++_lvalue_rvalue_ref_examples.png" alt="c++_lvalue_rvalue_ref_examples" width="500"/>


<span style="color:red">A reference yields an lvalue</span>  
<span style="color:blue">Why reference</span>: C++ has references so that overloaded operators can look just like build-in operators.
Lvalue properties we need, such as implement ++ operator

```C++

Month & operator++(month& x)
{
   return x = static_cast<month>(x + 1);
}

// Month & operator++(month x) will only increase a copy of x but not x itself
```

### pass parameter by value or reference
```C++
R f(T t)                 // by value: f has access only to a copy of x, not x itself

R f(T const &t)    // by reference to const, f's parameter is declared to be non-modifiable 
```

<img src="resource/pictures/c++_lvalue_rvalue_why_use_ref_const.png" alt="c++_lvalue_rvalue_why_use_ref_const" width="500"/>


### references and temporaries


<span>&#9673;</span><span style="color:red">A "pointer to T" can point only to an lvalue of type T</span>  
<span>&#9673;</span><span style="color:red">A "reference to T" binds only to an lvalue of type T</span>  
```
int *pi = &3;  // wrong, can't apply & to 3
int &ri = 3;   // wrong, 3 is a rvalue

int i;
double *pd = &i   // wrong, can't convert pointers
double &rd = i    // wrong, can't bind
```

#### Exceptions


<img src="resource/pictures/c++_lvalue_rvalue_exceptions_const.png" alt="c++_lvalue_rvalue_exceptions_const" width="500"/>


```
int const &ri = 3;
double const &rd = ri;
// converts ri from int to double -> create a temp double to hold result -> binds rd to temp -> when out scope rd will destroys the temp
```


### operator
operator+ yields an rvalue

```
string operator+(string const &lo, string const &ro)
s = s + string(",") + t // lvalue + rvalue + lvalue
string *p = &(s + t);   // wrong
```

***

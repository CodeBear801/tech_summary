# Frequent be used data structures of python

## Basic structure

### str

Built-in data structure, can be used directly without importing

```python
format(*args, **kwargs)  # Various usage, the algorithm can be used in the string form of the binary conversion
split(sep=None, maxsplit=-1) 
strip([chars])
join(iterable)
replace(old, new[, count])
count(substring[, start[, end]])
startswith(prefix[, start[, end]])
endswith(suffix[, start[, end]])
```

spliting/slice

```python
myString="PythonForBeginners"
x=myString[0:10:2]  # PtoFr
x=myString[0:10:3]  # PhFB

txt = "apple#banana#cherry#orange"
x = txt.split("#", 2)          #  ['apple', 'banana', 'cherry#orange']

```

### tuple

+ Tuples are immutable objects
+ The only functions provided are index and count
+ Tuples are often used for multivariable assignment and multivalued return, but they are usually used without parentheses to make python look good.
+ It also supports stitching, slicing, repetition and other operations.
+ Use `namedtuple` for lightweight, immutable data containers before you need the flexibility of a full class
  + You can’t specify default argument values for namedtuple classes. -> dataclass
  + The attribute values of namedtuple instances are still accessible using numerical indexes and iteration

```python
# example 1
from collections import namedtuple
Grade = namedtuple('Grade', ('score', 'weight'))


def __init__(self):
  self._grades = []

def report_grade(self, score, weight):
  self._grades.append(Grade(score, weight))

# example 2
namedtuple(typename, field_names, *, rename=False, defaults=None, module=None)


```

Be carefull abuot `Tuple` only contains one element
```python
tup = (1)
print(tup, type(tup))  # not a tuple, () <class 'tuple'>

tup = (1,)
print(tup, type(tup)) # (1,) <class 'tuple'>
```

### list


| function | operation | example
|:---:| --- | --- 
| `cmp(list1,list2)`  | Compare the size of the two lists | 
| `max(list)` | Maximum value in the list | 
| `min(list)` | Minimum value in the list | 
| `list.append (data)` | Insert data at the end of the list | 
| `list.extend(seq)	` | Append multiple values from another sequence at the end of the list at one time | 
| `list.pop()	` | Remove an element and default to footer | 
| `list. index (data)	` | Find the first place in the list that matches the data | 
| `list. insert (location, data)	` | Inserts an element at the specified location | 
| `list. count (data)	` | Count the number of occurrences of this data in the list | 
| `` |  | 
| `` |  | 


Overide comparison magic methods

```python
class A(object):

    def __init__(self, value):
        self.value = value

    def __lt__(self, other):
        return self.value < other.value

    def __le__(self, other):
        return self.value <= other.value

    def __eq__(self, other):
        return self.value == other.value

    def __ne__(self, other):
        return self.value != other.value

    def __gt__(self, other):
        return self.value > other.value

    def __ge__(self, other):
        return self.value >= other.value

    def __str__(self):
        return str(self.value)

a = A(10)
b = A(20)
min(a, b)
```


`sort`

The list has a sort method, which is used to reorder the original list. The key parameter is specified.  The key is an anonymous function, and the item is a dictionary element in the list. We sort the list according to the age in the dictionary. By default, it is sorted in ascending order, and the specified value is usedreverse=True In descending order

```python
items = [{'name': 'Homer', 'age': 39},
         {'name': 'Bart', 'age': 10},
         {"name": 'cater', 'age': 20}]
items.sort(key=lambda item: item.get("age"), reverse=True)

#  [{'name': 'Homer', 'age': 39}, {'name': 'cater', 'age': 20}, {'name': 'Bart', 'age': 10}]

items = [{'name': 'Homer', 'age': 39},
         {'name': 'Bart', 'age': 10},
         {"name": 'cater', 'age': 20}]

new_items = sorted(items, key=lambda item: item.get("age"))

# items will not changed, new_items contains the sorted list
```


`iteration`

```python
for index, item in enumerate(items):
    print(index, "-->", item)
```

`append` vs `extend`
+ `Append` means to append a data to the bottom of the list as a new element. Its parameters can be any object
+ The parameter of `extend` must be an iterative object, which means that all elements in the object are appended to the back of the list one by one

```python
x = [1, 2, 3]
y = [4, 5]

x.append(y) # x: [1, 2, 3, [4, 5]]
# reset
x.extend(y) # [1, 2, 3, 4, 5]
```
Check is list empty
```python
if not items:
    Print ("empty list")
```

Use internal defined functions
```python

class User(object):
    def __init__(self, age):
        self.age = age
​
    def __repr__(self):
        return 'User(%d)' % self.age
​
    def __add__(self, other):
        age = self.age + other.age
        return User(age)
​
user_a = User(10)
user_b = User(20)
​
c = user_a + user_b
​
print(c)
​
>>>
User(30)

```

## Unordered objects

### dict

| function | operation | example
|:---:| --- | --- 
| `dict.get(key)	` | Returns the value of the specified key | 
| `dict.items()	` | Return traversable (key, value) in listtuplearray  | 
| `dict.has_key(key)	` | Determine whether the key is in the dictionary | 
| `dict.keys()	` | Returns all keys in the dictionary | 
| `dict.pop(key)	` | Delete the key value pair corresponding to the given key in the dictionary | 
| `` |  | 
| `` |  | 



`defaultdict`

```python

from collections import defaultdict
d = defaultdict(lambda : value)
# When a non-existent key is taken, value is returned.
```

[sort dict by value](https://stackoverflow.com/questions/613183/how-do-i-sort-a-dictionary-by-value)

```python
x = {1: 2, 3: 4, 4: 3, 2: 1, 0: 0}
{k: v for k, v in sorted(x.items(), key=lambda item: item[1])} #{0: 0, 2: 1, 1: 2, 4: 3, 3: 4}
```

[sort dict for tuple](https://www.geeksforgeeks.org/python-sort-dictionary-by-tuple-key-product/)

```python
# method 1
# Using dictionary comprehension + lambda + sorted()

# initializing dictionary
test_dict = {(5, 6) : 3, (2, 3) : 9, (8, 4): 10, (6, 4): 12}
  
# printing original dictionary
print("The original dictionary is : " + str(test_dict))
  
# sorted() over lambda computed product 
# dictionary comprehension reassigs dictionary by order 
res = {key: test_dict[key] for key in sorted(test_dict.keys(), key = lambda ele: ele[1] * ele[0])}
  
# printing result 
print("The sorted dictionary : " + str(res))   # The sorted dictionary : {(2, 3): 9, (6, 4): 12, (5, 6): 3, (8, 4): 10}


# method 2
# Method #2 : Using dict() + sorted() + lambda
# initializing dictionary
test_dict = {(5, 6) : 3, (2, 3) : 9, (8, 4): 10, (6, 4): 12}
  
# printing original dictionary
print("The original dictionary is : " + str(test_dict))
  
# sorted() over lambda computed product 
# dict() used instead of dictionary comprehension for rearrangement
res = dict(sorted(test_dict.items(), key = lambda ele: ele[0][1] * ele[0][0]))
  
# printing result 
print("The sorted dictionary : " + str(res)) 

```

Use `sortedcontainers `

```python
# https://www.geeksforgeeks.org/python-sorted-containers-an-introduction/
from sortedcontainers import SortedList, SortedSet, SortedDict


# declare a dictionary of tuple with student data
data = {('bhanu', 10): 'student1', 
        ('uma', 12): 'student4', 
        ('suma', 11): 'student3', 
        ('ravi', 11): 'student2',
        ('gayatri', 9): 'student5'}
  
# sort student dictionary of tuple based 
# on items using OrderedDict
print(OrderedDict(sorted(data.items(), key = lambda elem : elem[1])))
```


### heap

Override `__it__` for user defined class ([more info](https://stackoverflow.com/questions/3954530/how-to-make-heapq-evaluate-the-heap-off-of-a-specific-attribute))

```python
# import required module
import heapq as hq
  
# class definition
class employee:
  
  # constructor
    def __init__(self, n, d, yos, s):
        self.name = n
        self.des = d
        self.yos = yos
        self.sal = s
  
  # function for customized printing
    def print_me(self):
        print("Name :", self.name)
        print("Designation :", self.des)
        print("Years of service :", str(self.yos))
        print("salary :", str(self.sal))
  
   # override the comparison operator
    def __lt__(self, nxt):
        return self.sal < nxt.sal
  
  
# creating objects
e1 = employee('Anish', 'manager', 3, 10000)
e2 = employee('kathy', 'programmer', 2, 15000)
e3 = employee('Rina', 'Analyst', 5, 30000)
e4 = employee('Vinay', 'programmer', 1, 10000)
e5 = employee('Bob', 'programmer', 1, 50000)
  
# list of employee objects
emp = [e1, e2, e3, e4, e5]
  
# converting to min-heap
# based on yos
hq.heapify(emp)

while len(emp) != 0:
    tmp = hq.heappop(emp)
    tmp.print_me()
    print()

```

Use user defined function ([link](https://pythontic.com/algorithms/heapq/merge))

```python
import heapq

# Circuit class definition
class Circuit:
    def __init__(self, name, distance):
        self.name = name
        self.distance = distance

# Create sorted lists of circuit instances       
c0 = Circuit("Circuit0", 10)
c1 = Circuit("Circuit1", 30)
c2 = Circuit("Circuit2", 40)
i1 = [c0, c1, c2]

c3 = Circuit("Circuit3", 15)
c4 = Circuit("Circuit4", 25)
c5 = Circuit("Circuit5", 35)
i2 = [c3, c4, c5]

# Key function used for comparison while sorting
# sorting how to: https://docs.python.org/3/howto/sorting.html 
def keyfunc(circuit):
    return circuit.distance

# Merge elements from two Python iterables whose elements are already in sorted order
merged = heapq.merge(i1, i2, key=keyfunc)

# Print the merged sequence
print("Merged sequence:")
for item in merged:
    print(item.distance)
```


## Reference
+ [The 10 most common list operations in Python](https://developpaper.com/the-10-most-common-list-operations-in-python/)
+ Effective python
+ [Python: detailed notes of list, tuple, dictionary and string (super detailed)](https://developpaper.com/python-detailed-notes-of-list-tuple-dictionary-and-string-super-detailed)

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





## Reference
+ [The 10 most common list operations in Python](https://developpaper.com/the-10-most-common-list-operations-in-python/)
+ Effective python
+ [Python: detailed notes of list, tuple, dictionary and string (super detailed)](https://developpaper.com/python-detailed-notes-of-list-tuple-dictionary-and-string-super-detailed)

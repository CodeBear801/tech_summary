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

+ The only functions provided are index and count
+ Tuples are often used for multivariable assignment and multivalued return, but they are usually used without parentheses to make python look good.
+ It also supports stitching, slicing, repetition and other operations.
+ Use `namedtuple` for lightweight, immutable data containers before you need the flexibility of a full class
  + You can’t specify default argument values for namedtuple classes. -> dataclass
  + The attribute values of namedtuple instances are still accessible using numerical indexes and iteration

```python
from collections import namedtuple
Grade = namedtuple('Grade', ('score', 'weight'))


def __init__(self):
  self._grades = []

def report_grade(self, score, weight):
  self._grades.append(Grade(score, weight))
```

### list

```python
lst.sort(*, key=None, reverse=False)
lst.append(val)  # It can also be lst = lst + [val]
lst.clear()
lst.count(val)
lst.extend(t)  # or s += t  # += Actually, the _iadd_ method is called.
lst.pop(val=lst[-1])
lst.remove(val)
lst.reverse() # this changes orginal list
lst.insert(i, val)
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







## Reference
+ [The 10 most common list operations in Python](https://developpaper.com/the-10-most-common-list-operations-in-python/)
+ Effective python


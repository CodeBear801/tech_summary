# Debugging Morden C++ with GDB
Wendesday, Nov 13, 2019<br/>
7:00pm - 8:45pm<br/>
@[Mercedes-Benz Research & Development North America] By Senthilkumar Selvaraj


## Notes from lecture
- The most frequent useful tool is cout + assert
- debug tool is used for validate your assumption
- debug hints: smart small, binary search


### Frequent use commends

```
gdb [prog][core|procID]

gdbinit

s step
r run

set args
attach progid
detach progid

break function
break offset
break filename:function
Info break


watch expr
info watchpoints
catch


step
next
continue
util

```

### Highlights

<img src="resources/imgs/debug_cpp_gdb_first.png" alt="debug_cpp_gdb_first" width="600"/>
<br/>

<img src="resources/imgs/debug_cpp_gdb_how_gdb_works.png" alt="debug_cpp_gdb_how_gdb_works" width="600"/>
<br/>

<img src="resources/imgs/debug_cpp_gdb_gdb_compiling.png" alt="debug_cpp_gdb_gdb_compiling" width="600"/>
<br/>

<img src="resources/imgs/debug_cpp_gdb_stack.png" alt="debug_cpp_gdb_stack" width="600"/>
<br/>

x addr examine address
<img src="resources/imgs/debug_cpp_gdb_address.png" alt="debug_cpp_gdb_address" width="600"/>
<br/>


## My exploration

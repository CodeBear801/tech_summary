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

### Why GDB could debug your program

There's one thing to know about Linux processes: parent processes can get additional information about their children, in particular the ability to `ptrace` them. And, you can guess, the debugger is the parent of the debuggee process (or it becomes, processes can adopt a child in Linux :-).

### How is ptrace implemented
Linux [ptrace](http://man7.org/linux/man-pages/man2/ptrace.2.html) API allows a (debugger) process to access low-level information about another process (the debuggee). 
- read and write the debuggee's memory: PTRACE_PEEKTEXT, PTRACE_PEEKUSER, PTRACE_POKE...
- read and write the debuggee's CPU registers: PTRACE_GETREGSET, PTRACE_SETREGS,
- be notified of system events: PTRACE_O_TRACEEXEC, PTRACE_O_TRACECLONE, PTRACE_O_EXITKILL, PTRACE_SYSCALL (you can recognize the exec syscall, clone, exit, and all the other syscalls)
- control its execution: PTRACE_SINGLESTEP, PTRACE_KILL, PTRACE_INTERRUPT, PTRACE_CONT (notice the CPU single-stepping here)
- alter its signal handling: PTRACE_GETSIGINFO, PTRACE_SETSIGINFO

### What ptrace helps

We can see here that all the low-level mechanisms required to implement a debugger are there, provided by this ptrace API:
- Catch the exec syscall and block the start of the execution,
- Query the CPU registers to get the process's current instruction and stack location,
- Catch for clone/fork events to detect new threads,
- Peek and poke data addresses to read and alter memory variables.


### Breakpoints
how to set a breakpoint at a given address:

- The debugger reads (ptrace peek) the binary instruction stored at this address, and saves it in its data structures.
- It writes an invalid instruction at this location. What ever this instruction, it just has to be invalid.
- When the debuggee reaches this invalid instruction (or, put more correctly, the processor, setup with the debuggee memory context), the it won't be able to execute it (because it's invalid).
- In modern multitask OSes, an invalid instruction doesn't crash the whole system, but it gives the control back to the OS kernel, by raising an interruption (or a fault).
- This interruption is translated by Linux into a SIGTRAP signal, and transmitted to the process ... or to it's parent, as the debugger asked for.
- The debugger gets the information about the signal, and checks the value of the debuggee's instruction pointer (i.e., where the trap occurred). If the IP address is in its breakpoint list, that means it's a debugger breakpoint (otherwise, it's a fault in the process, just pass the signal and let it crash).
- Now that the debuggee is stopped at the breakpoint, the debugger can let its user do what ever s/he wants, until it's time to continue the execution.
- To continue, the debugger needs to 1/ write the correct instruction back in the debuggee's memory, 2/ single-step it (continue the execution for one CPU instruction, with ptrace single-step) and 3/ write the invalid instruction back (so that the execution can stop again next time). And 4/, let the execution flow normally.


### Conditional breakpoints
Conditional breakpoints are normal breakpoints, except that, internally, the debugger checks the conditions before giving the control to the user. If the condition is not matched, the execution is silently continued.

### Step into/step out

the stack trace is "unwinded" from the current frame ($sp and $bp/#fp) upwards, one frame at a time. Functions' name, parameters and local variables are found in the debug information.

### watchpoints
watchpoints are implemented (if available) with the help of the processor: write in its registers which addresses should be monitored, and it will raise an exception when the memory is read or written. If this support is not available, or if you request more watchpoints than the processor supports ... then the debugger falls back to "hand-made" watchpoints: execute the application instruction by instruction, and check if the current operation touches a watchpointed address. Yes, that's very slow!



## More info
- [How Does a C Debugger Work?](https://blog.0x972.info/?d=2014/11/13/10/40/50-how-does-a-debugger-work)


### Chapter 6: Processes

## 6.1
* Binary format specificaiton (ELF for linux, rest is old garbage)
* machine-language instructions (code section)
* program entrypoint address -> first instruction to be executed
* data -> variable initialization and constants
* symbol/relocation tables -> locations and names of functions inside program. enables runtime resolution of symbols
* shared lib/ dynamic linking information
* other info

## 6.2

* pid_t getpid(void); => returns process id
* pid_t getppid(void); => gets parent pid

## 6.3 - Memory layout

* Text segment => machine language instructins
* initialized data segment => global/static initialized variables
* uninitialized data segment => uninitialized global and static variables
* stack
* heap => dynamically allocated memory

## 6.4 - Virtual memory management

* spacial locality => generally close memory addresses are accesses together often (Sequential data)
* temporal locality => genenrally accesses same memory addresses that in recent times (loops)

Memory layout

---------------------
Kernel
-------------------- 0xc00000000
argv, environ
-------------------
stack (grows downwards)
 .
 .
 \/
--------------------

Unallocated memory

--------------------
/\
 .
 .
Heap (grows upwards)
-----------------------
uninitialized data (bss)
----------------------
initialized data
---------------------
Text (code)
-------------------- 0x08048000

------------------- 0x00000000

* locality allows us to keep only small pieces of memory in RAM (pages)
* pages are small subdivisions of virtual memory
* process virtual address space => page table => physical memory (RAM)
* virt address not in page table => SIGSEGV

## 6.5

* stack/stack frames => trivial

## 6.6

* command line args: argc e argv
* The command-line arguments of any process can be read via the Linux-specific
/proc/PID/cmdline file

## 6.7 - environment

* char *getenv(const char *name);
* int putenv(char *string);
* int setenv(const char *name, const char *value, int overwrite);
* int unsetenv(const char *name);

## 6.8 - non local jump
* int setjmp(jmp_buf env); => Calling setjmp() establishes a target for a later jump performed by longjmp()
* void longjmp(jmp_buf env, int val); => void longjmp(jmp_buf env, int val);
* The initial setjmp() returns 0, while the later “faked” return supplies whatever value is specified in the val argument of the longjmp() call. By using different
values for the val argument, we can distinguish jumps to the same target from different points in the program.



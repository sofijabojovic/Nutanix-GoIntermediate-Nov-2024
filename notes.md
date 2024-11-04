# Intermediate Go

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
- Session-01    : 01:30 hrs
- Tea Break     : 00:15 mins
- Session-02    : 01:30 hrs
- Lunch Break   : 00:45 mins
- Session-03    : 01:30 hrs
- Tea Break     : 00:15 mins
- Session-04    : 01:30 hrs

## Methodology
- No powerpoints
- Discussion & Code
- Floor is open for Q&A

## Software Requirements
- Go Tools (https://go.dev/dl)
- Any Editor
- Any git client

## Repository
- https://github.com/tkmagesh/Nutanix-GoIntermediate-Nov-2024

## Prerequisites
- Familiarity with the following
    - Date Types
    - Variables, Constants & iota
    - if else, switch case & for
    - Functions
        - Function Types
        - Higher Order Functions
    - Pointers
    - Structs & Struct composition
    - Methods & Method Overriding
    - Interfaces
    - Error Handling
    - Panic & Recovery
    - Modules & Packages

## Recap
- None

## Concurrency
- Ability to have more than one execution path
- Typically achieved using OS threads
- OS Threads are costly
    - ~ 2MB RAM per thread (in linux)
    - Thread context switch 
- In Go Applications
    - Application binary is embedded with "GC" & "Go Scheduler"
    - "Go Scheduler" starts when the application is started
    - "Go Scheduler" will in turn schedule the execution of the "main()" function
    - All other concurrent operations are scheduled through the "Go Scheduler"
- Go Routines
    - A goroutine is used to represent a "concurrent" operation
    - Are cheap (~4KB of RAM)
    - GoScheduler ensures that the best use of the OS Threads are achieved

### Go Concurrency
- Concurrency support is built in the language itself
    - go keyword
    - channel "data type"
    - channel "operator" (<-)
    - range keyword
    - select-case construct
    - close() function
- API support
    - "sync" package
    - "sync/atomic" package
#### sync.WaitGroup
    - semaphore based counter
    - has the ability to block the excution of a function until the counter becomes 0




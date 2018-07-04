# Spinlock

*Spinlock implementation in Go and inline assembler.*

## Overview

Package spinlock provides a low level implementation of spinlock in assembly.
Also, it provides fallback to implementation based on built-in atomics.

On my private laptop benchmark results are the following. Measured time is wall
time.

```
    goos: linux
    goarch: amd64
    pkg: github.com/daskol/spinlock
    BenchmarkMutex/1-4                       50000000	        37.7 ns/op
    BenchmarkMutex/2-4                       30000000	        52.4 ns/op
    BenchmarkMutex/4-4                       20000000	        64.4 ns/op
    BenchmarkMutex/8-4                       20000000	        83.8 ns/op
    BenchmarkMutex/16-4                      20000000	        91.2 ns/op
    BenchmarkSpinlockInAsm/1-4              100000000	        22.1 ns/op
    BenchmarkSpinlockInAsm/2-4              100000000	        22.1 ns/op
    BenchmarkSpinlockInAsm/4-4              100000000	        22.1 ns/op
    BenchmarkSpinlockInAsm/8-4              100000000	        22.8 ns/op
    BenchmarkSpinlockInAsm/16-4             100000000	        23.6 ns/op
    BenchmarkSpinlockInGo/1-4                50000000	        28.0 ns/op
    BenchmarkSpinlockInGo/2-4                50000000	        27.8 ns/op
    BenchmarkSpinlockInGo/4-4                50000000	        28.1 ns/op
    BenchmarkSpinlockInGo/8-4                50000000	        28.7 ns/op
    BenchmarkSpinlockInGo/16-4               50000000	        29.5 ns/op
    BenchmarkSpinlockThin/1-4               100000000	        17.8 ns/op
    BenchmarkSpinlockThin/2-4               100000000	        17.5 ns/op
    BenchmarkSpinlockThin/4-4               100000000	        17.6 ns/op
    BenchmarkSpinlockThin/8-4               100000000	        18.0 ns/op
    BenchmarkSpinlockThin/16-4              100000000	        18.4 ns/op
    PASS
    ok  	github.com/daskol/spinlock	39.539s
```

The fastest implementation based on free function Lock()/Unlock(). However, it
is two times slower than spinlock implemented in modern C++.

```
    Run on (4 X 3400 MHz CPU s)
    CPU Caches:
      L1 Data 32K (x2)
      L1 Instruction 32K (x2)
      L2 Unified 256K (x2)
      L3 Unified 4096K (x1)
    ---------------------------------------------------------------
    Benchmark                        Time           CPU Iterations
    ---------------------------------------------------------------
    run/real_time/threads:1          9 ns          9 ns   74377642
    run/real_time/threads:2          9 ns         19 ns   74705518
    run/real_time/threads:4         11 ns         44 ns   64348172
    run/real_time/threads:8         10 ns         43 ns   7931257
```

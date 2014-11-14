formam-benchmark
================

Benchmark about formam package (compared with built-in JSON, [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema))

In a Mac i7 2,8Ghz 4 cores, 8 GB 1067Mhz DDR3

### test 1

NOTE: schema not support map...

```
BenchmarkAJGForm	   50000	     44224 ns/op
BenchmarkFormam	      100000	     20741 ns/op
BenchmarkJSON	      100000	     17001 ns/op
```

### test 2

```
BenchmarkAJGForm	  100000	     21551 ns/op
BenchmarkSchema	      100000	     22500 ns/op
Benchmark2Formam	  200000	      9226 ns/op
BenchmarkJSON	      200000	      9879 ns/op
```

Conclusion
----------

Formam is twice faster than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema). And equal to or slightly slower than built-in/json.

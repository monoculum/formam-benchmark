formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a Mac i7 2,8Ghz 4 cores, 8 GB 1067Mhz DDR3

### test 1

**NOTE**: `gorilla/schema` not support map...

```
BenchmarkAJGForm	    50000	     44667 ns/op	    6170 B/op	     100 allocs/op
BenchmarkFormam	       100000	     20835 ns/op	    2070 B/op	      59 allocs/op
BenchmarkJSON	       100000	     18654 ns/op	    1793 B/op	      27 allocs/op
```

### test 2

```
BenchmarkAJGForm	  100000	     24384 ns/op	    3169 B/op	      53 allocs/op
BenchmarkSchema	      100000	     24574 ns/op	    4020 B/op	      66 allocs/op
Benchmark2Formam	  200000	      9647 ns/op	     873 B/op	      30 allocs/op
BenchmarkJSON	      200000	      9874 ns/op	     870 B/op	      13 allocs/op
```

Conclusion
----------

Formam is twice faster than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and equal to or slightly slower than [built-in/json](http://golang.org/pkg/encoding/json/).
Furthermore, Formam allocates in memory twice less than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and slightly more than [built-in/json](http://golang.org/pkg/encoding/json/).

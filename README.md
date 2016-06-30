formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form), [gorilla/schema](https://github.com/gorilla/schema), [go-playground/form](https://github.com/go-playground/form) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a iMac 2.8GHz i7, 8GB 1067MHz DDR3 and Go 1.7beta1


SCENARIOS
---------

The follow scenarios are real examples that can be seen in many projects.

### SIMPLE

```
BenchmarkAJGFormTestSIMPLE-8   	   50000	     26575 ns/op	    3153 B/op	      66 allocs/op
BenchmarkSchemaTestSIMPLE-8    	  200000	      9879 ns/op	    1376 B/op	      48 allocs/op
BenchmarkFormamTestSIMPLE-8    	  200000	      7031 ns/op	     612 B/op	      42 allocs/op
BenchmarkFormTestSIMPLE-8      	  200000	      6899 ns/op	     904 B/op	      29 allocs/op

BenchmarkJSONTestSIMPLE-8      	  200000	      9685 ns/op	     848 B/op	      17 allocs/op
```

formam is three times faster than [ajg/form](https://github.com/ajg/form), two times faster than [gorilla/schema](https://github.com/gorilla/schema), slightly more slowly than [go-playground/form](https://github.com/go-playground/form), and faster than [built-in/json](http://golang.org/pkg/encoding/json/). 
[go-playground/form](https://github.com/go-playground/form) has the least allocations.

### MEDIUM

**NOTE**: `gorilla/schema` not support map...

```
BenchmarkAJGFormTestMEDIUM-8   	   30000	     44443 ns/op	    5907 B/op	     120 allocs/op
BenchmarkFormamTestMEDIUM-8    	  100000	     17434 ns/op	    5770 B/op	      84 allocs/op
BenchmarkFormTestMEDIUM-8      	   50000	     26513 ns/op	    3977 B/op	     136 allocs/op

BenchmarkJSONTestMEDIUM-8      	  100000	     16324 ns/op	    1696 B/op	      32 allocs/op
```

formam is the fastest and has the least allocations.

### COMPLEX

```
BenchmarkFormamTestCOMPLEX-8   	    5000	    278635 ns/op	   42564 B/op	    1797 allocs/op
BenchmarkFormTestCOMPLEX-8     	    3000	    568477 ns/op	   90193 B/op	    2028 allocs/op
```

formam is the fastest and has the least allocations.

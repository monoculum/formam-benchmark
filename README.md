formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form), [gorilla/schema](https://github.com/gorilla/schema), [go-playground/form](https://github.com/go-playground/form) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a iMac 2.8GHz i7, 8GB 1067MHz DDR3 and Go 1.7beta1


SCENARIOS
---------

The follow scenarios are real examples that can be seen in many projects.

### SIMPLE

```
BenchmarkAJGFormTestSIMPLE-8   	   50000	     24670 ns/op	    3153 B/op	      66 allocs/op
BenchmarkSchemaTestSIMPLE-8    	  200000	      9791 ns/op	    1376 B/op	      48 allocs/op
BenchmarkFormamTestSIMPLE-8    	  200000	      8094 ns/op	     692 B/op	      49 allocs/op
BenchmarkFormTestSIMPLE-8      	  200000	      7515 ns/op	    1832 B/op	      35 allocs/op

BenchmarkJSONTestSIMPLE-8      	  200000	     10058 ns/op	     848 B/op	      17 allocs/op
```

formam is three times faster than [ajg/form](https://github.com/ajg/form), two times faster than [gorilla/schema](https://github.com/gorilla/schema), slightly more slowly than [go-playground/form](https://github.com/go-playground/form), and faster than [built-in/json](http://golang.org/pkg/encoding/json/). 
[go-playground/form](https://github.com/go-playground/form) has the least allocations.

### MEDIUM

**NOTE**: `gorilla/schema` not support map...

```
BenchmarkAJGFormTestMEDIUM-8   	   30000	     46213 ns/op	    5897 B/op	     120 allocs/op
BenchmarkFormamTestMEDIUM-8    	  100000	     22306 ns/op	    2075 B/op	     150 allocs/op
BenchmarkFormTestMEDIUM-8      	   50000	     29148 ns/op	    5832 B/op	     148 allocs/op

BenchmarkJSONTestMEDIUM-8      	  100000	     16204 ns/op	    1696 B/op	      32 allocs/op
```

formam is the fastest.
[ajg/form](https://github.com/ajg/form) has the least allocations.

### COMPLEX

```
BenchmarkFormamTestCOMPLEX-8   	    5000	    315707 ns/op	   40483 B/op	    2063 allocs/op
BenchmarkFormTestCOMPLEX-8     	    2000	    631384 ns/op	  107138 B/op	    2108 allocs/op
```

formam is the fastest and has the least allocations.

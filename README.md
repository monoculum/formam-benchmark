formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form), [gorilla/schema](https://github.com/gorilla/schema), [go-playground/form](https://github.com/go-playground/form) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a iMac 2.8GHz i7, 8GB 1067MHz DDR3 and Go 1.7beta1

### test 1 (COMPLEX)

**NOTE**: `gorilla/schema` not support map...

```
BenchmarkAJGFormTest1-8   	   30000	     46088 ns/op	    5843 B/op	     113 allocs/op
BenchmarkFormamTest1-8    	   50000	     25366 ns/op	    2316 B/op	     104 allocs/op
BenchmarkFormTest1-8      	   50000	     31089 ns/op	    6776 B/op	     165 allocs/op
BenchmarkJSONTest1-8      	  100000	     16695 ns/op	    1696 B/op	      32 allocs/op
```

### test 2 (SIMPLE)

```
BenchmarkAJGFormTest2-8   	   50000	     25056 ns/op	    3152 B/op	      66 allocs/op
BenchmarkSchemaTest2-8    	  100000	     17638 ns/op	    3543 B/op	      80 allocs/op
BenchmarkFormamTest2-8    	  200000	     11217 ns/op	     932 B/op	      54 allocs/op
BenchmarkFormTest2-8      	  200000	      8609 ns/op	    2504 B/op	      45 allocs/op
BenchmarkJSONTest2-8      	  200000	      9650 ns/op	     848 B/op	      17 allocs/op
```

Conclusion
----------

Formam is twice faster than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and equal to or slower than [built-in/json](http://golang.org/pkg/encoding/json/).
Furthermore, Formam allocates in memory twice less than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and slightly more than [built-in/json](http://golang.org/pkg/encoding/json/).

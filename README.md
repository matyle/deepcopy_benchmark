# deepcopy_benchmark
三种深拷贝结构体指针的方法

- 通过 gob 进行序列化和反序列化
- 通过 json 进行序列化和反序列化（json std和 sonic）
- 定制copy（针对具体结构体每个成员拷贝）


## Benchmark

```bash
goos: darwin
goarch: amd64
pkg: deepcopy
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkGob-8      	   31314	     38937 ns/op
BenchmarkJson-8     	  199303	      5705 ns/op
BenchmarkSonic-8    	  597932	      1772 ns/op
BenchmarkCustom-8   	 2943228	       469.0 ns/op
PASS
ok  	deepcopy	6.367s
```

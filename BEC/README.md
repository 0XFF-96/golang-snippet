

### What is BCE in golang ?

两种边界检查的型
1、a[i]
2、a[i:j]

### Whcih boulds
1、Duplicate checks
2、Constant slice size with masked index
3、Constant index
4、Trivial bound checks
5、a[i:j] generates two bounds checks: for 0 <= i <= j and for 0 <= j <= cap(a). Sometimes we can remove one of them:
6、Decreasing constant indexes

### Reference
1、https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit#heading=h.jhgvmxgsg5ol
2、https://go-review.googlesource.com/c/go/+/21008

====
3、https://go.googlesource.com/go/+/master/test/prove.go
4、https://go.googlesource.com/go/+/master/test/loopbce.go

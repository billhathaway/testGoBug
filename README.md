Different behavior observed with Go 1.2.2 and 1.3 on OSX
==

// the test passes on 1.3 if chars is changed to var OR the length of
// the chars string is 1,2,4,7,8  fails being var with len 3,5,6,9
// const chars = "ABC"  // fails
// var chars = "ABC"    // passes
// const chars = "ABCD" //passes

```
$ pwd
~/go/src/github.com/billhathaway/testGoBug
$  GOROOT=~/go1.2.2 PATH=~/go1.2.2/bin:$PATH go test -v
=== RUN TestV1vsV2
--- PASS: TestV1vsV2 (0.00 seconds)
PASS
ok  	github.com/billhathaway/testGoBug	0.026s
$  GOROOT=~/go1.3 PATH=~/go1.3/bin:$PATH go test -v
=== RUN TestV1vsV2
--- FAIL: TestV1vsV2 (0.00 seconds)
FAIL
exit status 1
FAIL	github.com/billhathaway/testGoBug	0.013s
```

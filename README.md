Different behavior observed with Go 1.2.2 and 1.3 on OSX
==

// go test passes on 1.3 if chars in main.go is changed to var OR the length of
// the chars string is 1,2,4,7,8  and fails with char being a const with len 3,5,6,9
// it passes on golang 1.2.2 in all cases

// const chars = "ABC"  // fails
// var chars = "ABC"    // passes
// const chars = "ABCD" //passes

```
$ pwd
~/go/src/github.com/billhathaway/testGoBug
$  GOROOT=~/go1.2.2 PATH=~/go1.2.2/bin:$PATH go test -v
=== RUN TestLetters
--- PASS: TestLetters (0.00 seconds)
PASS
ok  	github.com/billhathaway/testGoBug	0.023s
$  GOROOT=~/go1.3 PATH=~/go1.3/bin:$PATH go test -v
=== RUN TestLetters
--- FAIL: TestLetters (0.00 seconds)
FAIL
exit status 1
FAIL	github.com/billhathaway/testGoBug	0.015s
```

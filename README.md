Different behavior observed with Go 1.2.2 and 1.3 on OSX
==
The only difference between makeRandomMessageV1 and makeRandomMessageV2 functions is assigning a variable that is then assigned to _ and so should be a noop.  
On 1.2.2 both functions work the same, on 1.3, I get unexpected values back from V1  


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

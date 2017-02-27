Stella Artosis
--------------

![Stella](artosis.jpg)

Stella Artosis is a simple program used to quickly hash files on your
filesystem for use in forensic analysis.

Usage:

```
go get github.com/abraithwaite/stella-artosis
go install github.com/abraithwaite/stella-artosis/cmd/artosis
export GOMAXPROCS=`sysctl -n hw.ncpu`
artosis -include /usr/lib
```

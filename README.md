# teedate

`tee(1)` but with a date prefix for each line.

# usage

```
ping localhost | teedate --format="%Y/%m/%d" --append output1.txt output2.txt
```

result for example:

```
2015/07/09 PING localhost (127.0.0.1): 56 data bytes
2015/07/09 64 bytes from 127.0.0.1: icmp_seq=0 ttl=64 time=0.073 ms
2015/07/09 64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.156 ms
2015/07/09 64 bytes from 127.0.0.1: icmp_seq=2 ttl=64 time=0.155 ms
```

# install

```
go get github.com/kbinani/teedate
```

# license

The MIT License

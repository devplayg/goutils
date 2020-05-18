# goutils

[![Build Status](https://travis-ci.org/devplayg/goutils.svg?branch=master)](https://travis-ci.org/devplayg/goutils)

Something that enriches the Go


### Crypto

```go
EncAES(data, key []byte) ([]byte, error) 
DecAES128(data, key []byte) ([]byte, error) 
DecAES192(data, key []byte) ([]byte, error) 
DecAES256(data, key []byte) ([]byte, error) 
EncAES128(data, key []byte) ([]byte, error) 
EncAES192(data, key []byte) ([]byte, error) 
EncAES256(data, key []byte) ([]byte, error) 
```

### Compress

```go
Gunzip(s []byte) ([]byte, error) 
Gzip(data []byte) ([]byte, error) 
```

### Network

```go
IntToIPv4(nn uint32) net.IP 
IPv4ToInt(ip net.IP) uint32 
```


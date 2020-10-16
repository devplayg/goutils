# goutils

[![Build Status](https://travis-ci.org/devplayg/goutils.svg?branch=master)](https://travis-ci.org/devplayg/goutils)

Something that enriches the Go


### Crypto

```go
EncAesCbc(data, key []byte) ([]byte, error) 
DecAesCbc(data, key []byte) ([]byte, error) 
EncAes128Cbc(data, key []byte) ([]byte, error) 
EncAes192Cbc(data, key []byte) ([]byte, error) 
EncAes256Cbc(data, key []byte) ([]byte, error) 
DecAes128Cbc(data, key []byte) ([]byte, error) 
DecAes192Cbc(data, key []byte) ([]byte, error) 
DecAes256Cbc(data, key []byte) ([]byte, error)  

EncAesGcm(data, key, nonce []byte) ([]byte, error)
DecAesGcm(data, key, nonce []byte) ([]byte, error)
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


### Introduction

Given the packet sizes, this application can calculate the number of packs we need to ship to the customer

### Usage

##### define package sizes

```
curl -X POST -H "Content-Type: application/json" -d '{"packSizes":[23,31,53]}' "http://localhost:8080/pack"
```

##### query packet size

get packet size based on order size
```
curl -X GET -H "Content-Type: application/json" "http://localhost:8080/pack?orderSize=263"
```


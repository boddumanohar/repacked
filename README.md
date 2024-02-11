### Introduction

Given the packet sizes, this application can calculate the number of packs we need to ship to the customer

### start

`make dockerrun`

### Usage

##### define package sizes

we first will need to define packet size so that we can query

```
curl -X POST -H "Content-Type: application/json" -d '{"packSizes":[23,31,53]}' "http://localhost:8080/pack"
```

##### query packet size

get packet size based on order size
```
curl -X GET -H "Content-Type: application/json" "http://localhost:8080/pack?orderSize=263"
```

this will output
```
{"packSizes":[23,31,53],"packs":[2,7,0]}
```

meaning: 
```
2 x 23
7 x 31
```

#### start the application

```
docker-compose up -d --build
```

you could access the UI at http://localhost:3000

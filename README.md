### Introduction

Given the packet sizes, this application can calculate the number of packs we need to ship to the customer

### start

`make dockerrun`

### Usage

##### define package sizes

we first will need to define packet size so that we can query

```
curl -X POST -H "Content-Type: application/json" -d '{"packSizes":[250,500,1000,2000,5000]}' "http://13.58.115.118:8080/pack"
```

##### query packet size

get packet size based on order size
```
curl -X GET -H "Content-Type: application/json" "http://13.58.115.118:8080/pack?orderSize=501"
```

this will output
```
{"packSizes":[250,500,1000,2000,5000],"packs":[1,1,0,0,0]}
```

meaning: 
```
1 x 250
1 x 500
```

#### start the application

```
docker-compose up -d --build
```

### deployment

you could access the UI at http://13.58.115.118:3000


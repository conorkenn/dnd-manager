# DND manager 

Create and List characters MVP

```
$ curl -X POST http://localhost:8080/characters -H "Content-Type: application/json" -d '{"name":"Suel","race":"Tiefling","class":"Bard","level":1}'
{"name":"Suel","race":"Tiefling","class":"Bard","level":1}
```

```
$ curl -X GET http://localhost:8080/characters
[{"name":"Suel","race":"Tiefling","class":"Bard","level":1},{"name":"Suelchi","race":"Orc","class":"Monk","level":1}]
```

Dice 

```
$ curl -X POST http://localhost:8080/roll -H "Content-Type: application/json" -d '{"sides":6, "num_rolls":4}'
{"message":"Successfully rolled dice","num_rolls":4,"results":[3,4,2,3],"sides":6}
```





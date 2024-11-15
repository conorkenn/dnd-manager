# DND manager 

Create and List characters MVP

```
Invoke-RestMethod -Uri http://localhost:8080/characters -Method Post -Headers @{ "Content-Type" = "application/json" } -Body '{"name":"Suel","race":"Tiefling","class":"Bard","level":1}'

name race     class level
---- ----     ----- -----
Suel Tiefling Bard      1

Invoke-WebRequest -Uri http://localhost:8080/characters -Method Get

StatusCode        : 200
StatusDescription : OK
Content           : [{"name":"Suel","race":"Tiefling","class":"Bard","level":1}]
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 60
                    Content-Type: application/json; charset=utf-8
                    Date: Fri, 15 Nov 2024 23:18:46 GMT

                    [{"name":"Suel","race":"Tiefling","class":"Bard","level":1}]
Forms             : {}
Headers           : {[Content-Length, 60], [Content-Type, application/json; charset=utf-8], [Date, Fri, 15 Nov 2024 23:18:46 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 60

```



# CRUDhttp
Simple To Do List 

## Base requests

`http://localhost:[port]/tasks` (`GET`)  
`http://localhost:[port]/task/{id:[0-9]+}` (`GET`)  
`http://localhost:[port]/task/{id:[0-9]+}` (`POST`)
`http://localhost:[port]/tasks/{id:[0-9]+}` (`UPDATE`)  
`http://localhost:[port]/tasks/{id:[0-9]+}` (`DELETE`)

## Example Post request 
`http://localhost:[port]/task/{id:1}` 
In body of request you shloud add 
```
"name" : "Simple Task"
"status" :  "In Progress"
```

Same you need to do for Update request


## For start project

`make help` - Show list commands.  
`make build` - Create `server.exe`.  
`make start` - Create and Run `server.exe`. (dependence `build`)  
`make clean` - Remove `server.exe`.  
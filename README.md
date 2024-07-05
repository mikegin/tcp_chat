# Go TCP chat server

### run the tcp server
```
go build -o server/server server/main.go
./server/server
```


### Example Chat Clients

#### Alice
```
telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Welcome to budgetchat! What shall I call you?
alice
* bob has entered the room
* charlie has entered the room
hey guys!
[bob] hows it goign
[charlie] im out of here
* charlie has left the room
[bob] same
* bob has left the room
ok :'(
^]

telnet> quit
Connection closed.
```
#### Bob
```
telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Welcome to budgetchat! What shall I call you?
bob
* The room contains: alice
* charlie has entered the room
[alice] hey guys!
hows it goign
[charlie] im out of here
* charlie has left the room
same
^]

telnet> quit
Connection closed.
```

#### Charlie
```
telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Welcome to budgetchat! What shall I call you?
charlie
* The room contains: alice, bob
[alice] hey guys!
[bob] hows it goign
im out of here
^]

telnet> quit
Connection closed.
```
# Go TCP chat server

https://protohackers.com/problem/3

### run the tcp server
```
go build .
./tcp_chat
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
* The room contains: 
* bob has entered the room
hey bob!
* charlie has entered the room
[bob] hey alice, hey charlie!
[charlie] hey guys!
[charlie] gotta go!
* charlie has left the room
[bob] same, see ya!
* bob has left the room
off I go!
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
[alice] hey bob!
* charlie has entered the room
hey alice, hey charlie!
[charlie] hey guys!
[charlie] gotta go!
* charlie has left the room
same, see ya!
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
[bob] hey alice, hey charlie!
hey guys!
gotta go!
^]

telnet> quit
Connection closed.
```
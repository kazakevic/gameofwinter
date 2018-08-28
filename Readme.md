**Game "Winter is coming"**

[logo]: https://www.xootr.com/blog/wp-content/uploads/2015/09/winter-is-coming-300x221.jpg "Game "Winter is coming"

##### Some information:

* Project based on gorilla package, chat example: https://github.com/gorilla/websocket/tree/master/examples/chat
* App is using port 8080 for communication with socket, so be sure it's open
  
#### How to start:

* You can start bu running executable on windows -> client.exe
* Or if you have Golang installed you can use command go run *.go from main project directory
* When client is started you can connect to server using websocket, endpoint is: ws://localhost:8080/ws
* Or just use simple client added to project files (client.html)
* 
* After connection to the server you are able to use this commands
    * /start playerName - where plaeyerName is your name
    * shoot x y - where x,y are coordinates (max x=10, max y = 30)
  
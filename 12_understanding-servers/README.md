# Web Servers

Web Servers are basically simple computer programs that dispense the web page when they are requested using the web client. The machines on which this program run are usually called as a server, with both the names web server and server almost used interchangeably.

## Web Programming Synonymous Terms
* router
* request router
* multiplexer
* mux
* servemux
* server
* http router
* http request router
* http multiplexer
* http mux
* http servemux
* http server

## Clients and servers

Computers connected to the web are called **clients** and **servers**. 

In general, all of the machines on the Internet can be categorized as two types: servers and clients. Those machines that provide services (like Web servers or FTP servers) to other machines are servers. And the machines that are used to connect to those services are clients.

A simplified diagram of how they interact might look like this:

![Clients and Servers](img/Client-server.jpg)

* Clients are the typical web user's internet-connected devices (for example, your computer connected to your Wi-Fi, or your phone connected to your mobile network) and web-accessing software available on those devices (usually a web browser like Firefox or Chrome).

* Servers are computers that store webpages, sites, or apps. When a client device wants to access a webpage, a copy of the webpage is downloaded from the server onto the client machine to be displayed in the user's web browser.

### What happens when you try to reach a website behind the scenes
 
 The browser broke the URL into three parts:
 
 * The protocol ("http")
 * The server name ("www.example.com")
 * The file name ("web-server.htm")
 
 The browser communicated with a name server to translate the server name("www.example.com") into an IP Address, which it uses to connect to the server machine.The browser then formed a connection to the server at that IP address on port 80.
 
 Following the HTTP protocol, the browser sent a GET request to the server, asking for the file "https://www.example.com/web-server.htm."
 
 The server then sent the HTML text for the Web page to the browser. 
 
 The browser read the HTML tags and formatted the page onto your screen.
 
  
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
 
  ## Request And Response
  
  Request and response messages are similar. Both messages consist of:
  
  * a request/response line
  
  * zero or more header lines
  
  * a blank line (ie, CRLF by itself)
  
  * an optional message body
  
  ### HTTP Request
  
  
Request
- request line
- headers
- optional message body

Request-Line
- Method SP Request-URI SP HTTP-Version CRLF (SP = space)

example request line:
- GET /path/to/file/index.html HTTP/1.1

### HTTP response

Reponse
- status line
- headers
- optional message body

Status-Line
- HTTP-Version SP Status-Code SP Reason-Phrase CRLF

example status line:
- HTTP/1.1 200 OK
  
 ---
  
 ## RFC 7230
  
  HTTP was created for the World Wide Web (WWW) architecture and has
  evolved over time to support the scalability needs of a worldwide
  hypertext system.  Much of that architecture is reflected in the
  terminology and syntax productions used to define HTTP.
  
  ### 2.1.  Client/Server Messaging
  
  HTTP is a stateless request/response protocol that operates by
  exchanging messages across a reliable TRANSPORT- or
  SESSION-layer "CONNECTION".
  
  
  An HTTP "CLIENT" is a program that establishes a CONNECTION to a server
  for the purpose of sending one or more HTTP requests.
  
  An HTTP "SERVER" is a program that accepts CONNECTIONS in order to service HTTP requests by sending HTTP responses.
  
  The terms "client" and "server" refer only to the roles that these
  programs perform for a particular connection.  The same program might
  act as a client on some connections and a server on others.  The term
  "user agent" refers to any of the various client programs that
  initiate a request, including (but not limited to) browsers, spiders
  (web-based robots), command-line tools, custom applications, and
  mobile apps.  The term "origin server" refers to the program that can
  originate authoritative responses for a given target resource.  The
  terms "sender" and "recipient" refer to any implementation that sends
  or receives a given message, respectively.
  
  HTTP relies upon the Uniform Resource Identifier (URI) standard
  to indicate the target resource and relationships between resources.
  Messages are passed in a format similar to that used by Internet mail
  and the Multipurpose Internet Mail Extensions (MIME) (see Appendix A of
  [RFC7231] for the differences between HTTP and MIME messages).
  
  Most HTTP communication consists of a retrieval request (GET) for a
  representation of some resource identified by a URI.  In the simplest
  case, this might be accomplished via a single bidirectional
  CONNECTION (===) between the user agent (UA) and the origin
  server (O).
  
  	request   >
  UA ======================================= O
  						   <   response
  
  A client sends an HTTP REQUEST to a server in the form of a REQUEST
  MESSAGE, beginning with a REQUEST-LINE that includes a method, URI,
  and protocol version, followed by header fields
  containing request modifiers, client information, and representation
  metadata, an empty line to indicate the end of the
  header section, and finally a message body containing the payload
  body (if any, Section 3.3).
  
  A server responds to a client's request by sending one or more HTTP
  RESPONSE MESSAGES, each beginning with a STATUS LINE that includes
  the protocol version, a success or error code, and textual reason
  phrase, possibly followed by header fields containing
  server information, resource metadata, and representation metadata,
  an empty line to indicate the end of the header
  section, and finally a message body containing the payload body (if
  any, Section 3.3).
  
  A connection might be used for multiple REQUEST/RESPONSE exchanges,
  as defined in Section 6.3.
  
  The following example illustrates a typical message exchange for a
  GET request on the URI
  "http://www.example.com/hello.txt":
  
  #### Client request:
  
  GET /hello.txt HTTP/1.1
  
  User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7l zlib/1.2.3
  
  Host: www.example.com
  
  Accept-Language: en, mi
  
  
  #### Server response:
  
  HTTP/1.1 200 OK
  
  Date: Mon, 27 Jul 2009 12:28:53 GMT
  
  Server: Apache
  
  Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
  
  ETag: "34aa387-d-1568eb00"
  
  Accept-Ranges: bytes
  
  Content-Length: 51
  
  Vary: Accept-Encoding
  
  Content-Type: text/plain
  
  Hello World! My payload includes a trailing CRLF.
  
  ### 2.2.  Implementation Diversity
  
  When considering the design of HTTP, it is easy to fall into a trap
  of thinking that all user agents are general-purpose browsers and all
  origin servers are large public websites.  That is not the case in
  practice.  Common HTTP USER AGENTS include household appliances,
  stereos, scales, firmware update scripts, command-line programs,
  mobile apps, and communication devices in a multitude of shapes and
  sizes.  Likewise, common HTTP ORIGIN SERVERS include home automation
  units, configurable networking components, office machines,
  autonomous robots, news feeds, traffic cameras, ad selectors, and
  video-delivery platforms.
  
  The term "USER AGENT" does not imply that there is a human user
  directly interacting with the software agent at the time of a
  request.  In many cases, a user agent is installed or configured to
  run in the background and save its results for later inspection (or
  save only a subset of those results that might be interesting or
  erroneous).  Spiders, for example, are typically given a start URI
  and configured to follow certain behavior while crawling the Web as a
  hypertext graph.
  
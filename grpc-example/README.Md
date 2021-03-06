gRPC uses HTTP/2 for transport, Protocol Buffers as the interface description language. With gRPC we can define our service once in a .proto file and implement clients and servers. 

Protocol buffers, usually referred as Protobuf, is a protocol developed to allow serialization and deserialization of structured data, with the goal to provide a faster way to make systems communicate as compared to xml. 

Once you've defined your messages, you run the protocol buffer compiler for your application's language on your .proto file to generate data access classes. Generated classes provide accessors for each field (like name() and set_name()) as well as methods to serialize/parse the whole structure to/from raw bytes

When this message is serialized (encoded to the protocol buffer binary format) it would probably be 28 bytes long and take around 100-200 nanoseconds to parse. The XML version is at least 69 bytes if you remove whitespace, and would take around 5,000-10,000 nanoseconds to parse.

```
Installation:

 go get github.com/golang/protobuf/protoc-gen-go  
 export PATH="$PATH:$(go env GOPATH)/bin" 
```

Working with grpc:

1. Define the datastructures(protocol buffer message types) and APIs required in the proto file 

2. Once you've defined your messages, you run the protocol buffer compiler

```
To compile the proto file and generate code:
    * install protoc-gen
    * create a directory (chat) where we want the files generated by protoc to go into.
    * Run the command to compile chat.proto and generate code
    
        protoc --go_out=plugins=grpc:chat chat.proto
        
        General syntax : protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/chat.proto
```
3 Implement the chat service

 **chat.go** in package chat is the implementation of the ChatServer, the apis defined in proto file
 Uses the interface methods, register chat methods with grpc server
 
4. Start the grpc server to provide handlers for the services (apis) defined above (**server.go**)

    ~~~
    Start grpc server 
    In server.go -> 
        The grpc server is started
        Listeners are added for chat service (register the chat service defined in the proto file)
        
    ~~~
 
5. Send messages to the server from a client:

  **client.go** initializes a client and sends a message using the method defined in the proto file (sayHello)
 (Refer chat.pb.go for available methods and data structures)


References

* https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
* https://developers.google.com/protocol-buffers/docs/gotutorial

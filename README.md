# Go Microservices

Codes in this repo is taken from [this](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) video series.

## Ep.1

Creating a basic HTTP server.

## Ep.2 

Better project structure and gracefully shutdown.
 
## Ep.3 

- Usage of JSON in Golang
- Power of the custom types

## Ep.4 

I understand why standard http package is not good for big project.

## Ep.5 

Refactor code with Gorilla framework.

## Ep.6 

Add some JSON validation.

## Ep.7

I loved documenting with Swagger. I should definetely go deep dive into it.

## Ep.8 

Auto-generating HTTP client with swagger.

## Ep.9

Allow CORS. Here is a good [article](https://medium.com/@baphemot/understanding-cors-18ad6b478e2b) about CORS.

## Ep.10-11

Uploading files and Multipart

## Ep.12 

Add gzip encoding support.

## Ep.13

Protocol Buffers:
- [Google - Protocol Buffers Overview](https://developers.google.com/protocol-buffers/docs/overview)
> gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

## Ep.14 

Communicating with other Microservices is very powerfull. I accidentally broke one of my Microservices and the other one worked without an error. This is awesome

## Ep.15 

"No one writes good code at first time, usually". Importance of refactoring.

## Ep.16-17

Bi-directional connections. "AWESOME"

## Ep.18-19

Richer error messages in gRPC.

# Conclusion

I learned so many things in this video series. Here are some of them:

- Making mistakes is a part of "SUCCESS": Nic is an experienced developer and I watched him making mistakes and fixing those mistakes.
- No one can write good code at the first time: Refactoring is very important. As in every expertise in the world, writing good codes wants some time.
- Streaming.
- Power of the microservices.
- Protocol Buffers: I believe we are using unnecessarily some human-readable protocols for communicating (computer-to-computer). If we don't have to read, we should use serialized protocols like ProtoBufs.
- gRPC: I wasn't imaging gRPC something like this. I'll learn this deeply.
- Building better documentations.

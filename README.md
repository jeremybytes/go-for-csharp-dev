CodeTour: Go for the C# Developer
=======================
The aim of this project is to provide an overview of Go (golang) language features for developers familiar with the C# language.

CodeTour  
----------
This repository contains several CodeTours to walk you through the various parts of the project code. CodeTour is an extension for Visual Studio Code that can be found here: **[CodeTour](https://marketplace.visualstudio.com/items?itemName=vsls-contrib.codetour)**.  

To run the CodeTour, open the root folder of this repo in Visual Studio Code, expand the "CodeTour" item in the Explorer, and click the "Start Tour" button (the green arrow) next to the tour.  

**Tours**
* *1 - Go Syntax and Conventions*  
This tour goes through various syntax rules and programming conventions for the Go language. It emphasizes similarities and differences with C#.  
* *2 - Get Data from a Service*  
This tour walks through the function to get data from a web service. This includes calling the service, parsing the JSON result, and handling errors.  
* *3 - Parallel Programming*  
This tour walks through the main function that makes service calls in parallel. This includes goroutines, using channels for communication, and using WaitGroup.

More information on CodeTour: **[CodeTour and Visual Studio Code](https://jeremybytes.blogspot.com/2020/08/codetour-and-visual-studio-code.html)**

Motivation
----------
Learning other programming languages enhances our work in our primary language. From the perspective of a C# developer, the Go language (golang) has many interesting ideas. Go is opinionated on some things (such as where curly braces go and what items are capitalized). Declaring an unused variable causes a compile failure; the use of "blank identifiers" (or "discards" in C#) are common. Concurrency is baked right in to the language through goroutines and channels. Programming by exception is discouraged; it's actually called a "panic" in Go. Instead, errors are treated as states to be handled like any other data state. We'll explore these features (and others) by building an application that uses concurrent operations to get data from a service. These ideas make us think about the way we program and how we can improve our day-to-day work (in C# or elsewhere).  

*Video Walkthrough: If you prefer a video walkthrough of the code, take a look at this video on YouTube: [A Tour of Go (golang) for the C# Developer](https://www.youtube.com/watch?v=NW-8WpnGQtE).*

Project Layout
--------------
You can go through the CodeTour without having Go installed on your machine.  

If you want to build and run the code, you will need to have both Go and .NET 5 installed on your machine.

This project assumes that you have both "go" (this was created with version 1.14.5) and "dotnet" (this was created with 5.0.101) installed. Visit [https://golang.org/doc/install](https://golang.org/doc/install) and [https://dotnet.microsoft.com/download](https://dotnet.microsoft.com/download) to download the tools. In addition, I highly recommned using the "Go" extension with Visual Studio Code: [https://marketplace.visualstudio.com/items?itemName=golang.go](https://marketplace.visualstudio.com/items?itemName=golang.go). Samples should work on all platforms supported by the runtimes (Windows, macOS, and Linux).

**/async** contains the Go program  
**/net-people-service** contains a .NET 5 service (used by the Go program)  
**/go-people-service** contains a Go service (equivalent to the .NET 5 service). This is mainly here to show how the service can be re-created using Go.

The Go program is a console application that calls the .NET service and displays the output. In order to show concurrency, the application gets each record individually.

Running the Service
-------------------  
The **.NET service** can be started from the command line by navigating to the ".../net-people-service" directory and typing `dotnet run`. This provides endpoints at the following locations:

* http://localhost:9874/people  
Provides a list of "Person" objects. This service will delay for 3 seconds before responding. Sample result:

```json
[{"id":1,"givenName":"John","familyName":"Koenig","startDate":"1975-10-17T00:00:00-07:00","rating":6,"formatString":null},  
{"id":2,"givenName":"Dylan","familyName":"Hunt","startDate":"2000-10-02T00:00:00-07:00","rating":8,"formatString":null}, 
{...}]
```

* http://localhost:9874/people/ids  
Provides a list of "id" values for the collection. Sample:  

```json
[1,2,3,4,5,6,7,8,9]
```

* http://localhost:9874/people/1  
Provides an individual "Person" record based on the "id" value. This service will delay for 1 second before responding. Sample record:

```json
{"id":1,"givenName":"John","familyName":"Koenig","startDate":"1975-10-17T00:00:00-07:00","rating":6,"formatString":null}
```

The Go Service
---------------------
The **Go service** can be started from the command line by navigating to the ".../go-people-service" directory and typing `go build` to build the service, then typing `.\go-people-service.exe` to run the service in Windows, or `./go-people-service` to run the service in Linux or macOS.  

Note: This service uses the same endpoints as the .NET service (on localhost:9874). You will want to run the .NET service **or** the Go service, but not both.

**CodeTour**  
At some point, I may put together a CodeTour that walks through the Go service, but that is not in place at this time.

The Go Sample Program
---------------------
The **/async** folder contains the "main.go" file which is the completed project.  

If you would like a step-by-step walkthrough of how this application was built, check out the video walkthrough ([A Tour of Go (golang) for the C# Developer](https://youtu.be/NW-8WpnGQtE)) and the accompanying GitHub repo ([https://github.com/jeremybytes/video-go-for-csharp-dev](https://github.com/jeremybytes/video-go-for-csharp-dev))

Other Topics
------------
This is by no means an exhaustive look at Go. Additional topics and topics to look into further include packages, exports, project structure, types, interfaces, and pointers.  

Here are a few articles (written by me):  
* [Go (golang) Loops - A Unified "for"](https://jeremybytes.blogspot.com/2021/01/go-golang-loops-unified-for.html)  
* [Go (golang) defer - A Better "finally"](https://jeremybytes.blogspot.com/2021/01/go-golang-defer-making-sure-something.html)  
* [Go (golang) Error Handling - A Different Philosophy](https://jeremybytes.blogspot.com/2021/01/go-golang-error-handling-different.html)  
* [Go (golang) Multiple Return Values - Different from C# Tuples](https://jeremybytes.blogspot.com/2021/01/go-golang-multiple-return-values.html)  
* [Go (golang) Goroutines - Running Functions Asynchronously](https://jeremybytes.blogspot.com/2021/01/go-golang-goroutines-running-functions.html)  
* [Go (golang) Channels - Moving Data Between Concurrent Processes](https://jeremybytes.blogspot.com/2021/01/go-golang-channels-moving-data-between.html)  
* [Go (golang) WaitGroup - Signal that a Concurrent Operation is Complete](https://jeremybytes.blogspot.com/2021/02/go-golang-waitgroup-signal-that.html)
* [Go and Interfaces](https://jeremybytes.blogspot.com/2020/07/go-and-interfaces.html)  

Other Resources:
* Learn Go interactively: [A Tour Of Go](https://tour.golang.org/)  
* The official site: [golang.org](https://golang.org/)  
* Installation: [Getting Started](https://golang.org/doc/install)  
* Visual Studio Code extension: [Go Language Extension](https://code.visualstudio.com/docs/languages/go)  
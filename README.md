# volume_fi

This project repo is designed for volume.finance exam.

## Basic Design

This code is organized in a standard way for golang APIs served on a simple golang server.

    1. / API: A default API to just serve a message whenever a  localhost:8080 is hit.
    2. /calculate API: This API reads a JSON request of slice of slices of strings, e.g. [["A","B"],["B","C"]], which are assumed to be source and destinations of a flight and it responds back with a slice of strings e.g. ["A","B","C"] which is the calculated flight path with the first entry in the slice as source and the last entry in the slice as destination.

## Code Organization

This code is organized under $GOPATH/src/volume_fi. The packages are as follows:

    -- volume_fi
        |-- api
            |-- api_handler.go
        |-- utilities
            |- utils.go
        |-- main.go
        |-- go.mod
        |-- Readme.md
        |-- build
        |-- clean
        |-- test

## Getting Started

Using a git clone command should allow you to clone this repo

    git clone <url>

After cloning the repo, change the diretory to volume_fi. Then you can use the build script and build the code.

        ./build

At this moment, it should build and work seamlessly. If any of the tests fail, the build would fail. You can alter that behavior in the script but we ***highly recommend NOT to do so.***

## Test, Build and Clean (Testing and Coverage)

To build and run it in a simple way, use

    ./build
    ./volume_fi

Or

    go build -race -o volume_fi
    ./volume_fi

To directly run, use

    go run main.go

There are three bash files in the repo. It is suggested that you use "build" bash file, which would run all the tests first and then gives you a binary which you can run. You can always use 

    ./clean

to clean any generated files or binaries.

NOTE:- please make sure that you run chmod command to make all the bash files executable, if they're not already. You can do so as:

    chmod +x test

OR

    chmod 777 test

Similar statments for build and clean.

## Usage

### / default API

Go to a browser or Postman and run a simple GET request with localhost:8080/ or 127.0.0.1:8080/ and that should serve you the default message

    `{"message":"The server is up. Use /calculate endpoint for flight tracking"}`

### /calculate API

Go to Postman and create a POST request with the following JSON input:

    {
        "flight_list":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
    }

and send the request. You should receive the response:

    {
        "flight_path": [
            "SFO",
            "ATL",
            "GSO",
            "IND",
            "EWR"
        ]
    }

## Prerequisites

There are no special prerequisites for this code to run. Just need to include this under your src folder in $GOPATH, and it should work smoothly.

## Scripts Heirarchy

The 'build' script calls 'test' script by default as we want to implement TDD as our standard practice. Now, the 'test' script call the 'clean' script within itself to clean all / any temporary files generated during testing. Make sure if you add your test cases and those generate some temporary files, you add them in 'clean' script to get them removed. Also, you can use both 'test' and 'clean' script individually according to your needs.

NOTE :- The code has ample comments to make life easy.

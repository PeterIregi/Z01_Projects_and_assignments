# TCPChat

A NetCat-inspired TCP group chat server written in Go.

This project recreates the behavior of the `nc` command in server mode using a client-server architecture. The server supports multiple concurrent TCP clients, enforces connection limits, maintains message history, and broadcasts join/leave events.

---

## Features

- TCP server supporting multiple clients (1 → many)
- Maximum 10 simultaneous connections
- Mandatory non-empty username
- Broadcast messaging to all connected clients
- Join and leave notifications
- Message history sent to newly connected clients
- Timestamped messages in required format:

  ```
  [2006-01-02 15:04:05][username]:message
  ```

- Empty messages are not broadcasted
- Server continues running when a client disconnects
- Concurrency implemented using:
  - Goroutines
  - `sync.Mutex`

---

## Architecture

The project follows a clean, modular structure with two binaries:

```
net-cat/
│
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── client/ # NOTE - Will be a later implementation
│       └── main.go
│
├── internal/
│   ├── server/
│   │   ├── server.go # Start, HandleConnection
│   │   ├── client.go # ReadLoop
│   │   ├── chat.go # AddClient, RemoveClient ,Broadcast
│   │   └── message.go # FormatMessage
│   │   └── conn-handlers.go # promptUserForName, GetValidUserName, CheckNameExists, SendMsgHistoryToNewConn
│   │
│   └── utils/
│       └── welcome-ascii.go
|       └── parse-os-args.go
│
├── tests/
│   └── server_test.go
│
├── go.mod
└── README.md
```

## State Flow Summary

### When Client Connects

1. HandleConnection
2. Validate name
3. AddClient
4. Send history
5. Broadcast join message
6. Start ReadLoop

### When Client Sends Message

1. ReadLoop receives text
2. Ignore empty
3. Create Message
4. Append to history
5. Broadcast

### When Client Disconnects

1. RemoveClient
2. Broadcast leave message
3. Close connection

## To start project TCP server

```bash
go run cmd/server/main.go
```

Will start a server on localhost and listening at default port `8989`

OR

```bash
go run cmd/server/main.go 2525
```

Will start a server on localhost and listening at the specified port 2525

## Connecting Clients

You can use NetCat to connect:

```bash
nc localhost 8989 -v
```

On connection, the server responds with:

```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:
```

- Name must not be empty.
- If server is full (10 clients), connection is rejected.

---

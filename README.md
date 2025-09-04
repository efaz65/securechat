# 🔒 Secure Ephemeral Chat

![Go](https://img.shields.io/badge/Language-Go-blue) ![License](https://img.shields.io/badge/License-MIT-green)

**Secure Ephemeral Chat** is a lightweight, end-to-end **AES-256-GCM encrypted chat system** written in Go. Designed for **temporary, highly secure communication**: messages are never stored, and all session data vanishes when the program closes.

---

## Features
- 🔐 **AES-256-GCM Encryption** – Strong end-to-end encryption.  
- 🕵️ **Ephemeral Sessions** – No logs, no disk writes; memory cleared on exit.  
- 🗝️ **Manual Key Input** – Each session requires a unique 32-byte key.  
- 🖥️ **Cross-platform** – Works on Windows, Linux, macOS.  
- ✨ **Screen Wipe** – Terminal buffer automatically cleared on exit.  
- 🌐 **LAN or Internet** – Remote access via port forwarding (TCP 9000).  

---

## Components
| File | Description |
|------|-------------|
| `server.go` | Server program, listens for incoming client connections. |
| `client.go` | Client program, connects to server for secure messaging. |
| `crypto.go` | AES-256-GCM encryption/decryption helpers. |
| `keygen.go` | Generates strong random 32-byte keys. |


---

## Quick Start

```bash
# 1. Generate a random 32-byte key
go run keygen.go

# 2. Start the server
go run server.go crypto.go 

# 3. Start the client (on same or different machine)
go run client.go crypto.go 

# 4. Chat Securely
```text
# Example
You: Hello!
Client: Hi, this is encrypted!
```
All messages are encrypted end-to-end.

Closing the program erases all traces.


**Security Notes**

- Use a unique key for each session.
- Run the server only when needed.
- Terminal buffers are cleared automatically.
Forward TCP port 9000 for remote connections.

**Future Improvements**

- Auto session timeout / idle shutdown.
- NAT punch-through / auto-discovery for remote connections.
- GUI interface for easier messaging.

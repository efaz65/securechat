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
| `wipe.go` | Clears terminal screen buffers on exit. |

---

## Quick Start

```bash
# 1. Generate a random 32-byte key
go run keygen.go

# 2. Start the server
go run server.go crypto.go wipe.go

# 3. Start the client (on same or different machine)
go run client.go crypto.go wipe.go

# 4. Chat Securely
```text
# Example
You: Hello!
Client: Hi, this is encrypted!

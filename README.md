Secure Ephemeral Chat 🔒

Secure Ephemeral Chat is a lightweight, end-to-end AES-256-GCM encrypted chat system written in Go. It is designed for temporary, highly secure communication: no messages are stored, and all session data vanishes when the program closes.

Features

🔐 AES-256-GCM Encryption – Military-grade end-to-end encryption.

🕵️ Ephemeral Sessions – No logs or disk storage; memory cleared on exit.

🗝️ Manual Key Input – Each session requires a fresh 32-byte key.

🖥️ Cross-platform – Works on Windows, Linux, and macOS.

✨ Screen Wipe – Terminal buffer cleared on exit.

🌐 LAN or Internet – Remote access via port forwarding (TCP 9000).

Components
File	Description
server.go	Server program, listens for incoming client connections.
client.go	Client program, connects to server for secure messaging.
crypto.go	AES-256-GCM encryption/decryption helpers.
keygen.go	Generates a strong random 32-byte key.
wipe.go	Clears terminal screen buffers on exit.
Usage
1. Generate a Key
go run keygen.go


Copy the generated 32-byte hex key.

2. Start the Server
go run server.go crypto.go wipe.go


Enter the key when prompted.

3. Start the Client
go run client.go crypto.go wipe.go


Enter the same key and server IP (127.0.0.1 for local, public IP for remote).

4. Chat Securely

Messages are encrypted end-to-end.

Close the program to erase all traces.

Security Notes

Each session must use a unique key for maximum security.

Server should only run when actively in use.

Terminal buffers are cleared automatically on exit.

Recommended to forward TCP port 9000 for remote connections.

Future Improvements

Auto session timeout / idle shutdown.

NAT punch-through or auto-discovery for remote connections.

GUI interface for easier messaging.

License

This project is licensed under MIT License – see LICENSE for details.

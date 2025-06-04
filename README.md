[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/license/mit/)

# Simple TCP Tunnel application

We were tired of complicated `netsh`,`iptables`,`socat` commands, so we wrote this simple TCP tunnel application in Go.




## Description
This TCP tunnel application allows you to forward TCP connections from one port to another. It can be useful for debugging or accessing services that are not directly exposed or proxy between two networks.

The application has no external dependencies.

It supports both IPv4 and IPv6 connections.

## Usage
To run the application as standalone:

```bash
./gotunnel -listen <source-address>:<source-port> -connect <target-address>:<target-port>
```

## Command Line Options
- `-listen <source-address>:<source-port>`: The address and port to listen for incoming connections.
- `-connect <target-address>:<target-port>`: The address and port to connect to for forwarding the traffic.
- `-help`: Show the help message with usage instructions.

if `source-address` or `target-address` is not specified, it defaults to all interfaces.

### Example
To forward traffic from port 8080 on all interfaces to port 80 on `example.com`, you can run:

```bash
./gotunnel -listen :8080 -connect example.com:80
```



## Systemd Service

Example of a systemd service file to run the application as a service (`gotunnel.service`):

```ini
[Unit]
Description=Simple TCP Tunnel Service
After=network.target

[Service]
ExecStart=/path/to/gotunnel -listen <source-address>:<source-port> -connect <target-address>:<target-port>
Restart=always

[Install]
WantedBy=multi-user.target

```

Place the `gotunnel.service` file in `/etc/systemd/system/` and enable it with:

```bash
sudo systemctl enable gotunnel.service
sudo systemctl start gotunnel.service
```




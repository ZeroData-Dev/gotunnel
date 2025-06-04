# Simple TCP Tunnel application

## Description
This is a simple TCP tunnel application that allows you to forward TCP connections from one port to another. It can be useful for debugging or accessing services that are not directly exposed or proxy between two networks.

The application has no external dependencies.

It supports both IPv4 and IPv6 connections.

## Usage
To run the application as standalone:

```bash
./gotunnel -listen <source-address>:<source-port> -connect <target-address>:<target-port>
```

### Systemd Service
```ini
[Unit]
Description=Simple TCP Tunnel Service
After=network.target

[Service]
ExecStart=/path/to/gotunnel -listen <source-address>:<source-port> -connect <target-address>:<target-port>
Restart=always

[Install]
WantedBy=multi-user.target
To run the application as a systemd service, you can create a service file named `gotunnel.service` with the following content:

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

To run the application as a service, you can use the provided systemd service file. Place the `gotunnel.service` file in `/etc/systemd/system/` and enable it with:

```bash
sudo systemctl enable gotunnel.service
sudo systemctl start gotunnel.service
```




# Networking Tips

## IP

### Get global IP address of my machine

```sh
curl https://httpbin.org/ip
```

> [!NOTE]
> `curl ifconfig.me` is deprecated.

## Port

### Port Scanning

```sh
nc -vz <hostname> <port>
```

e.g.
```sh
# single port
nc -vz localhost 22

# range of port
nc -vz localhost 1000-2000
```

### List all linstening ports

```sh
sudo ss --listening --numeric --tcp --udp --processes
```
```
       -l, --listening
              Display only listening sockets (these are omitted by default).

       -n, --numeric
              Do not try to resolve service names. Show exact bandwidth values, instead of human-readable.

       -t, --tcp
              Display TCP sockets.

       -u, --udp
              Display UDP sockets.

       -p, --processes
              Show process using socket.
```

## Proxy

### Get proxy's URL on PowerShell

```ps1
([System.Net.WebRequest]::GetSystemWebproxy()).GetProxy('http://google.com')
```
You can use the value of `OriginalString` as `$http_proxy` and `$https_proxy` varialbe.

# Networking Tips

## IP

### Get global IP address of my machine
Run `curl ifconfig.me`

## Port Scanning
```sh
nc -vz <hostname> <port>
```
eg.  
```sh
# single port
nc -vz localhost 22

# range of port
nc -vz localhost 1000-2000
```

## Proxy

## Get proxy's URL on PowerShell
```ps
([System.Net.WebRequest]::GetSystemWebproxy()).GetProxy('http://google.com')
```
You can use the value of `OriginalString` as `$http_proxy` and `$https_proxy` varialbe.

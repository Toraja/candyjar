# Tips for Linux

## Table of Contents
<!-- vim-markdown-toc GFM -->

* [Proxy](#proxy)
  * [curl](#curl)
  * [apt-get](#apt-get)
  * [yum](#yum)
  * [pip/easy_install](#pipeasy_install)
* [User](#user)
  * [Add user to sudoers](#add-user-to-sudoers)
* [Networking](#networking)
  * [Assign static IP address](#assign-static-ip-address)
* [CLI only mode](#cli-only-mode)
  * [Increase font size](#increase-font-size)
  * [Supperss visual bell](#supperss-visual-bell)
  * [Increase command line width](#increase-command-line-width)
  * [Copy and paste text on console](#copy-and-paste-text-on-console)
* [Ubuntu](#ubuntu)
  * [Change timezone](#change-timezone)
* [CentOS](#centos)
  * [Update package management cached](#update-package-management-cached)
* [Xubuntu](#xubuntu)
  * [Execute script after logout](#execute-script-after-logout)
  * [Add startup application](#add-startup-application)

<!-- vim-markdown-toc -->

## Proxy
<mark>NOTE</mark> <span>Use URL encoding for special characters (e.g. %20 for Spaces)</span>  

Setting environment variable as below generally works.
```sh
export http_proxy='http://[<user>:<password>]@<proxyhost>[:<port>]'
export https_proxy='https://[<user>:<password>]@<proxyhost>[:<port>]'
```
However, for commands that requires `sudo`, since `sudo` will be executed in not sub-session but
totally separate session (_this is my guess, citation necessary_), use `sudo -E`.  
(`-E` option preserve environment of the current session)  
### curl
Use option  
```sh
# -x or --proxy
curl -x 'http[s]://[<user>:<password>]@<proxyhost>[:<port>]'
```

### apt-get
```aptconf
# add below to /etc/apt/apt.conf
Acquire::http::Proxy "http://[<user>:<password>]@<proxyhost>[:<port>]";
Acquire::https::Proxy "https://[<user>:<password>]@<proxyhost>[:<port>]";
```

### yum
```ini
# add below to /etc/yum.conf
[main]
------
proxy=http://<Proxy-Server-IP-Address>:<Proxy_Port>
proxy_username=<Proxy-User-Name>
proxy_password=<Proxy-Password>
------
```

### pip/easy_install
Same as [curl](#curl)  
(On Ubuntu, `--proxy` option does not work apparently.)  

## User
### Add user to sudoers  
`usermod -aG wheel username`


## Networking
### Assign static IP address  
Edit `/etc/sysconfig/network-scripts/ifcfg-*interface name`  

Reference:  
- [11.1 About Network Interfaces](https://docs.oracle.com/cd/E37670_01/E41138/html/ol_about_netconf.html)  
- [FAQ/CentOS6 - CentOS Wiki](https://wiki.centos.org/FAQ/CentOS6)  
- [13.2. Interface Configuration Files](https://www.centos.org/docs/5/html/Deployment_Guide-en-US/s1-networkscripts-interfaces.html)  


## CLI only mode
### Increase font size
Run `setfont <font>`  
Available fonts and my favorite font for each OS are:

|   OS   |          Path         | Favorite |
|:------:|:---------------------:|:--------:|
| CentOS | /lib/kbd/consolefonts | sun12x22 |

In `.bashrc`, add the code below as `$TERM` is *linux* while cygwin via ssh access is *xterm*.  
```sh
if [[ $TERM = "linux" ]]; then
  setfont <font>
fi
```

### Supperss visual bell
`set bell-style none`  
(executing on the command line does not work but putting in inputrc does)  

### Increase command line width
<mark>NOTE</mark> <span style="color: ">This might differ in OS other than CentOS</span>  
- Open `/etc/default/grub` and add `vga=791` to `GRUB_CMDLINE_LINUX`  
  (791 means 1024x768x16. This is the largest resolution supported)  
- Run `grub2-mkconfig -o /boot/grub2/grub.cfg` (make sure to backup original cfg file with `cp -a`)  
- Reboot  

### Copy and paste text on console
use `screen` command  

## Ubuntu
### Change timezone
```sh
# change
sudo timedatectl set-timezone <time-zone>
# list possible time zone options
timedatectl list-timezones
```

## CentOS
### Update package management cached
`yum makecache fast`  

## Xubuntu
### Execute script after logout
<mark>NOTE</mark> <span style="color: ">Not before the logout so that controlling application cannot be done via this method</span>
1. write a script which you want to be executed
1. add "session-cleanup-script=/path/to/script" to the below
1. /etc/xdg/xdg-xubuntu/lightdm/lightdm-gtk-greeter.conf
<source>
http://askubuntu.com/questions/293312/execute-a-script-upon-logout-reboot-shutdown-in-ubuntu

### Add startup application
1. add application to /etc/xdg/autostart/
2. Add an initscript.
Create a new script in /etc/init.d/myscript


[Reference](http://askubuntu.com/questions/228304/how-do-i-run-a-script-at-start-up)

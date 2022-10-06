id
which curl
curl edge
ps x
id
cat /etc/hosts
which ping
which curl
ping 172.20.0.2
ping 172.20.0.4
ping 172.20.0.3
ping 172.20.0.0
ping 172.20.0.255
for i in $(seq 1.255); do echo $i; done
for i in $(seq 1 255); do echo $i; done
ifconfig
for i in $(seq 1 255); do echo ping -w1 -c1 172.20.0.$i; done
for i in $(seq 1 255); do ping -w1 -c1 172.20.0.$i; done
curl 172.20.0.2
curl 172.20.0.2:8080
curl 172.20.0.2:8000
curl 172.20.0.2:8000/
curl 172.19.0.2:8000
curl 172.19.0.2:8000/files
curl 172.19.0.2:8000/files/
curl 172.19.0.2:8000/files/flag
ifconfig
env
ls /config
cat /etc/ssh/sshd_config 
cd /etc/ssh
rgrep 2222
grep -r 2222
ls
grep 222 sshd_config
grep 22 sshd_config
ps x
ps ax

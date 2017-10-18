周期性的探测外网VIP的端口可用性
第一步：写好脚本
#!/bin/bash

for ip in $(cat totalvip.txt)
do
    curl --connect-timeout 2 $ip:443 >/dev/null 2>&1
    if [ $? -ne 0 ];then
        send_rtx.pl "dongfeng" "$ip:The port 443 is not unreachable!"
    fi
done

for ip in $(cat totalvip.txt)
do
    curl --connect-timeout 2 $ip:80 >/dev/null 2>&1
    if [ $? -ne 0 ];then
        send_rtx.pl "dongfen" "$ip:The port 80 is not unreachable!"
    fi
done

第二步：写好定时器
crontab -e

*/5 * * * * (./vipmonitor.sh) >/dev/null 2>&1

# Srcdomain_monitor 
## SRC子域名资产监控
> 采用go造了个轮子，release里面有编译好的程序，直接就能跑啦，目前有mac和linux的。~~就是重复的造了一个个轮子~~
---

1. 下载release对应版本
2. 设置config.yaml里的domain参数 
如 
`domain: ["baidu.com","xiami.com","其他你想要的域名"] `
3. 添加config.yaml里的server酱的apikey

执行一遍程序，顺利的话，微信会收到通知啦～

注：使用crontab等**计划任务**即可达成每日监控并推送～ 
```bash 
crontab -e 

37 10 * * * /home/ubuntu/linux_adm64_SRCdomian_monitor/cron.sh  >/var/log/cron.log 2>&1
```

有问题欢迎提issue，如果能帮到您的话，麻烦给个star🌟，谢啦。

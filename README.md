# f-server 
基于fasthttp简单文件服务器，若你部署了[traefik](https://github.com/containous/traefik)现代反向代理，需要一个简单的静态文件服务器。你可能需要本服务。

# 部署

## docker

```
docker run -d --name fserver \
-l traefik.frontend.rule=Host:<你的域名>\
-v <你的文件夹>:/var/www \
hub.faas.pro/fserver
```

用你的实际信息替换以上<**>部分。


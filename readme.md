# ali_eci_exporter
Alibaba Cloud ECI Exporter For Prometheus

阿里云ECI(弹性容器实例)Prometheus Exporter

## 背景
本人正在使用阿里云ECI服务。但是ECI目前并没有提供监控的服务，但是有对外开放API。

为了保证服务稳定性,特此开发此插件。将ECI监控数据导入Prometheus。并可以通过Grafana完成图表绘制

## 安装
### Docker
推荐使用Docker安装
~~~bash
docker run -p 8080:8080 -d -e ALICLOUD_ACCESSKEY=阿里云ACCSSKEY -e ALICLOUD_ACCESSSECRET=阿里云AccessSecret -e ALICLOUD_REGION=ECI所在地域   qiqizjl/ali_eci_exporter
~~~
然后在Prometheus注册该exporter即可。在这里不在说明，可查Prometheus官网文档哦

# WebHooks

> 通过配置 Yaml 文件的 WebHooks 部署脚本
>
> 支持 Github、Gitee、Coding、Gitlab、Gitea 及 自定义事件

## 什么是 WebHooks

作为一名运维的同学，前端的同学提了一个需求：

> ”最近刚学 Jquery，想和后端同学做一个笔记平台，你可以帮我部署一下嘛？“

你很开心，自己终于能帮上同学的忙了：

> ”好的，你们先把代码 push 到 git 平台上，我部署好了告诉你！“

随着 git 使用的越发熟练，你的工作量越来越大，你想到了一个办法：

> ”可以使用脚本自动部署到服务器上，这样我就不用每次都去部署了！“

于是你编写一个全自动的 shell 脚本，使用 cron 将脚本添加到定时任务中，每隔一段时间就会自动 clone 一次代码：

> ”这就是自动化运维吗？终于可以去摸鱼了，太棒了！“

直到有一天，你发现自己 64c 128g 的服务器怎么这么越卡，一看性能监控：

> ”卧槽，为什么每隔十分钟 CPU 占用就突然 120%，五分钟后才恢复到正常水平！“

接着你就被卷到了，后端同学居然在一个笔记平台上使用这么多技术：

> ”这些都是什么鬼？“
>
> `Kafka、Hadoop、Spark、Hive、Hbase、Flink、Zookeeper、Elasticsearch、Redis、Mysql、MongoDB、PostgreSQL、Neo4j、Cassandra、RabbitMQ、Memcached、Nginx、Tomcat、Jetty、ActiveMQ、Solr、Haproxy、Zookeeper、Docker、Kubernetes、Jenkins、Gitlab、Sonarqube、Nexus、Jira、Confluence、Rundeck、Zabbix、Grafana、Prometheus、Kibana、Elasticsearch、Graylog、Nacos、Sentinel、RocketMQ、SkyWalking、Seata、Dubbo、ShardingSphere`

每次提交都自动部署一次，就算神威·太湖之光来了恐怕也不够用，你只能选择放弃自动部署，思考一个更好的解决方案：

> ”能不能提交 tag 以后才能自动部署一次？“

于是你了解到了 WebHooks，它可以根据 git 操作，通过 http 请求触发服务器上配置的事件，你决定使用 WebHooks 来实现自动化部署：

> ”欸，这个项目好像不错，那就用这个吧！“

就这样本项目多了一个 star。

## 快速开始

> 正在筹划 cli 的版本，可以直接使用命令行来部署，敬请期待！
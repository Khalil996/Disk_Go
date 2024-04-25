# Disk_Go

```使用到的命令
创建文件
goctl api new Disk
启动项目
go run disk.go -f etc/disk-api.yaml
更新api接口
goctl api go -api Disk.api -dir . -style go_zero
```

## elk配置

```配置内容
   logstash.yaml:
   http.host: "0.0.0.0"
   xpack.monitoring.elasticsearch.hosts: [ "http://172.29.112.1:9200" ]
   xpack.monitoring.enabled: true
   logstashkafal.conf:
   input {
  kafka {
    bootstrap_servers => "172.29.112.1:9200"
    topics => ["log_err_netdisk"]
    auto_offset_reset => "latest"
    add_field => { "source" => "err" }
  }
}

input {
  kafka {
    bootstrap_servers => "172.29.112.1:9200"
    topics => ["log_operation_netdisk"]
    auto_offset_reset => "latest"
    add_field => { "source" => "operation" }
  }
}

output {
  if [source] == "err" {
    elasticsearch {
      hosts => ["http://172.29.112.1:9200"]
      index => "log_err_netdisk-%{+YYYY.MM.dd}"
    }
  } else if [source] == "operation" {
    elasticsearch {
      hosts => ["http://172.29.112.1:9200"]
      index => "log_operation_netdisk-%{+YYYY.MM.dd}"
    }
  }
}


# zookeeper
docker run --name zookeeper -p 2181:2181 -d zookeeper

# kafka advertise_listeners 的 ip 要换
docker run -d --name kafka --link zookeeper:zookeeper -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://172.29.112.1:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka

# es
docker run --name elasticsearch -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-d docker.elastic.co/elasticsearch/elasticsearch:7.9.3

# kibana
docker run --name kibana -p 5601:5601 \
--link elasticsearch:elasticsearch \
-e "ELASTICSEARCH_URL=http://elasticsearch:9200" \
-d docker.elastic.co/kibana/kibana:7.9.3

# logstash 容器卷路径要换
docker run -d --name logstash -p 9600:9600 -v "D:/docker/logstash/logstash.yml:/usr/share/logstash/config/logstash.yml" -v "D:/docker/logstash/logstashkafka.conf:/usr/share/logstash/pipeline/logstash.conf" docker.elastic.co/logstash/logstash:7.9.3

```
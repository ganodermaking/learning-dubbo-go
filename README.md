# learning-dubbo-go
dubbo go example

## go调用java

### 启动zookeeper

```shell
docker run --privileged=true -d --name zookeeper --publish 2181:2181 -d zookeeper:latest
```

### 启动java

```shell
mvn clean install package -Dmaven.test.skip=true
java -jar learning-api-server/target/learning-api-server-0.0.1-SNAPSHOT.jar
```

### 启动go

```shell
export DUBBO_GO_CONFIG_PATH=./config/dubbogo.yaml
make http
```

### go调用java

```shell
curl http://127.0.0.1:5000/api/v1/hello
```

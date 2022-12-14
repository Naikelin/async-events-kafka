# Docker
We will use Docker Kafka packaged by Bitnami
https://hub.docker.com/r/bitnami/kafka


- Create a Network
    ```
    docker network create kafka-test --driver bridge
    ```

- Launch Zookeeper
    ```
    docker run -d --name zookeeper-server \
    --network kafka-test \
    -e ALLOW_ANONYMOUS_LOGIN=yes \
    bitnami/zookeeper:latest
    ```
- Launch Kafka
    ```
    docker run -d --name kafka-server \
        --network kafka-test \
        -e ALLOW_PLAINTEXT_LISTENER=yes \
        -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
        bitnami/kafka:latest
    ```

# Compose

## Single Broker

```Docker
version: "3"
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    ports:
      - '2181:2181'
    environment:
      - ALLOW_ANONYMOUS_LOGIN=yes
  kafka:
    image: 'bitnami/kafka:latest'
    ports:
      - '9092:9092'
    environment:
      - KAFKA_BROKER_ID=1
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
```

## Multiple Broker

```Docker
version: "3"
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    ports:
      - '2181:2181'
    environment:
      - ALLOW_ANONYMOUS_LOGIN=yes
  kafka1:
    image: 'bitnami/kafka:latest'
    ports:
      - '9092:9092'
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
  kafka2:
    image: 'bitnami/kafka:latest'
    ports:
      - '9093:9092'
    environment:
      - KAFKA_BROKER_ID=2
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
  kafka3:
    image: 'bitnami/kafka:latest'
    ports:
      - '9094:9092'
    environment:
      - KAFKA_BROKER_ID=3
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
```

## Execute commands

```
docker exec -it <container-id> bash
```

or

```
docker-compose exec kafka1 bash
```

Inside of the container, need to search */opt/bitnami/kafka/bin* path. Contains scripts.

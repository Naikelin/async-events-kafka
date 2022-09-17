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
      - '9092:9092'
    environment:
      - KAFKA_BROKER_ID=2
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
  kafka3:
    image: 'bitnami/kafka:latest'
    ports:
      - '9092:9092'
    environment:
      - KAFKA_BROKER_ID=3
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
```

### Topic replicated

```
kafka-topics.sh --create --bootstrap-server localhost:9092 --topic mytopic --partitions 3 --replication-factor 3
```

### Cheatsheet

#### Topic

- ID -> id of the topic
- Number of *Partitions* -> Fragments that divide the topic, across the brokers. If we have 3 brokers, could only use 3 partitions.
- Replication *Factor* -> Replications of *Partitions* across the brokers. (*Note: Use 3 in production*)


- First item
- Second item
- Third item
- Fourth item 
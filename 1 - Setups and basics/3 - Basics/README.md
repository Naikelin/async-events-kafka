# Basics

## Topic

- ID -> id of the topic
- Number of *Partitions* -> Fragments that divide the topic, across the brokers. (unit of parallelism)
- Replication *Factor* -> Replications of *Partitions* across the brokers. (*Note: Use min. 3 in production*)

### Create Topic replicated

```
kafka-topics.sh --create --bootstrap-server localhost:9092 --topic mytopic --partitions 3 --replication-factor 3
```

### List Topics

```
kafka-topics.sh --list --bootstrap-server localhost:9092
```

### Describe Topic

```
kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic mytopic
```

## Consume - Produce


### Create a consumer (C1)

```
kafka-console-consumer.sh --bootstrap-server :9092 --topic mytopic --consumer-property group.id=CG1
```

### Create a producer attached to partition 0 (**-p 0**)

```
docker run  --interactive \
            --network 2-kafkadocker_default \
            confluentinc/cp-kafkacat \
            kafkacat -P -b kafka1:9092 \
                    -t mytopic \
                    -p 0 -K:
```

## Consumer groups

https://dev.to/de_maric/what-is-a-consumer-group-in-kafka-49il
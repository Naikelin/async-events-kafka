# Basics

## Create Topic replicated

```
kafka-topics.sh --create --bootstrap-server localhost:9092 --topic mytopic --partitions 3 --replication-factor 3
```

## List Topics

```
kafka-topics.sh --list --bootstrap-server localhost:9092
```

## Describe Topic

```
kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic mytopic
```

# Cheatsheet

### Topic

- ID -> id of the topic
- Number of *Partitions* -> Fragments that divide the topic, across the brokers. (unit of parallelism )
- Replication *Factor* -> Replications of *Partitions* across the brokers. (*Note: Use min. 3 in production*)
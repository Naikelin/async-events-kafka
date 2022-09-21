#!/bin/sh
docker-compose kafka exec /opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server :9092 --topic register-topic
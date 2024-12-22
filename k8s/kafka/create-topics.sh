#!/bin/bash

# Wait for Kafka to start
sleep 60

# Create topics
kafka-topics.sh --create --topic user_performance --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 --if-not-exists
kafka-topics.sh --create --topic challenge_data --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 --if-not-exists
kafka-topics.sh --create --topic code_execution_stats --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 --if-not-exists


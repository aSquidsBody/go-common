# Kafka

## CMAK
Kafka dashboard

### Deploy
helm upgrade -i --wait --create-namespace -n cmak cmak-operator oci://ghcr.io/eshepelyuk/helm/cmak-operator


## Kafka CLI
List topics
`kafka-topics.sh --list --bootstrap-server localhost:9092`

View messages
`kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test-topic --from-beginning`

Create topic
`bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor <number_of_replicas> --partitions <number_of_partitions> --topic <name_of_your_topic>`

List topic information
`kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group test-group --describe`
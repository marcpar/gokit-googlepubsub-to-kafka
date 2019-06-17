.PHONY: run
run:
	telepresence --deployment test --docker-run --rm -it -p 8080:8080 -v /tmp/key:/tmp/key --env-file .env asia.gcr.io/solaire-resort-ops/real-time-event-processing/pubsub-router:latest

.PHONY: build
build:
	docker build . -t asia.gcr.io/solaire-resort-ops/real-time-event-processing/pubsub-router:latest

.PHONY: push
push:
	docker push asia.gcr.io/solaire-resort-ops/real-time-event-processing/pubsub-router:latest

.PHONY: consumer
consumer:	
	kubectl exec -it -n kafka testclient -- bin/kafka-console-consumer.sh --bootstrap-server staging-kafka:9092 --topic econnect-data-stream

.PHONY: producer
producer:	
	kubectl exec -it -n kafka testclient -- bin/kafka-console-producer.sh --broker-list staging-kafka:9092 --topic econnect-data-stream

.PHONY: publish
publish:
	gcloud pubsub topics publish staging-kafka --message "sir rommel3 fdssssa" --attribute topicName=econnect-data-stream

.PHONY: kubetail
kubetail:
	kubetail pubsub-router
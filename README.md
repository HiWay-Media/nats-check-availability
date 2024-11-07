# NATS Check Availability

This repository provides a utility to check the availability of NATS (Neural Autonomic Transport System) servers. It aims to ensure that your NATS workload is running smoothly by allowing you to query and verify the status of your NATS servers.

### Environment Variables

- `NATS_SERVERS`: Comma-separated list of NATS server URLs.
- `NATS_JETSTREAM_SERVERS`: Comma-separated list of NATS JetStream server URLs.
- `NATS_REPLICA_SIZE`: Size of the replica for your NATS streams.
- `NATS_DEFAULT_CLUSTER`: Default cluster name to be used.
- `NATS_DEFAULT_STREAM`: Default stream name for publishing or subscribing.

env example: 

```env
LOG_LEVEL=debug
APP_ENV=LOCAL
NATS_SERVERS=nats://localhost:4222
NATS_JETSTREAM_SERVERS=
NATS_REPLICA_SIZE=3
NATS_DEFAULT_CLUSTER=default
NATS_DEFAULT_STREAM=test

```
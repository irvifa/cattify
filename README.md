# Cattify

There are several glossary that being created in this repo, there are [daemonset](http://kubernetes.io/docs/user-guide/deployments/), [service](http://kubernetes.io/docs/user-guide/services/), and [deployments](http://kubernetes.io/docs/user-guide/deployments/). For further understanding, please refer to the documentation provided by kubernetes.

## Filebeat
Filebeat being provided in each of available node in the cluster, thus using the daemonset.

## Logstash, Kibana, and Elasticsearch
Using deployment to make the management of the pods easier, this deployment stuff is an enchancement from [replica set](http://kubernetes.io/docs/user-guide/replicasets/). Service from each of end point is being presented by [kubernetes services](http://kubernetes.io/docs/user-guide/services/). For service that needed to being exposed to the public, change the type of those service into LoadBalancer.

## Milestone

1. Create a simulation service using the ELK Stack
2. Create a simulation for filtering in the Logstash configuration
3. Scale the systsem using the redis as a buffer

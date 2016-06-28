#Sample

This repository is provided us with several of `yaml` from the logging mechanism that being performed by addons from the kubernetes. Apparently, each of the nodes or previously named minions will have its own fluentd service and send the log data as a stream into the elastic search. Just look at the `yaml` files that being provided here to see, the config example of each of the node. There're actually several image that already being provided by google itself. 

The logging meschanism in the fluentd is a stream based logging mechanism, the behaviour of this could be seen from the sample that being provided in the kubernetes page, the counter.yaml, since it has been stated clearly there, there's not much that I want to write. 

## Customize the Config

Please read the conffiguration file available in the:

```bash
kubernetes/cluster/gce and kubernetes/cluster/ubuntu
```

change the default config available there for the KUBE_LOGGING_SYSTEM into

```bash
KUBE_LOGGING_DESTINATION=elasticsearch
```

This is another reference in the [Logging with Elasticsearch and Kibana] (http://kubernetes.io/docs/getting-started-guides/logging-elasticsearch/) section. For installation in the local please change the default config `KUBE_ENABLE_NODE_LOGGING` into:

```bash
KUBE_ENABLE_NODE_LOGGING= true
```

The configuration file is depended on which environment that we use, for exampele the conf file that being provided there is another example that i took from the directory `kubernetes/cluster/gce/`. Go for the directory `conf` to see the example. I've change the conf file so that we could use kibana instead of default google dashboard.


## Cluster Management With Kubernates

Please kindly this interesting slide [Cluster management with kubernetes] (http://www.slideshare.net/SatnamSingh67/2015-0605-cluster-management-with-kubernetes). It should be given us several example for the architecture of logging system using the elastic. However this slide also told us about the possibility of using the grafana for the monitoring proccess. 

## Miscellanous

Use this command to see the configuration that being used in the current cluster:

```bash
kubectl get <what-we-want-to-know> <name-space-if-applicable> <name> -o yaml
```

[!Sample](http://3.bp.blogspot.com/-8sYxFgrVZ3A/VGvL9PeTA4I/AAAAAAAAN00/Uua-54Nu6pQ/s1600/overview.png)

Reading material:

1. [Collecting Output of Containers in Kubernetes Pods] (http://blog.raintown.org/2014/11/logging-kubernetes-pods-using-fluentd.html)
2. [Fluentd runs on every node, we should find something that is more efficient] (https://github.com/kubernetes/kubernetes/issues/23782)
3. [Logging in Kubernetes with Fluentd and Elasticsearch] (http://www.dasblinkenlichten.com/logging-in-kubernetes-with-fluentd-and-elasticsearch/)
4. [Logging] (http://kubernetes.io/docs/getting-started-guides/logging/)
5. [Cluster Level Logging with Kubernetes] (http://blog.kubernetes.io/2015/06/cluster-level-logging-with-kubernetes.html)

NB:
The current issue is how to edit the current service and add the desired container to make the system works perhaps it was the best thing to us to understand the underlying principle of how the kubernetes when it was creating a new cluster.

I've found a way to solve it, just now, it's just like how exactly the `replication-controller` works. It actually for each of the `rc` the number of pods that being specified is just the number that being specifically mentioned in the `rc`. As we could see that the number when we edit the `rc` is actually the supposed number of replica it must have. So, we need to make the `rc` replace the current replica with the newest one. For that purpose, I simply delete all of the pods available in the kube-system namespace with the command of `kubectl delete --all --namespace=kube-system`.However, I deemed this mechanism is not really useful in the future, since we just create an incremental changes and patches for each of the replica set if we want to create a new service. 

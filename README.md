#Sample

This gist is provided us with several of json from the logging mechanism that being performed by addons from the kubernetes. Apparently, each of the nodes or previously named minions will have its own fluentd service and send the log data as a stream into the elastic search. Just look at the `json` files that being provided here to see, the config example of each of the node. There're actually several image that already being provided by google itself. 

The logging meschanism in the fluentd is a stream based logging mechanism, the behaviour of this could be seen from the sample that being provided in the kubernetes page, the counter.yaml, since it has been stated clearly there, there's not much that I want to write. I'm still trying to figure out could the system can be done for a new image, but I think I'll do it tomorrow. For now I'm just observing the pattern of how the logging mechanism that being performed in the addons itself.

Reading material:

1. [Collecting Output of Containers in Kubernetes Pods] (http://blog.raintown.org/2014/11/logging-kubernetes-pods-using-fluentd.html)
2. [Fluentd runs on every node, we should find something that is more efficient] (https://github.com/kubernetes/kubernetes/issues/23782)
3. [Logging in Kubernetes with Fluentd and Elasticsearch] (http://www.dasblinkenlichten.com/logging-in-kubernetes-with-fluentd-and-elasticsearch/)
4. [Logging] (http://kubernetes.io/docs/getting-started-guides/logging/)
5. [Cluster Level Logging with Kubernetes] (http://blog.kubernetes.io/2015/06/cluster-level-logging-with-kubernetes.html)


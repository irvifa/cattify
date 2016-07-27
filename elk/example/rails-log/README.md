#Test

```
export LOGSTASH_URI=tcp://localhost:11001  
```

run the application using `rails s` and then run this command:

```
docker run -it -p 11001:11001 gcr.io/kube-simulation/lsth-rb:0.7    
```

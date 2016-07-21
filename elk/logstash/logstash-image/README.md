# Configuration file

## Important notice

We must create a filter option to match available filter as fast as we could because we're battling with the IO option. In the logstash configuration define the tag that will be given to a log if the parse is become a failure. The defaulf option for this kind of log is _grokparsefailure. This tag could be dropped if we want to drop it. However to make things become easier just give a notice if the configuration file isn't matched with the available format. 

For example:

```
	I, [2016-01-26T23:21:44.581108 #27447]  INFO -- : Started GET "/login" for XXX.XXX.XXX.XXX at 2016-01-26 23:21:44 -0800
```

Could be matched with:

```bash
	I, \[%{TIMESTAMP_ISO8601:timestamp}\ #%{NUMBER:id}] %{LOGLEVEL:level} -- : Started %{WORD:method} (?<request>[^ ]+) for.*%{IP:ip} at.*%{DATE:date} %{TIME:time} %{NUMBER:gmt}
```



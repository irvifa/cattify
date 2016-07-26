#Configuration file

## Important notice

We must create a filter option to match available filter as fast as we could because we're battling with the IO option. In the logstash configuration define the tag that will be given to a log if the parse is become a failure. The defaulf option for this kind of log is _grokparsefailure. This tag could be dropped if we want to drop it. However to make things become easier just give a notice if the configuration file isn't matched with the available format. 

For example:

```
	I, [2016-01-26T23:21:44.581108 #27447]  INFO -- : Started GET "/login" for XXX.XXX.XXX.XXX at 2016-01-26 23:21:44 -0800
```

Could be matched with:

```bash
	I, \[%{TIMESTAMP_ISO8601:timestamp} #%{NUMBER:id}\] %{LOGLEVEL:level} -- : Started %{WORD:method} (?<request>[^ ]+) for.*%{IP:ip} at.*%{DATE:date} %{TIME:time} %{NUMBER:gmt}
```

## Multiline 

There's a problem since the data is not always in a single line, wo whe must think of how we could parse a multiline data. Here I assume all of the multiline data already being aggregared first and such become a one line.

```bash
	"I, \[%{TIMESTAMP_ISO8601:timestamp} #%{NUMBER:id}\]  %{LOGLEVEL:l
evel} -- : Started %{WORD:method} (?<request>[^ ]+) for.*%{IP:ip} at.*%{DATE:date} %{TIME:time}
 %{NUMBER:gmt} %{GREEDYDATA:message}"
```

In the input level we could parse this multiline input using a certain pattern. This pattern is using regex. If the pattern started with space then it belongs to the previous line. This principle is just the same with the ActionController line in the log output. After using the multiline we could use `mutate` and `gsub` to delete all of new lines in the multiline input before parsing it.

```
I, [2016-01-26T23:21:44.581108 #27447]  ERROR -- : Started GET "/login" for 192.168.0.1 at 2016-01-26 23:21:44 -0800 
 ActionController::RoutingError (No route matches [GET] "/palaver/info"):  
  actionpack (4.2.5.2) lib/action_dispatch/middleware/debug_exceptions.rb:21:in `call'
  newrelic_rpm (3.14.1.311) lib/new_relic/agent/instrumentation/middleware_tracing.rb:79:in `call'
  web-console (3.1.1) lib/web_console/middleware.rb:131:in `call_app'
  web-console (3.1.1) lib/web_console/middleware.rb:28:in `block in call'
  web-console (3.1.1) lib/web_console/middleware.rb:18:in `catch'
  web-console (3.1.1) lib/web_console/middleware.rb:18:in `call'
  newrelic_rpm (3.14.1.311) lib/new_relic/agent/instrumentation/middleware_tracing.rb:79:in `call'
  actionpack (4.2.5.2) lib/action_dispatch/middleware/show_exceptions.rb:30:in `call'
```


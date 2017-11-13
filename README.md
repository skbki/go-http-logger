# go-http-logger
Takes a generic JSON request and logs it to stdout. 

This is a super lighweight application to forward the bodies of web requests to stdout. Care should be taken if using this as-is as it does not discern between request sources, resources, formats or apply any kind of authorization protocol. Usage inside a K8s pod would be fairly safe though, as long as it is not exposed. 

## Use Case
In the vast majority of cases you'll code some logging functionality into your application to pipe logs to stdout or some other logging service using some standard format. In some rare cases you may want to use stdout to collect logs in a specific format like json, but the some factor makes this difficult. The use case I have found us Kong, which logs using the nginx format by default but can pipe http logs as json. Changing to json format is more difficult than a configuration change, so this is another option to writing a lua plugin. 

Me: This open source project is missing a feature, I should learn how to fix it and submit a PR<br/><br/>
Also me: Make another microservice
![](https://fsmedia.imgix.net/ff/19/fc/02/501c/4a8c/9fb5/0a93d3c05cb7.jpeg)
## Usage

Start the server, listing on port 8090. Using Kubernetes you can have this within the same pod as your application. Send any request with a body to this application on port 8090, and it will log the body to stdout. 

### Example

```
curl -v -d '{"log_level": "info", "error_text":"You done messed up A'ron"}' localhost:8090/whatever
```

This request would lead to the following text being printed out in the http logger application

```
{"log_level": "info", "error_text":"You done messed up A'ron"}
```

Inside the application whose logs you want to gather, simply send all json logs to the http logger via any URI other than `/health`, and whatever body is sent will be printed out as a separate line.

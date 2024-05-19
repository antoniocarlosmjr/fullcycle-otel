# fullcycle-otel

# How to run

* Start the containers using the docker command: `docker-compose up -d`

* Execute the server of `service_a` and `service-b` using the command in different terminals:
* `cd service_a && go run main.go`
* `cd service_b && go run main.go`

* After that, we can acces the `service_a` in the port 8080 and the `service_b` in the port 8081 and we can test
using the file `get_weather.http` in the folder `requests` where we can see examples of requests to the services.

* To see the tracing, we can access the url `http://localhost:9411/` in the browser.
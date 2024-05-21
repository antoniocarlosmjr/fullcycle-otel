# fullcycle-otel

# How to run

1. Start the containers using the docker command: `docker-compose up -d --build`

2. After that, we are listening the services `service_a` in the port 8080 and `service_b` in the port 8081.

3. After that, we can acces the `service_a` in the port 8080 and the `service_b` in the port 8081 and we can test
using the file `weather.http` in the folder `api` where we can see examples of requests to the services.

4. To see the tracing, we can access the url `http://localhost:9411/` in the browser.
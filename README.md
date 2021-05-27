# go-hello-port-range
A simple golang web server which listens on a range of ports. Useful for testing whether you've configured a load balancer range successfully.

To Build: `go build hello-port-range.go`

To Run: `./hello-port-range -start 8000 -end 8200`

You can then curl any of the bound ports to see that it's running: `curl localhost:8123`

**To Run in Docker**

Build the container: `docker build -t hello-port-range ./`

Run the container, mapping a small range of ports from your local machine so docker doesn't light your machine on fire:

`docker run -p 8000-8010:8000-8010 hello-port-range`

Then curl the docker container: `curl localhost:8004`

**Running the Tests**

Run the test locally with:

`go test ./hello-port-range.go hello-port-range_test.go -v`

Run the test in Docker with:

`docker run hello-port-range go test ./hello-port-range.go hello-port-range_test.go -v`

For a failing test, run hello-port-range_failing_test.go instead:

`go test ./hello-port-range.go hello-port-range_failing_test.go -v`

docker build -t wizz .
docker run --name WizzHomePage -d -p 8080:8080 wizz
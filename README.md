# Microservices Course

## Run services locally

`docker-compose up --build` - build each of services’ image before instantiating containers

`docker-compose up` - start containers from the image that was previously built

`docker-compose ps` - list running compose containers

`docker-compose stop` - shuts down the application (leave containers in the *****stopped***** state)

`docker-compose down` - remove running compose containers and return development workstation to a clean state (do this every time C-c is pressed, if not running `up -d` otherwise)

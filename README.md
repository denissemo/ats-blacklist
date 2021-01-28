# Small web server for adding new phone number to asterisk blacklist table

## Running with docker

### Build image

- `docker build -t <image-name> .`

### Run container

- `docker run -d -p 3000:3000 --name=<container-name> <image-name>`

### Show logs

- `docker logs <container-name>`

### Stop container

- `docker stop <container-name>`

### Connect to container for some changes

- `docker exec -t -i <container-name> /bin/bash`
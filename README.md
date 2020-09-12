# Devops Pre-interview Test

Containerize a Go endpoint and add CI with GitLab

# How to?

Build the containers:
```
docker build . -t endpoint
```

Run the image
```
docker run -p 8080:8080 endpoint:latest
```
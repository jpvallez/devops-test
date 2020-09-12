# Devops Pre-interview Test

Here we're containerizing a Go endpoint. Unit testing is run in the build container. 

Output (json):
* LastCommitSha: is the latest commit requested from Github API.
* Version: is the latest git tag.
* Description: hard coded.


# How to?

Build the containers:
```
docker build . -t endpoint
```

Run the image
```
docker run -p 8080:8080 endpoint:latest
```
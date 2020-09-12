# Devops Pre-interview Test

Here we're containerizing a Go endpoint. Unit testing is run in the build container. 

I'm using Make to build, test and add version information to the application from github.

Output (json):
* LastCommitSha: is the latest commit requested from Github API.
* Version: is the latest Git tag.
* Description: hard coded string.

### Example output:
```
{
    "LastCommitSha":    "102013709ec1dac0fa39d43f079d438ce54d7126",
    "Version":  "v0.1-0-g1020137",
    "Description":  "This is a pre-interview technical test"
}
```

# How to run it?

1. Build the containers:
```
docker build . -t endpoint
```

2. Run the image
```
docker run -p 8080:8080 endpoint:latest
```

3. Visit http://localhost:8080/version in your browser.
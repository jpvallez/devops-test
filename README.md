
## \* This is a copy of my Gitlab repo/pipeline at https://gitlab.com/jpvallez/devops-test
<br/>

___ 
___


# CI Pipeline Project (Devops Pre-interview Test)

This is a containerized Go app which presents a simple endpoint for version information about itself.

It is built, containerized and tested in a GitLab CI Pipeline.

Output (json):
* LastCommitSha: is the latest commit requested from Gitlab API. Not tied to version.
* Version: is the tag or last commit before build. 
     ```([tagename]-[number of commits after/away from tag]-[shortened last commit sha])```
* Description: hard coded string.

### Example output:
```
{
    "LastCommitSha":    "102013709ec1dac0fa39d43f079d438ce54d7126",
    "Version":  "v0.1-0-g1020137",
    "Description":  "This is a pre-interview technical test"
}
```

## Layers of the build

### Layer 1: Build Go Application
* The Go application can be built and run in local environment without containers. This will present the endpoint at localhost:8080/version.

### Layer 2: The Makefile. For Versioning and Unit Tests
* Unit Tests are run using Go's testing package. See: [main_test.go](../master/main_test.go)
* Go application is built (resulting binary) with version information (**latest tag**/fallback to abbreviated commit)
* command: make
* See: [Makefile](../master/Makefile)

### Layer 3: Containerize - Multi-stage build in Docker
* Build container runs the Makefile (Layer 2 inc Unit Tests) and includes all dependencies needed to build.
* The second stage is a lightweight alpine:latest, copies binary from build container and results in small purpose built deployable artefact.
* See: [Dockerfile](../master/Dockerfile)

## Gitlab CI Pipeline
* Each Git Push will kick off the GitLab CI PipeLine

### Stages:
* Using Gitlab Runners.
* **Build stage**: 
    * Runs all the layers of the build (build, unit test, containerize) as described above in docker:19.03.0 image using dind service.
    * SAST test run with Gitlab CI SAST. Analyses Go source code for known vulnerabilities and any coding quality issues such as unhandled exceptions.
    * Pushes built and unit tested container to Gitlab container registry for test stage to use.
* **Test stage**: 
    * Using an alpine image with bash, curl and jq.
    * Image which was pushed in **Build stage** runs as a _service_.
    * Run functional API tests (bash script) against endpoint container service. See: [container-test.sh](../master/container-test.sh)
* **Release Image stage**:
    * Creates a latest image from the successful test and pushes to Gitlab container registry.
* See: [.gitlab-ci.yml](../master/.gitlab-ci.yml)


# How to run the container?

### Option 1: Get it from the Gitlab container registry
1. Pull and Run latest image from the Gitlab container registry.

```
docker pull registry.gitlab.com/jpvallez/devops-test:latest
```
```
docker run registry.gitlab.com/jpvallez/devops-test:latest
```
2. Visit http://[container-address]:8080/version in your browser.

Or;

### Option 2: Build and run it locally
> :warning: Keep in mind this will skip the CI build, security and test process.

1. Clone this repository and cd to its location.

2. Build the container:
```
docker build . -t endpoint
```

3. Run the image
```
docker run -p 8080:8080 endpoint:latest
```

4. Visit http://localhost:8080/version in your browser.
<br/><br/><br/>

# Deployment

This container is deployed to a Kubernetes cluster hosted on GCE. The GKE cluster uses the most recent tagged image on my Gitlab container registry (registry.gitlab.com/jpvallez/)

To kick off an update of the app on GKE I run:

```
kubectl set image deployment/devops-test devops-test=registry.gitlab.com/jpvallez/devops-test:TAG-NAME
```

This kicks off a rolling update. During the rolling update the GKE cluster will ensure uptime by running a new pod with the updated image, and then swapping the load balancer to point at the new pod, and finally removing the old pod. If you have more than one pod it will incrementally replace them. 


## Release Process
My release process is as follows:
1. Tag new version. (e.g. v1.04)
2. Wait for CI Pipeline to complete (new tagged container pushed to Gitlab container repo)
3. Run above deployment command with new tage name on Google Cloud Console.



E.g.
```
kubectl set image deployment/devops-test devops-test=registry.gitlab.com/jpvallez/devops-test:v1-04
```
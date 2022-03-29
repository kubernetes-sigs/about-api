# about-api

A CRD for arbitrary properties about a cluster.
## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack](https://kubernetes.slack.com/messages/sig-multicluster)
- [Mailing List](https://groups.google.com/forum/#!forum/kubernetes-sig-multicluster)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).


### Run controller on the Cluster.

In order to build and run tests you would need access to an existing miniKube or kubernetes cluster.
Set the environment variable `TEST_USE_EXISTING_CLUSTER` to `true`

export TEST_USE_EXISTING_CLUSTER=true

Build and push your image to the location specified by IMG:

make docker-build docker-push IMG=<some-registry>/<project-name>:tag

Deploy the controller to the cluster with image specified by IMG:

make deploy IMG=<some-registry>/<project-name>:tag


### Run tests.

In order to run tests you would need access to an existing miniKube or kubernetes cluster.
Set the environment variable `TEST_USE_EXISTING_CLUSTER` to `true`

export TEST_USE_EXISTING_CLUSTER=true
go test ./...

# about-api

A CRD for arbitrary properties about a cluster.
## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack](https://kubernetes.slack.com/messages/sig-multicluster)
- [Mailing List](https://groups.google.com/forum/#!forum/kubernetes-sig-multicluster)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).


### How to run tests 

In order to run tests you would need access to either a miniKube or kubernetes cluster.
Set the environment variable `TEST_USE_EXISTING_CLUSTER` to `true`

export TEST_USE_EXISTING_CLUSTER=true
go test ./...

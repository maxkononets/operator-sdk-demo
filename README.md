# K8s demo controller 

### Prerequisites

- `golang 1.17`
- `docker`
- `minikube`
- `kubectl`
- `operators-sdk`

## Steps to deploy

Login to docker

    $> docker login [OPTIONS] [SERVER]

Build and push an operator image

    $> make docker-build docker-push

Deploy the operator

    $> make deploy

Apply the CR to default cluster

    $> kubectl apply -f config/samples/groupcontroller_v1alpha1_democontroller.yaml

### Main functionality

#### 1. Replicas threshold validation
Check the pods status. Take a look that here we have two pods of **democontroller-sample-***

    $> kubectl get pods
    NAME                                     READY   STATUS    RESTARTS   AGE
    democontroller-sample-6fb5d6db85-9rvkb   1/1     Running   0          55s
    democontroller-sample-6fb5d6db85-sjp66   1/1     Running   0          55s

Then try to scale deployment up to 3 replicas.

    $> kubectl scale --replicas=3 deployments democontroller-sample

See that DemoController do not allow creating more replicas than
specified at **groupcontroller_v1alpha1_democontroller.yaml**

    $> kubectl get pods
    NAME                                     READY   STATUS    RESTARTS   AGE
    democontroller-sample-6fb5d6db85-9rvkb   1/1     Running   0          90s
    democontroller-sample-6fb5d6db85-sjp66   1/1     Running   0          90s

#### 2. K8s event AttemptToExceedReplicasThreshold creation.
After scaling up extra replicas take a look at result of **get event** command and here we can see event says that there
was attempt to exceed **max replicas count** 

    $> kubectl get events
    LAST SEEN   TYPE      REASON                             OBJECT                                        MESSAGE
    68s         Normal    AttemptToExceedReplicasThreshold   deployment/democontroller-sample              Trying to set replicas to count to 3, but a maximum 2 allowed
    68s         Normal    ScalingReplicaSet                  deployment/democontroller-sample              Scaled up replica set democontroller-sample-6fb5d6db85 to 3
    67s         Normal    ScalingReplicaSet                  deployment/democontroller-sample              Scaled down replica set democontroller-sample-6fb5d6db85 to 2

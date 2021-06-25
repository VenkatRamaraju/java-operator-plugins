# Quickstart for Java Operators
### A simple set of instructions to set up and run a Java operator.

This guide walks through an example of building a simple memcached-operator using tools and libraries provided by the Operator SDK.

## Prerequisites

- [Operator SDK](https://sdk.operatorframework.io/docs/installation/) v1.8.0 or newer
- [Java](https://java.com/en/download/help/download_options.html) 11
- [Maven 3.6.3](https://maven.apache.org/install.html) or newer
- User authorized with `cluster-admin` permissions.
- [GNU Make](https://www.gnu.org/software/make/)

## Steps
1. Create a project directory for your project and initialize the project:

```sh
mkdir memcached-quarkus-operator
cd memcached-quarkus-operator
# we'll use a domain of example.com
# so all API groups will be <group>.example.com
operator-sdk init --plugins quarkus --domain example.com --project-name memcached-quarkus-operator
```

2. Create a simple Memcached API:
```console
$ operator-sdk create api --plugins quarkus --group cache --version v1 --kind Memcached
```


3. Manually create Custom Resource Definition

crd.yaml
```
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  name: memcacheds.cache.example.com
spec:
  group: cache.example.com
  names:
    kind: Memcached
    listKind: MemcachedList
    plural: memcacheds
    singular: memcached
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            type: string
          kind:
            type: string
          metadata:
            type: object
          spec:
            properties:
              size:
                format: int32
                type: integer
            type: object
          status:
            properties:
              nodes:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
```

4. Manually create the custom resource

memcached-sample.yaml
```
apiVersion: cache.example.com/v1
kind: Memcached
metadata:
  name: memcached-sample
spec:
  # Add fields here
  size: 1
```


## Running the operator 

The following steps will show how to run your operator in the cluster.

1. Build and push your operator's image:

```
make docker-build docker-push IMG=quay.io/YOURUSER/memcached-quarkus-operator:0.0.1
```


2. Install the CRD
```
kubectl apply -f crd.yaml
customresourcedefinition.apiextensions.k8s.io/memcacheds.cache.example.com created
```


3. Deploy the operator

```
make deploy
```


4. Apply the custom resource

```
$ kubectl apply -f memcached-sample.yaml
memcached.cache.example.com/memcached-sample created
```


5. Create and apply RBAC specifications

rbac.yaml
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: memcached-operator-admin
subjects:
- kind: ServiceAccount
  name: memcached-quarkus-operator-operator
  namespace: default
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: ""
```

```
kubectl apply -f rbac.yaml
```


6. Uninstall the operator

```
make undeploy
```

## Next Steps
Read the [full tutorial](https://github.com/VenkatRamaraju/java-operator-plugins/blob/Quickstart/docs/tutorial.md) for an in-depth walkthrough of building a Java Operator.


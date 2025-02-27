# Helm Support

## Runtime Configuration Guide for Dragonfly Helm Chart

When enable runtime configuration in dragonfly, you can skip [Configure Runtime](#configure-runtime-manually) manually.

### 1. Docker

> **We did not recommend to using dragonfly with docker in Kubernetes** due to many reasons: 1. no fallback image pulling policy. 2. deprecated in Kubernetes.
> Because the original `daemonset` in Kubernetes did not support `Surging Rolling Update` policy.
> When kill current dfdaemon pod, the new pod image can not be pulled anymore.

> If you can not change runtime from docker to others, remind to choose a plan when upgrade dfdaemon:
>     Option 1: pull newly dfdaemon image manually before upgrade dragonfly, or use [ImagePullJob](https://openkruise.io/docs/user-manuals/imagepulljob) to pull image automate.
>     Option 2: keep the image registry of dragonfly is different from common registries and add host in `containerRuntime.docker.skipHosts`.

Dragonfly helm supports config docker automatically.

Config cases:

**Case 1: [Preferred] Implicit registries support without restart docker**

Chart customize values.yaml:
```yaml
containerRuntime:
  docker:
    enable: true
    # -- Inject domains into /etc/hosts to force redirect traffic to dfdaemon.
    # Caution: This feature need dfdaemon to implement SNI Proxy, confirm image tag is greater than v2.0.0.
    # When use certs and inject hosts in docker, no necessary to restart docker daemon.
    injectHosts: true
    registryDomains:
    - "harbor.example.com"
    - "harbor.example.net"
```

This config enables docker pulling images from registries `harbor.example.com` and `harbor.example.net` via Dragonfly.
When deploying Dragonfly with above config, it's unnecessary to restart docker daemon.

Advantages:
* Support upgrade dfdaemon smoothness

> In this mode, when dfdaemon pod deleted, the `preStop` hook will remove all injected hosts info in /etc/hosts,
> all images traffic fallbacks to original registries.

Limitations:
* Only support implicit registries

**Case 2: Arbitrary registries support with restart docker**

Chart customize values.yaml:
```yaml
containerRuntime:
  docker:
    enable: true
    # -- Restart docker daemon to redirect traffic to dfdaemon
    # When containerRuntime.docker.restart=true, containerRuntime.docker.injectHosts and containerRuntime.registry.domains is ignored.
    # If did not want restart docker daemon, keep containerRuntime.docker.restart=false and containerRuntime.docker.injectHosts=true.
    restart: true
    skipHosts:
    - "127.0.0.1"
    - "docker.io" # Dragonfly use this image registry to upgrade itself, so we need skip it. Change it in real environment.
```

This config enables docker pulling images from arbitrary registries via Dragonfly.
When deploying Dragonfly with above config, dfdaemon will restart docker daemon.

Advantages:
* Support arbitrary registries

Limitations:
* Must enable live-restore feature in docker
* Need restart docker daemon
* When upgrade dfdaemon, new image must be pulled beforehand.

### 2. Containerd

The config of containerd has two version with complicated fields. These are many cases to consider:

**Case 1: Version 2 config with config_path**

There is `config_path` in `/etc/containerd/config.toml`:
```toml
[plugins."io.containerd.grpc.v1.cri".registry]
  config_path = "/etc/containerd/certs.d"
```

This case is very simple to enable multiple registry mirrors support.

Chart customize values.yaml:
```yaml
containerRuntime:
  containerd:
    enable: true
```

**Case 2: Version 2 config without config_path**

* Option 1 - Allow charts to inject config_path and restart containerd.

This option also enable multiple registry mirrors support.

> Caution: if there are already many other mirror config in config.toml, should not use this option, or migrate your config with `config_path`.

Chart customize values.yaml:
```yaml
containerRuntime:
  containerd:
    enable: true
    injectConfigPath: true
```

* Option 2 - Just mirror only one registry which `dfdaemon.config.proxy.registryMirror.url` is

Chart customize values.yaml:

```yaml
containerRuntime:
  containerd:
    enable: true
```

**Case 3: Version 1**

With version 1 config.toml, only support the registry which `dfdaemon.config.proxy.registryMirror.url` is.

Chart customize values.yaml:

```yaml
containerRuntime:
  containerd:
    enable: true
```

### 3. [WIP] CRI-O

> DON'T USE, Work in progress

Dragonfly helm supports config CRI-O automatically with drop-in registries.

Chart customize values.yaml:
```yaml
containerRuntime:
  crio:
    # -- Enable CRI-O support
    # Inject drop-in mirror config into /etc/containers/registries.conf.d.
    enable: true
    # Registries full urls
    registries:
    - "https://ghcr.io"
    - "https://quay.io"
    - "https://harbor.example.com:8443"
```

## Prepare Kubernetes Cluster

If there is no available Kubernetes cluster for testing, [minikube](https://minikube.sigs.k8s.io/docs/start/) is
recommended. Just run `minikube start`.

## Install Dragonfly

### Install with default configuration

```shell
helm repo add dragonfly https://dragonflyoss.github.io/helm-charts/
helm install --create-namespace --namespace dragonfly-system dragonfly dragonfly/dragonfly
```

### Install with custom configuration

Create the `values.yaml` configuration file. It is recommended to use external redis and mysql instead of containers.

The example uses external mysql and redis. Refer to the document for [configuration](https://artifacthub.io/packages/helm/dragonfly/dragonfly#values).

```yaml
mysql:
  enable: false

externalMysql:
  migrate: true
  host: mysql-host
  username: dragonfly
  password: dragonfly
  database: manager
  port: 3306

redis:
  enable: false

externalRedis:
  host: redis-host
  password: dragonfly
  port: 6379
```

Install dragonfly with `values.yaml`.

```shell
helm repo add dragonfly https://dragonflyoss.github.io/helm-charts/
helm install --create-namespace --namespace dragonfly-system dragonfly dragonfly/dragonfly -f values.yaml
```

### Install with an existing manager

Create the `values.yaml` configuration file. Need to configure the cluster id associated with scheduler and cdn.

The example is to deploy a cluster using the existing manager and redis. Refer to the document for [configuration](https://artifacthub.io/packages/helm/dragonfly/dragonfly#values).

```yaml
scheduler:
  config:
    manager:
      schedulerClusterID: 1

cdn:
  config:
    base:
      manager:
        cdnClusterID: 1

manager:
  enable: false

externalManager:
  enable: true
  host: "dragonfly-manager.dragonfly-system.svc.cluster.local"
  restPort: 8080
  grpcPort: 65003

redis:
  enable: false

externalRedis:
  host: redis-host
  password: dragonfly
  port: 6379

mysql:
  enable: false
```

## Wait Dragonfly Ready

Wait all pods running

```shell
kubectl -n dragonfly-system wait --for=condition=ready --all --timeout=10m pod
```

## Manager Console

The console page will be displayed on `dragonfly-manager.dragonfly-system.svc.cluster.local:8080`.

If you need to bind Ingress, you can refer to [configuration options](https://artifacthub.io/packages/helm/dragonfly/dragonfly#values) of Helm Charts, or create it manually.

Console features preview reference document [console preview](../../../design/manager.md).

## Configure Runtime Manually

Use Containerd with CRI as example, more runtimes can be found [here](../../../quick-start.md)

> This example is for single registry, multiple registries configuration is [here](../../../container-runtimes)

For private registry:

```toml
# explicitly use v2 config format, if already v2, skip the "version = 2"
version = 2
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."harbor.example.com"]
endpoint = ["http://127.0.0.1:65001", "https://harbor.example.com"]
```

For docker public registry:

```toml
# explicitly use v2 config format, if already v2, skip the "version = 2"
version = 2
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
endpoint = ["http://127.0.0.1:65001", "https://registry-1.docker.io"]
```

Add above config to `/etc/containerd/config.toml` and restart Containerd

```shell
systemctl restart containerd
```

## Using Dragonfly

After all above steps, create a new pod with target registry. Or just pull an image with `crictl`:

```shell
crictl harbor.example.com/library/alpine:latest
```

```shell
crictl pull docker.io/library/alpine:latest
```

After pulled images, find logs in dfdaemon pod:
```shell
# find pods
kubectl -n dragonfly-system get pod -l component=dfdaemon
# find logs
pod_name=dfdaemon-xxxxx
kubectl -n dragonfly-system exec -it ${pod_name} -- grep "peer task done" /var/log/dragonfly/daemon/core.log
```

Example output:
```
{"level":"info","ts":"2021-06-28 06:02:30.924","caller":"peer/peertask_stream_callback.go:77","msg":"stream peer task done, cost: 2838ms","peer":"172.17.0.9-1-ed7a32ae-3f18-4095-9f54-6ccfc248b16e","task":"3c658c488fd0868847fab30976c2a079d8fd63df148fb3b53fd1a418015723d7","component":"streamPeerTask"}
```
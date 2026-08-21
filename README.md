# gd-demo-1-hello — step 1: just deploy it

First of the five graduated demo repositories for
[git-deploy-operator](https://github.com/fsamin/git-deploy-operator). Each one
adds exactly one feature, and each one stands alone.

**What this repository proves:** a `Dockerfile` with `EXPOSE` is the entire
contract. There is no platform manifest here, and `git-deploy init` writes no
file into your repository.

## Demo

```sh
git-deploy -n demo init --name hello
git-deploy status                 # Pending → Building → Releasing → Exposing → Ready
git-deploy show --open            # https://hello.demo.<base-domain>
```

Then make the rolling update visible: bump `version` in `main.go`, commit, push,
and wait one poll interval. The page's version and pod name both change.

## The failure path

The `no-expose` branch carries the same application with the `EXPOSE` line
removed from the Dockerfile. The operator cannot discover a port and fails
**explicitly**, with a message saying what to do:

```sh
git-deploy -n demo init --name hello-noexpose --ref no-expose
git-deploy status --name hello-noexpose      # Failed: no EXPOSE and no port override
```

Recover it without touching the repository:

```sh
git-deploy edit --name hello-noexpose        # set spec.port: 8080
```

## Locally

```sh
go run .          # http://localhost:8080
```

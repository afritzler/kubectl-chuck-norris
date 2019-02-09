# kubectl-chuck-norris

Round-house kicking your Kubernetes cluster with the Chuck Norris `kubectl` plugin.

```bash
kubectl chucknorris get fact
Chuck Norris doesn't need containers, he ships with a round-house kick.
```

Feel free to contribute yours [here](https://github.com/afritzler/kubectl-chuck-norris/blob/master/cmd/get.go) and create a PR!

## Installation

First make sure you have [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl) installed on your machine.

### Linux

```bash
curl -LO https://github.com/afritzler/kubectl-chuck-norris/releases/download/v0.0.1/kubectl-chucknorris_linux_amd64 \
&& sudo mv kubectl-chucknorris_linux_amd64 /usr/local/bin/kubectl-chucknorris \
&& sudo chmod +x /usr/local/bin/kubectl-chucknorris
```

### MacOS

```bash
curl -LO https://github.com/afritzler/kubectl-chuck-norris/releases/download/v0.0.1/kubectl-chucknorris_linux_amd64 \
&& sudo mv kubectl-chucknorris_darwin_amd64 /usr/local/bin/kubectl-chucknorris \
&& sudo chmod +x /usr/local/bin/kubectl-chucknorris
```

## Usage

```bash
kubectl chucknorris get fact
```
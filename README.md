# kubectl-chuck-norris

Roundhouse kicking your Kubernetes cluster with the Chuck Norris `kubectl` plugin.

```bash
kubectl chucknorris get fact
Chuck Norris's computer has no 'backspace' button, Chuck Norris doesn't make mistakes.
```

Feel free to contribute and create a PR!

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
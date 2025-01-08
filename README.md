# Project Documentation

## Overview

This repository contains code in multiple languages (Go, JavaScript, Python) and Terraform configuration files for setting up a Kubernetes cluster using `kind`.

## Directory Structure

```
code/
    hello.go
    hello.js
    hello.py
infra/
    main.tf
    provider.tf
README.md
```

### Code

- **Go**: [code/hello.go](code/hello.go)
  ```go
  package main

  import (
     "net/http"
  )

  func Handler(w http.ResponseWriter, r *http.Request) {
      msg := "Hello World"
      w.Write([]byte(msg))
  }
  ```

- **JavaScript**: [code/hello.js](code/hello.js)
  ```js
  module.exports = async function(context) {
      return {
          status: 200,
          body: "hello, world!\n"
      };
  }
  ```

- **Python**: [code/hello.py](code/hello.py)
  ```py
  def main():
      return "Hello World, Parabéns! Sua função Fission foi executada com sucesso. 🚀"
  ```

### Infrastructure

- **Terraform Provider Configuration**: [infra/provider.tf](infra/provider.tf)
  ```tf
  terraform {
    required_providers {
      kind = {
        source  = "tehcyx/kind"
        version = "0.4.0"
      }
    }
  }

  provider "kind" {}
  ```

- **Terraform Main Configuration**: [infra/main.tf](infra/main.tf)
  ```tf
  resource "kind_cluster" "default" {
    name           = "kubelearn"
    wait_for_ready = true
    node_image     = "kindest/node:v1.27.1"
    kind_config {
      kind        = "Cluster"
      api_version = "kind.x-k8s.io/v1alpha4"

      node {
        role = "control-plane"
      }

      node {
        role = "worker"
      }

      node {
        role = "worker"
      }
    }
  }
  ```

## Usage

### Terraform

To apply the Terraform configuration, use:
```sh
cd infra
terraform init
terraform apply
```
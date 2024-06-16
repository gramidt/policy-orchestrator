# Table of Contents

- [Hexa Policy Orchestrator](#hexa-policy-orchestrator)
    * [Getting Started](#getting-started)
        + [Build the Hexa image](#task-build-the-hexa-orchestrator-image)
        + [Run the Policy Orchestrator](#task-run-the-policy-orchestrator)
    * [Application descriptions](#application-descriptions)
        + [Example workflow](#example-workflow)
    * [Getting involved](#getting-involved)

![hexa-logo](docs/hexa-logo.svg)

# Hexa Policy Orchestrator

[![Build results](https://github.com/hexa-org/policy-orchestrator/workflows/build/badge.svg)](https://github.com/hexa-org/policy-orchestrator/actions)
[![Go Report Card](https://goreportcard.com/badge/hexa-org/policy-orchestrator)](https://goreportcard.com/report/hexa-org/policy-orchestrator)
[![codecov](https://codecov.io/gh/hexa-org/policy-orchestrator/branch/main/graph/badge.svg)](https://codecov.io/gh/hexa-org/policy-orchestrator)
[![CodeQL](https://github.com/hexa-org/policy-orchestrator/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/hexa-org/policy-orchestrator/actions/workflows/codeql-analysis.yml)

Hexa Policy Orchestrator enables you to manage all of your policies consistently across software providers
so that you can unify access policy management. The below diagram describes the current provider architecture.

![Hexa Provider Architecture](docs/hexa-provider-architecture.svg "hexa provider architecture")

## Getting Started

The Hexa project contains two applications, and demonstrates use of applications
from [Policy-Opa](https://github.com/hexa-org/policy-opa)

- Policy Orchestrator with policy translations
- Demo Policy Administrator
- Demo web application (Policy-OPA)
- Hexa OPA Server and Bundle Server (Policy-OPA)

To get started with running these, clone or download the codebase from GitHub to your local machine:

```bash
cd $HOME/workspace # or similar
git clone git@github.com:hexa-org/policy-orchestrator.git
```

### Prerequisites

Install the following dependencies.

- [Go 1.22](https://go.dev) - Needed to compile and install
- [Docker Desktop](https://www.docker.com/products/docker-desktop) - Needed to run docker-compose configuration

### Task: Build the Hexa Orchestrator image

Build the Hexa Orchestrator Docker containers...

```bash
cd demo
sh ./build.sh
```

### Task: Run the Policy Orchestrator

Run all the applications with Docker Compose from within the `demo` directory

> ```bash
> source .env_development
> docker-compose up
> ```

## Application Descriptions

Docker runs the following applications:

- **hexa-orchestrator**

  Runs on [localhost:8885](http://localhost:8885/health). The main application service API
  that manages IDQL policy across various platforms and communicates with the
  various platform interfaces, converting IDQL policy to and from the respective
  platform types using the [Hexa Policy-Mapper](https://github.com/hexa-org/policy-mapper) SDK. This service uses TLS
  and is protected and uses JWT authorization tokens
  (see [RFC7519](https://datatracker.ietf.org/doc/html/rfc7519)) for authenticating access.

- **hexa-admin-ui**

  Runs on [localhost:8884](http://localhost:8884/). An example web application user-interface
  demonstrating the latest interactions with the policy orchestrator. Uses `hexa-orchestrator` to retrieve and provision
  policy.
  <p/>
  The Hexa Admin UI service uses the OAuth2 Client Credentials Grant Flow (see [RFC6749 Section 4.4](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4)) to authenticate to a token server (e.g. Keycloak) and obtain tokens for accessing
  the Hexa-Orchestrator API service using JWT tokens.

- **hexa-industry-demo-app**

  Runs on [localhost:8886](http://localhost:8886/). A demo web application used to
  highlight enforcing of both coarse and fine-grained policy. The application
  integrates with platform authentication/authorization proxies,
  [Google IAP](https://cloud.google.com/iap) for example, for coarse-grained
  access and the [Open Policy Agent (OPA)](https://www.openpolicyagent.org/)
  for fine-grained policy access. In the docker-compose configuration, the server uses the Hexa-OPA-Server

- **hexa-opa-agent**

  Runs on [localhost:8887](http://localhost:8887/). HexaOPA is
  an [IDQL extended](https://github.com/hexa-org/policy-opa) Open Policy Agent ([OPA](https://openpolicyagent.org))
  server used as an IDQL Policy Decision service. In this configuration, The `hexa-opa-agent` server retrieves policies
  from
  the `hexa-opaBundle-server` instance which is an OPA provisioning service.

- **hexa-opaBundle-server**

  Runs on [localhost:8889](http://localhost:8889/health). An OPA HTTP Bundle server from which the OPA server can
  download policy bundles configured
  by Hexa-Orchestrator. See [OPA bundles][opa-bundles] for more info.

- **keycloak** and **postgres**

  These two servers are used to provide an OAuth2 demonstration and testing environment for Hexa
  services. [Keycloak](https://www.keycloak.org) is pre-configured with a demonstration security "realm" called
  [Hexa-Orchestrator-Realm](http://localhost:8080/admin/master/console/#/Hexa-Orchestrator-Realm). This realm contains
  role definitions and Client credential enabling Hexa Admin UI service to access the Hexa Orchestrator service. By default,
  a bootstrap admin user id and password are configured (see .env_development). It is strongly recommended that these be
  changed.

## Example Workflow

Fine-grained policy management with OPA.

Using the **hexa-admin-ui** application available via `docker-compose`, upload an
OPA integration configuration file. The file describes the location of the IDQL
policy. An example integration configuration file may be found in
[deployments/opa-server/example](demo/deployments/hexaOpaServer/example).

Once configured, IDQL policy for the **hexa-demo** application can be modified
on the [Applications](http://localhost:8884/applications) page. The
**hexa-admin** communicates the changes to the **hexa-orchestrator**, or
"Policy Management Point (PMP)", which then updates the **hexa-demo-config** bundle
server, making the updated policy available to the OPA server.

OPA, the "Policy Decision Point (PDP)", periodically reads config from the
**hexa-demo-config** bundle server and allows or denies access requests based on
the IDQL policy. Decision enforcement is handled within the **hexa-demo**
application or "Policy Enforcement Point (PEP)".

The Hexa Demo architecture may be visualized as follows (note: POSTGRES is no longer used by Hexa-Orchestrator and now
uses a JSON based configuration file specified by `ORCHESTRATOR_CONFIG_FILE`):

![Hexa Demo Architecture](docs/hexa-demo-architecture.svg "hexa demo architecture")

## Getting involved

Take a look at our [product backlog](https://github.com/orgs/hexa-org/projects/1)
where we maintain a fresh supply of good first issues. In addition to
enhancement requests, feel free to post any bugs you may find.

- [Backlog](https://github.com/orgs/hexa-org/projects/1)

Here are a few additional resources for those interested in contributing to the
Hexa project:

- [Contributing](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Development](DEVELOPMENT.md)
- [Security](SECURITY.md)

This [repository](https://github.com/hexa-org/policy-orchestrator) also includes
[documentation](docs/infrastructure/README.md) for the current demo deployment
infrastructure.

[opa-bundles]: https://www.openpolicyagent.org/docs/latest/management-bundles/

<footer>
   <p>The Linux Foundation has registered trademarks and uses trademarks. For a list of trademarks of The Linux Foundation, 
         please see our <a href="https://www.linuxfoundation.org/legal/trademark-usage">Trademark Usage page</a>.
   </p>
</footer>
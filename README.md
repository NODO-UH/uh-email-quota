# uh-email-quota

## Description

This microservice is designed to provide information on the consumption of space in the email given a given user, giving the space consumed and the maximum possible.

## Table of Contents

- [uh-email-quota](#uh-email-quota)
  - [Description](#description)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Compiling](#compiling)
    - [Release](#release)
  - [Usage](#usage)

## Installation

Quota Scraper is developed in the Go language, so one of the ways to install it is by compiling the source code, or by downloading any of the previously compiled releases.

### Compiling

The only requirement to compile the code is to have Go. Once the repository is cloned and all the dependencies are in the `GOPATH`, it can be compiled from the file in`src/main.go`.

### Release

If you want to download a compilation directly, simply go to the releases in this repository and choose the desired version. If you can't find support for a certain architecture, feel free to open an issue to let us know.

## Usage

This microservice runs a REST API on port 8080, in addition, it requires the existence of the `doveadm` command, with which the consumption of each user is consulted. For security reasons, the header `X-API-Key` must be sent in each request with a unique key that only the administrator should know, the key is taken from the environment variable`EMAIL_QUOYA_API_KEY`.

For example, if you want to know the consumption of the user `user@domain.ext`. A request must be sent with the `GET` method and to the path`host:8080/quota?userEmail=user@domain.ext`, with the header `X-API-Key` containing the correct key. If the key matches the one stored in `EMAIL_QUOYA_API_KEY`, then we proceed to find the consumption using the`doveadm` command:

```bash
doveadm -f tab quota get -u user@domain.ext
```

Then it reads the output and returns the result through the REST API through a JSON of the form:

```json
{
    "value": "consumed in bytes",
    "limit": "limit in bytes"
}
```

For example:

```json
{
    "value": 1024,
    "limit": 2048
}
```

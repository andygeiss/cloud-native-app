<p align="center">
<img src="https://github.com/andygeiss/cloud-native-app/blob/main/logo.png?raw=true" />
</p>

# Cloud Native App

[![License](https://img.shields.io/github/license/andygeiss/cloud-native-app)](https://github.com/andygeiss/cloud-native-app/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/cloud-native-app)](https://github.com/andygeiss/cloud-native-app/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/cloud-native-app)](https://goreportcard.com/report/github.com/andygeiss/cloud-native-app)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/30de4509fb6d4b388e2dc52157afb0fd)](https://app.codacy.com/gh/andygeiss/cloud-native-app/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/30de4509fb6d4b388e2dc52157afb0fd)](https://app.codacy.com/gh/andygeiss/cloud-native-app/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

**Cloud Native App** automates the process of bootstrapping a repository.
It generates a project structure with everything you need to start coding right away.

## Project Motivation

The motivation behind **Cloud Native App** is to provide a practical example of implementing
a cloud-native application that adheres to best practices. The project aims to:

1. Automate the process of setting up a new repository.
2. Provide a project structure that follows best practices.
3. Offer a reference implementation for features like CI/CD, testing, and documentation.
4. Inspire developers to adopt cloud-native patterns in their projects.

## Project Setup and Run Instructions

Follow these steps to set up and run the **Cloud Native App**:

### Prerequisites

1. Install the latest version of [Go](https://golang.org/dl/).
2. Install the latest version of [Just](https://github.com/casey/just).
3. Include `$HOME/bin` in your `$PATH`.

### Installation

Clone the repository and install the `cloud-native-app` binary into `$HOME/bin`:

```bash
just install
```

### Usage

Change to your target directory and create a new demo module:

```bash
cloud-native-app demo
```

Navigate into the newly created demo directory and start the HTTP service on port 8080:

```bash
cd demo
just run-service
```

Open the UI in your browser:

[http://localhost:8080/ui](http://localhost:8080/ui])

![alt text](image.png)

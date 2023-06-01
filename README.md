# KVM Dashboard

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/IRONICBo/kvm-dashboard/blob/master/LICENSE)



## Introduction

KVM Dashboard is a web-based tool for monitoring and configuring virtual machines on hosts running [libvirt](https://libvirt.org/).

- Home

![Dashboard](https://github.com/IRONICBo/kvm-dashboard/assets/47499836/8ebdb50d-a3b1-499e-8978-cf050cffc5ab)

- Login

![Login](https://github.com/IRONICBo/kvm-dashboard/assets/47499836/de5a6bf6-245c-4d46-beed-1997a10bbb3a)

## Live

Here is live gif for dashboard running on host machine.

![Live Demo](https://github.com/IRONICBo/kvm-dashboard/assets/47499836/1b90d69b-5152-4c92-9128-700b8a3e73c2)

## How to run

**NOTE: This project is still in progress.**

**Prerequest:** Need to install `libvirtd`, `influxdb`in your local machine.

### 1. Install project

```bash
git clone https://github.com/IRONICBo/kvm-dashboard
cd kvm-dashboard
```

### 2. Modify config file

Modify config in `conf/app.yaml`.

### 3. Install packages

```bash
go mod tidy
```

### 4. Run backend

```bash
go run main.go
```

### 5. Run frontend

```bash
cd web/kvm-dashboard
npm i
npm run serve
```
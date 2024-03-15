// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linode

const tokenID = "7b2c56dd51edd90952c1b94c472b94b176f20c5c777e376849edd8ad1c6c03bb"

// Networking IPs constants
const networking_ips_us_east = `
{
  "page": 1,
  "pages": 1,
  "results": 9,
  "data": [
    {
      "address": "66.228.47.103",
      "gateway": "66.228.47.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li328-103.members.linode.com",
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "172.104.18.104",
      "gateway": "172.104.18.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li1832-104.members.linode.com",
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "192.168.148.94",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "192.168.170.51",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "96.126.108.16",
      "gateway": "96.126.108.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li365-16.members.linode.com",
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "45.33.82.151",
      "gateway": "45.33.82.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li1028-151.members.linode.com",
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "192.168.201.25",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "2600:3c03::f03c:92ff:fe1a:fb4c",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26837992,
      "region": "us-east",
      "public": true
    },
    {
      "address": "2600:3c03::f03c:92ff:fe1a:1382",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east",
      "public": true
    }
  ]
}`

const networking_ips_ca_central = `
{
  "page": 1,
  "pages": 1,
  "results": 2,
  "data": [
    {
      "address": "192.53.120.25",
      "gateway": "192.53.120.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li2216-25.members.linode.com",
      "linode_id": 26837938,
      "region": "ca-central"
    },
    {
      "address": "2600:3c04::f03c:92ff:fe1a:fb68",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26837938,
      "region": "ca-central",
      "public": true
    }
  ]
}`

const networking_ips_all = `
{
  "page": 1,
  "pages": 1,
  "results": 13,
  "data": [
    {
      "address": "192.53.120.25",
      "gateway": "192.53.120.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li2216-25.members.linode.com",
      "linode_id": 26837938,
      "region": "ca-central"
    },
    {
      "address": "66.228.47.103",
      "gateway": "66.228.47.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li328-103.members.linode.com",
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "172.104.18.104",
      "gateway": "172.104.18.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li1832-104.members.linode.com",
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "192.168.148.94",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26837992,
      "region": "us-east"
    },
    {
      "address": "192.168.170.51",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "96.126.108.16",
      "gateway": "96.126.108.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li365-16.members.linode.com",
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "45.33.82.151",
      "gateway": "45.33.82.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li1028-151.members.linode.com",
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "192.168.201.25",
      "gateway": null,
      "subnet_mask": "255.255.128.0",
      "prefix": 17,
      "type": "ipv4",
      "public": false,
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east"
    },
    {
      "address": "139.162.196.43",
      "gateway": "139.162.196.1",
      "subnet_mask": "255.255.255.0",
      "prefix": 24,
      "type": "ipv4",
      "public": true,
      "rdns": "li1359-43.members.linode.com",
      "linode_id": 26848419,
      "region": "eu-west"
    },
    {
      "address": "2600:3c04::f03c:92ff:fe1a:fb68",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26837938,
      "region": "ca-central",
      "public": true
    },
    {
      "address": "2600:3c03::f03c:92ff:fe1a:fb4c",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26837992,
      "region": "us-east",
      "public": true
    },
    {
      "address": "2600:3c03::f03c:92ff:fe1a:1382",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26838044,
      "region": "us-east",
      "public": true
    },
    {
      "address": "2a01:7e00::f03c:92ff:fe1a:9976",
      "gateway": "fe80::1",
      "subnet_mask": "ffff:ffff:ffff:ffff::",
      "prefix": 64,
      "type": "ipv6",
      "rdns": null,
      "linode_id": 26848419,
      "region": "eu-west",
      "public": true
    }
  ]
}`

// IPv6 Ranges constants
const networking_ipv6_ranges_us_east = `
{
  "data": [
    {
      "range": "2600:3c03:e000:123::",
      "prefix": 64,
      "region": "us-east",
      "route_target": "2600:3c03::f03c:92ff:fe1a:fb4c"
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 1
}`

const networking_ipv6_ranges_ca_central = `
{
  "data": [
    {
      "range": "2600:3c04:e001:456::",
      "prefix": 64,
      "region": "ca-central",
      "route_target": "2600:3c04::f03c:92ff:fe1a:fb68"
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 1
}`

const networking_ipv6_ranges_all = `
{
  "data": [
    {
      "range": "2600:3c03:e000:123::",
      "prefix": 64,
      "region": "us-east",
      "route_target": "2600:3c03::f03c:92ff:fe1a:fb4c"
    },
    {
      "range": "2600:3c04:e001:456::",
      "prefix": 64,
      "region": "ca-central",
      "route_target": "2600:3c04::f03c:92ff:fe1a:fb68"
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 2
}`

// Instances
const instances_us_east = `
{
  "data": [
    {
      "id": 26838044,
      "label": "prometheus-linode-sd-exporter-1",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:23:44",
      "updated": "2021-05-12T04:23:44",
      "type": "g6-standard-2",
      "ipv4": [
        "45.33.82.151",
        "96.126.108.16",
        "192.168.170.51",
        "192.168.201.25"
      ],
      "ipv6": "2600:3c03::f03c:92ff:fe1a:1382/128",
      "image": "linode/arch",
      "region": "us-east",
      "specs": {
        "disk": 81920,
        "memory": 4096,
        "vcpus": 2,
        "gpus": 0,
        "transfer": 4000
      },
      "alerts": {
        "cpu": 180,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    },
    {
      "id": 26837992,
      "label": "prometheus-linode-sd-exporter-4",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:22:06",
      "updated": "2021-05-12T04:22:06",
      "type": "g6-nanode-1",
      "ipv4": [
        "66.228.47.103",
        "172.104.18.104",
        "192.168.148.94"
      ],
      "ipv6": "2600:3c03::f03c:92ff:fe1a:fb4c/128",
      "image": "linode/ubuntu20.04",
      "region": "us-east",
      "specs": {
        "disk": 25600,
        "memory": 1024,
        "vcpus": 1,
        "gpus": 0,
        "transfer": 1000
      },
      "alerts": {
        "cpu": 90,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 2
}`

const instances_ca_central = `
{
  "data": [
    {
      "id": 26837938,
      "label": "prometheus-linode-sd-exporter-3",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:20:11",
      "updated": "2021-05-12T04:20:11",
      "type": "g6-standard-1",
      "ipv4": [
        "192.53.120.25"
      ],
      "ipv6": "2600:3c04::f03c:92ff:fe1a:fb68/128",
      "image": "linode/ubuntu20.04",
      "region": "ca-central",
      "specs": {
        "disk": 51200,
        "memory": 2048,
        "vcpus": 1,
        "gpus": 0,
        "transfer": 2000
      },
      "alerts": {
        "cpu": 90,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 1
}`

const instances_all = `
{
  "data": [
    {
      "id": 26838044,
      "label": "prometheus-linode-sd-exporter-1",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:23:44",
      "updated": "2021-05-12T04:23:44",
      "type": "g6-standard-2",
      "ipv4": [
        "45.33.82.151",
        "96.126.108.16",
        "192.168.170.51",
        "192.168.201.25"
      ],
      "ipv6": "2600:3c03::f03c:92ff:fe1a:1382/128",
      "image": "linode/arch",
      "region": "us-east",
      "specs": {
        "disk": 81920,
        "memory": 4096,
        "vcpus": 2,
        "gpus": 0,
        "transfer": 4000
      },
      "alerts": {
        "cpu": 180,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    },
    {
      "id": 26848419,
      "label": "prometheus-linode-sd-exporter-2",
      "group": "",
      "status": "running",
      "created": "2021-05-12T12:41:49",
      "updated": "2021-05-12T12:41:49",
      "type": "g6-standard-2",
      "ipv4": [
        "139.162.196.43"
      ],
      "ipv6": "2a01:7e00::f03c:92ff:fe1a:9976/128",
      "image": "linode/debian10",
      "region": "eu-west",
      "specs": {
        "disk": 81920,
        "memory": 4096,
        "vcpus": 2,
        "gpus": 0,
        "transfer": 4000
      },
      "alerts": {
        "cpu": 180,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    },
    {
      "id": 26837938,
      "label": "prometheus-linode-sd-exporter-3",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:20:11",
      "updated": "2021-05-12T04:20:11",
      "type": "g6-standard-1",
      "ipv4": [
        "192.53.120.25"
      ],
      "ipv6": "2600:3c04::f03c:92ff:fe1a:fb68/128",
      "image": "linode/ubuntu20.04",
      "region": "ca-central",
      "specs": {
        "disk": 51200,
        "memory": 2048,
        "vcpus": 1,
        "gpus": 0,
        "transfer": 2000
      },
      "alerts": {
        "cpu": 90,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    },
    {
      "id": 26837992,
      "label": "prometheus-linode-sd-exporter-4",
      "group": "",
      "status": "running",
      "created": "2021-05-12T04:22:06",
      "updated": "2021-05-12T04:22:06",
      "type": "g6-nanode-1",
      "ipv4": [
        "66.228.47.103",
        "172.104.18.104",
        "192.168.148.94"
      ],
      "ipv6": "2600:3c03::f03c:92ff:fe1a:fb4c/128",
      "image": "linode/ubuntu20.04",
      "region": "us-east",
      "specs": {
        "disk": 25600,
        "memory": 1024,
        "vcpus": 1,
        "gpus": 0,
        "transfer": 1000
      },
      "alerts": {
        "cpu": 90,
        "network_in": 10,
        "network_out": 10,
        "transfer_quota": 80,
        "io": 10000
      },
      "backups": {
        "enabled": false,
        "schedule": {
          "day": null,
          "window": null
        },
        "last_successful": null
      },
      "hypervisor": "kvm",
      "watchdog_enabled": true,
      "tags": [
        "monitoring"
      ]
    }
  ],
  "page": 1,
  "pages": 1,
  "results": 4
}`

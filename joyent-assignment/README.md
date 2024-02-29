# Joyent CDT code challenge
In order to assess your skills and approach please follow the instructions below.

## Assumptions
See the assumptions and commentary [here](./AssumptionsAndComments.md) 

## Go Application
Write a simple API in Go which contains as many endpoints as you wish but should have at least 3
* `/devices` this endpoint should return a list of devices.
* `/interfaces` this endpoint should return a list of interfaces.
* `/reports` this endpoint should use the below JSON Object as POST payload to populate the device and interface information.

### JSON Object
```
{
    "devinfo": {
        "vendor": "Arista, Inc.",
        "name": "arista_switch",
        "serial": "123456",
        "version": "1.2.3",
        "platform": "EOS",
        "LastUpdated": "2023-03-24 13:50:08.403945"
    },
    "intinfo": {
        "fortygige 1/6/1": {
            "name": "fortyGigE 1/6/1",
            "mac": "11:76:22:eb:be:42",
            "description": "a physical interface",
            "driver": "40000mbps",
            "arp": {},
            "ipinfo": [],
            "parent": "port-channel 6"
        },
        "vlan 2222": {
            "name": "Vlan 2222",
            "mac": "",
            "description": "L3 VLTi",
            "driver": "veth",
            "arp": {},
            "ipinfo": [{
                "address": "1.1.1.1",
                "bits": "28",
                "network": "1.1.1.0"
            }]
        }
    }
}
```

## Environment
Using a docker-compose.yaml file please spin up the following environment from scratch:
* Database: postgresql
* Build / Run: using a Dockerfile build your code and then deploy it to a docker image

### note:
The above environment should run with a simple `docker compose up` command. The code should be fully functional and be ready to ingest a report as described above.

## Share on Github
* Create a github repo
* Push all code necessary for completing this task
* Share the repo with user [bdelano](https://github.com/bdelano)


# TODO
- tests!
    - Investigate unit testing 
    - Add functional tests
    - Investigate automatic tests
- review code more. 
- Add CI 
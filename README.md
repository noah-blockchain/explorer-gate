<p align="center" style="text-align: center;">
    <a href="https://github.com/noah-blockchain/explorer-gate/blob/master/LICENSE">
        <img src="https://img.shields.io/packagist/l/doctrine/orm.svg" alt="License">
    </a>
    <img alt="undefined" src="https://img.shields.io/github/last-commit/noah-blockchain/explorer-gate.svg">
</p>

# NOAH Gate

The official repository of Noah Gate service.

Noah Gate is a service which provides to clients publish prepared transactions to Noah Network

Don't forget to read the [documentation](https://noah-blockchain.github.io/noah-gate-docs/)

_NOTE: This project in active development stage so feel free to send us questions, issues, and wishes_

## BUILD

- make create_vendor
- make build

## Configure Extender Service from Environment (example in .env.example)
1) Set up connect to Node which working in non-validator mode. 
2) Set up connect to Gate service. 

## RUN
./gate

_We recommend use our official docker image._
### Important Environments
Example for all important environments you can see in file .env.example.
Its config for connect Node API URL, Gate service and service mode (debug, prod).


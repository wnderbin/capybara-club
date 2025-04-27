# CapybaraClub


<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-capybaraclub">About CapybaraClub</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#project-status">Project status</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#api-documentation">API documentation</a></li>
        <li><a href="#dependencies">Dependencies</a></li>
        <li><a href="#installation-and-launch">Installation & Launch</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#author">Author</a></li>
  </ol>
</details>

## About CapybaraClub

Delivery service with microservice architecture

[![Go](https://github.com/wnderbin/capybara-club/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/wnderbin/capybara-club/actions/workflows/go.yml)

![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

### Built with


![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Alpine Linux](https://img.shields.io/badge/Alpine_Linux-%230D597F.svg?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)

## Project status

Here is information about how the project will work on its different versions.


- [ ] `Launch in containers`
- [X] `Launch workflows`
- [X] `Run locally via makefile`
- [X] `Dependencies`: Postgres & Redis
- [ ] `Microservices Status`: status of tests and health of microservices
  - [X] `User-Service`
  - [ ] `Reustaurant-Service`
  - [ ] `Order-Service`
  - [ ] `Refactoring microservices structure`

## Getting started

Instructions on how to run a project locally

### API documentation

You can find the documentation in the *docs/doc.md*

### Dependencies

#### Installing dependencies

```
go mod download
```


### Installation and Launch

```
git clone https://github.com/wnderbin/capybara-club
cd capybara-club
---
make go-run-service-user 
# launching a microservice that works with users
---
make go-run-service-restaurant
# launching a microservice that works with restaurants
---
make go-run-service-admin
# launching a microservice that works with admins & restaurants
---
```

## Project structure

**.github** - CI/CD \
**docs** - documentation about CapybaraClub \
**user-service** - user microservice \
**admin-service** - admin microservice \
**restaurant-service** - restaurant microservice

## License
Before using the project, it is recommended to read the license

## Author:
* wnderbin


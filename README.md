<div align="center">

  <h1>OVALFi GO-SDK</h1>

  <p>
    A Go SDK for Oval Finance's API Service 
  </p>


<!-- Badges -->
<p>
  <a href="https://github.com/ovalfi/go-sdk/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/Louis3797/awesome-readme-template" alt="contributors" />
  </a>
  <a href="https://github.com/ovalfi/go-sdk/network/members">
    <img src="https://img.shields.io/github/forks/Louis3797/awesome-readme-template" alt="forks" />
  </a>
  <a href="https://github.com/ovalfi/go-sdk/stargazers">
    <img src="https://img.shields.io/github/stars/Louis3797/awesome-readme-template" alt="stars" />
  </a>
  <a href="https://github.com/ovalfi/go-sdk/issues/">
    <img src="https://img.shields.io/github/issues/Louis3797/awesome-readme-template" alt="open issues" />
  </a>
  <a href="https://github.com/ovalfi/go-sdk/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/Louis3797/awesome-readme-template.svg" alt="license" />
  </a>
</p>

<h4>
    <a href="https://github.com/ovalfi/go-sdk">View Demo</a>
  <span> · </span>
    <a href="https://github.com/ovalfi/go-sdk">Documentation</a>
  <span> · </span>
    <a href="https://github.com/ovalfi/go-sdk/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/ovalfi/go-sdk/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->
# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
    * [Tech Stack](#space_invader-tech-stack)
    * [Environment Variables](#key-environment-variables)
- [Getting Started](#toolbox-getting-started)
    * [Prerequisites](#bangbang-prerequisites)
    * [Installation](#gear-installation)
    * [Running Tests](#test_tube-running-tests)
    * [Run Locally](#running-run-locally)
- [Roadmap](#compass-roadmap)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)



<!-- About the Project -->
## :star2: About the Project
This project is an sdk alternative to using OvalFi's public REST APIs. It is written in go and
uses `restyClient` to talk to the public REST APIs over HTTP.


<!-- TechStack -->
### :space_invader: Tech Stack


<details>
  <summary>Server</summary>
  <ul>
    <li><a href="https://go.dev/">Typescript</a></li>
    <li><a href="https://github.com/go-resty/resty">Go-Resty</a></li>
  </ul>
</details>

<!-- Env Variables -->
### :key: Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PUBLIC_KEY`

`BASE_URL`

`BEARER_TOKEN`

<!-- Getting Started -->
## 	:toolbox: Getting Started

<!-- Prerequisites -->
### :bangbang: Prerequisites

This project requires Go >= 1.17

```bash
 brew install go
```

<!-- Installation -->
### :gear: Installation

Install go-sdk with 

```bash
  cd go-sdk
  go install go-sdk
```

<!-- Running Tests -->
### :test_tube: Running Tests

To run tests, run the following command

```bash
  cd go-sdk
  go test
```

<!-- Run Locally -->
### :running: Run Locally

Clone the project

```bash
  git clone git@github.com:ovalfi/go-sdk.git
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Run the local version

Uncomment the lines in `main.go` and change your `BASE_URL` environment variables to
`https://sandbox-api.ovalfi-app.com/api/`
```bash
  go run main.go
```


<!-- Roadmap -->
## :compass: Roadmap

* [x] Customer APIs
* [x] Yield Offering APIs
* [ ] Deposit APIs
* [ ] Withdrawal APIs



<!-- License -->
## :warning: License
[License](https://github.com/ovalfi/go-sdk/blob/main/LICENSE)

Distributed under the GNU General Public License v2.0. See LICENSE.txt for more information.


<!-- Contact -->
## :handshake: Contact

Segun Mustafa - segun@ovalfi.com

Kehinde Odetola - kehinde@ovalfi.com

Olawale Oladapo - olawale@ovalfi.com

Chinonso Okoli - chinonso@ovalfi.com

Project Link: [https://github.com/ovalfi/go-sdk](https://github.com/ovalfi/go-sdk)


<!-- Acknowledgments -->
## :gem: Acknowledgements
- [Awesome README](https://github.com/matiassingers/awesome-readme)



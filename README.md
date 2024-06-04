<div align="center">

  <h1>OVALFi GO-SDK</h1>

  <p>
    A Go SDK for Oval Finance's API Service 
  </p>


<!-- Badges -->
<p>
  <a href="https://github.com/ovalfi/go-sdk/graphs/contributors">
    CONTRIBUTORS
  </a>
*
  <a href="https://github.com/ovalfi/go-sdk/network/members">
    FORKS
  </a>
*
  <a href="https://github.com/ovalfi/go-sdk/stargazers">
    STARS
  </a>
*
  <a href="https://github.com/ovalfi/go-sdk/issues/">
    ISSUES
  </a>
*
  <a href="https://github.com/ovalfi/go-sdk/blob/master/LICENSE">
    LICENSE
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
This project is an sdk alternative to using OvalFi's public [REST APIs](https://docs.ovalfi.com/docs). It is written in go and
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

**Also, we have a system in place to track API requests via the SDK. For every context you pass in the communication with our APIs, we require that you add a `requestID` of type uuid.UUID string to the context.
This must be passed in the context like below:**
````
{

    ctx := context.WithValue(context.Background(), "api_RequestIDContextKey", requestID),

}

Note:In our payload response to your API calls, we now have an header field like this: `X-Request-Id: 71fb13a7-595f-49b8-bdd3-2eb7dcf476c1'
````


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
  go get github.com/ovalfi/go-sdk
```


<!-- Example -->
## :bulb: Example

### Creating A Customer

```go
package main

import (
    "context"
	
    "github.com/ovalfi/go-sdk/api"
    "github.com/ovalfi/go-sdk/model"
)

func main() {
    logger := log.New() // Any logger of your choice 
    client := resty.New() // A REST API client of your choice
    apiCalls := api.New(&logger, client, config.PUBLIC_KEY, config.BEARER_TOKEN, config.BASE_URL)
    ctx := context.Background()
    
    customer, err := apiCalls.CreateCustomer(ctx, model.CreateCustomerRequest{
        Name:            "Nonso",
        Email:           "chinonso@ovalfinance.com",
        Reference:       "ref123",
        MobileNumber:    "09012345678",
        YieldOfferingID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
      }
    )
	
    if err != nil {
    	handleError(err) // Handle the error per your business logic
    	return
    }
	
    UseCustomer(customer) // Use the customers per your business logic
}
```

### Getting Business Portfolios

```go
package main

import (
    "context"
	
    "github.com/ovalfi/go-sdk/api"
    "github.com/ovalfi/go-sdk/model"
)

func main() {
    logger := log.New() // Any logger of your choice
    client := resty.New() // A REST API client of your choice
    apiCalls := api.New(&logger, client, config.PUBLIC_KEY, config.BEARER_TOKEN, config.BASE_URL)
    ctx := context.Background()
    
    portfolios, err := apiCalls.GetBusinessPortfolios(ctx)
	
    if err != nil {
    	handleError(err) // Handle the error per your business logic
    	return
    }
	
    UsePortfolios(portfolios) // Use the portfolios per your business logic
}
```


<!-- Roadmap -->
## :compass: Roadmap

* [x] Customer APIs
* [x] Payout APIs
* [x] Transfer APIs
* [x] Payment APIs
* [x] Currency Swap APIs
* [x] Transaction APIs
* [x] Beneficiary APIs


<!-- Issues -->
## :book: Issues

If you come across a bug or unexpected behaviour, create an issue [here](https://github.com/ovalfi/go-sdk/issues/).
Use the template below to file your complaints.
 - What happened
 - Expected behavior
 - Steps to reproduce
 - Versions


<!-- Contributions -->
## 	:running: Contributions
Contributions are always welcome.

<!-- Development -->
### Development

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

<!-- Running Tests -->
### :test_tube: Running Tests

To run tests, run the following command

```bash
  cd go-sdk
  go test
```


<!-- License -->
## :warning: License
[License](https://github.com/ovalfi/go-sdk/blob/main/LICENSE)

Distributed under the GNU General Public License v2.0. See LICENSE.txt for more information.


<!-- Acknowledgments -->
## :gem: Acknowledgements
- [Awesome README](https://github.com/matiassingers/awesome-readme)



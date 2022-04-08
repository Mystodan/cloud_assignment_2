# Assignment 2
#### By Daniel Hao Huynh

### Implementation
> All Advanced tasks, and required tasks are implemented

### Dependencies 
Here are the dependencies for the project:
- <a href = https://www.postman.com/downloads/> Postman </a> for interracting with the application
- An IDE, we used <a href = https://code.visualstudio.com/download> VSCODE </a>
- <a href = https://git-scm.com/downloads> Git bash package</a>
- <a href = https://go.dev/dl/> Golang </a>


#### Running the CODE~
Use `go run .\server\server.go` to run from `localhost:8080` 
> The default url would therefore be `localhost:8080/corona/v1/` 

## How to use...


**0.**&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Use `go run .\server\server.go` to run from `localhost:8080` or the already deployed `10.212.140.185:8080` <br>
-   **GET** requests using Postman <br>
    -   Use `/corona/v1/` with 
        -   `cases/{countryname, or alpha3 code}`, 
        -   `policy/{countryname, or alpha3 code} {?scope=YYYY-MM-DD}`
        -   `notifications/` in order to see all webhooks
        -   `notifications/{webhook_id}` in order to see a specific webhook
        
-   **POST** requests using Postman <br>
    -   Use `/corona/v1/` with `notification/` with [raw:JSON] in body in order to create a webhook<br>

            {
            "url": string,
            "country": string,
            "calls": int
            }
-   **DELETE** requests using Postman <br>
    -   Use `/corona/v1/` with `notification/{webhook_id}` in order to delete that webhook <br>


## Testing...
#### Stubbing was used for testing purposes
##### All testing is documented in the endpoint `corona/v1/stubbing` VIA **GET** REQUESTS
> Examples: (Using`localhost:8080` or the already deployed `10.212.140.185:8080`)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->** `/corona/v1/stubbing/cases/nor`  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->** `/corona/v1/stubbing/cases/norway` <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->**  `/corona/v1/stubbing/cases<br>`<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->** `/corona/v1/stubbing/policy/nor`<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->**  `/corona/v1/stubbing/policy/norway`<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**->**   `/corona/v1/stubbing/policy/<br>`




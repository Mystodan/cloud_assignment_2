# Assignment 2
#### By Daniel Hao Huynh

### Implementation
> All Advanced tasks, and required tasks are implemented
#### Difference between local and server deployment
One small change made to the web deployed service is that the constants which contain the directories for the required json files have an extra `.`<br>
> This is because of how its being deployed, the application searches from the file differently and therefore requires it to be so.<br>
However **This is not such a big change that requires it to have its own repository for deployment**

> **For recreational purposes** heres an example: localhost("./global_types/alpha3.json") -> 10.212.140.185("../global_types/alpha3.json")<br>

#### **!!!** Current state of Policy API
> **[08.04.2022 - 06.26]** Currently when writing this, the policy api might be online, **HOWEVER** the stringency data returns data unavailable for every date<br>
Therefore it seems like my application returns `Data unavailable` for the policy api, However on the **stub documentation** i have documented what happens on current date<br>
When the stub was recorded, which was on the [05.04.2022]

### Project Structure
-   All constants are contained within the **constants** folder<br>
    - Contains appConstants, dependConstants, serverConstants
-   All Global data types are contained within the **global_types** folder <br>
    - Alpha3 local library is also located here since its a library accessed globally<br>
-   All endpoints are contained within the **endpoints** folder<br>
    - The endpoints folder also contain common functions for all endpoints<br>
    - Testing is also documented within an endpoint `/corona/v1/stubbing/policy/`
-   All Server functions and caching is contained within the **server** folder <br>





### Dependencies 
Here are the dependencies for the project:
- <a href = https://www.postman.com/downloads/> Postman </a> for interracting with the application
- An IDE, we used <a href = https://code.visualstudio.com/download> VSCODE </a>
- <a href = https://git-scm.com/downloads> Git bash package</a>
- <a href = https://go.dev/dl/> Golang </a>
- FireBase

## Deployment
Deployed on NTNU's OpenStack<br>
serviceAccountKey for firebase was passed to the server using: <br> `scp -i MyKey.pem ./serviceAccountKey.json ubuntu@10.212.140.185:\assignment-2
serviceAccountKey.json ` 

#### Running the CODE~
Use `go run .\server\server.go` to run from `localhost:8080` 
> The default url would therefore be `localhost:8080/corona/v1/` 


## How to use...

#### **0a.**&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Special Requirement:
###### -    IT IS REQUIRED THAT YOU CREATE A FOLDER IN `assignment-2` FOLDER AND PUT IN YOUR FIREBASE SERVICE KEY AS `serviceAccountKey.json` IN ORDER TO RUN IT LOCALLY<br>

**0b.**&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Use `go run .\server\server.go` to run from `localhost:8080` or the already deployed `10.212.140.185:8080` <br>
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




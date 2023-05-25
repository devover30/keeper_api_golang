# Activity Flow

## Client and mobile login activity flow in htttp rest api

##### Guidlines for managing the authentication flow across platform like mobile app or web app

> 1.  Whatever be the means single source of truth for authorizing user is the mobile number of the said user.

<!-- -->

> 2.  User should register with his/her mobile number.

<!-- -->

> 3.  Upon registering mobile user will get otp for verification for the said mobile number.

<!-- -->

> 4.  Once the said **OTP** is verified the said user's mobile number **isvalid** flag in Database be updated to **TRUE**

<!-- -->

> 5.  Once user hits login resource server should validate the **mobile number** and also check **isvalid** flag in the database.

<!-- -->

> 6.  Authorization/authentication Server should issue a jwt token to be used by the mobile app or web app for accessing application resource.

<!-- -->

> 7.  Authorization/authentication Server should also expose a endpoint to verify the issued token and return appropriate reponse.

##### Authorization/authentication server http rest points with explaination.

- **Notes:**: this is a draft proposal and can change overtime according to software needs.

<details>
 <summary><code>POST</code><code><b>/members</b></code> <code>Register a new user </code></summary>

##### Body

```
{
   mobile: <10 digit mobile number>
}
```

##### Responses

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**
    ```
    {
        hash : <hash generated by server>
    }
    Note: hash generated by server need to be sent back with the otp received on mobile for validating the user.
    ```

- **Error Response:** \* **Code:** 500 SERVER ERROR <br />
**Content:**
`   { 
        error : "Internal Server Error....Try Again Later" 
      }
  `
</details>

<details>
 <summary><code>POST</code><code><b>/members/verify</b></code> <code>Verify register user.</code></summary>

##### Body

```
{
   mobile: <10 digit mobile number>,
   otp: <6 digit otp received through sms on mobile number>,
   hash: <hash received when register user response>
}
```

##### Responses

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```
    {
        id : <uuid>,
        mobile : <request body mobile no>
    }
    ```

- **Error Response:**

      | http code| content-type|response|
      |----------|-------------|--------|
      | `400`    | `application/json`   | `{"error":<error message>}`|
      | `500`    | `application/json`   | `{"error":<error message>}`|

  </details>

<details>
 <summary><code>POST</code><code><b>/authentication</b></code> <code>Login a user by his/her Mobile Number</code></summary>

##### Body

```
{
   mobile: <10 digit mobile number>
}
```

##### Responses

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```
    {
        hash : <hash generated by server>
    }
    Note: hash generated by server need to be sent back with the otp received on mobile for validating the user.
    ```

- **Error Response:**
| http code| content-type|response|
|----------|-------------|--------|
| `400` | `application/json` | `{"error":"Bad Request"}`|
| `401` | `application/json` | `{"error":"Invalid Mobile Number"}`|
| `500` | `application/json` | `{"error":"Server Error...Try Again Later"}`|
</details>

<details>
 <summary><code>POST</code><code><b>/authentication/verify</b></code> <code>Verify otp and hash</code></summary>

##### Body

```
{
   mobile: <10 digit mobile number>,
   otp: <6 digit otp received through sms on mobile number>,
   hash: <hash received when authenticate user response>
}
```

##### Responses

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```
    {
        token: <jwt token issed by server>
    }
    Note: hash generated by server need to be sent back with the otp received on mobile for validating the user.
    ```

- **Error Response:**
| http code| content-type|response|
|----------|-------------|--------|
| `400` | `application/json` | `{"error":"Bad Request"}`|
| `401` | `application/json` | `{"error":"Invalid Mobile Number"}`|
| `500` | `application/json` | `{"error":"Server Error...Try Again Later"}`|
</details>

<details>
 <summary><code>GET</code><code><b>/authentication/verify</b></code><code>Verify JWT token received</code></summary>

##### Header

```
   Authorization: Bearer <jwt token>
```

##### Responses

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```
    {
        id : <id stored in jwt token>,
        mobile : <user data stored in jwt token>,
    }
    ```

- **Error Response:**
| http code| content-type|response|
|----------|-------------|--------|
| `400` | `application/json` | `{"error":"Bad Request"}`|
| `401` | `application/json` | `{"error":"Invalid Token"}`|
| `500` | `application/json` | `{"error":"Server Error...Try Again Later"}`|
</details>
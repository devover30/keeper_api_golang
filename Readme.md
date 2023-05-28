# Rest API Flow

## Authentication and Resource Flow in http rest api

##### Guidelines for managing the rest api flow for usage by platforms like mobile app or web app

> 1.  Whatever be the means **IDP** is the data returned by the Auth Gateway Server.

<!-- -->

> 2. Each Request coming from Web App or Mobile App will be intercepted or pass through Middleware<br>which will check whether Request has authorization header and contains bearer token.

<!-- -->

> 3. If the above condition passes then there will be second interceptor/middleware which will authenticate<br>the token through **AUTH GATEWAY** server.

> 4.  User Mobile will be the single source of truth for all the accounts and it's transactions.

##### Keeper server http rest points with explanation.

- **Notes:**: this is a draft proposal and can change overtime according to software needs.

<details>
 <summary><code>POST</code><code><b>/credentials</b></code> <code>Create a new credential</code></summary>

##### Header

**`Authentication: Bearer <token>`**

##### Body

**`Content-Type:application/json`**

```
{
   platformName: <platform name>,
   username: <user name of website or app>,
   password: <password of website or app>
}
```

##### Responses

- **Success Response:**

  - **Code:** 201 <br />
    **`Content-Type:application/json`**
    ```
    {
      id: <newly created credential id>,
      platformName: <platform name>,
      username: <user name of website or app>,
      password: <password of website or app>,
      createdAt: <credential creation date>,
      modifiedAt: <credential modification date>
    }
    ```

- **Error Response:**

  | http code | content-type       | response                                    |
  | --------- | ------------------ | ------------------------------------------- |
  | `400`     | `application/json` | `{"error":"Invalid Request"`                |
  | `401`     | `application/json` | `{"error":"Invalid Token"`                  |
  | `500`     | `application/json` | `{"error":"Server Error...Try again later"` |

  </details>

<details>
 <summary><code>GET</code><code><b>/credentials</b></code> <code>Get all the credentials bind to the user.</code></summary>

##### Header

**`Authentication: Bearer <token>`**

##### Responses

**Success Response:**

- **Code:** 200 <br />
  **`Content-Type:application/json`**

  ```
  [
    {
        id: <credential id>,
        platformName: <platform name>,
        username: <user name of website or app>,
        password: <password of website or app>,
        createdAt: <credential creation date>,
        modifiedAt: <credential modification date>
    },
    ....
  ]
  ```

- **Error Response:**

  | http code | content-type       | response                                    |
  | --------- | ------------------ | ------------------------------------------- |
  | `400`     | `application/json` | `{"error":"Invalid Request"`                |
  | `401`     | `application/json` | `{"error":"Invalid Token"`                  |
  | `500`     | `application/json` | `{"error":"Server Error...Try again later"` |

</details>

<details>
 <summary><code>PUT</code><code><b>/credentials/:id</b></code> <code>Update a credential.</code></summary>

##### Header

**`Authentication: Bearer <token>`**

##### URL Param

**`:id = <credential id>`**

##### Body

```
{
   platformName: <platform name>,
   username: <user name of website or app>,
   password: <password of website or app>
}
```

##### Responses

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**

    ```
    {
        id: <credential id>,
        platformName: <platform name>,
        username: <user name of website or app>,
        password: <password of website or app>,
        createdAt: <credential creation date>,
        modifiedAt: <credential modification date>
    }

    ```

- **Error Response:**

  | http code | content-type       | response                                     |
  | --------- | ------------------ | -------------------------------------------- |
  | `400`     | `application/json` | `{"error":"Bad Request"}`                    |
  | `401`     | `application/json` | `{"error":"Invalid Token"}`                  |
  | `500`     | `application/json` | `{"error":"Server Error...Try Again Later"}` |

  </details>

<details>
 <summary><code>DELETE</code><code><b>/credentials/:id</b></code> <code>Delete a credential</code></summary>

##### Header

**`Authentication: Bearer <token>`**

##### URL Param

**`:id = <credential id>`**

##### Responses

- **Success Response:**

  - **Code:** 204 <br />

- **Error Response:**

  | http code | content-type       | response                                     |
  | --------- | ------------------ | -------------------------------------------- |
  | `400`     | `application/json` | `{"error":"Bad Request"}`                    |
  | `401`     | `application/json` | `{"error":"Invalid Token"}`                  |
  | `500`     | `application/json` | `{"error":"Server Error...Try Again Later"}` |

  </details>

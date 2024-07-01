# Create User

## Description
API to create user

## BaseURL
`https://localhost:8080`

## Endpoint
`[POST] /api/v1/user`

### Headers
```
Content-Type: application/json
Authorization: Basic {token}
```
### Query parameters
  - `type` (optional, string): 

**Examples**

Create student
```json
"?type=student"
```

Create teacher
```json
"?type=teacher"
```
### Body

  - `id` (optional, uint): ID of user
  - `username` (required, string): Username of user
  - `profile` (optional, object): Profile of user
    - `fullname` (optional, string): 
    - `birthday` (optional, string): 
    - `email` (optional, string): 
    - `address` (optional, string): 
  - `pets` (optional, array of object): List of pets that user has
    - `id` (optional, uint): 
    - `name` (optional, string): 
    - `type` (optional, string): 
    - `profile` (optional, object): 
      - `fullname` (optional, string): 
      - `birthday` (optional, string): 
      - `email` (optional, string): 
      - `address` (optional, string): 

**Examples**

Create user
```json
{
    "id": 1,
    "pets": [],
    "profile": {},
    "username": "vuhk"
}
```
### Response
Response when create user success
  - `meta` (optional, object): Metadata of response
    - `code` (optional, int): Status code
    - `message` (optional, string): Message of response
    - `request_id` (optional, string): ID of request
  - `data` (optional, string): UserID of created user

**Examples**

Success Response
```json
{
    "data": "1",
    "meta": {
        "code": 200,
        "message": "success",
        "request_id": "123456"
    }
}
```

Error Response
```json
{
    "meta": {
        "code": 400,
        "message": "error",
        "request_id": "123456"
    }
}
```
### Examples

Create user
```bash
curl -u
```

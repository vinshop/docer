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
```
### Query parameters
  - `type` (optional, string): 

Example
```json
"?type=student"
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

Example
```json
{
	"id": 1,
	"pets": [],
	"profile": {},
	"username": "vuhk"
}
```

### Response
- Success Response
```json
{
	"meta": {
		"message": "success",
		"status": 200
	}
}
```
- Error Response
```json
{
	"meta": {
		"message": "[error message]",
		"status": 400
	}
}
```

### Example
```bash
curl -o
```

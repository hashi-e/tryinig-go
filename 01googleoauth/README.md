# 01 Google OAuth Trying
Try to use Google OAuth and obtain email.

# How to try
## get Library
go get -u google.golang.org/api/oauth2/v2

## Prepare Credentials
Create Credentials on Google Developer Console
https://console.developers.google.com/apis/
And add http://localhost:5001/loginr into Authorised redirect URIs

## Replace source code
Replace "Your Client ID" and "Your Client Secret" to your credentials info

## Try
go run main.go
try to access http://localhost:5001/login

# Refer to
https://developers.google.com/identity/protocols/OAuth2
https://github.com/googleapis/google-api-go-client

# 01 Google OAuth Trying
Try to use Google OAuth and obtain email.

# How to try
## get Library
go get -u google.golang.org/api/oauth2/v2

## Create Credentials on Google Developer Console
And add http://localhost:5001/loginr into Authorised redirect URIs

## Replace client info
Replace "Your Client ID" and "Your Client Secret" to your credentials info

## try
go run main.go
try to access http://localhost:5001/login

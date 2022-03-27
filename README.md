# request-decryption-middleware

Siumulates interception of request , decryption (simply changing value in a request), before passing it on to the controller.

Giving a request as 
{
    "old_credential" : "123",
    "new_credential" : "Username"
}
to the endpoint 
http://localhost:8080/api/change-password

will print

**Decrypted credentials are  {decrypted_123 decrypted_Username 2022-03-28 01:22:02.605495 +0530 IST}**

it shows the values from the request are replaced by "decrypted_" before reaching the controller. It would simulate an actual decryption call.

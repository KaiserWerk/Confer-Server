# Confer Server

Confer allows you to edit multiple ASCII files on your remote servers with a
central UI. Confer Server is the server component of Confer.

### Usage

After building for your OS/Arch, just start the binary with the following 
command (usually on linux servers):
    
    ./confer-server -h "127.0.0.1:1663" -k "your auth key here"
    
* A TLS connection is not used here to keep it simple.
* The auth key is supposed to be some arbitrary, password-like string.
* You can omit the IP address.

Now, the listener is running and accepts requests for the path ``/file``.

A ``POST`` request containing a ``requestedFile`` struct (just with the file 
name) will return the content of the requested file.

A ``PUT`` request containing a `requestedFile struct` (containing both file 
name and new file content) will save the supplied content to file, 
effectively overwriting its existing content.

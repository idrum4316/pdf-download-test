# Test Server for Issue 1110937
This is a small web server I slapped together that replicates what is described in Chromium issue [1110937](https://bugs.chromium.org/p/chromium/issues/detail?id=1110937&q=pdf&can=2).

To replicate the issue, you need to serve a PDF as the reponse to a POST request (I only tested GET and POST), and the server needs to have an invalid SSL certificate.
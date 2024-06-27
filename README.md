**Go Encrypt**

A plaintext encryption application with a Golang backend and a NextJS frontend.

I implemented AES-256 encryption to secure plaintext data. After encryption, the resulting cipher text is stored in MongoDB.

To decrypt the cipher text, the key is passed to the backend, which executes the decryption algorithm and returns the plain text.

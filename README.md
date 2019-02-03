# Redirector
Redirects any incoming requests to a defined target.

## Environment variables:

`LISTEN`, the address to listen on.
`REDIRECT_TO`, the address to redirect to.

## Usage:

```
docker run -p 19489:19489 -e LISTEN=":19489" -e REDIRECT_TO="https://github.com/machiel/redirector" -it machiel/redirector
```

# URL Shortener

A simple URL Shortener App built in Go.

It's an API that can receive urls and return shortened ones which are then mapped to the originals for redirection.

The API is hosted [here](git@github.com:AlejandroJorge/url-shortener-go.git)

The client for interacting with the API is hosted [here](https://trim.up.railway.app/)

# How to run the app

## Running the API

The easiest way to run this app is to first clone the repository:

```
$ git clone https://github.com/AlejandroJorge/url-shortener-go
```

Then enter the API folder:

```
$ cd url-shortener-go/api
```

Download all the dependencies:

```
$ go get
```

And finally run the app:

```
$ go run .
```

> The default port is 3000, but it can be changed by the PORT environment variable:
>
> ```
> $ export PORT=5000 && go run .
> ```

## Running the Client

First, go to the client folder:

```
$ cd ../client
```

Download all dependencies:

```
$ npm install
```

Run the app like this:

```
$ export VITE_API_URL=http://localhost:3000/urls && npm run dev
```

> 3000 is the port, if you setted a different port for the API to run, you should set it here as well

# Endpoints

## POST /urls

This endpoint accepts the following request format as JSON in the body:

```
{
  "originalURL": "https://foo.bar"
}
```

It will respond with the following response:

```
{
  "shortenedURL": "https://localhost:5173/foo"
}
```

> Here foo is a random generated string for mapping the URL provided

## GET /:shortenedPath

This endpoint compares the shortenedPath accesed and, if it is mapped in the app, redirects to its original URL

# Implementation details

## Technologies used

- SQLite3 is used for database since it's a very small app
- GORM is used as an ORM for interacting with the database
- Gorilla/mux is used for easy and minimalist HTTP routing
- Vite and React are used for building the client
- Deployed using Docker for easier deployment in Railway

# Why I built this project

I primarily wanted to make a "real" backend project, with database interactions, routing, environment variables and basic architecture

# diccionario

A simple word list service control plane.

It has the following capabilities:

* Check if a word exists in the word list
* Add a new word to the word list
* Find a list of words that match a given prefix

The word list is stored in a flat file at [words.txt](./words.txt).

# Endpoints

## /exists/:word

This endpoint checks if a word exists in the word list.
This endpoint has some issues with it's implementation.

Expected functionality:
* It returns a 200 upon success
* It returns other status codes as appropriate (4XXs for input errors, 5XXs for internal server errors)
* The response body is a JSON object with a single field `exists` of type boolean
  * Example: `{ "exists": true }`
* It performs case insensitive matching to the words in the wordlist
* It only returns true if the word exists (exactly matches)in the wordlist

## /add

This endpoint adds a new word to the word list.
This endpoint needs to be implemented still.

Expected functionality:
* It returns a 204 upon success.
* It returns a 409 if the word already exists in the word list.
* It returns other status codes as appropriate (4XXs for input errors, 5XXs for internal server errors)
* The newly added word should persist for the life of the running Docker container.
* A word is considered a single string of unbroken alpha characters (no numbers or special characters)

## /matches/:prefix

This endpoint returns a list of words that match a given prefix.
This endpoint needs to be more performant.

Expected functionality:
* It returns a 200 upon success
* It returns other status codes as appropriate (4XXs for input errors, 5XXs for internal server errors)
* The response body is a JSON object with a single field `matches` of type string array
  * Example: `{ "matches": ["word1", "word2"] }`
* It performs case insensitive matching to the words in the wordlist
* It only returns words that match the given prefix

# Go Implementation

All of the Go-related code is in the `go` directory.

From the root of this repo run:

```sh
docker compose up --build go
```

The server will be available at http://localhost:8080.
It will automatically reload when you make changes to the source code
(it's using [air](https://github.com/air-verse/air)).

To stop the server, press Ctrl+C in the terminal where it's running.

To access the running container, run:

```sh
docker compose exec go bash
```

To exit the terminal session, press Ctrl+D in the terminal where it's running.

# Typescript Implementation

All of the Typescript-related code is in the `ts` directory.

From the root of this repo run:

```sh
docker compose up --build --renew-anon-volumes ts
```

The server will be available at http://localhost:5000.
It will automatically reload when you make changes to the source code.
(`--renew-anon-volumes` gives the container a fresh copy of the image's
`node_modules`, so dependency changes show up after a rebuild.)

To stop the server, press Ctrl+C in the terminal where it's running.

To access the running container, run:

```sh
docker compose exec ts bash
```

To exit the terminal session, press Ctrl+D in the terminal where it's running.

# golang-autocomplete-api

AutoComplete API is a packaged front and back-end software solution that can be integrated into existing web applications for intelligent autocompletion. The API tracks user typing and behavior, and builds a list of common words and phrases utilized by the user. In the background, the server builds up a Trie data structure using a MongoDB non-relational data store which can be rapidly queried to provide intelligent suggestion.

The API exposes a single endpoint:

```
/api/type
```

which recieves `POST` requests from the client on each keypress event in an input field.

Based on the state of the payload, the server will respond with a set of commonly-used words and sentances beginning with that prefix, or an empty set if a suggestion is not appropriate.

Using these intelligent suggestions, users can manage easy autocompletion on the front-end using a complimentary React library.

## Dependencies

This project uses Golang and MongoDB to supply lighning-fast API responses to common word prefixes. In order to use this project, please visit [Getting Started with Go](https://golang.org/doc/install) and the [MongoDB Installation Guide](https://docs.mongodb.com/manual/installation/) and follow their installation guides.

If integrating with our front-end library, [Node.js](https://nodejs.org/en/) and [NPM](https://docs.npmjs.com/) are also required installations.

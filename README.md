Overview
---
This is a chrome extension to parse google query result and pass it to the server

- Chrome extension using vanilla javascript
- Server using Go

How to install extension
---

- open chrome and type following url `chrome://extensions/`
- turn on the `Developer mode` toggle
- click `Load unpacked` button
- choose `chrome-ext` directory inside this repo
- turn on the extension toggle

How to run server
---
To run it we can use `make` command or manual run using `go run main.go`.   
Here is `make` command to run this application:   

- `make run` to run launch this application into port `:8080`, 
- `make test` to run the all test case and check the coverage  

This application support MySQL as a database, but we can implement a different database.  
By default, it will connect into our mySQL database with default host & port `localhost:3306` and database `users`  
To connect into a different database we need to set database information in environment variable.  


After setting the database information we only need to run the main.go file  
`go run main.go`  

#### API List & Payloads
Here is API and its payload:  

1. [POST] **/search**  
```json
[{
    "title":  "Title 1",
    "desc":   "Desc 1",
    "url":    "www.url1.com/url",
    "domain": "www.url1.com",
    "cite":   "Cite 1",
}, {
    "title":  "Title 2",
    "desc":   "Desc 2",
    "url":    "www.url2.com/url",
    "domain": "www.url2.com",
    "cite":   "Cite 2",
}]
```

Project Structure
---
Following is project structures explanation:

1. **chrome-ext**  
contains chrome extensions
2. **server/api**  
contains handler for API
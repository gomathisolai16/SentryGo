## Requirements

For development, you will only need Go installed in your environement. You can go to https://golang.org/ this site and install go.

If the installation was successful, you should be able to run the following command.

    $ go version
     go1.14.1
### Next you have to get these packages by running this command
- #### go get -u github.com/getsentry/sentry-go/gin
- #### go get -u github.com/getsentry/sentry-go


# How to Integrate Sentry
### Step 1:  Create a project in sentry 

1. Go to this site and create a project https://sentry.io by seleting GO 

2. You will get an sample code with your public dsn link in sentry

3. Secure that link for using in the project 

### Step 2: Import Packages

import 
(
    "github.com/gin-gonic/gin"
   sentry  "github.com/getsentry/sentry-go"
  sentrygin "github.com/getsentry/sentry-go/gin"
)

### Step 3: Write a small code for sentry integration

1. We are going to use gin framework for creating web server.

2. Define router and web server port in the main function.

3. Initialize the sentry in the router. Use the YOUR_PUBLIC_DSN key in Dsn argument.

4. You will get an error if sentry is not initialized properly.

5. As we are using gin, sentry has to integrated for all the routes in the application.

6. You can write a small function to test the error is capturing in sentry.
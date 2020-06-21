# Google Reseller CLI

Perform customer, subscription, user and site verification functions from the command line

**Authentication**
Follow the Google service account set up for authentication. Create a folder named 'creds' in the same directory as the binary with a filename of credentials.json and store the service_acount json file here.

Use this guide to ensure the service account credentials are created properly along with relevant permission steps:
https://developers.google.com/admin-sdk/reseller/v1/codelab/intro

This is also handy for reference
https://developers.google.com/identity/protocols/oauth2/service-account

**Google api docs**
This cli makes use of The Google Reseller API, Admin SDK Directory API and Site verification API. See the docs for each api here:

Reseller: https://developers.google.com/admin-sdk/reseller/v1/get-start/getting-started
Directory: https://developers.google.com/admin-sdk/directory/
Site verification: https://developers.google.com/site-verification

##CLI Usage

### Customers

***To see All Customer Commands Use:***
```
gsuite customer
```

***Fetch a customers details***
```
gsuite customer get --id CUSTOMERID
```
The id will be either the primaryDomain of the customer or the google Customer id 

***Create a new customer***

```
gsuite customer create --customer {jsonData}
```
The jsonData here must be escaped to work on the command line! The jsonData provided should be a customer object on the Google Reseller API:
https://developers.google.com/admin-sdk/reseller/v1/reference/customers

### Subscriptions
***To see All Subscription Commands Use:***
```
gsuite subscription
```

***Fetch a Subscriptions details***
```
gsuite subscription get --id CUSTOMERID
```
The id will be either the primaryDomain of the customer or the google Customer id 

***Create a new subscription***
```
gsuite subscription create --id CUSTOMERID --subscription {jsonData}
```
The jsonData here must be escaped to work on the command line! The jsonData provided should be a subscription object on the Google Reseller API:
https://developers.google.com/admin-sdk/reseller/v1/reference/subscriptions

***Suspend a Subscription***
```
gsuite subscription suspend --id CUSTOMERID
```

***Activate a Subscription***
```
gsuite subscription activate --id CUSTOMERID
```

### Users
***To see All User Commands Use:***
```
gsuite user
```

***Fetch a Users Details***
```
gsuite user get --userKey USERKEY
```
A user Key is either the primary email address of the user or the unique userId (Recommended)

***Create a new User***
```
gsuite user create --user {jsonData}
```
The jsonData here must be escaped to work on the command line! The jsonData provided should be a user object on the Google Directory API
https://developers.google.com/admin-sdk/directory/v1/reference/users

***Update a User***
```
gsuite user update --userKey USERKEY --user {jsonData}
```

***Make User super admin***
```
gsuite user make-admin --userKey USERKEY
```

***Delete a User***
```
gsuite user delete --userKey USERKEY
```

###Aliases

***Add an Alias to a User***
```
gsuite user add-alias --userKey USERKEY --userAlias ALIAS
```
The Alis here is a full email address e.g. <n/>alias@primary-domain.com

***Delete an Alias***
```
gsuite user delete-alias --userKey USERKEY --userAlias ALIAS
```

### Site verification
***To see All Verification Commands Use:***
```
gsuite verification
```

***Get Verification Token***
```
gsuite verification get-token --domain DOMAIN
```
The domain here is the users primary domain without www. e.g. <n/>primary-domain.com
This command currently defaults to DNS verification with a TXT record

***Verify***
```
gsuite verify --domain DOMAIN
```
This will attempt to verify a domain (ensure verification steps have been carried out to get a success response here)
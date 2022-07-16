# Manu CMS Docs

## General Info:

All requests require an Authorization cookie obtained from /api/login

## API Endpoints:

### /api/login

Type: GET
Query: username, password
Response: Nil
Sets Authorization Cookie.

### /api/adduser

Type: GET
Query: username, password
Response: Nil
Adds user to CMS with provided username and password.

### /api/addcollection

Type: GET
Query: name
Response: Nil
Creates Collection with provided name.

### /api/isauthenticated

Type: GET
Query: Nil
Response: Nil
Checks whether authorization cookie is valid/not expired, response header is 200 if so, 403 otherwise.

### /api/collections

Type: GET
Query: Nil
Response: [string]
Returns a string array of collection names in database.

### /api/getcollection

Type: GET
Query: name
Response: [document]
Returns a document array of documents in requested collection.

### /api/createdocument

Type: POST
Body: { collectionName: string, [map<string, string>] }
Response: Nil
Creates a document in collection name, provided an array of key value pairs.

### /api/getdocument

Type: GET
Query: documentid, collectionname
Response: document
Returns document with matching documentid from in collection corresponding to collectionname query paramater.

### /api/updatedocument

Type: POST
Body: { collectionName: string, documentId: string, doc: [map<string, string>] }
Response: Nil
Updates document with matching documentId from collection collectionName to doc.

### /api/deletedocument

Type: GET
Query: documentid, collectionname
Response: Nil
Deletes document of documentid from collection collectionName.


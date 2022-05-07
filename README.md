# RiseExercise

Welcome to phone contact app!

To run the app, run go contacts.go
The app will run on localhost:8000

While the app is running you can make the following requests:
http://localhost:8000/contacts to see the existing contacts - will show the 10 first contacts in the book
http://localhost:8000/contacts/page/{n} will show you the nth page (10 contacts per page) if there are enough contacts for the page to exist.
/contacts/page/1 yields the same result as /contacts

To add a contact you need to make a POST request to http://localhost:8000/createContact with json body with the following data:
FirstName, LastName, PhoneNumber

An example of the command to run in parallel terminal :
curl -i -X POST -H 'Content-Type: application/json' -d '{"FirstName": "Amir", "LastName": "Israel", "PhoneNumber": "0547643583"}' http://localhost:8000/createContact

This will add a new contact automatically with ID as an autoincrement.

To delete a contact, run http://localhost:8000/deleteContact/{id} where id is the autoincremented Id of the contact you wish to delete.

To edit a contact, run http://localhost:8000/updateContact/{id} as a POST request with json body with the updated data. Example:
curl -i -X POST -H 'Content-Type: application/json' -d '{"FirstName": "Amir", "LastName": "Israeli", "PhoneNumber": "0547644586"}' http://localhost:8000/updateContact/3

To search for a pattern, which can be either in the first name, the last name or the phone number, use http://localhost:8000/search/{pattern}

---- Ideas of features to improve ---
1. Add more fields to the contacts (email, address...)
2. Perform field sanity check (for example, check the phone number has only numbers or +)
3. Forbid duplicates (couple of first name and last name already existing or phone number already existing)
4. Save the contacts to a DB (right now, when you stop running the application, the data is lost)
5. Build a UI

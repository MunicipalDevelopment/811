# 811: NM One Call

NMOneCall (811) is a service which notifies subscribers when someone is going to dig. It allows us to line spot our utilities to make sure nothing is hit. They have
moved from TelDig to a new system. This new system has a different ticket format and an API. The code in the repo handles the ticket parsing, includes a front end, and sends results to the API.

Parse811.py is a python script to parse the XML tickets and send them to the DB. This should run at an interval.

Index.html is the main entrance to the front end. It is a table of tickets from the DB. Selecting a ticket will take you to detail.html. This page will display the location of the ticket and a form to add comments and the resolution.

Submitting the form connects to the GoServer and sends the data to NMOneCall. Currently, they do not allow CORS so it has to happen server side. I have put in a request to change this.

Application needs to update the DB on success with resolution and success. Then remove from the table on the main page (only display open tickets). Failures can be handle any way you would like.

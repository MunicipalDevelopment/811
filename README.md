# 811: NM One Call

Tickets sent from NMOneCall when someone wants to dig. Need to respond and possibly line spot the utilities to make sure they will not be hit.

The Parse811.py file will load tickets from a directory in to a database. Set this to run at interval.
You need a DB table to dump it in.

Publish the table as a REST Endpoint. This code will be available later.

Consume REST in webpage. The webpage has a list of tickets(index.html) and each links to detail.html. This shows a map, layers you need and a form to post your response/close ticket.

This system was built due to changes in the 811 ticketing system that made it incompatible with our system - changes to the file structure of tickets and new API for responding/closing.


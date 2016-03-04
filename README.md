# 811: NM One Call

Tickets sent from NMOneCall when someone wants to dig. Need to respond and possibly line spot the utilities to make sure they will not be hit.

The Parse811.py file will load tickets from a directory in to a database.

We then publish the table as a REST Endpoint and consume it in a webpage. The webpage has a form to post responses back to 811 to close the tickets with our response and comments.

This system was built due to changes in the 811 ticketing system that made it incompatible with our system - changes to the file structure of tickets and new API for responding/closing.

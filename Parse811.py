import xmltodict
import pyodbc
import os
from os.path import isfile, join


#CONNECTION AND CURSOR

cnxn = pyodbc.connect('DRIVER={SQL Server};SERVER=MyServer;DATABASE=MyDB;')
cursor = cnxn.cursor()



#GET ALL EXISTING TICKET NUMBERS TO MAKE SURE NOT DUPLICATE. 

have=[]
cursor.execute("select * from dbo.NMOneCall")
rows = cursor.fetchall()
for row in rows:
    have.append(str(row.number))



#READ DIRECTORY FOR TICKETS. EXCLUDE WHAT WE ALREADY HAVE - ADD FUCNTIONALITY TO DELETE ONCE READ AND SENT TO DB.

path="C:\Users\me\Desktop\811Tickets"
thefiles=[f for f in os.listdir(path) if isfile(join(path,f)) and f.endswith(".xml")]


for x in thefiles:
    if "Audit" not in x:
        with open(join(path,x)) as fd:
            doc = xmltodict.parse(fd.read())
            number = str(doc['ticket']['ticket']['number'])
            if  number not in have:
                s=[]
                s.append(str(doc['ticket']['ticket']['number']))
                s.append(str(doc['ticket']['ticket']['source']))
                s.append(str(doc['ticket']['ticket']['takenBy']))
                s.append(str(doc['ticket']['ticket']['sequence']))
                s.append(str(doc['ticket']['ticket']['excavator']['@type']))
                s.append(str(doc['ticket']['ticket']['excavator']['phone']))
                s.append(str(doc['ticket']['ticket']['excavator']['name']))
                s.append(str(doc['ticket']['ticket']['excavator']['address']))
                s.append(str(doc['ticket']['ticket']['excavator']['city']))
                s.append(str(doc['ticket']['ticket']['excavator']['state']))
                s.append(str(doc['ticket']['ticket']['excavator']['zip']))
                s.append(str(doc['ticket']['ticket']['excavator']['fax']))
                s.append(str(doc['ticket']['ticket']['excavator']['caller']))
                s.append(str(doc['ticket']['ticket']['excavator']['callerPhone']))
                s.append(str(doc['ticket']['ticket']['excavator']['callerPhoneExtension']))
                s.append(str(doc['ticket']['ticket']['excavator']['callerEmail']))
                s.append(str(doc['ticket']['ticket']['excavator']['contact']))
                s.append(str(doc['ticket']['ticket']['excavator']['contactPhone']))
                s.append(str(doc['ticket']['ticket']['excavator']['contactPhoneExtension']))
                s.append(str(doc['ticket']['ticket']['excavator']['contactEmail']))
                s.append(str(doc['ticket']['ticket']['excavator']['callback']))
                s.append(str(doc['ticket']['ticket']['emailCopyToCaller']))
                s.append(str(doc['ticket']['ticket']['faxCopyToCaller']))
                s.append(str(doc['ticket']['ticket']['ticketType']))
                s.append(str(doc['ticket']['ticket']['workType']))
                s.append(str(doc['ticket']['ticket']['workState']))
                s.append(str(doc['ticket']['ticket']['workCounty']))
                s.append(str(doc['ticket']['ticket']['workPlace']))
                s.append(str(doc['ticket']['ticket']['workOn']))
                s.append(str(doc['ticket']['ticket']['workStreetAddress']))
                s.append(str(doc['ticket']['ticket']['workStreetPrefix']))
                s.append(str(doc['ticket']['ticket']['workStreetName']))
                s.append(str(doc['ticket']['ticket']['workStreetType']))
                s.append(str(doc['ticket']['ticket']['workStreetSuffix']))
                s.append(str(doc['ticket']['ticket']['workIntersection']))
                s.append(str(doc['ticket']['ticket']['workSite']['callerSuppliedPoints']))
                s.append(str(doc['ticket']['ticket']['workSite']['latitude']))
                s.append(str(doc['ticket']['ticket']['workSite']['longitude']))
                s.append(str(doc['ticket']['ticket']['workSite']['secondaryLatitude']))
                s.append(str(doc['ticket']['ticket']['workSite']['secondaryLongitude']))
                s.append(str(doc['ticket']['ticket']['workSite']['viewAreaExtent']))
                s.append(str(doc['ticket']['ticket']['workSite']['workAreaExtent']))
                s.append(str(doc['ticket']['ticket']['workSite']['workArea']))
                s.append(str(doc['ticket']['ticket']['workSite']['workAreaBuffer']))
                s.append(str(doc['ticket']['ticket']['workDoneFor']))
    # This was in DOCS but not in final production tickets     s.append(str(doc['ticket']['ticket']['compliance']))
                s.append(str(doc['ticket']['ticket']['extent']))
                s.append(str(doc['ticket']['ticket']['drivingDirections']))
                s.append(str(doc['ticket']['ticket']['markingInstructions']))
                s.append(str(doc['ticket']['ticket']['remarks']))
                s.append(str(doc['ticket']['ticket']['isWhitePaint']))
                s.append(str(doc['ticket']['ticket']['isExplosives']))
                s.append(str(doc['ticket']['ticket']['isAddressInRemarks']))
                s.append(str(doc['ticket']['ticket']['isDirectionalBoring']))
                s.append(str(doc['ticket']['ticket']['isGridsCallerSupplied']))
                s.append(str(doc['ticket']['ticket']['creation']))
                cursor.execute("insert into dbo.NMOneCall(number,source,takenBy,sequence,excavator,ephone,ename,eaddress,ecity,estate,ezip,efax,ecaller,ecallerPhone,ecallerPhoneExtension,ecallerEmail,econtact,econtactPhone,econtactPhoneExtension,econtactEmail,ecallback,emailCopyToCaller,faxCopyToCaller,ticketType,workType,workState,workCounty,workPlace,workOn,workStreetAddress,workStreetPrefix,workStreetName,workStreetType,workStreetSuffix,workIntersection,callerSuppliedPoints,latitude,longitude,secondaryLatitude,secondaryLongitude,viewAreaExtent,workAreaExtent,workArea,workAreaBuffer,workDoneFor,extent,drivingDirections,markingInstructions,remarks,isWhitePaint,isExplosives,isAddressInRemarks,isDirectionalBoring,isGridCallerSupplied,creation) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",s)
                cnxn.commit()


            
            



package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "GET")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=MyServer;user id=ORG\\me;password=MyPassword;database=MyDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//probably not best idea, but it works and easy.
	var (
		number                 string
		source                 string
		takenBy                string
		sequence               string
		excavator              string
		ephone                 string
		ename                  string
		eaddress               string
		ecity                  string
		estate                 string
		ezip                   string
		efax                   string
		ecaller                string
		ecallerPhone           string
		ecallerPhoneExtension  string
		ecallerEmail           string
		econtact               string
		econtactPhone          string
		econtactPhoneExtension string
		econtactEmail          string
		ecallback              string
		emailCopyToCaller      string
		faxCopyToCaller        string
		ticketType             string
		workType               string
		workState              string
		workCounty             string
		workPlace              string
		workOn                 string
		workStreetAddress      string
		workStreetPrefix       string
		workStreetName         string
		workStreetType         string
		workStreetSuffix       string
		workIntersection       string
		callerSuppliedPoints   string
		latitude               string
		longitude              string
		secondaryLatitude      string
		secondaryLongitude     string
		viewAreaExtent         string
		workAreaExtent         string
		workArea               string
		workAreaBuffer         string
		workDoneFor            string
		extent                 string
		drivingDirections      string
		markingInstructions    string
		remarks                string
		isWhitePaint           string
		isExplosives           string
		isAddressInRemarks     string
		isDirectionalBoring    string
		isGridCallerSupplied   string
		creation               string
	)

	type Result struct {
		Number                 string `json:"number"`
		Source                 string `json:"source"`
		TakenBy                string `json:"takenby"`
		Sequence               string `json:"sequence"`
		Excavator              string `json:"excavator"`
		Ephone                 string `json:"ephone"`
		Ename                  string `json:"ename"`
		Eaddress               string `json:"eaddress"`
		Ecity                  string `json:"ecity"`
		Estate                 string `json:"estate"`
		Ezip                   string `json:"ezip"`
		Efax                   string `json:"efax"`
		Ecaller                string `json:"ecaller"`
		EcallerPhone           string `json:"ecallerphone"`
		EcallerPhoneExtension  string `json:"ecallerphoneextension"`
		EcallerEmail           string `json:"ecalleremail"`
		Econtact               string `json:"econtact"`
		EcontactPhone          string `json:"econtactphone"`
		EcontactPhoneExtension string `json:"econtactphoneextension"`
		EcontactEmail          string `json:"econtactemail"`
		Ecallback              string `json:"ecallback"`
		EmailCopyToCaller      string `json:"emailcopytocaller"`
		FaxCopyToCaller        string `json:"faxcopytocaller"`
		TicketType             string `json:"tickettype"`
		WorkType               string `json:"worktype"`
		WorkState              string `json:"workstate"`
		WorkCounty             string `json:"workcounty"`
		WorkPlace              string `json:"workplace"`
		WorkOn                 string `json:"workon"`
		WorkStreetAddress      string `json:"workstreetaddress"`
		WorkStreetPrefix       string `json:"workstreetprefix"`
		WorkStreetName         string `json:"workstreetname"`
		WorkStreetType         string `json:"workstreettype"`
		WorkStreetSuffix       string `json:"workstreetsuffix"`
		WorkIntersection       string `json:"workintersection"`
		CallerSuppliedPoints   string `json:"callersuppliedpoints"`
		Latitude               string `json:"latitude"`
		Longitude              string `json:"longitude"`
		SecondaryLatitude      string `json:"secondarylat"`
		SecondaryLongitude     string `json:"secondarylong"`
		ViewAreaExtent         string `json:"viewareaextent"`
		WorkAreaExtent         string `json:"workareaextent"`
		WorkArea               string `json:"workarea"`
		WorkAreaBuffer         string `json:"workareabuffer"`
		WorkDoneFor            string `json:"workdonefor"`
		Extent                 string `json:"extent"`
		DrivingDirections      string `json:"drivingdirections"`
		MarkingInstructions    string `json:"markinginstructions"`
		Remarks                string `json:"remarks"`
		IsWhitePaint           string `json:"iswhitepaint"`
		IsExplosives           string `json:"isexplosives"`
		IsAddressInRemarks     string `json:"isaddressinremarks"`
		IsDirectionalBoring    string `json:"isdirectionalboring"`
		IsGridCallerSupplied   string `json:"isgridcallersupplied"`
		Creation               string `json:"creation"`
	}

	rows, err := db.Query("select * from dbo.NMOneCall")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var output []Result
	for rows.Next() {
		err := rows.Scan(&number, &source, &takenBy, &sequence, &excavator, &ephone, &ename, &eaddress, &ecity, &estate, &ezip, &efax, &ecaller, &ecallerPhone, &ecallerPhoneExtension, &ecallerEmail, &econtact, &econtactPhone, &econtactPhoneExtension, &econtactEmail, &ecallback, &emailCopyToCaller, &faxCopyToCaller, &ticketType, &workType, &workState, &workCounty, &workPlace, &workOn, &workStreetAddress, &workStreetPrefix, &workStreetName, &workStreetType, &workStreetSuffix, &workIntersection, &callerSuppliedPoints, &latitude, &longitude, &secondaryLatitude, &secondaryLongitude, &viewAreaExtent, &workAreaExtent, &workArea, &workAreaBuffer, &workDoneFor, &extent, &drivingDirections, &markingInstructions, &remarks, &isWhitePaint, &isExplosives, &isAddressInRemarks, &isDirectionalBoring, &isGridCallerSupplied, &creation)
		if err != nil {
			log.Fatal(err)
		}

		d := Result{number, source, takenBy, sequence, excavator, ephone, ename, eaddress, ecity, estate, ezip, efax, ecaller, ecallerPhone, ecallerPhoneExtension, ecallerEmail, econtact, econtactPhone, econtactPhoneExtension, econtactEmail, ecallback, emailCopyToCaller, faxCopyToCaller, ticketType, workType, workState, workCounty, workPlace, workOn, workStreetAddress, workStreetPrefix, workStreetName, workStreetType, workStreetSuffix, workIntersection, callerSuppliedPoints, latitude, longitude, secondaryLatitude, secondaryLongitude, viewAreaExtent, workAreaExtent, workArea, workAreaBuffer, workDoneFor, extent, drivingDirections, markingInstructions, remarks, isWhitePaint, isExplosives, isAddressInRemarks, isDirectionalBoring, isGridCallerSupplied, creation}
		output = append(output, d)

	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		panic(err)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//code for post,delete and put can be provided upon request - or check https://github.com/PaulCrickard/Go/tree/master/RESTserver
func IndexPOST(w http.ResponseWriter, r *http.Request) {
	//set SQL query to insert.
}

func IndexDELETE(w http.ResponseWriter, r *http.Request) {
	//Set SQL delete
}

func IndexPUT(w http.ResponseWriter, r *http.Request) {

	//set SQL to update

}

const PageState = Object.freeze({"NOT_LOADED":1, "LOADING":2, "EXISTS":3, "DOES_NOT_EXIST":4, "SUBMITTED":5});

var pageState = PageState.NOT_LOADED;

var pageLoadSection = document.getElementById("page_load");
var loginSection = document.getElementById("please_log_in");
var formSection = document.getElementById("form");
var volunteerSubmittedSection = document.getElementById("volunteer_submitted");
var volunteerExistsSection = document.getElementById("volunteer_exists");
var nameBox = document.getElementById("volunteer-name");
var emailBox = document.getElementById("volunteer-email");
var submitButton = document.getElementById("submit-button");
var submitLoading = document.getElementById("submit-loading");

var dynamicDiv = document.getElementById('dynamic-data')
var dynamicBase = dynamicDiv.getAttribute('data-dynamicBase')

function signInOutHook(isSignedIn) {
	if (isSignedIn) {
		onLoad();
	} else {
		setPageState(PageState.NOT_LOADED);
	}
}

function setPageState(newPageState) {
	switch(newPageState) {
		case PageState.NOT_LOADED:
			pageLoadSection.style.display = 'none';
			loginSection.style.display = 'block';
			formSection.style.display = 'none';
			volunteerSubmittedSection.style.display = 'none';
			volunteerExistsSection.style.display = 'none';
			break;
		case PageState.LOADING:
			pageLoadSection.style.display = 'block';
			loginSection.style.display = 'none';
			formSection.style.display = 'none';
			volunteerSubmittedSection.style.display = 'none';
			volunteerExistsSection.style.display = 'none';
			break;
		case PageState.EXISTS:
			pageLoadSection.style.display = 'none';
			loginSection.style.display = 'none';
			formSection.style.display = 'none';
			volunteerSubmittedSection.style.display = 'none';
			volunteerExistsSection.style.display = 'block';
			break;
		case PageState.NOT_EXISTS:
			pageLoadSection.style.display = 'none';
			loginSection.style.display = 'none';
			formSection.style.display = 'block';
			volunteerSubmittedSection.style.display = 'none';
			volunteerExistsSection.style.display = 'none';
			break;
		case PageState.SUBMITTED:
			pageLoadSection.style.display = 'none';
			loginSection.style.display = 'none';
			formSection.style.display = 'none';
			volunteerSubmittedSection.style.display = 'block';
			volunteerExistsSection.style.display = 'none';
			break;
	}
	pageState = newPageState;
}

function onLoad(){
	if (pageState != PageState.NOT_LOADED) {
		return
	}

	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token
	if (typeof access_token === "undefined") {
		setPageState(PageState.NOT_LOADED);
		return;
	}

	setPageState(PageState.LOADING);

	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4 || req.status != 200) {
			return;
		}
		var resp 
		try {
			resp = JSON.parse(req.responseText);
		} catch(e) {
			if (req.responseText == "") {
				window.location.href = siteBase + "/error/?source_page=/volunteer&message=empty_volunteer_exists_body";
				return;
			}

			window.location.href = siteBase + "/error/?source_page=/volunteer&message="+req.responseText;
			return;
		}
		if (typeof resp.errors !== "undefined") { 
			window.location.href = siteBase + "/error/?source_page=/volunteer&message="+encodeURI(req.responseText);
			return;
		}
		if (resp.exists) {
			setPageState(PageState.EXISTS);
			return;
		}
		setPageState(PageState.NOT_EXISTS);

		return;
	}
	req.open("GET", dynamicBase + "/VolunteerExists", true);
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.send(null)
}

function submitForm() {
	if (pageState != PageState.NOT_EXISTS) {
		setPageState(pageState);
		return;
	}

	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token
	if (typeof access_token === "undefined") {
		setPageState(PageState.NOT_LOADED);
	}

	submitButton.disabled = true;
	submitLoading.style.display = 'block';

	var j = new Object();
	j.name = nameBox.value;
	j.email = emailBox.value;

	var jsonString = JSON.stringify(j);

	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4 || req.status != 200) {
			return;
		}
		var resp 
		try {
			resp = JSON.parse(req.responseText);
		} catch(e) {
			if (req.responseText == "") {
				window.location.href = siteBase + "/error/?source_page=/volunteer&message=empty_insert_volunteer_body";
				return;
			}

			window.location.href = siteBase + "/error/?source_page=/volunteer&message="+req.responseText;
			return;
		}

		submitButton.disabled = false;
		submitLoading.style.display = 'none';

		if (typeof resp.errors !== "undefined") { 
			if (Array.IsArray(resp.errors) && resp.errors.length == 1 && resp.errors[0].type == "ALREADY_EXISTS") {
				setPageState(PageState.EXISTS);
				return;
			}
			return;
		}
		setPageState(PageState.SUBMITTED)
		return;
	}
	req.open("POST", dynamicBase + "/InsertVolunteer", true)
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.setRequestHeader("Content-Type", "application/json")
	req.setRequestHeader("Accept", "application/json")
	req.send(jsonString)
}


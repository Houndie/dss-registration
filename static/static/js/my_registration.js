function parseDollar(intCost) {
	dollar = intCost.toString()
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

function isAlphaNumeric(str) {
  var code, i, len;

  for (i = 0, len = str.length; i < len; i++) {
    code = str.charCodeAt(i);
    if (!(code > 47 && code < 58) && // numeric (0-9)
        !(code > 64 && code < 91) && // upper alpha (A-Z)
        !(code > 96 && code < 123)) { // lower alpha (a-z)
      return false;
    }
  }
  return true;
};

function parseResponse(req) {
	try {
		var resp = JSON.parse(req.responseText);
		if (typeof resp.errors !== "undefined" && resp.errors.length != 0) {
			window.location.href = siteBase + "/error/?source_page=/my_registration&message="+encodeURI(responseText);
			return null;
		}

		if (req.status != 200) {
			window.location.href = siteBase + "/error/?source_page=/my_registration&message=status"+req.status;
			return null;
		}

		return resp;
	} catch(e) {
		if (req.status != 200) {
			window.location.href = siteBase + "/error/?source_page=/my_registration&message=status"+req.status;
			return null;
		}
		if (req.responseText == "") {
			window.location.href = siteBase + "/error/?source_page=/my_registration&message=empty_response_body";
			return null;
		}

		window.location.href = siteBase + "/error/?source_page=/my_registration&message="+req.responseText;
		return null;
	}
}

var current_tier;
var weekendPassSelector = document.getElementById("root_weekendPassType");
var workshopLevelBox = document.getElementById('root_workshopLevel');
var workshopLevelDiv = document.getElementById('dss-workshopLevel')
var danceOption = document.getElementById("dance_only_pass_option");
var fullWeekendOption = document.getElementById("full_weekend_pass_option");
var mixAndMatchBox = document.getElementById("root_mixAndMatch");
var mixAndMatchLabel = document.getElementById("mix_and_match_label");
var mixAndMatchRoleDiv = document.getElementById('dss-mixAndMatchRole')
var mixAndMatchRoleInput = document.getElementById('root_mixAndMatchRole')
var soloJazzBox = document.getElementById("root_soloJazz");
var soloJazzLabel = document.getElementById("solo_jazz_label");
var teamCompBox = document.getElementById("root_teamCompetition");
var teamCompLabel = document.getElementById("team_competition_label");
var teamNameDiv = document.getElementById('dss-teamName')
var teamNameInput = document.getElementById('root_teamName')
var tShirtBox = document.getElementById("root_tShirt");
var tShirtLabel = document.getElementById("tshirt_label");
var tShirtSizeDiv = document.getElementById('dss-tShirtSize')
var tShirtSizeInput = document.getElementById('root_tShirtSize')
var firstNameBox = document.getElementById('root_firstName');
var lastNameBox = document.getElementById('root_lastName');
var addressBox = document.getElementById('root_address');
var cityBox = document.getElementById('root_city');
var stateBox = document.getElementById('root_state');
var zipBox = document.getElementById('root_zip');
var emailBox = document.getElementById('root_email');
var homeSceneBox = document.getElementById('root_homeScene');
var studentBox = document.getElementById('root_homeScene');
var housingBox = document.getElementById('root_housingStatus');
var petAllergiesBox = document.getElementById('root_petAllergies');
var petAllergiesDiv = document.getElementById('dss-petAllergies');
var housingRequestDetailsBox = document.getElementById('root_housingRequestDetails');
var housingRequestDetailsDiv = document.getElementById('dss-housingRequestDetails');
var myPetsBox = document.getElementById('root_myPets');
var myPetsDiv = document.getElementById('dss-myPets');
var housingNumberBox = document.getElementById('root_housingNumber');
var housingNumberDiv = document.getElementById('dss-housingNumber');
var myHousingDetailsBox = document.getElementById('root_myHousingDetails');
var myHousingDetailsDiv = document.getElementById('dss-myHousingDetails');
var ordersDiv = document.getElementById('orders-table');
var ordersBody = document.getElementById('orders-body');

function onLoad() {
	urlparams = new URLSearchParams(window.location.search);

	var calls = 2
	var populateRes
	var myRegistrationRes
	var populateReq = new XMLHttpRequest();
	populateReq.onreadystatechange = function() {
		if (populateReq.readyState != 4) {
			return
		}
		populateRes = parseResponse(populateReq);
		if (!populateRes) {
			return
		}
		calls--;
		if (calls == 0) {
			populateForm(populateRes, myRegistrationRes)
		}
	}
	populateReq.open("GET", dynamicBase + "/PopulateForm", true)
	populateReq.send(null)
	var myRegistrationReq = new XMLHttpRequest();
	myRegistrationReq.open("GET", dynamicBase + "/GetUserRegistration?id="+urlparams.get('registration_id'), true)
	myRegistrationReq.onreadystatechange = function() {
		if (myRegistrationReq.readyState != 4) {
			return
		}
		myRegistrationRes = parseResponse(myRegistrationReq);
		if (!myRegistrationRes) {
			return
		}
		calls--;
		if (calls == 0) {
			populateForm(populateRes, myRegistrationRes)
		}
	}
	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token;
	myRegistrationReq.setRequestHeader("Authorization", "Bearer "+access_token)
	myRegistrationReq.setRequestHeader("Accept", "application/json")
	myRegistrationReq.send(null)
}

function populateForm(populateRes, myRegistrationRes) {
	switch (myRegistrationRes.registration.weekend_pass_type) {
	case "Full":
		fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + myRegistrationRes.registration.full_weekend.tier + ")";
		workshopLevelBox.value = myRegistrationRes.registration.full_weekend.level;
		workshopLevelBox.disabled = true;
	case "Dance":
		weekendPassSelector.value = myRegistrationRes.registration.weekend_pass_type;
		weekendPassShowHide();
		break;
	default:
		danceOption.innerHTML = "Dance Only Pass (" + parseDollar(populateRes.dance_pass_cost) + ")";
		fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + populateRes.weekend_pass_tier + " - " + parseDollar(populateRes.weekend_pass_cost) + ")";
		current_tier = populateRes.weekend_pass_tier;
		weekendPassSelector.disabled = false;
		studentBox.disabled = false;
	}

	if (myRegistrationRes.registration.mix_and_match) {
		mixAndMatchBox.checked = true;
		mixAndMatchRoleDiv.value = myRegistrationRes.registration.mix_and_match_role;
		mixAndMatchShowHide();
	} else {
		mixAndMatchBox.disabled = false;
		mixAndMatchRoleDiv.disabled = false;
		mixAndMatchLabel.innerHTML = "Mix And Match Competition (" + parseDollar(populateRes.mix_and_match_cost) + ")";
	}
	if (myRegistrationRes.registration.solo_jazz) {
		soloJazzBox.checked = true;
	} else {
		soloJazzBox.disabled = false;
		soloJazzLabel.innerHTML = "Solo Jazz Competition (" + parseDollar(populateRes.solo_jazz_cost) + ")";
	}
	if (myRegistrationRes.registration.team_competition) {
		teamCompBox.checked = true;
		teamNameInput.value = myRegistrationRes.registration.team_name;
		teamShowHide();
	} else {
		teamCompBox.disabled = false;
		teamNameInput.disabled = false;
		teamCompLabel.innerHTML = "Team Competition (" + parseDollar(populateRes.team_comp_cost) + ")";
	}
	if (myRegistrationRes.registration.tshirt) {
		tShirtBox.checked = true;
		tShirtSizeInput.value = myRegistrationRes.registration.tshirt_size;
		tShirtShowHide();
	} else {
		tShirtBox.disabled = false;
		tShirtSizeInput.disabled = false;
		tShirtLabel.innerHTML = "T-Shirt (" + parseDollar(populateRes.tshirt_cost) + ")";
	}

	firstNameBox.value = myRegistrationRes.registration.first_name;
	firstNameBox.disabled = false;
	lastNameBox.value = myRegistrationRes.registration.last_name;
	lastNameBox.disabled = false;
	addressBox.value = myRegistrationRes.registration.address;
	addressBox.disabled = false;
	cityBox.value = myRegistrationRes.registration.city;
	cityBox.disabled = false;
	stateBox.value = myRegistrationRes.registration.state;
	stateBox.disabled = false;
	zipBox.value = myRegistrationRes.registration.zip;
	zipBox.disabled = false;
	emailBox.value = myRegistrationRes.registration.email;
	emailBox.disabled = false;
	homeSceneBox.value = myRegistrationRes.registration.home_scene;
	homeSceneBox.disabled = false;
	studentBox.checked = myRegistrationRes.registration.student;
	studentBox.disabled = false;

	housingBox.value = myRegistrationRes.registration.housing_status;
	housingBox.disabled = false;
	switch (myRegistrationRes.registration.housing_status) {
		case "Require":
			petAllergiesBox.value = myRegistrationRes.registration.require_housing.pet_allergies;
			housingRequestDetailsBox.value = myRegistrationRes.registration.require_housing.housing_request_details;
			break;	
		case "Provide":
			myPetsBox.value = myRegistrationRes.registration.provide_housing.my_pets;
			housingNumberBox.value = myRegistrationRes.registration.provide_housing.housing_number;
			myHousingDetailsBox.value = myRegistrationRes.registration.provide_housing.my_housing_details;
			break;
	}
	housingShowHide();

	if (typeof myRegistrationRes.registration.orders !== 'undefined') {
		var sorted_orders = myRegistrationRes.registration.orders;
		sorted_orders.sort(function(a,b){
			var date1 = a.created_at;
			var date2 = b.created_at;
			return date1 > date2;
		});
		for (var i = 0; i < sorted_orders.length; i++) {
			var row = ordersBody.insertRow(i);
			row.insertCell(0).textContent = sorted_orders[i].items.join();
			row.insertCell(1).textContent = parseDollar(sorted_orders[i].cost);
			if (sorted_orders[i].paid) {
				row.insertCell(2).textContent = '\u2705';
			} else {
				// Be double extra sure of no html injection.  Should be fine because https, but hey let's not mess around here.
				if (!isAlphaNumeric(sorted_orders[i].id)) {
					window.location.href = siteBase + "/error/?source_page=/my_registration&message=id_not_alphanumeric"
					return
				}
				row.insertCell(2).innerHTML = '<button class="btn btn-info" onclick="payOrder(\''+sorted_orders[i].id+'\')">Pay Now</button>';
			}
		}
		ordersDiv.style.display = 'block';
	}

	document.getElementById('populate-loading').style.display = 'none';
}

function payOrder(orderId) {
	alert(orderId);
}

function weekendPassShowHide() {
	switch (weekendPassSelector.value) {
		case "Dance":
		case "None":
			workshopLevelDiv.style.display = 'none';
			workshopLevelBox.required = false;
			break;
		default:
			workshopLevelDiv.style.display = 'block';
			workshopLevelBox.required = true;
			break;
	}
}

function mixAndMatchShowHide() {
	if (mixAndMatchBox.checked) {
		mixAndMatchRoleDiv.style.display = 'block';
		mixAndMatchRoleInput.required = true;
	} else {
		mixAndMatchRoleDiv.style.display = 'none';
		mixAndMatchRoleInput.required = false;
	}
}

function teamShowHide() {
	if (teamCompBox.checked) {
		teamNameDiv.style.display = 'block';
		teamNameInput.required = true;
	} else {
		teamNameDiv.style.display = 'none';
		teamNameInput.required = false;
	}
}

function tShirtShowHide() {
	if (tShirtBox.checked) {
		tShirtSizeDiv.style.display = 'block';
		tShirtSizeInput.required = true
	} else {
		tShirtSizeDiv.style.display = 'none';
		tShirtSizeInput.required = false
	}
}

function housingShowHide() {
	switch (housingBox.value) {
		case "None":
			myPetsDiv.style.display = 'none';
			housingNumberDiv.style.display = 'none';
			housingNumberBox.required = false
			myHousingDetailsDiv.style.display = 'none';
			petAllergiesDiv.style.display = 'none';
			housingRequestDetailsDiv.style.display = 'none';
			break;
		case "Require":
			myPetsDiv.style.display = 'none';
			housingNumberDiv.style.display = 'none';
			housingNumberBox.required = false
			myHousingDetailsDiv.style.display = 'none';
			petAllergiesDiv.style.display = 'block';
			housingRequestDetailsDiv.style.display = 'block';
			break;
		default:
			myPetsDiv.style.display = 'block';
			housingNumberDiv.style.display = 'block';
			housingNumberBox.required = true
			myHousingDetailsDiv.style.display = 'block';
			petAllergiesDiv.style.display = 'none';
			housingRequestDetailsDiv.style.display = 'none';
			break;
	}
}

function submitRegistration() {
	document.getElementById('submit-button').disabled = true;
	document.getElementById('submit-loading').style.display = 'block';
	var j = new Object();
	j.id = urlparams.get('registration_id')
	j.first_name = firstNameBox.value;
	j.last_name = lastNameBox.value;
	j.address = addressBox.value;
	j.city = cityBox.value;
	j.state = stateBox.value;
	j.zip = zipBox.value;
	j.email = emailBox.value;
	j.home_scene = homeSceneBox.value;
	j.student = studentBox.checked;

	j.weekend_pass_type = weekendPassSelector.value;
	if (j.weekend_pass_type == 'Full') {
		j.full_weekend = new Object();
		j.full_weekend.level = workshopLevelBox.value;
		j.full_weekend.tier = current_tier;
	}

	j.mix_and_match = mixAndMatchBox.checked;
	if (j.mix_and_match) {
		j.mix_and_match_role = mixAndMatchRoleInput.value;
	}

	j.solo_jazz = soloJazzBox.checked;

	j.team_competition = teamCompBox.checked;
	if (j.team_competition) {
		j.team_name = teamNameInput.value;
	}

	j.tshirt = tShirtBox.checked;
	if (j.tshirt) {
		j.tshirt_size = tShirtSizeInput.value;
	}

	j.housing_status = housingBox.value;
	if (j.housing_status == "Require") {
		j.require_housing = new Object();
		j.require_housing.pet_allergies = petAllergiesBox.value;
		j.require_housing.housing_request_details = housingRequestDetailsBox.value;
	} else if (j.housing_status == "Provide") {
		j.provide_housing = new Object();
		j.provide_housing.my_pets = myPetsBox.value;
		j.provide_housing.housing_number = housingNumberBox.value;
		j.provide_housing.my_housing_details = myHousingDetailsBox.value;
	}
	j.redirect_url = siteBase+"/registration-complete"

	var jsonString = JSON.stringify(j);
	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4) {
			return
		}
		res = parseResponse(req);
		if (!res) {
			return
		}
		window.location.href = res.checkout_url;
	}
	req.open("POST", dynamicBase + "/UpdateRegistration", true)
	req.setRequestHeader("Content-Type", "application/json")
	req.setRequestHeader("Accept", "application/json")
	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.send(jsonString)
}

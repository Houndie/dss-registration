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

var existingCodes = [];
var current_tier;
var danceOption = document.getElementById("dance_only_pass_option");
var fullWeekendOption = document.getElementById("full_weekend_pass_option");
var mixAndMatchLabel = document.getElementById("mix_and_match_label");
var soloJazzLabel = document.getElementById("solo_jazz_label");
var teamCompLabel = document.getElementById("team_competition_label");
var tShirtLabel = document.getElementById("tshirt_label");
var firstNameBox = document.getElementById('root_firstName');
var lastNameBox = document.getElementById('root_lastName');
var addressBox = document.getElementById('root_address');
var cityBox = document.getElementById('root_city');
var stateBox = document.getElementById('root_state');
var zipBox = document.getElementById('root_zip');
var emailBox = document.getElementById('root_email');
var homeSceneBox = document.getElementById('root_homeScene');
var petAllergiesBox = document.getElementById('root_petAllergies');
var housingRequestDetailsBox = document.getElementById('root_housingRequestDetails');
var myPetsBox = document.getElementById('root_myPets');
var myHousingDetailsBox = document.getElementById('root_myHousingDetails');
var ordersDiv = document.getElementById('orders-div');
var ordersList = document.getElementById('orders-list');
var ordersCost = document.getElementById('orders-cost');
var submitButton = document.getElementById('submit-button');
var updatedRegistrationAlert = document.getElementById('updated-registration-alert');
var submitAlert = document.getElementById('submit-alert');
var submitLoading = document.getElementById('submit-loading');

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
		current_tier = myRegistrationRes.registration.full_weekend.tier;
		workshopLevelBox.disabled = true;
	case "Dance":
		weekendPassSelector.value = myRegistrationRes.registration.weekend_pass_type;
		weekendPassShowHide();
		break;
	default:
		danceOption.innerHTML = "Dance Only Pass (" + parseDollar(populateRes.dance_pass_cost) + ")";
		dancePassCost = populateRes.dance_pass_cost;
		fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + populateRes.weekend_pass_tier + " - " + parseDollar(populateRes.weekend_pass_cost) + ")";
		fullWeekendCost = populateRes.weekend_pass_cost
		current_tier = populateRes.weekend_pass_tier;
		weekendPassSelector.disabled = false;
		studentBox.disabled = false;
		workshopLevelBox.disabled = false;
		studentDiscount = popuateRes.student_discount;
	}


	if (myRegistrationRes.registration.mix_and_match) {
		mixAndMatchBox.checked = true;
		mixAndMatchRoleDiv.value = myRegistrationRes.registration.mix_and_match_role;
		mixAndMatchShowHide();
	} else {
		mixAndMatchBox.disabled = false;
		mixAndMatchRoleInput.disabled = false;
		mixAndMatchLabel.innerHTML = "Mix And Match Competition (" + parseDollar(populateRes.mix_and_match_cost) + ")";
		mixAndMatchCost = populateRes.mix_and_match_cost;
	}
	if (myRegistrationRes.registration.solo_jazz) {
		soloJazzBox.checked = true;
	} else {
		soloJazzBox.disabled = false;
		soloJazzLabel.innerHTML = "Solo Jazz Competition (" + parseDollar(populateRes.solo_jazz_cost) + ")";
		soloJazzCost = populateRes.solo_jazz_cost;
	}
	if (myRegistrationRes.registration.team_competition) {
		teamCompBox.checked = true;
		teamNameInput.value = myRegistrationRes.registration.team_name;
		teamShowHide();
	} else {
		teamCompBox.disabled = false;
		teamNameInput.disabled = false;
		teamCompLabel.innerHTML = "Team Competition (" + parseDollar(populateRes.team_comp_cost) + ")";
		teamCompCost = populateRes.team_comp_cost;
	}
	if (myRegistrationRes.registration.tshirt) {
		tShirtBox.checked = true;
		tShirtSizeInput.value = myRegistrationRes.registration.tshirt_size;
		tShirtShowHide();
	} else {
		tShirtBox.disabled = false;
		tShirtSizeInput.disabled = false;
		tShirtLabel.innerHTML = "T-Shirt (" + parseDollar(populateRes.tshirt_cost) + ")";
		tshirtCost = populateRes.tshirt_cost;
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

	if (typeof myRegistrationRes.registration.unpaid_items !== 'undefined') {
		ordersList.innerHTML = "";
		for (var i = 0; i < myRegistrationRes.registration.unpaid_items.items.length; i++) {
			var text = document.createTextNode(myRegistrationRes.registration.unpaid_items.items[i]);
			var li = document.createElement("LI");
			li.appendChild(text);
			ordersList.appendChild(li);
		}
		var text = document.createTextNode(parseDollar(myRegistrationRes.registration.unpaid_items.cost));
		ordersCost.innerHTML = "";
		ordersCost.appendChild(text);
		submitButton.value = "Update & Pay";
		ordersDiv.style.display = 'block';
	}

	if (typeof myRegistrationRes.registration.discounts !== "undefined") {
		for (var i = 0; i < myRegistrationRes.registration.discounts.length; i++) {
			addDiscount(myRegistrationRes.registration.discounts[i].code, myRegistrationRes.registration.discounts[i].discounts)
			existingCodes.push(myRegistrationRes.registration.discounts[i].code);
		}
	}

	recalculateTotal();
	document.getElementById('populate-loading').style.display = 'none';
	if (myRegistrationRes.registration.updated_tier) {
		updatedRegistrationAlert.textContent = 'Your unpaid full weekend pass tier has sold out, and your registration has been updated to the next tier';
		updatedRegistrationAlert.style.display = 'block';
	}
}

function submitRegistration() {
	submitButton.disabled = true;
	submitLoading.style.display = 'block';
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

	newCodes = [];
	for (var i = 0; i < discountCodes.length; i++) {
		if (!existingCodes.includes(discountCodes[i])) {
			newCodes.push(discountCodes[i])
		}
	}
	if (discountCodes.length > 0) {
		j.discount_codes = newCodes;
	}

	var jsonString = JSON.stringify(j);
	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4) {
			return
		}
		var registrationRes
		try {
			registrationRes = JSON.parse(req.responseText);
			if (typeof registrationRes.errors !== "undefined" && registrationRes.errors.length != 0 && (registrationRes.errors.length != 1 || registrationRes.errors[0].type != "OUT_OF_STOCK")) {
				window.location.href = siteBase + "/error/?source_page=/my_registration&message="+encodeURI(registrationResonseText);
				return;
			}

			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/my_registration&message=status"+req.status;
				return;
			}
		} catch(e) {
			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/my_registration&message=status"+req.status;
				return;
			}
			if (req.registrationResonseText == "") {
				window.location.href = siteBase + "/error/?source_page=/my_registration&message=empty_registrationResonse_body";
				return;
			}

			window.location.href = siteBase + "/error/?source_page=/my_registration&message="+req.responseText;
			return;
		}
		if (typeof registrationRes.errors !== "undefined") {
			// Awkward hack, but the fastest way to code the updated existing cost is to just reload the page start
			onLoad()
			//current_tier = registrationRes.errors[0].out_of_stock_details.next_tier;
			//fullWeekendCost = registrationRes.errors[0].out_of_stock_details.next_cost;
			//fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + current_tier + " - " + parseDollar(fullWeekendCost) + ")"
			submitAlert.textContent = 'Unfortunately, that tier is now sold out.  Your registration has been updated to the next available tier, submit your registration again to confirm.';
			submitAlert.style.display = 'block';
			submitButton.disabled = false;
			submitLoading.style.display = 'none';
		} else {
			window.location.href = registrationRes.checkout_url;
		}
	}
	req.open("POST", dynamicBase + "/UpdateRegistration", true)
	req.setRequestHeader("Content-Type", "application/json")
	req.setRequestHeader("Accept", "application/json")
	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.send(jsonString)
}

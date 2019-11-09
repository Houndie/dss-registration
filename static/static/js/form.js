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

var current_tier

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
var studentBox = document.getElementById('root_homeScene');
var petAllergiesBox = document.getElementById('root_petAllergies');
var housingRequestDetailsBox = document.getElementById('root_housingRequestDetails');
var myPetsBox = document.getElementById('root_myPets');
var myHousingDetailsBox = document.getElementById('root_myHousingDetails');
var submitButton = document.getElementById('submit-button');
var submitLoading = document.getElementById('submit-loading');

function onLoad() {
	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState == 4 && req.status == 200) {
			var resp = JSON.parse(req.responseText)
			danceOption.innerHTML = "Dance Only Pass (" + parseDollar(resp.dance_pass_cost) + ")"
			dancePassCost = resp.dance_pass_cost;
			fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + resp.weekend_pass_tier + " - " + parseDollar(resp.weekend_pass_cost) + ")"
			fullWeekendCost = resp.weekend_pass_cost
			current_tier = resp.weekend_pass_tier;
			mixAndMatchLabel.innerHTML = "Mix And Match Competition (" + parseDollar(resp.mix_and_match_cost) + ")"
			mixAndMatchCost = resp.mix_and_match_cost;
			soloJazzLabel.innerHTML = "Solo Jazz Competition (" + parseDollar(resp.solo_jazz_cost) + ")"
			soloJazzCost = resp.solo_jazz_cost;
			teamCompLabel.innerHTML = "Team Competition (" + parseDollar(resp.team_comp_cost) + ")"
			teamCompCost = resp.team_comp_cost;
			tShirtLabel.innerHTML = "T-Shirt (" + parseDollar(resp.tshirt_cost) + ")"
			resp.tshirtCost = resp.tshirt_cost;
			recalculateTotal();
			document.getElementById("populate-loading").style.display='none';
		}
	}
	req.open("GET", dynamicBase + "/PopulateForm", true)
	req.send(null)
}

function submitRegistration() {
	submitButton.disabled = true;
	submitLoading.style.display = 'block';
	var j = new Object();
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
		j.provide_housing.housing_number = parseInt(housingNumberBox.value, 10);
		j.provide_housing.my_housing_details = myHousingDetailsBox.value;
	}
	j.redirect_url = siteBase+"/registration-complete"
	if (discountCodes.length > 0) {
		j.discount_codes = discountCodes;
	}

	var jsonString = JSON.stringify(j);

	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4) {
			return;
		}

		var registrationRes = parseResponse(req)
		if (!registrationRes) {
			return
		}
		window.location.href = registrationRes.checkout_url;
	}

	req.open("POST", dynamicBase + "/AddRegistration", true)
	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token
	if (typeof access_token !== "undefined") {
		req.setRequestHeader("Authorization", "Bearer "+access_token)
	}
	req.setRequestHeader("Content-Type", "application/json")
	req.setRequestHeader("Accept", "application/json")
	req.send(jsonString)
}

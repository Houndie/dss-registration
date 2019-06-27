function parseDollar(intCost) {
	dollar = intCost.toString()
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

var dynamicBase
var siteBase
var current_tier

function onLoad(d, s) {
	dynamicBase = d
	siteBase = s
	var danceOption = document.getElementById("dance_only_pass_option")
	var fullWeekendOption = document.getElementById("full_weekend_pass_option")
	var mixAndMatch = document.getElementById("mix_and_match_label")
	var soloJazz = document.getElementById("solo_jazz_label")
	var teamComp = document.getElementById("team_competition_label")
	var tShirt = document.getElementById("tshirt_label")
	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState == 4 && req.status == 200) {
			var resp = JSON.parse(req.responseText)
			danceOption.innerHTML = "Dance Only Pass (" + parseDollar(resp.dance_pass_cost) + ")"
			fullWeekendOption.innerHTML = "Full Weekend Pass (Tier " + resp.weekend_pass_tier + " - " + parseDollar(resp.weekend_pass_cost) + ")"
			mixAndMatch.innerHTML = "Mix And Match Competition (" + parseDollar(resp.mix_and_match_cost) + ")"
			soloJazz.innerHTML = "Solo Jazz Competition (" + parseDollar(resp.solo_jazz_cost) + ")"
			teamComp.innerHTML = "Team Competition (" + parseDollar(resp.team_comp_cost) + ")"
			tShirt.innerHTML = "T-Shirt (" + parseDollar(resp.tshirt_cost) + ")"
			current_tier = resp.weekend_pass_tier
		}
	}
	req.open("GET", dynamicBase + "/PopulateForm", true)
	req.send(null)
}

function weekendPassShowHide() {
	var levelDiv = document.getElementById('dss-workshopLevel')
	var levelInput = document.getElementById('root_workshopLevel')
	switch (document.getElementById('root_weekendPassType').value) {
		case "Dance":
		case "":
			levelDiv.style.display = 'none';
			levelInput.required = false
			break;
		default:
			levelDiv.style.display = 'block';
			levelInput.required = true
			break;
	}
}

function mixAndMatchShowHide() {
	var mixAndMatchRole = document.getElementById('dss-mixAndMatchRole')
	var mixAndMatchRoleInput = document.getElementById('root_mixAndMatchRole')
	if (document.getElementById('root_mixAndMatch').checked) {
		mixAndMatchRole.style.display = 'block';
		mixAndMatchRoleInput.required = true
	} else {
		mixAndMatchRole.style.display = 'none';
		mixAndMatchRoleInput.required = false
	}
}

function teamShowHide() {
	var teamName = document.getElementById('dss-teamName')
	var teamNameInput = document.getElementById('root_teamName')
	if (document.getElementById('root_teamCompetition').checked) {
		teamName.style.display = 'block';
		teamNameInput.required = true
	} else {
		teamName.style.display = 'none';
		teamNameInput.required = false
	}
}

function tShirtShowHide() {
	var tShirtSize = document.getElementById('dss-tShirtSize')
	var tShirtSizeInput = document.getElementById('root_tShirtSize')
	if (document.getElementById('root_tShirt').checked) {
		tShirtSize.style.display = 'block';
		tShirtSizeInput.required = true
	} else {
		tShirtSize.style.display = 'none';
		tShirtSizeInput.required = false
	}
}

function housingShowHide() {
	var myPets = document.getElementById('dss-myPets');
	var housingNumber = document.getElementById('dss-housingNumber');
	var housingNumberInput = document.getElementById('root_housingNumber');
	var myHousingDetails = document.getElementById('dss-myHousingDetails');
	var petAllergies = document.getElementById('dss-petAllergies')
	var housingRequestDetails = document.getElementById('dss-housingRequestDetails')
	switch (document.getElementById('root_housingStatus').value) {
		case "None":
			myPets.style.display = 'none';
			housingNumber.style.display = 'none';
			housingNumberInput.required = false
			myHousingDetails.style.display = 'none';
			petAllergies.style.display = 'none';
			housingRequestDetails.style.display = 'none';
			break;
		case "Require":
			myPets.style.display = 'none';
			housingNumber.style.display = 'none';
			housingNumberInput.required = false
			myHousingDetails.style.display = 'none';
			petAllergies.style.display = 'block';
			housingRequestDetails.style.display = 'block';
			break;
		default:
			myPets.style.display = 'block';
			housingNumber.style.display = 'block';
			housingNumber.required = true
			myHousingDetails.style.display = 'block';
			petAllergies.style.display = 'none';
			housingRequestDetails.style.display = 'none';
			break;
	}
}

function submitRegistration() {
	var j = new Object();
	j.first_name = document.getElementById('root_firstName').value;
	j.last_name = document.getElementById('root_lastName').value;
	j.address = document.getElementById('root_address').value;
	j.city = document.getElementById('root_city').value;
	j.state = document.getElementById('root_state').value;
	j.zip = document.getElementById('root_zip').value;
	j.email = document.getElementById('root_email').value;
	j.home_scene = document.getElementById('root_homeScene').value;
	j.student = document.getElementById('root_homeScene').checked;

	j.weekend_pass_type = document.getElementById('root_weekendPassType').value;
	if (j.weekend_pass_type == 'Full') {
		j.full_weekend = new Object();
		j.full_weekend.level = document.getElementById('root_workshopLevel').value;
		j.full_weekend.tier = current_tier;
	}

	j.mix_and_match = document.getElementById('root_mixAndMatch').checked;
	if (j.mix_and_match) {
		j.mix_and_match_role = document.getElementById('root_mixAndMatchRole').value;
	}

	j.solo_jazz = document.getElementById('root_soloJazz').checked;

	j.team_competition = document.getElementById('root_teamCompetition').checked;
	if (j.team_competition) {
		j.team_name = document.getElementById('root_teamName').value;
	}

	j.tshirt = document.getElementById('root_tShirt').checked;
	if (j.tshirt) {
		j.tshirt_size = document.getElementById('root_tShirtSize').value;
	}

	j.housing_status = document.getElementById('root_housingStatus').value;
	if (j.housing_status == "Require") {
		j.require_housing = new Object();
		j.require_housing.pet_allergies = document.getElementById('root_petAllergies').value;
		j.require_housing.housing_request_details = document.getElementById('root_housingRequestDetails').value;
	} else if (j.housing_status == "Provide") {
		j.provide_housing = new Object();
		j.provide_housing.my_pets = document.getElementById('root_myPets').value;
		j.provide_housing.housing_number = document.getElementById('root_housingNumber').value;
		j.provide_housing.my_housing_details = document.getElementById('root_myHousingDetails').value;
	}

	var jsonString = JSON.stringify(j);

	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4) {
			return;
		}

		try {
			var resp = JSON.parse(req.responseText);
			if (typeof resp.errors !== "undefined" && resp.errors.length != 0) {
				window.location.href = siteBase + "error/?source_page=/registration&error="+encodeURI(req.responseText)
				return;
			}

			if (req.status != 200) {
				window.location.href = siteBase + "error/?source_page=/registration&error=status"+req.status
			}

			window.location.href = resp.checkout_url;
		} catch(e) {
			window.location.href = siteBase + "error/?source_page=/registration&error="+e.name+"%3A%20"+e.message
		}
	}

	req.open("POST", dynamicBase + "/AddRegistration", true)
	req.setRequestHeader("Content-Type", "application/json")
	req.setRequestHeader("Accept", "application/json")
	req.send(jsonString)
}

function parseDollar(intCost) {
	dollar = intCost.toString()
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

function onLoad(dynamicBase) {
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

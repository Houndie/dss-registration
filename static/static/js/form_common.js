var discountCodes = [];

var fullWeekendCost = 0;
var dancePassCost = 0;
var soloJazzCost = 0;
var mixAndMatchCost = 0;
var teamCompCost = 0;
var tshirtCost = 0;
var studentDiscount = 0;

var fullWeekendPercentDiscounts = [];
var fullWeekendDollarDiscounts = [];
var dancePassPercentDiscounts = [];
var dancePassDollarDiscounts = [];
var soloJazzPercentDiscounts = [];
var soloJazzDollarDiscounts = [];
var mixAndMatchPercentDiscounts = [];
var mixAndMatchDollarDiscounts = [];
var teamCompPercentDiscounts = [];
var teamCompDollarDiscounts = [];
var tshirtPercentDiscounts = [];
var tshirtDollarDiscounts = [];

var studentBox = document.getElementById("root_student");
var weekendPassSelector = document.getElementById("root_weekendPassType");
var workshopLevelBox = document.getElementById('root_workshopLevel');
var workshopLevelDiv = document.getElementById('dss-workshopLevel')
var mixAndMatchBox = document.getElementById("root_mixAndMatch");
var mixAndMatchRoleDiv = document.getElementById('dss-mixAndMatchRole')
var mixAndMatchRoleInput = document.getElementById('root_mixAndMatchRole')
var soloJazzBox = document.getElementById("root_soloJazz");
var teamCompBox = document.getElementById("root_teamCompetition");
var teamNameDiv = document.getElementById('dss-teamName')
var teamNameInput = document.getElementById('root_teamName')
var tShirtBox = document.getElementById("root_tShirt");
var tShirtSizeDiv = document.getElementById('dss-tShirtSize')
var tShirtSizeInput = document.getElementById('root_tShirtSize')
var housingBox = document.getElementById('root_housingStatus');
var petAllergiesDiv = document.getElementById('dss-petAllergies');
var housingRequestDetailsDiv = document.getElementById('dss-housingRequestDetails');
var housingNumberBox = document.getElementById('root_housingNumber');
var housingNumberDiv = document.getElementById('dss-housingNumber');
var myPetsDiv = document.getElementById('dss-myPets');
var myHousingDetailsDiv = document.getElementById('dss-myHousingDetails');
var couponList = document.getElementById('coupon-list');
var couponBox = document.getElementById('root_coupon');
var couponAlert = document.getElementById('coupon-alert');
var totalCount = document.getElementById('totalCount');

function parseDollar(intCost) {
	dollar = intCost.toString()
	while(dollar.length < 3) {
		dollar = "0" + dollar;
	}
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

function weekendPassShowHide() {
	switch (weekendPassSelector.value) {
		case "Dance":
		case "None":
			workshopLevelDiv.style.display = 'none';
			workshopLevelBox.required = false
			break;
		default:
			workshopLevelDiv.style.display = 'block';
			workshopLevelBox.required = true
			break;
	}
	recalculateTotal();
}

function mixAndMatchShowHide() {
	if (mixAndMatchBox.checked) {
		mixAndMatchRoleDiv.style.display = 'block';
		mixAndMatchRoleInput.required = true
	} else {
		mixAndMatchRoleDiv.style.display = 'none';
		mixAndMatchRoleInput.required = false
	}
	recalculateTotal();
}

function teamShowHide() {
	if (teamCompBox.checked) {
		teamNameDiv.style.display = 'block';
		teamNameInput.required = true
	} else {
		teamNameDiv.style.display = 'none';
		teamNameInput.required = false
	}
	recalculateTotal();
}

function tShirtShowHide() {
	if (tShirtBox.checked) {
		tShirtSizeDiv.style.display = 'block';
		tShirtSizeInput.required = true
	} else {
		tShirtSizeDiv.style.display = 'none';
		tShirtSizeInput.required = false
	}
	recalculateTotal();
}

function soloJazzShowHide() {
	recalculateTotal();
}

function studentDiscountShowHide() {
	recalculateTotal();
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

function addDiscount(code, d) {
	// Sanatize "code"
	tmp = document.createElement('DIV');
	tmp.textContent = code;
	sanCode = tmp.innerHTML;

	var listId = 'discounts-list-'+sanCode
	if (document.getElementById(listId)) {
		couponAlert.style.display = 'block';
		couponAlert.textContent = 'coupon already applied';
		return null; 
	}

	var newListItem = document.createElement('LI');
	newListItem.classList.add('list-group-item');
	newListItem.classList.add('d-flex');
	newListItem.classList.add('justify-content-between');
	newListItem.id = listId;

	var items = ""
	for (var i = 0; i < d.length; i++) {
		var thisDiscount = d[i];

		var itemDiscount = new Object();
		itemDiscount.code = code;
		
		items += '<p class="mb-1">'+thisDiscount.applied_to + ': '
		switch(thisDiscount.type) {
			case 'percent':
				items += thisDiscount.percent + '%'
				var p = parseFloat(thisDiscount.percent);
				p = 1 - (p/100.0);
				itemDiscount.percent = p;

				switch (thisDiscount.applied_to) {
					case "Full Weekend":
						fullWeekendPercentDiscounts.push(itemDiscount);
						break;
					case "Dance Only":
						dancePassPercentDiscounts.push(itemDiscount);
						break;
					case "Mix And Match":
						mixAndMatchPercentDiscounts.push(itemDiscount);
						break;
					case "Solo Jazz":
						soloJazzPercentDiscounts.push(itemDiscount);
						break;
					case "Team Competition":
						teamCompPercentDiscounts.push(itemDiscount);
						break;
					case "TShirt":
						tshirtPercentDiscounts.push(itemDiscount);
						break;
				}
				break;
			case 'dollar':
				items += parseDollar(thisDiscount.dollar)
				itemDiscount.dollar = thisDiscount.dollar
				switch (thisDiscount.applied_to) {
					case "Full Weekend":
						fullWeekendDollarDiscounts.push(itemDiscount);
						break;
					case "Dance Only":
						dancePassDollarDiscounts.push(itemDiscount);
						break;
					case "Mix And Match":
						mixAndMatchDollarDiscounts.push(itemDiscount);
						break;
					case "Solo Jazz":
						soloJazzDollarDiscounts.push(itemDiscount);
						break;
					case "Team Competition":
						teamCompDollarDiscounts.push(itemDiscount);
						break;
					case "TShirt":
						tshirtDollarDiscounts.push(itemDiscount);
						break;
				}
				break;
		}
		items += ' off</p>'
	}
	newListItem.innerHTML = '<div class="d-flex flex-column"><h5 class="mb-1">Code: "'+sanCode+'"</h5>' + items + '</div>';

	couponList.appendChild(newListItem);
	return newListItem;
}

function submitDiscount() {
	var code = couponBox.value;

	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState != 4) {
			return;
		}

		var discountRes
		try {
			discountRes = JSON.parse(req.responseText);
			if (typeof discountRes.errors !== "undefined" && discountRes.errors.length != 0) {
				if (discountRes.errors.length == 1 && 
					discountRes.errors[0].type == "BAD_PARAMETER" && 
					discountRes.errors[0].bad_parameter_details.parameter_name == "code" &&
					discountRes.errors[0].bad_parameter_details.reason == "discount with this name does not exist") {
					couponAlert.style.display = 'block';
					couponAlert.textContent = 'coupon code "' + code + '" is invalid';
					return;
				}
				window.location.href = siteBase + "/error/?source_page=/registration&message="+encodeURI(req.responseText);
				return;
			}

			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/registration&message=status"+req.status;
				return;
			}
		} catch(e) {
			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/registration&message=status"+req.status;
				return;
			}
			if (req.responseText == "") {
				window.location.href = siteBase + "/error/?source_page=/registration&message=empty_response_body";
				return;
			}

			window.location.href = siteBase + "/error/?source_page=/registration&message="+req.responseText;
		}

		if (typeof discountRes.discounts === "undefined" || discountRes.discounts.length == 0) {
			couponAlert.style.display = 'block';
			couponAlert.textContent = 'coupon code "' + code + '" is invalid';
			return;
		}
		discountCodes.push(code)
		couponAlert.style.display = 'none';

		var newListItem = addDiscount(code, discountRes.discounts)
		if (!newListItem) {
			return;
		}

		var closeButton = document.createElement('BUTTON');
		closeButton.innerHTML = '<span aria-hidden="true">&times;</span>'
		closeButton.classList.add('close');
		closeButton.setAttribute('aria-label', 'close');
		closeButton.type = 'button';
		closeButton.onclick = function() {
			closeDiscount(code);
		}

		newListItem.appendChild(closeButton);

		couponBox.value = "";
		recalculateTotal();
	}
	req.open("GET", dynamicBase+"/GetDiscount?code="+code, true)
	req.setRequestHeader("Accept", "appliction/json")
	req.send(null)
}

function removeDiscountFromList(code, discountList) {
	for (var i = 0; i < discountList.length; i++) {
		if (discountList[i].code == code) {
			discountList.splice(i, 1);
		}
	}
}

function closeDiscount(code) {
	// Sanatize "code"
	tmp = document.createElement('DIV');
	tmp.textContent = code;
	sanCode = tmp.innerHTML;

	var listId = 'discounts-list-'+sanCode
	var listItem = document.getElementById(listId);
	listItem.parentNode.removeChild(listItem);

	for (var i = 0; i < discountCodes.length; i++) {
		if (discountCodes[0] == code) {
			discountCodes.splice(i, 1)
		}
	}

	removeDiscountFromList(code, fullWeekendDollarDiscounts);
	removeDiscountFromList(code, fullWeekendPercentDiscounts);
	removeDiscountFromList(code, dancePassDollarDiscounts);
	removeDiscountFromList(code, dancePassPercentDiscounts);
	removeDiscountFromList(code, soloJazzDollarDiscounts);
	removeDiscountFromList(code, soloJazzPercentDiscounts);
	removeDiscountFromList(code, teamCompDollarDiscounts);
	removeDiscountFromList(code, teamCompPercentDiscounts);
	removeDiscountFromList(code, tshirtDollarDiscounts);
	removeDiscountFromList(code, tshirtPercentDiscounts);

	recalculateTotal();
}

function calculateCost(base, percentDiscounts, dollarDiscounts) {
	var result = base;

	for (var i = 0; i < percentDiscounts.length; i++) {
		result *= percentDiscounts[i].percent;
	}
	for (var i = 0; i < dollarDiscounts.length; i++) {
		result -= dollarDiscounts[i].dollar;
	}
	if (result < 0) {
		result = 0;
	}

	return result;
}

function recalculateTotal() {
	var total = 0;
	switch (weekendPassSelector.value) {
		case "Dance":
			total += calculateCost(dancePassCost, dancePassPercentDiscounts, dancePassDollarDiscounts)
			break;
		case "Full":
			dollardiscounts = [...fullWeekendDollarDiscounts]
			if (studentBox.checked) {
				var itemDiscount = new Object();
				itemDiscount.dollar = studentDiscount
				dollardiscounts.push(itemDiscount)
			}
			total += calculateCost(fullWeekendCost, fullWeekendPercentDiscounts, dollardiscounts)
			break;
	}
	if (mixAndMatchBox.checked) {
		total += calculateCost(mixAndMatchCost, mixAndMatchPercentDiscounts, mixAndMatchDollarDiscounts)
	}
	if (teamCompBox.checked) {
		total += calculateCost(teamCompCost, teamCompPercentDiscounts, teamCompDollarDiscounts)
	}
	if (tShirtBox.checked) {
		total += calculateCost(tshirtCost, tshirtPercentDiscounts, tshirtDollarDiscounts)
	}
	if (soloJazzBox.checked) {
		total += calculateCost(soloJazzCost, soloJazzPercentDiscounts, soloJazzDollarDiscounts)
	}

	totalCount.textContent = parseDollar(total);
}

function couponEnter(e) {
	if (e.which != 13) {
		return;
	}
	e.preventDefault();	
	submitDiscount();
}

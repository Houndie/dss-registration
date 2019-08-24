function onLoad() {
	var access_token = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().access_token;
	if (typeof access_token === "undefined") {
		window.location.href = siteBase;
	}
	var req = new XMLHttpRequest();
	req.onreadystatechange = function () {
		if (req.readyState != 4) {
			return;
		}
		var resp 
		try {
			resp = JSON.parse(req.responseText);
			if (typeof resp.errors !== "undefined" && resp.errors.length != 0) {
				window.location.href = siteBase + "/error/?source_page=/my_registrations&message="+encodeURI(req.responseText)
				return;
			}

			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/my_registrations&message=status"+req.status
				return
			}

		} catch(e) {
			if (req.status != 200) {
				window.location.href = siteBase + "/error/?source_page=/my_registrations&message=status"+req.status
				return
			}
			if (req.responseText == "") {
				window.location.href = siteBase + "/error/?source_page=/my_registrations&message=empty_response_body"
				return
			}

			window.location.href = siteBase + "/error/?source_page=/my_registrations&message="+req.responseText
			return
		}
		var body = document.getElementById("registrations-body")
		for (var i = 0; i < resp.registrations.length; i++) {
			var row = body.insertRow(i)
			var a = document.createElement('a');
			var linkText = document.createTextNode(resp.registrations[i].created_at);
			a.appendChild(linkText);
			a.href = siteBase+"/my_registration?registration_id="+resp.registrations[i].registration_id;
			row.insertCell(0).appendChild(a);
			row.insertCell(1).textContent = resp.registrations[i].first_name;
			row.insertCell(2).textContent = resp.registrations[i].last_name;
			row.insertCell(3).textContent = resp.registrations[i].email;
			row.insertCell(4).textContent = resp.registrations[i].paid;
		}
	}
	req.open("GET", dynamicBase +"/ListUserRegistrations", true);
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.setRequestHeader("Accept", "application/json")
	req.send(null)
}

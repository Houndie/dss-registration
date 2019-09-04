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
		if (resp.registrations.length > 0) {
			var body = document.getElementById("registrations-body")
			var sorted_registrations = resp.registrations
			sorted_registrations.sort(function(a,b){
				var date1 = new Date(a.created_at)
				var date2 = new Date(b.created_at)
				return date1 > date2;
			});
			for (var i = 0; i < sorted_registrations.length; i++) {
				var row = body.insertRow(i)
				var a = document.createElement('a');
				var date = new Date(sorted_registrations[i].created_at)
				var linkText = document.createTextNode(date.toString());
				a.appendChild(linkText);
				a.href = siteBase+"/my_registration?registration_id="+sorted_registrations[i].registration_id;
				row.insertCell(0).appendChild(a);
				row.insertCell(1).textContent = sorted_registrations[i].first_name;
				row.insertCell(2).textContent = sorted_registrations[i].last_name;
				row.insertCell(3).textContent = sorted_registrations[i].email;
				row.insertCell(4).textContent = sorted_registrations[i].paid;
			}
			document.getElementById('populate-loading').style.display = 'none';
			document.getElementById('registration-table').style.display = 'block';
		} else {
			document.getElementById('populate-loading').style.display = 'none';
			document.getElementById('no-registrations').style.display = 'block';
		}
	}
	req.open("GET", dynamicBase +"/ListUserRegistrations", true);
	req.setRequestHeader("Authorization", "Bearer "+access_token)
	req.setRequestHeader("Accept", "application/json")
	req.send(null)
}

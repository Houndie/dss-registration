// Client ID and API key from the Developer Console
var CLIENT_ID = '166144116294-c115t8bqllktva4qp6tvjjeqe7mdggu3.apps.googleusercontent.com';
var API_KEY = 'AIzaSyAJaUR7I6ADbch4OX-WdkjlYsnOrhBx3xU';


// Array of API discovery doc URLs for APIs used by the quickstart
var DISCOVERY_DOCS = ["https://www.googleapis.com/discovery/v1/apis/people/v1/rest"];

// Authorization scopes required by the API; multiple scopes can be
// included, separated by spaces.
var SCOPES = "https://www.googleapis.com/auth/userinfo.email";

var authorizeButton = document.getElementById('authorize_button');
var signoutButton = document.getElementById('signout_button');
var myRegistrationsButton = document.getElementById('my_registrations_button');

/**
*  On load, called to load the auth2 library and API client library.
*/
function handleClientLoad() {
	gapi.load('client:auth2', function(){
		gapi.client.init({
			apiKey: API_KEY,
			clientId: CLIENT_ID,
			discoveryDocs: DISCOVERY_DOCS,
			scope: SCOPES
		}).then(function () {
			// Listen for sign-in state changes.
			gapi.auth2.getAuthInstance().isSignedIn.listen(updateSigninStatus);

			// Handle the initial sign-in state.
			updateSigninStatus(gapi.auth2.getAuthInstance().isSignedIn.get());
			authorizeButton.onclick = handleSigninClick;
			signoutButton.onclick = handleSignoutClick;

			if (typeof onLoad === 'function') {
				onLoad();
			}
		}, function(error) {
			alert(JSON.stringify(error,null,2));
		});
	})
}

/**
*  Initializes the API client library and sets up sign-in state
*  listeners.
*/
function initClient() {
}

/**
*  Called when the signed in status changes, to update the UI
*  appropriately. After a sign-in, the API is called.
*/
function updateSigninStatus(isSignedIn) {
	if (isSignedIn) {
		authorizeButton.style.display = 'none';
		signoutButton.style.display = 'block';
		myRegistrationsButton.style.display = 'block';
	} else {
		authorizeButton.style.display = 'block';
		signoutButton.style.display = 'none';
		myRegistrationsButton.style.display = 'none';
	}

	if (typeof signInOutHook === 'function') {
		signInOutHook(isSignedIn);
	}
}

/**
*  Sign in the user upon button click.
*/
function handleSigninClick(event) {
	gapi.auth2.getAuthInstance().signIn();
}

/**
*  Sign out the user upon button click.
*/
function handleSignoutClick(event) {
	gapi.auth2.getAuthInstance().signOut();
}

function error_onload() {
	urlparams = new URLSearchParams(window.location.search);
	document.getElementById('error_source_page').textContent= urlparams.get('source_page')
	document.getElementById('error_message').textContent = urlparams.get('message')

}

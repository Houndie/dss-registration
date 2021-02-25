const ImportRecaptchaEffect = () => {
	const script = document.createElement('script');

	script.src = "https://www.google.com/recaptcha/api.js?onload=onloadCallback&render=explicit";
	script.async = true;
	script.defer = true;

	document.body.appendChild(script);

	return () => {
	 document.body.removeChild(script);
	}
}

export { ImportRecaptchaEffect }

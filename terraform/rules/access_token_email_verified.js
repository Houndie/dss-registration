function(user, context, callback) {
	context.accessToken["${namespace}" + '/email_verified'] = user.email_verified;
	callback(null, user, context);
}

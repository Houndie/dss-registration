function (user, context, callback) {
	  if(user.email_verified || context.request.query.prompt === "none" || context.stats.loginsCount <= 1) {
		      return callback(null, user, context);
		    }
	  
	  var ManagementClient = require('auth0@2.9.1').ManagementClient;
	  
	  var management = new ManagementClient({
		      token: auth0.accessToken,
		      domain: auth0.domain
		    });

		management.sendEmailVerification({user_id: user.user_id}, () => {
			    return callback(null, user, context);
			  });
}

import React from 'react';
import { Auth0Provider } from '@auth0/auth0-react';
import { navigate } from 'gatsby';

interface RootElementProps {
	element: React.ReactNode
}

export default ({ element }: RootElementProps) => (
	<Auth0Provider
		domain={`${process.env.GATSBY_AUTH0_DOMAIN}`}
		clientId={`${process.env.GATSBY_CLIENT_ID}`}
		redirectUri={window.location.origin}
		onRedirectCallback={(appState) => {
			navigate(appState?.returnTo || '/', { replace: true });
		}}
		audience={`${process.env.GATSBY_AUTH0_AUDIENCE}`}
	>
		{element}
	</Auth0Provider>
)

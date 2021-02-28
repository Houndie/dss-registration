import React from 'react'
import {Clients, AuthedClients, AuthStatus} from './auth'

interface AuthGuardProps {
	clients: Clients
	children: (arg0: AuthedClients) => React.ReactNode
}

export default ({clients, children}: AuthGuardProps) => {
	switch(clients.status) {
		case AuthStatus.SignedIn:
			return <>{children({
				status: AuthStatus.SignedIn,
				discount: clients.discount,
				registration: clients.registration,
				forms: clients.forms
			})}</>
		case AuthStatus.SignedOut:
			return <p>You must be logged in to view this page!</p>
	}
	return <></>
}

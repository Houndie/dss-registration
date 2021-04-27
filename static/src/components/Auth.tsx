import React from 'react'
import { Auth0Provider, useAuth0 } from '@auth0/auth0-react'
import {createDiscount, createRegistration, createForms} from './twirp'
import {dss as dssDiscount} from "../rpc/discount.pb"
import {dss as dssRegistration} from "../rpc/registration.pb"
import {dss as dssForms} from "../rpc/forms.pb"

export enum AuthStatus {
	NotLoaded = 0,
	SignedOut,
	SignedIn
}

export type Clients = {
	status: AuthStatus
	discount: () => Promise<dssDiscount.Discount>
	registration: () => Promise<dssRegistration.Registration>
	forms: () => Promise<dssForms.Forms>
}

export type AuthedClients = {
	status: AuthStatus.SignedIn
	discount: () => Promise<dssDiscount.Discount>
	registration: () => Promise<dssRegistration.Registration>
	forms: () => Promise<dssForms.Forms>
}

export type AuthResult = {
	clients: Clients
	signIn: () => void
	signOut: () => void
}

interface AuthInnerProps{
	children: (arg0: AuthResult) => React.ReactNode
}

const AuthInner = ({children}: AuthInnerProps) => {
	const {isLoading, error, loginWithPopup, logout, isAuthenticated, getAccessTokenSilently} = useAuth0()
	console.log("HERE")
	console.log(isAuthenticated)
	return <>
		{children({
			signIn: loginWithPopup,
			signOut: logout,
			clients: (!isLoading && isAuthenticated ? {
				status: AuthStatus.SignedIn,
				discount: () => {
					return getAccessTokenSilently().then( token => {
						return createDiscount(token)
					})
				},
				registration: () => {
					return getAccessTokenSilently().then( token => {
						return createRegistration(token)
					})
				},
				forms: () => {
					return getAccessTokenSilently().then( token => {
						return createForms(token)
					})
				},
			} : {
				status: (isLoading ? AuthStatus.NotLoaded : AuthStatus.SignedOut),
				discount: () => {
					return Promise.resolve(createDiscount())
				},
				registration: () => {
					return Promise.resolve(createRegistration())
				},
				forms: () => {
					return Promise.resolve(createForms())
				},
			})
		})}
	</>
}

interface AuthProps{
	frontend: string,
	domain: string,
	clientId: string,	
	children: (arg0: AuthResult) => React.ReactNode
}

export default ({children, domain, frontend, clientId}: AuthProps) => (
	<Auth0Provider
		domain={domain}
		clientId={clientId}
		redirectUri={window.location.origin}
	>
		<AuthInner>{children}</AuthInner>
	</Auth0Provider>
)

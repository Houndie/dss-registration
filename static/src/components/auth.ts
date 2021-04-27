import {useState, useEffect} from 'react'
import {useGoogleLogin, useGoogleLogout, GoogleLoginResponse} from 'react-google-login'
import {createDiscount as twirpCreateDiscount} from "../rpc/discount.twirp"
import {dss as dssDiscount} from "../rpc/discount.pb"
import {createRegistration as twirpCreateRegistration} from "../rpc/registration.twirp"
import {dss as dssRegistration} from "../rpc/registration.pb"
import {createForms as twirpCreateForms} from "../rpc/forms.twirp"
import {dss as dssForms} from "../rpc/forms.pb"
import axios, {AxiosTransformer} from 'axios'
import {createDiscount, createRegistration, createForms} from './twirp'

const isGoogleLoginResponse = (res: any): res is GoogleLoginResponse => {
	return (res as GoogleLoginResponse).accessToken !== undefined
}

type BackendFunc = ((baseURL: string, options: any) => dssDiscount.Discount) | ((baseURL: string, options: any) => dssRegistration.Registration) | ((baseURL: string, options: any) => dssForms.Forms)

export enum AuthStatus {
	NotLoaded = 0,
	SignedOut,
	SignedIn
}

export type Clients = {
	status: AuthStatus
	discount: dssDiscount.Discount
	registration: dssRegistration.Registration
	forms: dssForms.Forms
}

export type AuthedClients = {
	status: AuthStatus.SignedIn
	discount: dssDiscount.Discount
	registration: dssRegistration.Registration
	forms: dssForms.Forms
}

export type AuthResult = {
	clients: Clients
	signIn: () => void
	signOut: () => void
}

type SignedInStatus = {
	status: AuthStatus.SignedIn
	timer: NodeJS.Timeout
}

type SignedOutStatus = {
	status: AuthStatus.SignedOut
}

type NotLoadedStatus = {
	status: AuthStatus.NotLoaded
}

type ClientPin = {
	status: SignedInStatus | SignedOutStatus | NotLoadedStatus
	discount: dssDiscount.Discount
	registration: dssRegistration.Registration
	forms: dssForms.Forms
}

const setReloadTimer = (expires_in: number, reload: GoogleLoginResponse["reloadAuthResponse"], setClients: (arg0: ClientPin) => void) => {
	return setTimeout(() => {
		reload().then(newAuth => {
			const nextTimer = setReloadTimer(newAuth.expires_in, reload, setClients)
			setClients({
				status: {
					status: AuthStatus.SignedIn,
					timer: nextTimer
				},
				discount: createDiscount(newAuth.access_token),
				registration: createRegistration(newAuth.access_token),
				forms: createForms(newAuth.access_token)
			})
		}).catch(e => {
			console.error("failure refreshing token")
			setClients({
				status: {
					status: AuthStatus.SignedOut
				},
				discount: createDiscount(),
				registration: createRegistration(),
				forms: createForms()
			})
		})
	}, expires_in*1000)
}

export const useAuth = (clientID: string) => {
	const [clients, setClients] = useState<ClientPin>({
		status: { status: AuthStatus.NotLoaded },
		discount: createDiscount(),
		registration: createRegistration(),
		forms: createForms()
	})

	useEffect(() => {
		return () => {
			if(clients.status.status == AuthStatus.SignedIn) {
				clearTimeout(clients.status.timer)
			}
		}
	}, [clients])

	const gLogin = useGoogleLogin({
		clientId: clientID,
		onSuccess: (newAuth) => {
			if(!isGoogleLoginResponse(newAuth)) {
				console.error("got offline auth somehow")
				return
			}
			const timer = setReloadTimer(newAuth.getAuthResponse(true).expires_in, newAuth.reloadAuthResponse, setClients)
			setClients({
				status: {
					status: AuthStatus.SignedIn,
					timer: timer
				},
				discount: createDiscount(newAuth.accessToken),
				registration: createRegistration(newAuth.accessToken),
				forms: createForms(newAuth.accessToken)
			})
		},
		onFailure: () => console.error('login failure'),
		onAutoLoadFinished: (success) => {
			if(success) {
				return
			}
			setClients({
				status: { status: AuthStatus.SignedOut },
				discount: createDiscount(),
				registration: createRegistration(),
				forms: createForms()
			})
		},
		cookiePolicy: 'single_host_origin',
		isSignedIn: true
	})

	const gLogout = useGoogleLogout({
		clientId: clientID,
		onLogoutSuccess: () => {
			if(clients.status.status == AuthStatus.SignedIn) {
				clearTimeout(clients.status.timer)
			}
			setClients({
				status: {status: AuthStatus.SignedOut},
				discount: createDiscount(),
				registration: createRegistration(),
				forms: createForms()
			})
		}
	})

	return {
		signIn: gLogin.signIn,
		signOut: gLogout.signOut,
		clients: {
			status: clients.status.status,
			discount: clients.discount,
			registration: clients.registration,
			forms: clients.forms
		}
	}
}

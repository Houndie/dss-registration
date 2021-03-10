import {createDiscount as twirpCreateDiscount} from "../rpc/discount.twirp"
import {dss as dssDiscount} from "../rpc/discount.pb"
import {createRegistration as twirpCreateRegistration} from "../rpc/registration.twirp"
import {dss as dssRegistration} from "../rpc/registration.pb"
import {createForms as twirpCreateForms} from "../rpc/forms.twirp"
import {dss as dssForms} from "../rpc/forms.pb"
import { useAuth0 } from '@auth0/auth0-react';

type BackendFunc<T> = (baseURL: string, options?: any) => T

const create = <T>(backend: BackendFunc<T>, isAuthenticated: boolean, getAccessTokenSilently: () => Promise<string>) => {
	if( !isAuthenticated ) {
		return Promise.resolve(backend(`${process.env.GATSBY_BACKEND}`))
	}
	return getAccessTokenSilently().then(token => {
		const opts = {
			headers: {
				"Authorization": `Bearer ${token}`
			},
		}
		return Promise.resolve(backend(`${process.env.GATSBY_BACKEND}`, opts))
	})
	
}

export default () => {
	const { isAuthenticated, getAccessTokenSilently } = useAuth0()
	return {
		discount: () => create(twirpCreateDiscount, isAuthenticated, getAccessTokenSilently),
		registration: () => create(twirpCreateRegistration, isAuthenticated, getAccessTokenSilently),
		forms: () => create(twirpCreateForms, isAuthenticated, getAccessTokenSilently)
	}
}

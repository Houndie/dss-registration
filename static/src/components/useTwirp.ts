import {createDiscount as twirpCreateDiscount} from "../rpc/discount.twirp"
import {createRegistration as twirpCreateRegistration} from "../rpc/registration.twirp"
import {createForms as twirpCreateForms} from "../rpc/forms.twirp"
import {createVaccine as twirpCreateVaccine} from "../rpc/vaccine.twirp"
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
		forms: () => create(twirpCreateForms, isAuthenticated, getAccessTokenSilently),
		vaccine: () => create(twirpCreateVaccine, isAuthenticated, getAccessTokenSilently)
	}
}

import {createDiscount as twirpCreateDiscount} from "../rpc/discount.twirp"
import {dss as dssDiscount} from "../rpc/discount.pb"
import {createRegistration as twirpCreateRegistration} from "../rpc/registration.twirp"
import {dss as dssRegistration} from "../rpc/registration.pb"
import {createForms as twirpCreateForms} from "../rpc/forms.twirp"
import {dss as dssForms} from "../rpc/forms.pb"
import { GoogleLoginResponse } from "react-google-login"

type BackendFunc<T> = (baseURL: string, options: any) => T

const create = <T>(backend: BackendFunc<T>, accessToken?: string, options = {}) => {
	const opts = (accessToken ? {
		headers: {
			"Authorization": `Bearer ${accessToken}`
		},
		...options
	} : options)
	
	return backend(`${process.env.GATSBY_BACKEND}`, opts)
}

export const createDiscount = (accessToken?: string, options = {}): dssDiscount.Discount => {
	return create(twirpCreateDiscount, accessToken, options)
}

export const createRegistration = (accessToken?: string, options = {}): dssRegistration.Registration => {
	return create(twirpCreateRegistration, accessToken, options)
}

export const createForms = (accessToken?: string, options = {}): dssForms.Forms => {
	return create(twirpCreateForms, accessToken, options)
}

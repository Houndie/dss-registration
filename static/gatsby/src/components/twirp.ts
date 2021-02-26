import {createDiscount as twirpCreateDiscount} from "../rpc/discount.twirp"
import {dss as dssDiscount} from "../rpc/discount.pb"
import {createRegistration as twirpCreateRegistration} from "../rpc/registration.twirp"
import {dss as dssRegistration} from "../rpc/registration.pb"
import {createForms as twirpCreateForms} from "../rpc/forms.twirp"
import {dss as dssForms} from "../rpc/forms.pb"
import { GoogleLoginResponse } from "react-google-login"

type BackendFunc = (baseURL: string, options: any) => dssDiscount.Discount | dssRegistration.Registration | dssForms.Forms

const create = (backend: BackendFunc, gAuth: GoogleLoginResponse | null = null, options = {}) => {
	const opts = (gAuth ? {
		headers: {
			"Authorization": `Bearer ${gAuth.accessToken}`
		},
		...options
	} : options)
	
	return backend(`${process.env.GATSBY_BACKEND}`, opts)
}

export const createDiscount = (gAuth: GoogleLoginResponse | null = null, options = {}): dssDiscount.Discount => {
	return create(twirpCreateDiscount, gAuth, options) as dssDiscount.Discount
}

export const createRegistration = (gAuth: GoogleLoginResponse | null = null, options = {}): dssRegistration.Registration => {
	return create(twirpCreateRegistration, gAuth, options) as dssRegistration.Registration
}

export const createForms = (gAuth: GoogleLoginResponse | null = null, options = {}): dssForms.Forms => {
	return create(twirpCreateForms, gAuth, options) as dssForms.Forms
}

import React, {useState, useEffect} from 'react'
import Page from '../components/Page'
import {Formik} from 'formik'
import {createRegistration} from "../rpc/registration.twirp"
import {dss} from "../rpc/registration.pb"
import {createDiscount} from "../rpc/discount.twirp"
import { v4 as uuidv4 } from 'uuid';
import useTwirp from "../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import RegistrationForm, {hasUnpaidItems, FormHousingOption, toProtoRegistration, FormWeekendPassOption, FormFullWeekendPassLevel, FormStyle, FormRole, RegistrationFormState} from "../components/RegistrationForm"

const initialState: RegistrationFormState  = {
	firstName: '', 
	lastName: '',
	streetAddress: '',
	city: '',
	state: '',
	zipCode: '',
	email: '',
	homeScene: '',
	isStudent: false,
	passType: FormWeekendPassOption.noPassOption,
	level: FormFullWeekendPassLevel.NotSelected,
	mixAndMatch: false,
	role: FormRole.NotSelected,
	soloJazz: false,
	teamCompetition: false,
	teamName: '',
	tshirt: false,
	style: FormStyle.NotSelected,
	housing: FormHousingOption.noHousingOption,
	pets: '',
	quantity: 0,
	provideDetails: '',
	petAllergies: '',
	requireDetails: '',
	vaccine: undefined,
	discounts: []
}

const Registration = () => {
	const [prices, setPrices] = useState<dss.RegistrationPricesRes | null>(null)
	const {registration} = useTwirp()

	useEffect(() => {
		registration().then(client => {
			return client.prices({})
		}).then(res => {
			setPrices(res);
		}, err => {
			console.error(err);
		});
	}, [])

	const { isLoading, isAuthenticated, loginWithRedirect } = useAuth0()

	return (
		<Page title="Registration">
			{() =>  {
				if(isLoading) {
					return <></>
				}
				if( !isAuthenticated ){
					return <p>You must be logged in to register! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
				}
				return (
					<Formik
						initialValues={initialState}
						onSubmit={(values, { setSubmitting }) => {
							if(!prices) {
								console.error("prices is null?")
								setSubmitting(false)
								return
							}

							const clientReg = toProtoRegistration(values, prices.weekendPassTier)

							return registration().then(client => {
								return client.add({
									registration: clientReg
								}).then(createRes => {
									if( !createRes.registration) {
										throw "No registration returned";
									}

									if (!values.vaccine) {
										return createRes.registration
									}

									return client.uploadVaxImage({
										id: createRes.registration.id,
										filesize: values.vaccine.size
									}).then(uploadRes => {
										if(!uploadRes.url) {
											throw "No upload url returned"
										}

										return fetch(uploadRes.url, {
											method: "PUT",
											body: values.vaccine
										})
									}).then(() => {
										if( !createRes.registration) {
											throw "No registration returned";
										}

										return createRes.registration
									})
								}).then( r => {
									const redirectURL = `${process.env.GATSBY_FRONTEND}/registration-complete`

									if (!hasUnpaidItems(r)) {
										return redirectURL
									}

									return client.pay({
										id: r.id,
										idempotencyKey: uuidv4(),
										redirectUrl: redirectURL,
									}).then((res) => {
										return res.checkoutUrl
									})
								})
							}).then(redirectTo => {
								window.location.href = redirectTo;	
							}).catch(err => {
								console.error(err);
							});
						}}
					>
					{({values, isSubmitting, handleSubmit, setFieldValue}) => prices != null ? (
						<RegistrationForm prices={prices}/>
					) : null}
					</Formik>
				)
			}}
		</Page>
	);
};

export default Registration

import React, {useState, useEffect} from 'react'
import Page from '../components/Page'
import {Formik} from 'formik'
import {createRegistration} from "../rpc/registration.twirp"
import {dss} from "../rpc/registration.pb"
import {createDiscount} from "../rpc/discount.twirp"
import { v4 as uuidv4 } from 'uuid';
import useTwirp from "../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import LoadingPage from "../components/LoadingPage"
import RegistrationForm, {isPaid, FormHousingOption, toProtoRegistration, FormWeekendPassOption, FormFullWeekendPassLevel, FormStyle, FormRole, RegistrationFormState, formValidate} from "../components/RegistrationForm"
import PleaseVerifyEmail from "../components/PleaseVerifyEmail"

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
	weekendPassOverride: false,
	danceOnlyOverride: false,
	mixAndMatch: false,
	role: FormRole.NotSelected,
	mixAndMatchOverride: false,
	soloJazz: false,
	soloJazzOverride: false,
	teamCompetition: false,
	teamName: '',
	teamCompetitionOverride: false,
	tshirt: false,
	tshirtOverride: false,
	style: FormStyle.NotSelected,
	housing: FormHousingOption.noHousingOption,
	pets: '',
	quantity: 0,
	provideDetails: '',
	petAllergies: '',
	requireDetails: '',
	vaccine: undefined,
	discounts: {},
	enabled: true
}

const Registration = () => {
	const [prices, setPrices] = useState<dss.RegistrationPricesRes | null>(null)
	const {registration, vaccine} = useTwirp()

	useEffect(() => {
		registration().then(client => {
			return client.prices({})
		}).then(res => {
			setPrices(res);
		}, err => {
			console.error(err);
		});
	}, [])

	const { isLoading, isAuthenticated, loginWithRedirect, user } = useAuth0()

	return (
		<Page title="Registration">
			{() =>  {
				if(process.env.GATSBY_ACTIVE !== 'true') {
					return <p>Registration is not open yet!</p>
				}
				if(isLoading) {
					return <LoadingPage/>
				}
				if(!prices) {
					return <LoadingPage/>
				}
				if( !isAuthenticated ){
					return (<>
						<p>You must be logged in to register! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
						<p>Note: If you have having issues logging in on the safari browser, please try using chrome or firefox.</p>
						</>)
				}
				if(!user?.email_verified) {
					return <PleaseVerifyEmail/>
				}
				return (
					<Formik
						initialValues={initialState}
						validate={formValidate}
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

									return vaccine().then(vaxClient => {
										return vaxClient.upload({
											id: createRes.registration?.id,
											filesize: values.vaccine?.size
										})
									}).then(uploadRes => {
										if(!uploadRes.url) {
											throw "No upload url returned"
										}

										return fetch(uploadRes.url, {
											method: "PUT",
											body: values.vaccine
										})
									}).then((res) => {
										if(!res.ok) {
											throw "error uploading vaccine card:  " + res.statusText
										}

										if( !createRes.registration) {
											throw "No registration returned";
										}

										return createRes.registration
									})
								}).then( r => {
									const redirectURL = `${process.env.GATSBY_FRONTEND}/registration-complete`

									if (isPaid(r)) {
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
					{({values, isSubmitting, handleSubmit, setFieldValue}) => (
						<RegistrationForm weekendPassTier={prices.weekendPassTier} admin={false}/>
					)}
					</Formik>
				)
			}}
		</Page>
	);
};

export default Registration

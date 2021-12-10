import React, {useEffect, useState} from 'react'
import Page from '../components/Page'
import Alert from 'react-bootstrap/Alert'
import WithAlert, {ResponseKind} from '../components/WithAlert'
import useTwirp from "../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import { v4 as uuidv4 } from 'uuid';
import {dss} from "../rpc/registration.pb"
import RegistrationForm, {isPaid, RegistrationFormState, toProtoRegistration, formWeekendPassOptionFromProto, fromProtoHousingOption, FormFullWeekendPassLevel, FormRole, FormStyle, fromProtoPassLevel, fromProtoRole, fromProtoStyle, formValidate} from "../components/RegistrationForm"
import {VaccineInfoEnum, VaccineInfo, fromProtoVaccine} from "../components/vaccine"
import {Formik} from 'formik'
import LoadingPage from "../components/LoadingPage"
import PleaseVerifyEmail from "../components/PleaseVerifyEmail"
import {Router} from "@reach/router"

export default () => { 
	const vaccineRef = React.useRef<HTMLInputElement>()
	const [prices, setPrices] = useState<dss.RegistrationPricesRes | null>(null)
	const [myRegistration, setMyRegistration] = useState<dss.IRegistrationInfo | null>(null)
	const [myVaccine, setMyVaccine] = useState<VaccineInfo | null>(null)
	const {registration, vaccine} = useTwirp()
	const { isLoading, isAuthenticated, loginWithRedirect, user } = useAuth0()
	var id: string | null = null
	useEffect(() => {
		const params = new URLSearchParams(window.location.search)
		id = params.get("id")
	})

	useEffect(() => {
		registration().then(client => {
			return client.prices({})
		}).then(res => {
			setPrices(res);
		}, err => {
			console.error(err);
		});
	}, [])

	useEffect(() => {
		if(isLoading || !isAuthenticated || !user?.email_verified || !id){
			setMyRegistration(null)
			return
		}

		registration().then(client => {
			return client.get({
				id: id
			})
		}).then(res => {
			if(!res.registration) {
				console.error("no registration?")
				return
			}
			setMyRegistration(res.registration);
		}, err => {
			console.error(err);
		});
	}, [isLoading, isAuthenticated, user])

	useEffect(() => {
		if(isLoading || !isAuthenticated || !user?.email_verified){
			setMyVaccine(null)
			return
		}

		vaccine().then(client => {
			return client.get({
				id: id
			})
		}).then(res => {
			setMyVaccine(fromProtoVaccine(res))
		}, err => {
			console.error(err);
		});
	}, [isLoading, isAuthenticated, user])

	const urlSearchParams = (typeof window !== "undefined" ? new URLSearchParams(window.location.search) : undefined);

	return (
		<Page title="Registration">{() => (
			<WithAlert>{(setResponse) => {
				if (!prices) {
					return <LoadingPage/>
				}

				if(isLoading) {
					return <LoadingPage/>
				}

				if( !isAuthenticated ){
					return <p>You must be logged in to view this page! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
				}

				if(!user?.email_verified) {
					return <PleaseVerifyEmail/>
				}

				if (!myRegistration) {
					return <LoadingPage/>
				}

				if (!myVaccine) {
					return <LoadingPage/>
				}

				if (!myRegistration.enabled)  {
					return <p>This registration is disabled.</p>
				}

				return (
					<>
						{!isPaid(myRegistration) && (
							<Alert variant="warning">
								<p>You registered for the following items, but haven't paid yet:</p>
									<ul>
										{(myRegistration.fullWeekendPass && !myRegistration.fullWeekendPass.squarePaid) && (<li>Full Weekend Pass</li>)}
										{(myRegistration.danceOnlyPass && !myRegistration.danceOnlyPass.squarePaid) && (<li>Dance Only Pass</li>)}
										{(myRegistration.mixAndMatch && !myRegistration.mixAndMatch.squarePaid) && (<li>Mix and Match</li>)}
										{(myRegistration.soloJazz && !myRegistration.soloJazz.squarePaid) && (<li>Solo Jazz</li>)}
										{(myRegistration.teamCompetition && !myRegistration.teamCompetition.squarePaid) && (<li>Team Competition</li>)}
										{(myRegistration.tshirt && !myRegistration.tshirt.squarePaid) && (<li>T-Shirt</li>)}
									</ul>
								<p>You will be taken to the payment page after updating your registration</p>
							</Alert>
						)}
						{(urlSearchParams && urlSearchParams.has("referenceId")) && (
							<Alert variant="success">Registration updated successfully!</Alert>
						)}
						<Formik
							initialValues={{
								firstName: (myRegistration.firstName ? myRegistration.firstName : ""),
								lastName: (myRegistration.lastName ? myRegistration.lastName : ""),
								streetAddress: (myRegistration.streetAddress ? myRegistration.streetAddress : ""),
								city: (myRegistration.city ? myRegistration.city : ""),
								state: (myRegistration.state ? myRegistration.state : ""),
								zipCode: (myRegistration.zipCode ? myRegistration.zipCode : ""),
								email: (myRegistration.email ? myRegistration.email : ""),
								homeScene: (myRegistration.homeScene ? myRegistration.homeScene : ""),
								isStudent: (myRegistration.isStudent ? myRegistration.isStudent : false),
								passType: formWeekendPassOptionFromProto(myRegistration),
								level: (myRegistration.fullWeekendPass && myRegistration.fullWeekendPass.level != null ? fromProtoPassLevel(myRegistration.fullWeekendPass.level) : FormFullWeekendPassLevel.NotSelected),
								weekendPassOverride: Boolean(myRegistration.fullWeekendPass && myRegistration.fullWeekendPass.adminPaymentOverride),
								danceOnlyOverride: Boolean(myRegistration.danceOnlyPass && myRegistration.danceOnlyPass.adminPaymentOverride),
								mixAndMatch: Boolean(myRegistration.mixAndMatch),
								role: (myRegistration.mixAndMatch && myRegistration.mixAndMatch.role != null? fromProtoRole(myRegistration.mixAndMatch.role) : FormRole.NotSelected),
								mixAndMatchOverride: Boolean(myRegistration.mixAndMatch && myRegistration.mixAndMatch.adminPaymentOverride),
								soloJazz: Boolean(myRegistration.soloJazz),
								soloJazzOverride: Boolean(myRegistration.soloJazz && myRegistration.soloJazz.adminPaymentOverride),
								teamCompetition: Boolean(myRegistration.teamCompetition),
								teamName: (myRegistration.teamCompetition && myRegistration.teamCompetition.name != null ? myRegistration.teamCompetition.name : ''),
								teamCompetitionOverride: Boolean(myRegistration.teamCompetition && myRegistration.teamCompetition.adminPaymentOverride),
								tshirt: Boolean(myRegistration.tshirt),
								style: (myRegistration.tshirt && myRegistration.tshirt.style != null ? fromProtoStyle(myRegistration.tshirt.style) : FormStyle.NotSelected),
								tshirtOverride: Boolean(myRegistration.tshirt && myRegistration.tshirt.adminPaymentOverride),
								housing: fromProtoHousingOption(myRegistration),
								pets: (myRegistration.provideHousing && myRegistration.provideHousing.pets ? myRegistration.provideHousing.pets : ""),
								quantity: (myRegistration.provideHousing && myRegistration.provideHousing.quantity ? myRegistration.provideHousing.quantity : 0),
								provideDetails: (myRegistration.provideHousing && myRegistration.provideHousing.details ? myRegistration.provideHousing.details : ""),
								petAllergies: (myRegistration.requireHousing && myRegistration.requireHousing.petAllergies ? myRegistration.requireHousing.petAllergies : ""), requireDetails: (myRegistration.requireHousing && myRegistration.requireHousing.details ? myRegistration.requireHousing.details : ""), 
								vaccine: undefined,
								discounts: (myRegistration.discountCodes ? myRegistration.discountCodes : []),
								enabled: Boolean(myRegistration.enabled)
							}}
							validate={formValidate}
							onSubmit={(values: RegistrationFormState, {setSubmitting, setFieldValue}) => {
								setResponse(null)

								const tier = (myRegistration.fullWeekendPass && myRegistration.fullWeekendPass.tier ? myRegistration.fullWeekendPass.tier : prices.weekendPassTier)

								const clientReg = toProtoRegistration(values, tier)
								clientReg.id = myRegistration.id
								clientReg.createdAt = myRegistration.createdAt

								return registration().then(client => {
									return client.update({
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
										}).then(res => {
											if(!res.ok) {
												throw "error uploading vaccine card:  " + res.statusText
											}

											if( !createRes.registration) {
												throw "No registration returned";
											}

											return createRes.registration
										})
									}).then( r => {
										const redirectURL = `${process.env.GATSBY_FRONTEND}/registration/${r.id}`

										if (isPaid(r)) {
											setFieldValue("vaccine", undefined)
											if(vaccineRef.current){
												vaccineRef.current.value = ""
											}

											setResponse({
												kind: ResponseKind.Good,
												message: "Registration updated successfully!"
											})

											return vaccine().then(client => {
												return client.get({
													id: id
												})
											}).then(res => {
												setMyVaccine(fromProtoVaccine(res))
											})
										}

										return client.pay({
											id: r.id,
											idempotencyKey: uuidv4(),
											redirectUrl: redirectURL,
										}).then(res => {
											window.location.href = res.checkoutUrl
											return
										})
									})
								}).catch(err => {
									setResponse({
										kind: ResponseKind.Bad,
										message: err.toString(),
									})
								});
							}}
						>
							<RegistrationForm 
								weekendPassTier={prices.weekendPassTier}
								previousRegistration={myRegistration}
								admin={false}
								vaccineUpload={myVaccine}
								vaccineRef={vaccineRef}
							/>
						</Formik>
					</>
				)
			}}</WithAlert>
		)}</Page>
	)
}

import React, {useEffect, useState} from 'react'
import Page from '../components/Page'
import Alert from 'react-bootstrap/Alert'
import WithAlert, {ResponseKind} from '../components/WithAlert'
import useTwirp from "../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import { v4 as uuidv4 } from 'uuid';
import {dss} from "../rpc/registration.pb"
import RegistrationForm, {hasUnpaidItems, RegistrationFormState, toProtoRegistration, formWeekendPassOptionFromProto, fromProtoHousingOption, FormFullWeekendPassLevel, FormRole, FormStyle, fromProtoPassLevel, fromProtoRole, fromProtoStyle} from "../components/RegistrationForm"
import {Formik} from 'formik'

type UserRegistrationProps = {
	id: string
}

export default ({id}: UserRegistrationProps) => { 
	const [prices, setPrices] = useState<dss.RegistrationPricesRes | null>(null)
	const [myRegistration, setMyRegistration] = useState<dss.IRegistrationInfo | null>(null)
	const {registration} = useTwirp()
	const { isLoading, isAuthenticated, loginWithRedirect } = useAuth0()

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
		if(isLoading || !isAuthenticated){
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
	}, [isLoading, isAuthenticated])

	const urlSearchParams = (typeof window !== "undefined" ? new URLSearchParams(window.location.search) : undefined);

	return (
		<Page title="Registration">{() => (
			<WithAlert>{(setResponse) => {
				if (!prices) {
					return <></>
				}

				if(isLoading) {
					return <></>
				}

				if( !isAuthenticated ){
					return <p>You must be logged in to view this page! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
				}

				if (!myRegistration) {
					return <></>
				}

				return (
					<>
						{hasUnpaidItems(myRegistration) && (
							<Alert variant="warning">
								<p>You registered for the following items, but haven't paid yet:</p>
									<ul>
										{(myRegistration.fullWeekendPass && !myRegistration.fullWeekendPass.paid) && (<li>Full Weekend Pass</li>)}
										{(myRegistration.danceOnlyPass && !myRegistration.danceOnlyPass.paid) && (<li>Dance Only Pass</li>)}
										{(myRegistration.mixAndMatch && !myRegistration.mixAndMatch.paid) && (<li>Mix and Match</li>)}
										{(myRegistration.soloJazz && !myRegistration.soloJazz.paid) && (<li>Solo Jazz</li>)}
										{(myRegistration.teamCompetition && !myRegistration.teamCompetition.paid) && (<li>Team Competition</li>)}
										{(myRegistration.tshirt && !myRegistration.tshirt.paid) && (<li>T-Shirt</li>)}
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
								level: (myRegistration.fullWeekendPass && myRegistration.fullWeekendPass.level ? fromProtoPassLevel(myRegistration.fullWeekendPass.level) : FormFullWeekendPassLevel.NotSelected),
								mixAndMatch: Boolean(myRegistration.mixAndMatch),
								role: (myRegistration.mixAndMatch && myRegistration.mixAndMatch.role ? fromProtoRole(myRegistration.mixAndMatch.role) : FormRole.NotSelected),
								soloJazz: Boolean(myRegistration.soloJazz),
								teamCompetition: Boolean(myRegistration.teamCompetition),
								teamName: (myRegistration.teamCompetition && myRegistration.teamCompetition.name ? myRegistration.teamCompetition.name : ''),
								tshirt: Boolean(myRegistration.tshirt),
								style: (myRegistration.tshirt && myRegistration.tshirt.style ? fromProtoStyle(myRegistration.tshirt.style) : FormStyle.NotSelected),
								housing: fromProtoHousingOption(myRegistration),
								pets: (myRegistration.provideHousing && myRegistration.provideHousing.pets ? myRegistration.provideHousing.pets : ""),
								quantity: (myRegistration.provideHousing && myRegistration.provideHousing.quantity ? myRegistration.provideHousing.quantity : 0),
								provideDetails: (myRegistration.provideHousing && myRegistration.provideHousing.details ? myRegistration.provideHousing.details : ""),
								petAllergies: (myRegistration.requireHousing && myRegistration.requireHousing.petAllergies ? myRegistration.requireHousing.petAllergies : ""),
								requireDetails: (myRegistration.requireHousing && myRegistration.requireHousing.details ? myRegistration.requireHousing.details : ""),
								vaccine: undefined,
								discounts: (myRegistration.discountCodes ? myRegistration.discountCodes : [])
							}}
							onSubmit={(values: RegistrationFormState, {setSubmitting}) => {
								setResponse(null)

								const tier = (myRegistration.fullWeekendPass && myRegistration.fullWeekendPass.tier ? myRegistration.fullWeekendPass.tier : prices.weekendPassTier)

								const clientReg = toProtoRegistration(values, tier)
								clientReg.id = myRegistration.id

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
										const redirectURL = `${process.env.GATSBY_FRONTEND}/registration/${r.id}`

										if (!hasUnpaidItems(r)) {
											setResponse({
												kind: ResponseKind.Good,
												message: "Registration updated successfully!"
											})
											return
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
										message: err,
									})
								});
							}}
						>
							<RegistrationForm 
								prices={prices} 
								disables={{
									passType: !Boolean(myRegistration.noPass),
									mixAndMatch: Boolean(myRegistration.mixAndMatch),
									soloJazz: Boolean(myRegistration.soloJazz),
									teamCompetition: Boolean(myRegistration.teamCompetition),
									tShirt: Boolean(myRegistration.tshirt)
								}}
							/>
						</Formik>
					</>
				)
			}}</WithAlert>
		)}</Page>
	)
}

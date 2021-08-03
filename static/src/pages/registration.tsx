import React, {useState, useEffect} from 'react'
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import Col from "react-bootstrap/Col"
import Row from "react-bootstrap/Row"
import Page from '../components/Page'
import FormField from '../components/FormField'
import FormSelect from '../components/FormSelect'
import FormCheck from '../components/FormCheck'
import {Formik} from 'formik'
import {createRegistration} from "../rpc/registration.twirp"
import {dss} from "../rpc/registration.pb"
import {createDiscount} from "../rpc/discount.twirp"
import { v4 as uuidv4 } from 'uuid';
import useTwirp from "../components/useTwirp"
import parseDollar from "../components/parseDollar"
import { useAuth0 } from '@auth0/auth0-react';

const fullWeekendPassOption = "Full";
const danceOnlyPassOption = "Dance";
const noPassOption = "None";

const provideOption = "Provide";
const requireOption = "Require";
const noHousingOption = "None";

const registrationClient = createRegistration(`${process.env.GATSBY_BACKEND}`)
const discountClient = createDiscount(`${process.env.GATSBY_BACKEND}`)

enum FormFullWeekendPassLevel {
	NotSelected = "not_selected",
	Level1 = "level_1",
	Level2 = "level_2",
	Level3 = "level_3"
}

const toProtoPassLevel = (level: FormFullWeekendPassLevel) => {
	switch(level){
		case FormFullWeekendPassLevel.Level1:
			return dss.FullWeekendPassLevel.Level1
		case FormFullWeekendPassLevel.Level2:
			return dss.FullWeekendPassLevel.Level2
		case FormFullWeekendPassLevel.Level3:
			return dss.FullWeekendPassLevel.Level3
		default:
			throw "cannot convert unselected pass level"
	}
}

enum FormRole {
	NotSelected = "not_selected",
	Follower = "follower",
	Leader = "leader"
}

const toProtoRole = (role: FormRole) => {
	console.log(typeof role)
	console.log(typeof FormRole.Leader)
	switch(role){
		case FormRole.Leader:
			console.log("A1")
			return dss.MixAndMatch.Role.Leader
		case FormRole.Follower:
			return dss.MixAndMatch.Role.Follower
		default:
			console.log("A2")
			throw "cannot convert unselected role"
	}
}

enum FormStyle {
	NotSelected = -1,
	UnisexS = "unisex_s",
	UnisexM = "unisex_m",
	UnisexL = "unisex_l",
	UnisexXL = "unisex_xl",
	Unisex2XL = "unisex_2xl",
	Unisex3XL = "unisex_3xl",
	BellaS = "bella_s",
	BellaM = "bella_m",
	BellaL = "bella_l",
	BellaXL = "bella_xl",
	Bella2XL = "bella_2xl",
}

const toProtoStyle = (style: FormStyle) => {
	switch(style){
		case FormStyle.UnisexS:
			return dss.TShirt.Style.UnisexS
		case FormStyle.UnisexM:
			return dss.TShirt.Style.UnisexM
		case FormStyle.UnisexL:
			return dss.TShirt.Style.UnisexL
		case FormStyle.UnisexXL:
			return dss.TShirt.Style.UnisexXL
		case FormStyle.Unisex2XL:
			return dss.TShirt.Style.Unisex2XL
		case FormStyle.Unisex3XL:
			return dss.TShirt.Style.Unisex3XL
		case FormStyle.BellaS:
			return dss.TShirt.Style.BellaS
		case FormStyle.BellaM:
			return dss.TShirt.Style.BellaM
		case FormStyle.BellaL:
			return dss.TShirt.Style.BellaL
		case FormStyle.BellaXL:
			return dss.TShirt.Style.BellaXL
		case FormStyle.Bella2XL:
			return dss.TShirt.Style.Bella2XL
		default:
			throw "cannot convert unselected style"
	}
}

const fromProtoTier = (tier: number | Long) => {
	switch(tier) {
		case 0:
			return "Tier 1"
		case 1:
			return "Tier 2"
		case 2:
			return "Tier 3"
		case 3:
			return "Tier 4"
		case 4:
			return "Tier 5"
		default:
			throw "unknown tier"
	}
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


	const square_data = JSON.parse(`${process.env.GATSBY_SQUARE_DATA}`)

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
						initialValues={{
							firstName: '', 
							lastName: '',
							streetAddress: '',
							city: '',
							state: '',
							zipCode: '',
							email: '',
							homeScene: '',
							isStudent: false,
							passType: noPassOption,
							level: FormFullWeekendPassLevel.NotSelected,
							mixAndMatch: false,
							role: FormRole.NotSelected,
							soloJazz: false,
							teamCompetition: false,
							teamName: '',
							tshirt: false,
							style: FormStyle.NotSelected,
							housing: noHousingOption,
							pets: '',
							quantity: 0,
							provideDetails: '',
							petAllergies: '',
							requireDetails: '',
							discounts: []
						}}
						onSubmit={(values, { setSubmitting }) => {
							if(!prices) {
								console.error("prices is null?")
								setSubmitting(false)
								return
							}
							const clientReg: dss.IRegistrationInfo = {
								firstName: values.firstName,
								lastName: values.lastName,
								streetAddress: values.streetAddress,
								city: values.city,
								state: values.state,
								zipCode: values.zipCode,
								email: values.email,
								homeScene: values.homeScene,
								isStudent: values.isStudent
							}

							switch (values.passType) {
								case fullWeekendPassOption:
									clientReg.fullWeekendPass = {
										tier: prices.weekendPassTier,
										level: toProtoPassLevel(values.level)
									}
									break
								case danceOnlyPassOption:
									clientReg.danceOnlyPass = {}
									break
								default:
									clientReg.noPass = {}
									break
							}

							if (values.mixAndMatch) {
								console.log(values.role)
								clientReg.mixAndMatch = { role: toProtoRole(values.role) }
							}

							if (values.soloJazz) {
								clientReg.soloJazz = {}
							}

							if (values.teamCompetition) {
								clientReg.teamCompetition = { name: values.teamName }
							}

							if (values.tshirt) {
								clientReg.tshirt = { style: toProtoStyle(values.style) }
							}

							switch (values.housing) {
								case provideOption: {
										clientReg.provideHousing = {
											pets: values.pets,
											quantity: values.quantity,
											details: values.provideDetails
										}
									}
									break
								case requireOption: {
										clientReg.requireHousing = {
											petAllergies: values.petAllergies,
											details: values.requireDetails,
										}
									}
									break
								default:
									clientReg.noHousing = {}
									break
							}

							return registration().then(client => {
								return client.add({
									registration: clientReg
								}).then(res => {
									if( !res.registration) {
										throw "No registration returned";
									}
									return client.pay({
										id: res.registration.id,
										idempotencyKey: uuidv4(),
										redirectUrl: `${process.env.GATSBY_FRONTEND}/registration-complete`,
									})
								})
							}).then(res => {
								window.location.href = res.checkoutUrl;	
							}).catch(err => {
								console.error(err);
							});
						}}
					>
					{({values, isSubmitting, handleSubmit, setFieldValue}) => prices != null ? (
						<Form onSubmit={handleSubmit}>
							<fieldset>
								<h2>Personal Information</h2>
								<Row><Col>
									<FormField label="First Name" name="firstName" type="text" />
								</Col><Col>
									<FormField label="Last Name" name="lastName" type="text" />
								</Col></Row>
								<FormField label="Street Address" name="streetAddress" type="text" />
								<Row><Col>
									<FormField label="City" name="city" type="text" />
								</Col><Col xs={1}>
									<FormSelect label="State" name="state">
										<option aria-label="none" value=""></option>
										<option value="AK">AK</option>
										<option value="AL">AL</option>
										<option value="AR">AR</option>
										<option value="AZ">AZ</option>
										<option value="CA">CA</option>
										<option value="CO">CO</option>
										<option value="CT">CT</option>
										<option value="DE">DE</option>
										<option value="FL">FL</option>
										<option value="GA">GA</option>
										<option value="HI">HI</option>
										<option value="IA">IA</option>
										<option value="ID">ID</option>
										<option value="IL">IL</option>
										<option value="IN">IN</option>
										<option value="KS">KS</option>
										<option value="KY">KY</option>
										<option value="LA">LA</option>
										<option value="MA">MA</option>
										<option value="MD">MD</option>
										<option value="ME">ME</option>
										<option value="MI">MI</option>
										<option value="MN">MN</option>
										<option value="MO">MO</option>
										<option value="MS">MS</option>
										<option value="MT">MT</option>
										<option value="NC">NC</option>
										<option value="ND">ND</option>
										<option value="NE">NE</option>
										<option value="NH">NH</option>
										<option value="NJ">NJ</option>
										<option value="NM">NM</option>
										<option value="NV">NV</option>
										<option value="NY">NY</option>
										<option value="OH">OH</option>
										<option value="OK">OK</option>
										<option value="OR">OR</option>
										<option value="PA">PA</option>
										<option value="RI">RI</option>
										<option value="SC">SC</option>
										<option value="SD">SD</option>
										<option value="TN">TN</option>
										<option value="TX">TX</option>
										<option value="UT">UT</option>
										<option value="VA">VA</option>
										<option value="VT">VT</option>
										<option value="WA">WA</option>
										<option value="WI">WI</option>
										<option value="WV">WV</option>
										<option value="WY">WY</option>
									</FormSelect>
								</Col><Col xs={3}>
									<FormField label="Zip Code" name="zipCode" type="text" />
								</Col></Row>
								<Row><Col>
									<FormField label="Email" name="email" type="email" />
								</Col><Col>
									<FormField label="Home Scene" name="homeScene" type="text" />
								</Col></Row>
								<FormCheck name="isStudent" label="I am a student" />
								<hr />
							</fieldset>
							<fieldset>
								<h2>Purchase</h2>
								<Row><Col xs={6}>
									<FormSelect label="Weekend Pass Type" name="passType">
										<option aria-label="no pass" value={noPassOption} />
										<option value={fullWeekendPassOption}>{"Full Weekend Pass - "+dss.FullWeekendPassTier[prices.weekendPassTier]+" ("+parseDollar(square_data.purchase_items.full_weekend_pass[fromProtoTier(prices.weekendPassTier)])+")"}</option>
										<option value={danceOnlyPassOption}>{"Dance Only Pass ("+parseDollar(square_data.purchase_items.dance_only_pass)+")"}</option>
									</FormSelect>
								</Col></Row>
								{values.passType === fullWeekendPassOption && (
									<Row><Col xs={1}></Col><Col xs={6}>
										<FormSelect label="Level" name="level">
											<option aria-label="none" value={FormFullWeekendPassLevel.NotSelected} />
											<option value={FormFullWeekendPassLevel.Level1}>Level 1</option>
											<option value={FormFullWeekendPassLevel.Level2}>Level 2</option>
											<option value={FormFullWeekendPassLevel.Level3}>Level 3</option>
										</FormSelect>
								</Col></Row>
								)}
								<FormCheck name="mixAndMatch" label={"Mix & Match Competition ("+parseDollar(square_data.purchase_items.mix_and_match)+")"} />
								{values.mixAndMatch && (
									<Row><Col xs={1}></Col><Col xs={6}>
										<FormSelect label="Role" name="role">
											<option aria-label="none" value={FormRole.NotSelected} />
											<option value={FormRole.Follower}>Follower</option>
											<option value={FormRole.Leader}>Leader</option>
										</FormSelect>
									</Col></Row>
								)}
								<FormCheck name="soloJazz" label={"Solo Jazz Competition ("+parseDollar(square_data.purchase_items.solo_jazz)+")"} />
								<FormCheck name="teamCompetition" label={"Team Competition ("+parseDollar(square_data.purchase_items.team_competition)+")"} />
								{values.teamCompetition && (
									<Row><Col xs={1}></Col><Col xs={6}>
										<FormField label="Team Name" name="teamName" type="text" />
									</Col></Row>
								)}
								<FormCheck name="tshirt" label={"T-Shirt ("+parseDollar(square_data.purchase_items.t_shirt)+")"} />
								{values.tshirt && (
									<Row><Col xs={1}></Col><Col xs={6}>
										<FormSelect label="T-Shirt Size/Style" name="style">
											<option aria-label="none" value={FormStyle.NotSelected}></option>
											<option value={FormStyle.UnisexS}>Unisex S</option>
											<option value={FormStyle.UnisexM}>Unisex M</option>
											<option value={FormStyle.UnisexL}>Unisex L</option>
											<option value={FormStyle.UnisexXL}>Unisex XL</option>
											<option value={FormStyle.Unisex2XL}>Unisex 2XL</option>
											<option value={FormStyle.Unisex3XL}>Unisex 3XL</option>
											<option value={FormStyle.BellaS}>Bella S</option>
											<option value={FormStyle.BellaM}>Bella M</option>
											<option value={FormStyle.BellaL}>Bella L</option>
											<option value={FormStyle.BellaXL}>Bella XL</option>
											<option value={FormStyle.Bella2XL}>Bella 2XL</option>
										</FormSelect>
									</Col></Row>
								)}
								<hr />
							</fieldset>
							<fieldset>
								<h2>Housing</h2>
								<Row><Col>
									<FormSelect label="Housing Status" name="housing">
										<option value={noHousingOption}>I neither require nor can provide housing</option>
										<option value={provideOption}>I can provide housing</option>
										<option value={requireOption}>I require housing</option>
									</FormSelect>
								</Col></Row>
								<Row><Col xs={1}></Col><Col>
									{values.housing === "Provide" ? (
										<>
										<FormField label="I have the following pets (cats, dogs, etc)" name="pets" type="text" />
										<FormField label="I can house this many people" name="quantity" type="number" />
										<FormField as="textarea" label="Any other details the organizers should know about my house" name="provideDetails" />
										</>
									) : (values.housing === "Require" && (
										<>
										<FormField label="I am allergic to the following pets" name="petAllergies" type="text" />
										<FormField as="textarea" label="Anything else I would like to say about my housing request" name="requireDetails" />
										</>
									))}
								</Col></Row>
							</fieldset>
							<hr/>
							<fieldset>
								<h2>Discounts</h2>
								<Formik
									initialValues={{
										newDiscount: ""
									}}
									onSubmit={(innerValues, { setSubmitting }) => {
										discountClient.get({
											code: innerValues.newDiscount
										}).then((res) => {
											setFieldValue('discounts', [...values.discounts, res.bundle])
											setSubmitting(false)
										}).catch((e) => {
											console.log(e)
											setSubmitting(false)
										})
									}}
								>
									{(innerProps) => (
										<Row><Col>
											<FormField label="Add Discount Code" name="newDiscount" type="text" />
										</Col><Col xs={1}>
											<br style={{"lineHeight": "200%"}} /><Button disabled={innerProps.isSubmitting} onClick={innerProps.submitForm}>Add</Button>
										</Col></Row>
									)}
								</Formik>
							</fieldset>
							<Button type="submit" disabled={isSubmitting}>Submit</Button>
						</Form>
					) : null}
					</Formik>
				)
			}}
		</Page>
	);
};

export default Registration

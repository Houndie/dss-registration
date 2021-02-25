import React, {useState, useEffect} from 'react'
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import Col from "react-bootstrap/Col"
import Page from '../components/Page'
import FormField from '../components/FormField'
import FormCheck from '../components/FormCheck'
import {Formik} from 'formik'
import {createRegistration} from "../rpc/registration.twirp"
import {dss} from "../rpc/registration.pb"
import {createDiscount} from "../rpc/discount.twirp"
import { v4 as uuidv4 } from 'uuid';
import parseDollar from "../components/parseDollar"

const fullWeekendPassOption = "Full";
const danceOnlyPassOption = "Dance";
const noPassOption = "None";

const provideOption = "Provide";
const requireOption = "Require";
const noHousingOption = "None";

const registrationClient = createRegistration(`${process.env.GATSBY_BACKEND}`)
const discountClient = createDiscount(`${process.env.GATSBY_BACKEND}`)

enum FormFullWeekendPassLevel {
	NotSelected = -1,
	Level1 = dss.FullWeekendPassLevel.Level1,
	Level2 = dss.FullWeekendPassLevel.Level2,
	Level3 = dss.FullWeekendPassLevel.Level3,
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
	NotSelected = -1,
	Follower = dss.MixAndMatch.Role.Follower,
	Leader = dss.MixAndMatch.Role.Leader
}

const toProtoRole = (role: FormRole) => {
	switch(role){
		case FormRole.Leader:
			return dss.MixAndMatch.Role.Leader
		case FormRole.Follower:
			return dss.MixAndMatch.Role.Follower
		default:
			throw "cannot convert unselected role"
	}
}

enum FormStyle {
	NotSelected = -1,
	UnisexS = dss.TShirt.Style.UnisexS,
	UnisexM = dss.TShirt.Style.UnisexM,
	UnisexL = dss.TShirt.Style.UnisexL,
	UnisexXL = dss.TShirt.Style.UnisexXL,
	Unisex2XL = dss.TShirt.Style.Unisex2XL,
	Unisex3XL = dss.TShirt.Style.Unisex3XL,
	BellaS = dss.TShirt.Style.BellaS,
	BellaM = dss.TShirt.Style.BellaM,
	BellaL = dss.TShirt.Style.BellaL,
	BellaXL = dss.TShirt.Style.BellaXL,
	Bella2XL = dss.TShirt.Style.Bella2XL
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

const Registration = () => {
	const [prices, setPrices] = useState<dss.RegistrationPricesRes | null>(null)
	useEffect(() => {
		registrationClient.prices({}).then(res => {
			setPrices(res);
		}, err => {
			console.error(err);
		});
	}, [])

	return (
		<Page title="Registration">
			{(gAuth) => 
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
						const registration: dss.IRegistrationInfo = {
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
								registration.fullWeekendPass = {
									tier: prices.weekendPassTier,
									level: toProtoPassLevel(values.level)
								}
							case danceOnlyPassOption:
								registration.danceOnlyPass = {}
							default:
								registration.noPass = {}
						}

						if (values.mixAndMatch) {
							registration.mixAndMatch = { role: toProtoRole(values.role) }
						}

						if (values.soloJazz) {
							registration.soloJazz = {}
						}

						if (values.teamCompetition) {
							registration.teamCompetition = { name: values.teamName }
						}

						if (values.tshirt) {
							registration.tshirt = { style: toProtoStyle(values.style) }
						}

						switch (values.housing) {
							case provideOption: {
									registration.provideHousing = {
										pets: values.pets,
										quantity: values.quantity,
										details: values.provideDetails
									}
								}
								break
							case requireOption: {
									registration.requireHousing = {
										petAllergies: values.petAllergies,
										details: values.requireDetails,
									}
								}
								break
							default:
								registration.noHousing = {}
						}

						return registrationClient.add({
							idempotencyKey: uuidv4(),
							redirectUrl: `${process.env.GATSBY_FRONTEND}/registration-complete`,
							registration: registration
						}).then(
							res => {
								window.location.href = res.redirectUrl;	
							}, err => {
								console.error(err);
							}
						);
					}}
				>
				{({values, isSubmitting, handleSubmit, setFieldValue}) => prices != null ? (
					<Form onSubmit={handleSubmit}>
						<fieldset>
							<h2>Personal Information</h2>
							<Form.Row><Col>
								<FormField label="First Name" name="firstName" type="text" />
							</Col><Col>
								<FormField label="Last Name" name="lastName" type="text" />
							</Col></Form.Row>
							<FormField label="Street Address" name="streetAddress" type="text" />
							<Form.Row><Col>
								<FormField label="City" name="city" type="text" />
							</Col><Col xs={1}>
								<FormField as="select" label="State" name="state">
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
								</FormField>
							</Col><Col xs={3}>
								<FormField label="Zip Code" name="zipCode" type="text" />
							</Col></Form.Row>
							<Form.Row><Col>
								<FormField label="Email" name="email" type="email" />
							</Col><Col>
								<FormField label="Home Scene" name="homeScene" type="text" />
							</Col></Form.Row>
							<FormCheck name="isStudent" label="I am a student" />
							<hr />
						</fieldset>
						<fieldset>
							<h2>Purchase</h2>
							<Form.Row><Col xs={6}>
								<FormField as="select" label="Weekend Pass Type" name="passType">
									<option aria-label="no pass" value={noPassOption} />
									<option value={fullWeekendPassOption}>{"Full Weekend Pass - "+dss.FullWeekendPassTier[prices.weekendPassTier]+" ("+parseDollar(prices.weekendPassCost)+")"}</option>
									<option value={danceOnlyPassOption}>{"Dance Only Pass ("+parseDollar(prices.dancePassCost)+")"}</option>
								</FormField>
							</Col></Form.Row>
							{values.passType === fullWeekendPassOption && (
								<Form.Row><Col xs={1}></Col><Col xs={6}>
									<FormField as="select" label="Level" name="level">
										<option aria-label="none" value={FormFullWeekendPassLevel.NotSelected} />
										<option value={FormFullWeekendPassLevel.Level1}>Level 1</option>
										<option value={FormFullWeekendPassLevel.Level2}>Level 2</option>
										<option value={FormFullWeekendPassLevel.Level3}>Level 3</option>
									</FormField>
							</Col></Form.Row>
							)}
							<FormCheck name="mixAndMatch" label={"Mix & Match Competition ("+parseDollar(prices.mixAndMatchCost)+")"} />
							{values.mixAndMatch && (
								<Form.Row><Col xs={1}></Col><Col xs={6}>
									<FormField as="select" label="Role" name="role">
										<option aria-label="none" value={FormRole.NotSelected} />
										<option value={FormRole.Follower}>Follower</option>
										<option value={FormRole.Leader}>Leader</option>
									</FormField>
								</Col></Form.Row>
							)}
							<FormCheck name="soloJazz" label={"Solo Jazz Competition ("+parseDollar(prices.soloJazzCost)+")"} />
							<FormCheck name="teamCompetition" label={"Team Competition ("+parseDollar(prices.teamCompetitionCost)+")"} />
							{values.teamCompetition && (
								<Form.Row><Col xs={1}></Col><Col xs={6}>
									<FormField label="Team Name" name="teamName" type="text" />
								</Col></Form.Row>
							)}
							<FormCheck name="tshirt" label={"T-Shirt ("+parseDollar(prices.tshirtCost)+")"} />
							{values.tshirt && (
								<Form.Row><Col xs={1}></Col><Col xs={6}>
									<FormField as="select" label="T-Shirt Size/Style" name="style">
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
									</FormField>
								</Col></Form.Row>
							)}
							<hr />
						</fieldset>
						<fieldset>
							<h2>Housing</h2>
							<Form.Row><Col>
								<FormField as="select" label="Housing Status" name="housing">
									<option value={noHousingOption}>I neither require nor can provide housing</option>
									<option value={provideOption}>I can provide housing</option>
									<option value={requireOption}>I require housing</option>
								</FormField>
							</Col></Form.Row>
							<Form.Row><Col xs={1}></Col><Col>
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
							</Col></Form.Row>
						</fieldset>
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
									<Form.Row><Col>
										<FormField label="Add Discount Code" name="newDiscount" type="text" />
									</Col><Col xs={1}>
										<br style={{"lineHeight": "200%"}} /><Button disabled={innerProps.isSubmitting} onClick={innerProps.submitForm}>Add</Button>
									</Col></Form.Row>
								)}
							</Formik>
						</fieldset>
						<Button type="submit" disabled={isSubmitting}>Submit</Button>
					</Form>
				) : null}
				</Formik>
			}
		</Page>
	);
};

export default Registration

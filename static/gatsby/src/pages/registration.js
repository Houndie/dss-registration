import React, {useState, useEffect} from 'react'
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import Col from "react-bootstrap/Col"
import Page from '../components/Page.js'
import FormField from '../components/FormField.js'
import FormCheck from '../components/FormCheck.js'
import {Formik} from 'formik'
import registrationTwirp from "../rpc/registration_pb_twirp.js"
import discountTwirp from "../rpc/discount_pb_twirp.js"
import { v4 as uuidv4 } from 'uuid';

const fullWeekendPassOption = "Full";
const danceOnlyPassOption = "Dance";
const noPassOption = "None";

const provideOption = "Provide";
const requireOption = "Require";
const noHousingOption = "None";

const parseDollar = (intCost) => {
	let dollar = intCost.toString()
	while(dollar.length < 3) {
		dollar = "0" + dollar;
	}
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

const tierToString = (tier) => {
	switch (tier) {
	case registrationTwirp.FullWeekendPassTier.TIER1:
		return "Tier 1";
	case registrationTwirp.FullWeekendPassTier.TIER2:
		return "Tier 2";
	case registrationTwirp.FullWeekendPassTier.TIER3:
		return "Tier 3";
	case registrationTwirp.FullWeekendPassTier.TIER4:
		return "Tier 4";
	case registrationTwirp.FullWeekendPassTier.TIER5:
		return "Tier 5";
	default:
		return "";
	}
}

const registrationClient = registrationTwirp.createRegistrationClient(`${process.env.GATSBY_BACKEND}`)
const discountClient = discountTwirp.createDiscountClient(`${process.env.GATSBY_BACKEND}`)

const Registration = () => {
	const [prices, setPrices] = useState(null)
	useEffect(() => {
		registrationClient.prices(new registrationTwirp.RegistrationPricesReq()).then(res => {
			setPrices(res);
		}, err => {
			console.error(err);
		});
	}, [])

	return (
		<Page title="Registration">
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
					level: '',
					mixAndMatch: false,
					role: '',
					soloJazz: false,
					teamCompetition: false,
					teamName: '',
					tshirt: false,
					style: '',
					housing: noHousingOption,
					pets: '',
					quantity: 0,
					provideDetails: '',
					petAllergies: '',
					requireDetails: '',
					discounts: []
				}}
				onSubmit={(values, { setSubmitting }) => {

					const r = new registrationTwirp.RegistrationInfo();
					r.setFirstName(values.firstName);
					r.setLastName(values.lastName);
					r.setStreetAddress(values.streetAddress);
					r.setCity(values.city);
					r.setState(values.state);
					r.setZipCode(values.zipCode);
					r.setEmail(values.email);
					r.setHomeScene(values.homeScene);
					r.setIsStudent(values.isStudent);

					switch (values.passType) {
						case fullWeekendPassOption:
							const f = new registrationTwirp.FullWeekendPass();
							f.setTier(prices.weekendPassTier);
							f.setLevel(parseInt(values.level));
							r.setFullWeekendPass(f);
							break
						case danceOnlyPassOption:
							r.setDanceOnlyPass(new registrationTwirp.DanceOnlyPass());
							break
						default:
							r.setNoPass(new registrationTwirp.NoPass());
					}

					if (values.mixAndMatch) {
						const m = new registrationTwirp.MixAndMatch();
						m.setRole(parseInt(values.role));
						r.setMixAndMatch(m);
					}

					if (values.soloJazz) {
						r.setSoloJazz(new registrationTwirp.SoloJazz());
					}

					if (values.teamCompetition) {
						const t = new registrationTwirp.TeamCompetition();
						t.setName(values.teamName);
						r.setTeamCompetition(t);
					}

					if (values.tshirt) {
						const t = new registrationTwirp.TShirt();
						t.setStyle(parseInt(values.style));
						r.setTshirt(t);
					}

					switch (values.housing) {
						case provideOption: {
								const h = new registrationTwirp.ProvideHousing();
								h.setPets(values.pets);
								h.setQuantity(values.quantity);
								h.setDetails(values.provideDetails);
								r.setProvideHousing(h);
							}
							break
						case requireOption: {
								const h = new registrationTwirp.RequireHousing();
								h.setPetAllergies(values.petAllergies);
								h.setDetails(values.requireDetails);
								r.setRequireHousing(h);
							}
							break
						default:
							r.setNoHousing(new registrationTwirp.NoHousing());
					}

					const req = new registrationTwirp.RegistrationAddReq();
					req.setIdempotencyKey(uuidv4());
					req.setRedirectUrl(`${process.env.GATSBY_FRONTEND}/registration-complete`);
					req.setRegistration(r);

					return registrationClient.add(req).then(
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
								<option value={fullWeekendPassOption}>{"Full Weekend Pass - "+tierToString(prices.weekendPassTier)+" ("+parseDollar(prices.weekendPassCost)+")"}</option>
								<option value={danceOnlyPassOption}>{"Dance Only Pass ("+parseDollar(prices.dancePassCost)+")"}</option>
							</FormField>
						</Col></Form.Row>
						{values.passType === fullWeekendPassOption && (
							<Form.Row><Col xs={1}></Col><Col xs={6}>
								<FormField as="select" label="Level" name="level">
									<option aria-label="none" value="" />
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL1}>Level 1</option>
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL2}>Level 2</option>
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL3}>Level 3</option>
								</FormField>
						</Col></Form.Row>
						)}
						<FormCheck name="mixAndMatch" label={"Mix & Match Competition ("+parseDollar(prices.mixAndMatchCost)+")"} />
						{values.mixAndMatch && (
							<Form.Row><Col xs={1}></Col><Col xs={6}>
								<FormField as="select" label="Role" name="role">
									<option aria-label="none" value="" />
									<option value={registrationTwirp.MixAndMatch.Role.FOLLOWER}>Follower</option>
									<option value={registrationTwirp.MixAndMatch.Role.LEADER}>Leader</option>
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
									<option aria-label="none" value=''></option>
									<option value={registrationTwirp.TShirt.Style.UNISEXS}>Unisex S</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXM}>Unisex M</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXL}>Unisex L</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXXL}>Unisex XL</option>
									<option value={registrationTwirp.TShirt.Style.UNISEX2XL}>Unisex 2XL</option>
									<option value={registrationTwirp.TShirt.Style.UNISEX3XL}>Unisex 3XL</option>
									<option value={registrationTwirp.TShirt.Style.BELLAS}>Bella S</option>
									<option value={registrationTwirp.TShirt.Style.BELLAM}>Bella M</option>
									<option value={registrationTwirp.TShirt.Style.BELLAL}>Bella L</option>
									<option value={registrationTwirp.TShirt.Style.BELLAXL}>Bella XL</option>
									<option value={registrationTwirp.TShirt.Style.BELLA2XL}>Bella 2XL</option>
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
								console.log("A1")
								const d = new discountTwirp.DiscountGetReq()
								d.AddCode(innerValues.newDiscount)
								console.log("A2")
								discountClient.Get(d).then((res) => {
								console.log("A3")
									setFieldValue('discounts', [...values.discounts, res.bundle])
									setSubmitting(false)
								}).catch((e) => {
								console.log("A4")
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
		</Page>
	);
};

export default Registration

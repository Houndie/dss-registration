import React, {useState, useEffect} from 'react'
import {dss as twirpRegistration} from "../rpc/registration.pb"
import {dss as twirpVaccine} from "../rpc/vaccine.pb"
import parseDollar from "../components/parseDollar"
import Form from "react-bootstrap/Form"
import {Formik, useFormikContext} from 'formik'
import Button from "react-bootstrap/Button"
import Col from "react-bootstrap/Col"
import Row from "react-bootstrap/Row"
import FormField from '../components/FormField'
import FormFile from '../components/FormFile'
import FormSelect from '../components/FormSelect'
import FormCheck from '../components/FormCheck'
import useTwirp from "../components/useTwirp"
import {VaccineInfoEnum, VaccineInfo, fromProtoVaccine} from "../components/vaccine"

export type RegistrationFormState = {
	firstName: string, 
	lastName: string,
	streetAddress: string,
	city: string,
	state: string,
	zipCode: string,
	email: string,
	homeScene: string,
	isStudent: boolean,
	passType: FormWeekendPassOption,
	level: FormFullWeekendPassLevel,
	weekendPassOverride: boolean,
	danceOnlyOverride: boolean,
	mixAndMatch: boolean,
	role: FormRole,
	mixAndMatchOverride: boolean,
	soloJazz: boolean,
	soloJazzOverride: boolean,
	teamCompetition: boolean,
	teamName: string,
	teamCompetitionOverride: boolean,
	tshirt: boolean,
	style: FormStyle,
	tshirtOverride: boolean,
	housing: FormHousingOption,
	pets: string,
	quantity: number | Long,
	provideDetails: string,
	petAllergies: string,
	requireDetails: string,
	vaccine: File|undefined,
	discounts: string[]
}

export const toProtoRegistration = (values: RegistrationFormState, tier: number, previous?: twirpRegistration.IRegistrationInfo) => {
	const clientReg: twirpRegistration.IRegistrationInfo = {
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
		case FormWeekendPassOption.fullWeekendPassOption:
			clientReg.fullWeekendPass = {
				tier: tier,
				level: toProtoPassLevel(values.level),
				squarePaid: previous?.fullWeekendPass?.squarePaid,
				adminPaymentOverride: values.weekendPassOverride
			}
			break
		case FormWeekendPassOption.danceOnlyPassOption:
			clientReg.danceOnlyPass = {
				squarePaid: previous?.danceOnlyPass?.squarePaid,
				adminPaymentOverride: values.weekendPassOverride
			}
			break
		default:
			clientReg.noPass = {}
			break
	}

	if (values.mixAndMatch) {
		clientReg.mixAndMatch = { 
			role: toProtoRole(values.role),
			squarePaid: previous?.mixAndMatch?.squarePaid,
			adminPaymentOverride: values.mixAndMatchOverride
		}
	}

	if (values.soloJazz) {
		clientReg.soloJazz = {
			squarePaid: previous?.soloJazz?.squarePaid,
			adminPaymentOverride: values.soloJazzOverride
		}
	}

	if (values.teamCompetition) {
		clientReg.teamCompetition = { 
			name: values.teamName,
			squarePaid: previous?.teamCompetition?.squarePaid,
			adminPaymentOverride: values.teamCompetitionOverride
		}
	}

	if (values.tshirt) {
		clientReg.tshirt = { 
			style: toProtoStyle(values.style),
			squarePaid: previous?.tshirt?.squarePaid,
			adminPaymentOverride: values.tshirtOverride
		}
	}

	switch (values.housing) {
		case FormHousingOption.provideOption: {
				clientReg.provideHousing = {
					pets: values.pets,
					quantity: values.quantity,
					details: values.provideDetails
				}
			}
			break
		case FormHousingOption.requireOption: {
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

	return clientReg
}

export enum FormWeekendPassOption {
	fullWeekendPassOption = "Full",
	danceOnlyPassOption = "Dance",
	noPassOption = "None"
}

export const formWeekendPassOptionFromProto = (r: twirpRegistration.IRegistrationInfo) => {
	if(r.fullWeekendPass) {
		return FormWeekendPassOption.fullWeekendPassOption
	}
	if(r.danceOnlyPass) {
		return FormWeekendPassOption.danceOnlyPassOption
	}
	if(r.noPass) {
		return FormWeekendPassOption.noPassOption
	}

	throw "No weekend pass type found"
}

export enum FormHousingOption {
	provideOption = "Provide",
	requireOption = "Require",
	noHousingOption = "None"
}

export const fromProtoHousingOption = (r: twirpRegistration.IRegistrationInfo) => {
	if(r.provideHousing) {
		return FormHousingOption.provideOption
	}

	if(r.requireHousing) {
		return FormHousingOption.requireOption
	}

	if(r.noHousing) {
		return FormHousingOption.noHousingOption
	}

	throw 'unable to convert housing option'
}

export enum FormFullWeekendPassLevel {
	NotSelected = "not_selected",
	Level1 = "level_1",
	Level2 = "level_2",
	Level3 = "level_3"
}

export const fromProtoPassLevel = (level: twirpRegistration.FullWeekendPassLevel) => {
	switch(level){
		case twirpRegistration.FullWeekendPassLevel.Level1:
			return FormFullWeekendPassLevel.Level1
		case twirpRegistration.FullWeekendPassLevel.Level2:
			return FormFullWeekendPassLevel.Level2
		case twirpRegistration.FullWeekendPassLevel.Level3:
			return FormFullWeekendPassLevel.Level3
		default:
			throw 'cannot convert level to form type'
	}
}

export const toProtoPassLevel = (level: FormFullWeekendPassLevel) => {
	switch(level){
		case FormFullWeekendPassLevel.Level1:
			return twirpRegistration.FullWeekendPassLevel.Level1
		case FormFullWeekendPassLevel.Level2:
			return twirpRegistration.FullWeekendPassLevel.Level2
		case FormFullWeekendPassLevel.Level3:
			return twirpRegistration.FullWeekendPassLevel.Level3
		default:
			throw "cannot convert unselected pass level"
	}
}

export enum FormRole {
	NotSelected = "not_selected",
	Follower = "follower",
	Leader = "leader"
}

export const toProtoRole = (role: FormRole) => {
	switch(role){
		case FormRole.Leader:
			return twirpRegistration.MixAndMatch.Role.Leader
		case FormRole.Follower:
			return twirpRegistration.MixAndMatch.Role.Follower
		default:
			throw "cannot convert unselected role"
	}
}

export const fromProtoRole = (role: twirpRegistration.MixAndMatch.Role) => {
	switch(role){
		case twirpRegistration.MixAndMatch.Role.Leader:
			return FormRole.Leader
		case twirpRegistration.MixAndMatch.Role.Follower:
			return FormRole.Follower
		default:
			throw "cannot convert proto role"
	}
}

export enum FormStyle {
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

export const toProtoStyle = (style: FormStyle) => {
	switch(style){
		case FormStyle.UnisexS:
			return twirpRegistration.TShirt.Style.UnisexS
		case FormStyle.UnisexM:
			return twirpRegistration.TShirt.Style.UnisexM
		case FormStyle.UnisexL:
			return twirpRegistration.TShirt.Style.UnisexL
		case FormStyle.UnisexXL:
			return twirpRegistration.TShirt.Style.UnisexXL
		case FormStyle.Unisex2XL:
			return twirpRegistration.TShirt.Style.Unisex2XL
		case FormStyle.Unisex3XL:
			return twirpRegistration.TShirt.Style.Unisex3XL
		case FormStyle.BellaS:
			return twirpRegistration.TShirt.Style.BellaS
		case FormStyle.BellaM:
			return twirpRegistration.TShirt.Style.BellaM
		case FormStyle.BellaL:
			return twirpRegistration.TShirt.Style.BellaL
		case FormStyle.BellaXL:
			return twirpRegistration.TShirt.Style.BellaXL
		case FormStyle.Bella2XL:
			return twirpRegistration.TShirt.Style.Bella2XL
		default:
			throw "cannot convert unselected style"
	}
}

export const fromProtoStyle = (style: twirpRegistration.TShirt.Style) => {
	switch(style) {
		case twirpRegistration.TShirt.Style.UnisexS:
			return FormStyle.UnisexS
		case twirpRegistration.TShirt.Style.UnisexM:
			return FormStyle.UnisexM
		case twirpRegistration.TShirt.Style.UnisexL:
			return FormStyle.UnisexL
		case twirpRegistration.TShirt.Style.UnisexXL:
			return FormStyle.UnisexXL
		case twirpRegistration.TShirt.Style.Unisex2XL:
			return FormStyle.Unisex2XL
		case twirpRegistration.TShirt.Style.Unisex3XL:
			return FormStyle.Unisex3XL
		case twirpRegistration.TShirt.Style.BellaS:
			return FormStyle.BellaS
		case twirpRegistration.TShirt.Style.BellaM:
			return FormStyle.BellaM
		case twirpRegistration.TShirt.Style.BellaL:
			return FormStyle.BellaL
		case twirpRegistration.TShirt.Style.BellaXL:
			return FormStyle.BellaXL
		case twirpRegistration.TShirt.Style.Bella2XL:
			return FormStyle.Bella2XL
		default:
			throw 'unable to convert proto style'
	}
}

export const fromProtoTier = (tier: number | Long) => {
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

type RegistrationFormProps = {
	weekendPassTier: number
	previousRegistration?: twirpRegistration.IRegistrationInfo
	admin: boolean
	vaccineUpload?: VaccineInfo
	vaccineRef?: React.MutableRefObject<HTMLInputElement|undefined>
	setMyVaccine?: (arg0: VaccineInfo) => void
}

export const isPaid = (r: twirpRegistration.IRegistrationInfo) => 
	((!r.fullWeekendPass || r.fullWeekendPass.squarePaid || r.fullWeekendPass.adminPaymentOverride) &&
		(!r.danceOnlyPass || r.danceOnlyPass.squarePaid || r.danceOnlyPass.adminPaymentOverride) &&
		(!r.soloJazz || r.soloJazz.squarePaid || r.soloJazz.adminPaymentOverride) &&
		(!r.mixAndMatch || r.mixAndMatch.squarePaid || r.mixAndMatch.adminPaymentOverride) &&
		(!r.teamCompetition || r.teamCompetition.squarePaid || r.teamCompetition.adminPaymentOverride) &&
		(!r.tshirt || r.tshirt.squarePaid || r.tshirt.adminPaymentOverride))

export default ({weekendPassTier, previousRegistration, admin, vaccineUpload, vaccineRef, setMyVaccine}: RegistrationFormProps) => {
	const {discount} = useTwirp()
	const square_data = JSON.parse(`${process.env.GATSBY_SQUARE_DATA}`)
	const {values, isSubmitting, handleSubmit, setFieldValue} = useFormikContext<RegistrationFormState>()

	if(weekendPassTier === null || weekendPassTier === undefined) {
		return null
	}

	return (
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
					<FormSelect label="Weekend Pass Type" name="passType" disabled={Boolean(previousRegistration && (previousRegistration.fullWeekendPass || previousRegistration.danceOnlyPass)) }>
						<option aria-label="no pass" value={FormWeekendPassOption.noPassOption} />
						<option value={FormWeekendPassOption.fullWeekendPassOption}>{"Full Weekend Pass - "+twirpRegistration.FullWeekendPassTier[weekendPassTier]+" ("+parseDollar(square_data.purchase_items.full_weekend_pass[fromProtoTier(weekendPassTier)])+")"}</option>
						<option value={FormWeekendPassOption.danceOnlyPassOption}>{"Dance Only Pass ("+parseDollar(square_data.purchase_items.dance_only_pass)+")"}</option>
					</FormSelect>
				</Col></Row>
				{values.passType === FormWeekendPassOption.fullWeekendPassOption && (
					<>
						<Row><Col xs={1}></Col><Col xs={6}>
							<FormSelect label="Level" name="level">
								<option aria-label="none" value={FormFullWeekendPassLevel.NotSelected} />
								<option value={FormFullWeekendPassLevel.Level1}>Level 1</option>
								<option value={FormFullWeekendPassLevel.Level2}>Level 2</option>
								<option value={FormFullWeekendPassLevel.Level3}>Level 3</option>
							</FormSelect>
						</Col></Row>
						{admin && (
							<Row><Col xs={1}></Col><Col xs={6}>
								<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.fullWeekendPass && previousRegistration.fullWeekendPass.squarePaid).toString()}</p>
								<FormCheck label="Admin Free Pass" name="weekendPassOverride" />
							</Col></Row>
						)}
					</>
				)}
				{values.passType == FormWeekendPassOption.danceOnlyPassOption && admin && (
					<Row><Col xs={1}></Col><Col xs={6}>
						<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.danceOnlyPass && previousRegistration.danceOnlyPass.squarePaid).toString()}</p>
						<FormCheck label="Admin Free Pass" name="danceOnlyOverride" />
					</Col></Row>
				)}
				<FormCheck name="mixAndMatch" label={"Mix & Match Competition ("+parseDollar(square_data.purchase_items.mix_and_match)+")"} disabled={Boolean(previousRegistration && previousRegistration.mixAndMatch)}/>
				{values.mixAndMatch && (
					<>
						<Row><Col xs={1}></Col><Col xs={6}>
							<FormSelect label="Role" name="role">
								<option aria-label="none" value={FormRole.NotSelected} />
								<option value={FormRole.Follower}>Follower</option>
								<option value={FormRole.Leader}>Leader</option>
							</FormSelect>
						</Col></Row>
						{admin && (
							<Row><Col xs={1}></Col><Col xs={6}>
								<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.mixAndMatch && previousRegistration.mixAndMatch.squarePaid).toString()}</p>
								<FormCheck label="Admin Free Pass" name="mixAndMatchOverride" />
							</Col></Row>
						)}
					</>
				)}
				<FormCheck name="soloJazz" label={"Solo Jazz Competition ("+parseDollar(square_data.purchase_items.solo_jazz)+")"} disabled={Boolean(previousRegistration && previousRegistration.soloJazz)}/>
				{values.soloJazz && admin && (
					<Row><Col xs={1}></Col><Col xs={6}>
						<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.soloJazz && previousRegistration.soloJazz.squarePaid).toString()}</p>
						<FormCheck label="Admin Free Pass" name="soloJazzOverride" />
					</Col></Row>
				)}
				<FormCheck name="teamCompetition" label={"Team Competition ("+parseDollar(square_data.purchase_items.team_competition)+")"} disabled={Boolean(previousRegistration && previousRegistration.teamCompetition)} />
				{values.teamCompetition && (
					<>
						<Row><Col xs={1}></Col><Col xs={6}>
							<FormField label="Team Name" name="teamName" type="text" />
						</Col></Row>
						{admin && (
							<Row><Col xs={1}></Col><Col xs={6}>
								<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.teamCompetition && previousRegistration.teamCompetition.squarePaid).toString()}</p>
								<FormCheck label="Admin Free Pass" name="teamCompetitionOverride" />
							</Col></Row>
						)}
					</>
				)}
				<FormCheck name="tshirt" label={"T-Shirt ("+parseDollar(square_data.purchase_items.t_shirt)+")"} disabled={Boolean(previousRegistration && previousRegistration.tshirt)} />
				{values.tshirt && (
					<>
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
						{admin && (
							<Row><Col xs={1}></Col><Col xs={6}>
								<p>Currently Paid: {Boolean(previousRegistration && previousRegistration.tshirt && previousRegistration.tshirt.squarePaid).toString()}</p>
								<FormCheck label="Admin Free Pass" name="tshirtOverride" />
							</Col></Row>
						)}
					</>
				)}
				<hr />
			</fieldset>
			<fieldset>
				<h2>Housing</h2>
				<Row><Col>
					<FormSelect label="Housing Status" name="housing">
						<option value={FormHousingOption.noHousingOption}>I neither require nor can provide housing</option>
						<option value={FormHousingOption.provideOption}>I can provide housing</option>
						<option value={FormHousingOption.requireOption}>I require housing</option>
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
						discount().then(client => {
							return client.get({
								code: innerValues.newDiscount
							})
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
			<hr/>
			<fieldset>
				<h2>Vaccination Card</h2>
				<VaccineBlock 
					registrationID={previousRegistration?.id} 
					vaccineUpload={vaccineUpload} 
					vaccineRef={vaccineRef}
					setMyVaccine={setMyVaccine}
					admin={admin}
				/>
			</fieldset>
			<hr />
			<h2>Submit Registration</h2>
			<Button type="submit" disabled={isSubmitting}>Submit</Button>
		</Form>
	)
}

type VaccineBlockProps = {
	registrationID: string | null | undefined
	vaccineUpload?: VaccineInfo
	vaccineRef?: React.MutableRefObject<HTMLInputElement|undefined>
	setMyVaccine?: (arg0: VaccineInfo) => void
	admin?: boolean
}

enum VaccineBlockFormAction {
	approveAction = "approve",
	rejectAction = "reject",
	noAction = "none"
}

const VaccineBlock = ({registrationID, vaccineUpload, vaccineRef, setMyVaccine, admin}: VaccineBlockProps) => {
	const {vaccine} = useTwirp()
	const approve = () => {
		vaccine().then(client => {
			if(!registrationID) {
				throw "no registration id"
			}

			return client.approve({
				id: registrationID
			}).then(() => {
				if(!registrationID) {
					throw "no registration id"
				}
				
				return client.get({
					id: registrationID
				})
			})
		}).then(newVaccine => {
			if(!setMyVaccine) {
				throw "vaccine setter not set"
			}

			setMyVaccine(fromProtoVaccine(newVaccine))
		}).catch(err => {
			console.log(err)
		})
	}

	const reject = (reason: string) => {
		vaccine().then(client => {
			if(!registrationID) {
				throw "no registration id"
			}

			return client.reject({
				id: registrationID,
				reason: reason
			}).then(() => {
				if(!registrationID) {
					throw "no registration id"
				}
				
				return client.get({
					id: registrationID
				})
			})
		}).then(newVaccine => {
			if(!setMyVaccine) {
				throw "vaccine setter not set"
			}

			setMyVaccine(fromProtoVaccine(newVaccine))
		}).catch(err => {
			console.log(err)
		})
	}

	return (
		<>
			{
				(() => {
					if(!vaccineUpload || vaccineUpload.type === VaccineInfoEnum.NoVaxProofSuppliedEnum) {
						return (
							<>
								<p>Dayton Swing Smackdown is an event only for dancers who have the Covid-19 Vaccine.  You can upload an image of your vaccine card now in order to provide proof of vaccination.  This image will only be stored as long as it takes to verify your vaccine information, and will be subsequently deleted.</p>
								<p>If you do not wish to upload an image of your vaccine card, or do not yet have one, you can also present your vaccine information at the door.  Keep in mind that if you do not have your card (or an image of your card) you will be turned away</p>
								<FormFile label="Add Vaccine Card" name="vaccine" />
							</>
						)
					}


					if (vaccineUpload.type === VaccineInfoEnum.VaxApprovalPendingEnum) {
						return (
							<>
								<p>Your vaccine information is uploaded and pending approval.</p>
								<a href={vaccineUpload.URL}><img src={vaccineUpload.URL} width={300} /></a>
								<p>Upload different information:</p>
								<FormFile label="Add Vaccine Card" name="vaccine" myref={vaccineRef}/>
							</>
						)
					}

					return <p>Your vaccine information has been approved!</p>
				})()
			}
			{admin && (
				<Formik
					initialValues={{
						reason: "",
						action: VaccineBlockFormAction.noAction
					}}
					onSubmit={(values, {setSubmitting, setFieldValue}) => {
						switch(values.action) {
						case VaccineBlockFormAction.approveAction:
							approve()
							break;
						case VaccineBlockFormAction.rejectAction:
							reject(values.reason)
							break;
						default:
							console.log( "no action selected")
						}

						setFieldValue("reason", "")
						setFieldValue("action", VaccineBlockFormAction.noAction)
						setSubmitting(false)
					}} 
				>
					{({values, isSubmitting, submitForm, setFieldValue}) => {
						useEffect(() => {
							if(values.action != VaccineBlockFormAction.noAction){
								submitForm()
							}
						}, [values.action])

						return (
							<>
								<br/>
								<FormField as="textarea" label="Rejection Reason" name="reason" />
								<br/>
								<Row><Col>
									<Button disabled={isSubmitting} onClick={() => {
										setFieldValue("action", VaccineBlockFormAction.approveAction)
									}}>Approve</Button>
								</Col><Col>
									<Button disabled={isSubmitting} onClick={() => {
										setFieldValue("action", VaccineBlockFormAction.rejectAction)
									}}>Reject</Button>
								</Col></Row>
							</>
						)
					}}
				</Formik>
			)}
		</>
	)
}

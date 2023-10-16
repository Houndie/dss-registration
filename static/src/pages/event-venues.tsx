import React from 'react'
import Page from '../components/Page'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Image from 'react-bootstrap/Image'

interface VenueProps {
	title: string
	address: string[]
	days: string[]
	google: string
	children: React.ReactNode
	image: string
}

const Venue = ({ title, address, days, google, children, image }: VenueProps) => (
	<>
		<h2>{title}</h2>
		<Container>
			<Row><Col>
				<ul>{days.map((day, idx) => (<li key={idx}>{day}</li>))}</ul>
				<p>{address.map(a => <>{a}</>).reduce((accumulator, currentValue) => (<>{accumulator}<br />{currentValue}</>))}</p>
				{children}
			</Col><Col>
					<Image src={image} width="200px" />
				</Col></Row>
		</Container>
		<iframe title={title} className="embed-responsive-item" src={google} width="800" height="450" frameBorder="0" style={{ border: 0 }} allowFullScreen />
	</>
)

export default () => (
	<Page title="Event Venues">
		<Venue
			title="The Baum Opera House"
			days={[
				"Friday Night Dance",
				"Saturday Afternoon Lessons",
				"Saturday Night Dance"
			]}
			address={["15 South 1st Street", "Miamisburg, OH"]}
			image="/images/baum.jpg"
			google="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d12289.462909795679!2d-84.286574!3d39.641483!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x8840628c7d6aa843%3A0xc8ce9bb439e95bcd!2s15+S+1st+St%2C+Miamisburg%2C+OH+45342!5e0!3m2!1sen!2sus!4v1560288685713!5m2!1sen!2sus"
		>
			<p>The Baum Opera House is a registered National Historic Site. Built in 1884, it features two levels, the top level having a large ballroom and stage. The lower level features a smaller floor for a more intimate setting.  Every year it is host to theater production, teas, weddings and wedding receptions, Christmas parties, musical events and many other activities and functions.  This year the Baum Opera House will be host to two of our 3 level tracks, as well as the Friday Night and Saturday Night dances.</p>
			<p>Parking is available adjacent to the building, as well as across the street.</p>
		</Venue>
		<Venue
			title="Star City Social Hall"
			days={[
				"Saturday Afternoon Lessons",
			]}
			address={["16 Water St", "Miamisburg, OH"]}
			image="/images/scs.jpg"
			google="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3072.3608818129837!2d-84.29236282402715!3d39.641592071574635!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x884063ff4ea128f5%3A0xda4d03b11b301ea2!2sStar%20City%20Social!5e0!3m2!1sen!2sus!4v1697157779434!5m2!1sen!2sus"
		>
			<p>Our former venue of B&B Riverfont Hall is under new management and is now Star City Social Hall!  Located one block from the opera house, this venue is perfect for our second instruction track!  Entrances to the ballroom can be found on both Water Street and Main Street. Both doorways have a sign for the ballroom.</p>
		</Venue>
		<Venue
			title="Elegance In Dance"
			days={[
				"Friday Late Night",
				"Saturday Late Night",
				"Sunday Afternoon"
			]}
			address={["8967 Kingsridge Drive", "Miamisburg, OH 45342"]}
			image="/images/elegance.jpg"
			google="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d12290.777643719184!2d-84.2169!3d39.634084!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x88408901f3db4d17%3A0x3e1e66e57e673ed4!2s8967+Kingsridge+Dr%2C+Dayton%2C+OH+45458!5e0!3m2!1sen!2sus!4v1560292661280!5m2!1sen!2sus>"
		>
			<p>Elegance in Dance has been a staple in the Dayton swing dance scene for years.  With a large, sprung hardwood floor, a raised DJ booth and a bar (BYOB), it’s the perfect intimate setting. For more information, <a href="http://www.eleganceindance.com/">check out their website</a>. The studio is located at the end of the strip mall. Near Marion’s Pizza.</p>
			<p>Parking is available in front of the building.</p>
		</Venue>
	</Page>
)

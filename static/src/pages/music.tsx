import React from 'react'
import Page from '../components/Page'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

export default () => (
	<Page title="Music">
		{() => <>
			<h2>Live Band:  The Midwest All-Stars</h2>
			<Row className="align-items-center"><Col className="col-4"><img src="/images/Band_Flyer.png" width="100%" object-fit="scale-down"/></Col></Row>
			<h2>DJs</h2>
			<h3>Dave Berry</h3>
			<Row className="align-items-center">
				<Col className="col-4"><img src="/images/DJs/Dave.png" width="100%" object-fit="scale-down"/></Col>
				<Col><p>A staple in the Southeast since 2015, Dave is constantly searching through thrift stores, music shops, and the bowels of the internet to find new (old) tunes to bring into the community. In the booth his sets have “oomph”— starting with 30's era swing and ending with a groovy twist. This usually means a hefty amount of Johnny Hodges, Billie Holiday, Basie's Decca years and an impressive variety of rare and lesser known artists that Dave finds through his side hustle, The Jazz Garden. At his apartment in Atlanta, Dave gives warm homes to abandoned and forgotten vinyl records and CDs, and has serious opinions about his fantasy Big Band lineup.</p></Col>
			</Row>
			<h3>Emma Durham</h3>
			<Row className="align-items-center">
				<Col><p>Emma Durham started swing dancing at a college club at Savannah College of Art & Design in Georgia in 2008 and fell in love with the movement and then the music. She continued while living in Atlanta and New Orleans before moving back to her hometown of Cincinnati, where she started DJing at local and regional events. She has a soft spot for drummer-led bands, female vocalists, and anything that makes her want to swing out.</p></Col>
				<Col className="col-4"><img src="/images/DJs/Emma.jpg" width="100%" object-fit="scale-down"/></Col>
			</Row>
			<h3>Ann Sychterz</h3>
			<Row className="align-items-center">
				<Col className="col-4"><img src="/images/DJs/Ann.jpg" width="100%" object-fit="scale-down"/></Col>
				<Col><p>Ann began swing dancing in 2008 in Waterloo, Canada (not far from Toronto) and now lives in the US. Since getting her first songs, she has DJed in her home scenes including the one she founded in Switzerland called Swingtime Lausanne. She has had the fortune of playing swing jazz including head and competition DJ for events around the world. Digging through old stuff and new (soft spot for the original stuff), small group to big band to Newport, North American and European, she aims to create an atmosphere worthy of creativity and playing around with the dance.</p></Col>
			</Row>
			<h3>Aliceann Talley</h3>
			<Row className="align-items-center">
				<Col><p>Aliceann has been playing music and partner dancing for most of her life, and is always inspired in new ways by the music.  Believing the music is what drives these dances, she continues to dive into the history of vintage jazz and blues to collect new music that inspires dancers to move. She keeps two things in mind while DJing - there should be enough variation that everyone will want to dance throughout the night, and everyone should have enough time to finish hugging before the next song starts.</p></Col>
				<Col className="col-4"><img src="/images/DJs/Aliceann.jpg" width="100%" object-fit="scale-down"/></Col>
			</Row>
			<h3>Kemper Talley</h3>
			<Row className="align-items-center">
				<Col className="col-4"><img src="/images/DJs/Kemper.jpg" width="100%" object-fit="scale-down"/></Col>
				<Col><p>Kemper favors music that has authentic, historic roots in the Swing Era. He loves to play music from swing's early days to modern recreations, but heavily favors the small combos and big bands of the 1930s and 40s. Outside of DJing, he enjoys leading his classic swing small combo, the Jump Shop Sextet, as well as teaching local and regional dancers more about this thing we call swing.</p></Col>
			</Row>
		</>}
	</Page>
)

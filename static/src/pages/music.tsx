import React from 'react'
import Page from '../components/Page'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

enum Side {
	LEFT,
	RIGHT,
}

type DJ = {
	name: string,
	picture: string,
	bio: string,
}

const DJCard: React.FC<{
	dj: DJ,
	side: Side,
}> = ({ dj, side }) => {
	const picture = <Col className="col-4"><img src={dj.picture} width="100%" object-fit="scale-down" /></Col>
	const bio = <Col>{dj.bio.split("\n\n").map(line => (<p>{line}</p>))}</Col>

	return (
		<>
			<h2>{dj.name}</h2>
			<Row className="align-items-center">
				{(() => {
					switch (side) {
						case Side.LEFT:
							return (<>{picture}{bio}</>)
						case Side.RIGHT:
							return (<>{bio}{picture}</>)
					}
				})()}
			</Row>
		</>
	)
}

export default () => {
	const djs = [
		{
			name: "Mark Calkins",
			picture: "/images/DJs/mark.jpg",
			bio: `Mark’s DJing was once described as wonderfully whimsical (or tipsy) by a Yale faculty member. With 20 years of experience as a Swing DJ he strives to provide the best virtuosos and deep-cut selections to inspire dancers on and off the floor. He is excited to be part of the Dayton Swing Smackdown DJ team, this year as one of the competition DJs!`,
		}, {
			name: "Shelby Johnson",
			picture: "/images/DJs/shelby.jpg",
			bio: `Shelby Johnson is a self-proclaimed music snob, who fell in love with jazz at an early age listening to the sounds of Louis Armstrong, Duke Ellington and Count Basie. He really fell in love with swing music in the 6th grade thanks to a music appreciation class, and later started playing jazz in the 8th grade after 3 years in band as a drummer and percussionist. In high school, Shelby started playing the vibraphone and would transpose his parts for most pieces, as many pieces didn't provide written music. This developed Shelby's ear for music that was a huge catalyst for him catching the swing dance bug in 2011 in Orlando, Florida. And BOI, what kind of journey that has started! 

2 years later, Shelby became one of the premier DJs in the Southeast, being featured at exchanges in Florida and South Carolina. In 2013 Shelby moved to Dallas, TX, where he continues to deliver awesome syncopated tunes for Lindy Hop, Charleston, Balboa and Collegiate Shag. Shelby's passion for social dancing, teaching and competing landed him the 2016 Frankie Manning Ambassador Scholarship to Herrang Dance Camp where he expanded his knowledge on Jazz Dance and Music History. Now Shelby is back in Columbus, OH ready to continue bringing people groovy and rhythmic tunes to dance their best to.`
		}, {
			name: "Rachel Kreher",
			picture: "/images/DJs/rachel.jpg",
			bio: `Rachel fell in love with swing dancing at her very first lesson in 2012, and has since spent incalculable weekends traveling to learn, compete, and meet new friends from all over the country. Best described as an exuberant, not-so-mildly awkward teacher, Rachel loves to share her passion for dance with anyone willing to learn. As a DJ, she loves to play the jump blues vibes of her hometown right alongside her favorite Count Basie classics while highlighting the modern bands you can still see live at events. Whatever it is -- you can count on her to bring the energy to the floor!`
		}, {
			name: "Aliceann Talley",
			picture: "/images/DJs/aliceann.png",
			bio: `Aliceann has been playing music and partner dancing for most of her life, and is always inspired in new ways by the music.  Believing the music is what drives these dances, she continues to dive into the history of vintage jazz and blues to collect new music that inspires dancers to move. 

She keeps two things in mind while DJing - there should be enough variation that everyone will want to dance throughout the night, and everyone should have enough time to finish hugging before the next song starts.`,
		}, {
			name: "Kemper Talley",
			picture: "/images/DJs/kemper.png",
			bio: `Kemper favors music that has authentic, historic roots in the Swing Era. He loves to play music from swing's early days to modern recreations, but heavily favors
the small combos and big bands of the 1930s and 40s. Outside of DJing, he enjoys leading his classic swing small combo, the Jump Shop Sextet, as well as teaching local and regional dancers more about this thing we call swing.`
		}, {
			name: "Sam Zorich",
			picture: "/images/DJs/sam.jpg",
			bio: `Sam Zorich is thrilled to be part of the DJ team at Dayton this year! She began dancing in college, but fell in love with swing era and later jazz music in middle school when she started playing saxophone. Since then, she has DJd at home in Chicago and at regional lindy hop (and balboa!) events around the Midwest. While she's not much of a musician anymore, she loves to highlight the saxophone while playing music that makes people get on the floor and stay there!`,
		}
	]

	return (
		<Page title="Music">
			{djs.map((dj, index) => (
				<DJCard dj={dj} side={index % 2 == 0 ? Side.LEFT : Side.RIGHT} />
			))}
		</Page>
	)
}

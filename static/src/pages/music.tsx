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
	/*const djs = [
		{
			name: "Mark Calkins",
			picture: "/images/DJs/mark.jpg",
			bio: `Mark’s DJing was once described as wonderfully whimsical (or tipsy) by a Yale faculty member. With 20 years of experience as a Swing DJ, he strives to provide the best virtuosos and deep-cut selections to inspire the dancers on the dance floor. He is thrilled to be part of the DJ team once again and is no stranger to Dayton Swing Smackdown as he has served as DJ, judge, and instructor for many years and returns this year as a co-choreographer along with Shannon Varner for Team SwingColumbus! Cbus Ho!`
		}, {
			name: "Naomi Carmeli-Shabtay",
			picture: "/images/DJs/naomi.jpg",
			bio: `Naomi Carmeli-Shabtay began dancing lindy hop and blues in 2014. Since 2018 she has DJed at events around the Midwest from Milwaukee to St Louis, including at Whistle Stop & CincyHop, and appears regularly at dances in Chicago. 

At live music events Naomi’s been known to say “I love this song!!” roughly every other song. It’s always true, and as a DJ she strives to elicit the same response from the dance floor. Naomi has a love for vocals, guitar, and pulling out an oddball song or two that you might just fall in love with.`
		}, {
			name: "Emma Durham",
			picture: "/images/DJs/emma.jpg",
			bio: `Emma Durham started swing dancing at a college club at Savannah College of Art & Design in Georgia in 2008 and fell in love with the movement and then the music. She continued while living in Atlanta and New Orleans before moving back to her hometown of Cincinnati, where she started DJing at local and regional events, including Smackdown, School of Hard Knox, CincyHop, and LaFLX. She has a soft spot for drummer-led bands, female vocalists, and anything that really makes her want to swing out. If she’s at the DJ table, chances are good there’s a Pomeranian named Cheezit on the table too.`
		}, {
			name: "Shelby Johnson",
			picture: "/images/DJs/shelby.jpg",
			bio: `Shelby Johnson is a self-proclaimed music snob, who fell in love with jazz at an early age listening to the sounds of Louis Armstrong, Duke Ellington and Count Basie. He really fell in love with swing music in the 6th grade thanks to a music appreciation class, and later started playing jazz in the 8th grade after 3 years in band as a drummer and percussionist. In high school, Shelby started playing the vibraphone and would transpose his parts for most pieces, as many pieces didn't provide written music. This developed Shelby's ear for music that was a huge catalyst for him catching the swing dance bug in 2011 in Orlando, Florida. And BOI, what kind of journey that has started! 

2 years later, Shelby became one of the premier DJs in the Southeast, being featured at exchanges in Florida and South Carolina. In 2013 Shelby moved to Dallas, TX, where he continues to deliver awesome syncopated tunes for Lindy Hop, Charleston, Balboa and Collegiate Shag. Shelby's passion for social dancing, teaching and competing landed him the 2016 Frankie Manning Ambassador Scholarship to Herrang Dance Camp where he expanded his knowledge on Jazz Dance and Music History. Now Shelby is back in Columbus, OH ready to continue bringing people groovy and rhythmic tunes to dance their best to.`
		}, {
			name: "Rachel Kreher",
			picture: "/images/DJs/rachel.jpg",
			bio: `Rachel fell in love with swing dancing at her very first lesson in 2012, and has since spent incalculable weekends traveling to learn, compete, and meet new friends from all over the country. Best described as an exuberant, not-so-mildly awkward teacher, Rachel loves to share her passion for dance with anyone willing to learn. As a DJ, she loves to play the jump blues vibes of her hometown right alongside her favorite Count Basie classics while highlighting the modern bands you can still see live at events. Whatever it is -- you can count on her to bring the energy to the floor!`
		}]*/

	return (
		<Page title="Music">
			<p>DJs and Bands will be announced soon!</p>
		</Page>
	)
}

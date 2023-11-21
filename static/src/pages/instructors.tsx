import React from 'react'
import Page from '../components/Page'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

type Instructor = {
	name: string,
	picture: string,
	bio: string,
}

const InstructorsCards: React.FC<{
	instructor1: Instructor,
	instructor2: Instructor,
}> = ({ instructor1, instructor2 }) => (
	<>
		<h2>{instructor1.name} & {instructor2.name}</h2>
		<Row className="align-items-center">
			<Col className="col-4"><img src={instructor1.picture} width="100%" object-fit="scale-down" /></Col>
			<Col>{instructor1.bio.split("\n\n").map(line => (<p>{line}</p>))}</Col>
		</Row>
		<Row className="align-items-center">
			<Col>{instructor2.bio.split("\n\n").map(line => (<p>{line}</p>))}</Col>
			<Col className="col-4"><img src={instructor2.picture} width="100%" object-fit="scale-down" /></Col>
		</Row>
	</>
)


const Instructors = () => {
	const instructors = [
		[{
			name: "Kate Hedin (She/Her)",
			picture: "/images/instructors/Kate Hedin.jpg",
			bio: `Following in the path of her great ancestor, Swedish explorer Sven Hedin, Kate has traveled to the greatest unknown regions of following and footwork. Aside from a love of classically inspired Lindy Hop, she is most known for her unique and elegant style of Balboa.

The result is an impressive resume. Kate holds championship titles in almost every major competition, including the American Classic Balboa Championships, the International Lindy Hop Championships, the Euro Balboa Cup, and more. She is also a highly sought-after competition judge. In 2012, she released two highly praised DVDs specifically for Balboa followers on technique and aesthetics. As a teacher, Kate is known and loved for her logic-based language, her unique methods of teaching technique, and, overall, never leaving student followers disappointed.

She is proud to have held, along with her friends at Get Hep Swing in Cleveland, the Guinness World Record for longest dance party ever (52 hours, 3 minutes). Her fashion is admired by many, her lines desired as far as the Orient. She also does trapeze.`
		}, {
			name: "Bobby White (He/Him)",
			picture: "/images/instructors/Bobby White.jpg",
			bio: `Bobby teaches traditional swing dances around the world, and holds championship titles and placements in Balboa, Lindy Hop, and Solo Jazz. With each, Bobby strives to innovate and create a new voice, while still capturing the spirit of the original dancers. He is the co-director of the Lindy Hop performance group the Harvest Moon Hoppers, which specialize in dancing in the performance style of Whitey's Lindy Hoppers. He is the author of the popular swing dance blog Swungover*, and the book "Practice Swing."  As a dancer, he is widely recognized for his floppy hair, and as a teacher, by his sound effects.`,
		}],
		[{
			name: "Dave Barry (He/Him)",
			picture: "/images/instructors/David Barry.jpg",
			bio: `Dave hasn't said no to a dance since finding Lindy Hop in 2012 (seriously). As a teacher, Dave focuses on simple, driving rhythms and deep applicability of foundational concepts. His classes place high value on dynamic following, and inclusive listening from leaders with body mechanics always at the root of his lead/ follow focus. Musically driven with an expansive knowledge of early jazz, RnB, and blues, Dave’s classes incorporate structural musical concepts into every step. Dave has won or placed in major competitions at national and international events including ILHC, Lindyfest, ULHS, Midwest Lindyfest, Arctic Lindy Exchange, and Atlanta Varsity Showdown. At home in Atlanta Dave gives warm homes to wayward CDs and records.`
		}, {
			name: "Mandy Hogan (She/Her)",
			picture: "/images/instructors/Mandy Hogan.jpg",
			bio: `Mandy Hogan began her swing dancing journey at the Anderson University’s swing club in 2007. Over the years, she has traveled all over the Midwest and beyond attending workshops, social dancing, performing, competing, and teaching. On the social dance floor, she is known for her connection, playful energy, and creativity. As an instructor, she currently specializes in Lindy hop, Charleston, vernacular solo jazz and St. Louis Shag.`,
		}],
		[{
			name: "Christina Cacciatore (She/Her/Hers)",
			picture: "/images/instructors/Christina Cacciatore.jpg",
			bio: `Christina's dance adventure started in her Nani and Papa's kitchen long before she can remember. When she found Lindy Hop (again) in college, she found joy in connecting with her peers and more deeply with her family history. In the 12 years that have passed since her first formal lesson, Christina has continued to find connection and belonging in the dance - as well as a true love for teaching. Christina's primary focus is in teaching first time dancers and inspiring them to learn, explore, and play more in the dance. Her primary focus as a dancer is creating conversation within partnered dance and moving authentically while dancing solo. Let's explore and play in class!`
		}, {
			name: "Jony Navarro (He/They)",
			picture: "/images/instructors/Jony Navarro.jpg",
			bio: `Jony fell in love with jazz in early 2011 and has since never been the same.  He has a constant drive to make art completely his own, which is fueled by a desire for learning.  His fascination began with dancing, moved onto jazz music, then continued to black history and culture and beyond.

From the very beginning, Jony has always had very unique ideas when it comes to making art.  He excels at inspiring others to think about dance in a different way than they’re used to, often expanding on familiar topics.  While conceptual learning drives a lot of his classes, movement and technique play a large part in the learning environment he creates.

Solo movement of every kind has been a staple in Jony’s dancing since they began.  Their first love was Charleston, then grew into more of the solo jazz dances.  Their love for movement did not stop at jazz, They also cross-trained in tap, ballet, modern, salsa, bachata, and hip-hop, mostly in college courses.  Jony encourages everyone to come get weird with him as we explore all the possibilities jazz has to offer.`,
		}]
	]

	return (
		<Page title="Instructors">
			{instructors
				.map(([instructor1, instructor2]) =>
					(<InstructorsCards instructor1={instructor1} instructor2={instructor2} />)
				)
			}
		</Page >
	)
}

export default Instructors

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
			name: "Irina Amzashvili",
			picture: "/images/instructors/Irina Amzashvili.jpg",
			bio: `Dancing in all forms has been a lifelong passion for Irina (she/her). In 2009, at the age of 15, she stumbled upon Swing dancing and it took only one night of social dancing for her to fall head over heels in love with it. Though Lindy Hop was her gateway into the world of Swing, Irina's passion quickly expanded to include Balboa, Collegiate Shag, and solo jazz. She is proud to have spent her formative years of dancing in Southern California, home to some of the best Swing dancers, teachers, and music in the world. Irina has since received numerous accolades from events such as Camp Hollywood, The International Lindy Hop Championships, and Lindy Focus and has taught these dances all across the U.S., Asia, Europe, and Australia. Additionally, she is honored to have been inducted into the Camp Hollywood Hall of Fame in 2019 and the California Swing Dance Hall of Fame in 2022.

Irina's teaching style can best be described as patient and encouraging. Her goal is to guide her students through building solid foundations while inspiring individuality in personal movement, as well as helping build confidence through dance.

Photo by Jerry Almonte (http://jsalmonteproductions.com/)`,
		}, {
			name: "Rafal Pustelny",
			picture: "/images/instructors/Rafal Pustelny.jpg",
			bio: `Rafal Pustelny is a vibrant dancer and instructor who is passionate about traditional dances from all around the world! Rafal moved to NYC in 2010 and there studied American Jazz dance forms, including Lindy Hop, Charleston, and Balboa. He quickly established himself as one of the top dancers in NYC, and has since earned acclaim in the international swing dance community.

His wide knowledge and skill in various styles has helped him win top placements in international competitions and has created opportunities for him to teach around the world. Rafal’s expertise and passion for dance shine in the classroom; he uses his technical prowess, upbeat sense of humor, and innate ability to connect with people to distill high-level content into accessible pieces that elevate multiple aspects of his students’ dancing.

Rafal continues his own dance education by studying Tap, Clogging, Hungarian folk dances, and Polish folk dances. As a performer, he works with several regional groups, including “Gatsby Entertainment,” “The Vanaver Caravan,” “Project Connect”, and “Harvest Moon Hoppers”, as well as community events and corporate parties. A Frankie Manning Foundation Ambassador, he teaches and performs locally with “You Should Be Dancing” studios. You can also find him at international events where he works with dance communities across the globe.`
		}],
		[{
			name: "Ana Lisa Sutherland",
			picture: "/images/instructors/Ana Lisa Sutherland.jpg",
			bio: `Ana Lisa Sutherland (she/they) is a powerhouse of energy packed into a short frame, and she brings it all to the dance floor! She began her dance journey early in life in California, where she competed in Ballroom and Hip Hop before (accidentally) falling in love with Lindy Hop. She started as a scrappy newbie pestering her dance idols for knowledge and grew to a scrappy quirky professional and wild card in the dance scene. Since 2008, she has traveled across the country to compete and learn swing dancing from original legends and current-day champions. Ana Lisa holds multiple titles from prestigious weekend events like Northwest Balboa Fest, Lindy Bout, Camp Jitterbug, Lindy Focus, and Lindy On The Rocks. Some of her favorite performance highlights include Lincoln Center’s Midsummer Night Swing and local cabaret shows in her current city of Seattle.

Ana Lisa believes that swing dancing is a cultural experience that allows participants to connect with the roots of this African American historical dance, contribute to its legacy, and shine. There truly is a place for everyone within it. Consequently, she loves drawing inspiration from historical sources and the legends that have paved the way, while also incorporating newer dance forms from her ongoing explorations in the world of dance. She is a competitor, a performer, a budding street dancer, a retired Latin ballroom champion, a showgirl...she is a kaleidoscope of the swing dance world.

In the classroom, she strives to inspire other dancers to authentically express themselves on the floor and encourages everyone to embrace their uniqueness. Ana Lisa helps her students approach each dance and class with the mindset that anything could happen, unlocking the unique conversation they can have with their partner and the music.`
		}, {
			name: "Jon Tigert",
			picture: "/images/instructors/Jon Tigert.jpg",
			bio: `Jon Tigert is a world recognized Swing Dancer and musician. Specializing in Vintage dances such as Lindy Hop, Balboa, Blues, and Charleston, he has won titles from the International Lindy Hop Championships, The Canadian Swing Dance Championships, Lindy Focus,  and many more. Swing Dancing since 2005, Jon has made a home in the swing community as a teacher, an MC, a performer, an organizer, and a musician.  As a teacher, Jon is well known for his historical and musical knowledge, clear and concise communication, and endless supply of terrible dad jokes.

Jon now lives in Baltimore where he teaches at the Mobtown Ballroom and sings and plays drums with The Corner Pocket Jazz Band.`
		}],
		[{
			name: "Doug Bae",
			picture: "/images/instructors/Doug Bae.jpg",
			bio: `Doug Bae (he/him) stumbled into Lindy Hop in 2016 and it quickly became an enormous part of his life, traveling across the country to learn, teach, and dance! Now in Pittsburgh he is an organizer for PittStop as well the local scene as Doug loves nothing more than to share the joy he experiences while dancing. As a teacher he focuses on connection and energy, responding to your partner, and delving into core body mechanics in both solo and partnered movement.

On the social floor (where you’ll most easily find him), Doug is constantly exploring playful, musical, and comfortable dances with anyone and everyone. He’s also an avid baker and is always interested to hear good bakery or recipe recommendations!`
		}, {
			name: "Val Kutchko",
			picture: "/images/instructors/Val Kutchko.jpg",
			bio: `Val Kutchko (they/she) started swingin' out 9 years ago in jazzy Kansas City. Val quickly became enamored by the energy and rhythms of jazz music and Lindy Hop. Val has since moved to Pittsburgh, PA, where they teach, organize, and DJ, sometimes playing multiple versions of a song in the same night to see how folks' dancing changes with each version.

Val delights in both following and leading, and their favorite dances are full of laughter. Informed by their enjoyment of both roles, as an instructor Val emphasizes communication, comfort, and playfully disrupting gendered dynamics of leading and following.`
		}]
	]
	return (
		<Page title="Instructors">
			{instructors.map(instructorPair => (
				<InstructorsCards instructor1={instructorPair[0]} instructor2={instructorPair[1]} />
			))}
		</Page >
	)
}

export default Instructors

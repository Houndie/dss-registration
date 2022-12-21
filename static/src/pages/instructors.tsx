import React from 'react'
import Page from '../components/Page'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

const Instructors = () => (
	<Page title="Instructors">
		{() => (
			<>
				<h2>Natalia Eristavi & Christian Frommelt</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/Natalia.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col><p>Natalia (she/her) is a visual artist, jazz dancer, teacher and performer based in San Diego, California. Her swing and jazz adventures began in late 2010 although she had been dancing since the age of 4 and considers movement to music akin to breathing. Since her first steps into the professional swing dance circuit around 2012, she has taught, competed and performed at many notable national and international dance events, acquiring numerous awards along the way.</p>
					<p>While performing and teaching are a great source of passion for her, Natalia feels the most fulfilled when creating meaningful connections with her students and colleagues on the social dance floor.</p></Col>
				</Row>
				<Row className="align-items-center">
					<Col><p><b>Hello, I’m Christian Frommelt</b> (he/him/his), a swing dancer and musician from St. Louis, Missouri. As a teacher, organizer, and social dancer, I’ve dedicated much of my life to bringing people together through music and dance. The African-American vernacular dance values of spontaneity, rhythmic propulsion, and communal improvisation captivated me as a teenager, and I’ve been bent on learning jazz-based artforms ever since. It has become a privilege to also share the stories, sounds, and skills I’ve gained along the way as a teacher, writer, and performer.</p> 
					<p>I’m proud to have been recognized with competitive accolades at the Ultimate Lindy Hop Showdown and the International Lindy Hop Championships, which include awards for choreography with Jenny Shirar. For nine years I served as co-organizer of the renowned Nevermore Jazz Ball & St. Louis Swing Dance Festival and Cherokee Street Jazz Crawl in the effort of collective expression. I’ve played piano for the Gaslight Squares jazz band since 2019. </p></Col>
					<Col className="col-4"><img src="/images/instructors/Christian.jpg" width="100%" object-fit="scale-down"/></Col>
				</Row>
				<h2>Kerry Kapaku & Shannon Varner</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/Kerry.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col><p>Kerry Kapaku comes from a long history of training, performing, and teaching ballet and contemporary dance professionally, and holds a BFA in Dance Performance and Education from The Ohio State University. Kerry was introduced to the wonderful world of lindy hop while attending OSU, and has been teaching, competing, and performing across the US since. In the classroom, Kerry strives to cultivate an upbeat and encouraging environment for all of her students, while emphasizing the importance of efficient muscular engagement, physical and musical connection, and self expression.  Kerry holds competitive titles from events such as ILHC, Nevermore Jazz Ball, and Lindy Focus, Camp Hollywood, and more. At home, Kerry serves as an instructor for Naptown Stomp, and is the owner and director of DanceWorks Indy, Indianapolis's only dance and fitness studio for adults of all ages and abilities.</p></Col>
				</Row>
				<Row className="align-items-center">
					<Col><p>Shannon Varner discovered Lindy Hop while living in New York City ages ago. Since then she been fortunate to live in Lindy Hop-centric cities ever since. Whether Washington DC, Chicago, or Columbus, OH, they all have provided rich opportunities to hone her dance skills, and influence her dance styles. She enjoys sharing her passion for these dances with others and truly believes dancing is for everyone.</p>
					<p>Shannon has won numerous awards, placing in such competitions as the American Lindy Hop Championships, Beantown Dance Camp, Nevermore Jazz Ball, SwingIN, Southern Swing Challenge, and Midsummer Night Swing in NYC.</p>
					<p>Shannon also co-founded the award-winning competitive dance team SwingColumbus and served as co-choreographer from 2008 – 2017.</p></Col>
					<Col className="col-4"><img src="/images/instructors/Shannon.png" width="100%" object-fit="scale-down"/></Col>
				</Row>
				<h2>Dave Barry & Mimi Liu</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/Dave.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col>
						<p>Dave hasn't said no to a dance since finding Lindy Hop in 2012 (seriously). As a teacher, Dave focuses on simple, driving rhythms and deep applicability of foundational concepts. His classes place high value on dynamic following, and inclusive reactions from leaders with body mechanics always at the root of his lead/ follow focus. At home in Atlanta Dave gives warm homes to wayward CDs and records.</p>
					</Col>
				</Row>
				<Row className="align-items-center">
					<Col>
						<p>Ever since she was a wee baby in a stroller, Mimi has been moving and bopping to the beats, but it took until 2010 to find her true dance love, Lindy Hop, in a riverside bar in Cambodia. She loves playful and musical dances and exploring weird ideas. Her favorite part of teaching is helping students, especially followers, get creative on the dance floor while being grounded on solid and comfy partner connection.</p>
					</Col>
					<Col className="col-4"><img src="/images/instructors/Mimi.jpg" width="100%" object-fit="scale-down"/></Col>
				</Row>
			</>
		)
		}
	</Page>
)

export default Instructors

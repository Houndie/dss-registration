import React from 'react'
import Page from '../components/Page'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

const Instructors = () => (
	<Page title="Instructors">
		{() => (
			<>
				<h2>Gaby Cook & Jon Tigert</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/Gaby.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col><p>Gaby Cook is an esteemed, active professional in the global lindy hop community. For nearly 20 years, she has established her career, teaching and performing worldwide for events such as Herräng Dance Camp, Paris Jazz Roots and Lindy Focus. In the classroom, she is playful, informative and honest. She prioritizes equality for followers and leaders in the classroom space — and has championed a movement first approach to teaching partnered dance content.</p></Col>
				</Row>
				<Row className="align-items-center">
					<Col><p>Jon Tigert found swing dancing at the tender age of 15 and never looked back. Nowsomewhat older than 15, Jon shares his knowledge of the dance all over the United States and the World. You may know him from his roles at events such as Lindy Focus, Lindy Fest, The Canadian Swing Dance Championships, The Chinese Lindy Hop Championships, Herrang Dance Camp, The International Lindy Hop Championships and many more. Known for his concise language, historical knowledge and and infinite supply of dad jokes, Jon's classes are immensely informative and entertaining. Beyond teaching, Jon is often seen behind the microphone as a well known MC and host, behind the computer as a DJ, or behind the drum kit playing and singing with various musical groups, including the Corner Pocket Jazz Band!</p></Col>
					<Col className="col-4"><img src="/images/instructors/Jon.jpg" width="100%" object-fit="scale-down"/></Col>
				</Row>
				<h2>Jon Holmstrom & Kerry Kapaku</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/John.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col><p>John Holmstrom discovered Lindy Hop in Chicago in 2005. After college in Champaign-Urbana where he served as President of the Illini Swing Society, he returned to Chicago in 2010 and fully immersed himself in the Lindy Hop community teaching at Big City Swing, DJing, and organizing, founding Traffic Jam Swing and Second City Swingout. John has performed at various events for muggles and Lindy Hoppers alike from historic Chicago venues like Ravinia and the Double Door to Lindy Focus. John is an avid team competitor, coach and choreographer with wins at Dayton Swing Smackdown (2010 - Hoosier Hotshots), Hawkeye Swing Festival, Southern Swing Challenge, and Camp Jitterbug. As a teacher, John loves to focus on the technical elements in his classes, giving students similar moves with crucial differences for a choose your own adventure experience.</p></Col>
				</Row>
				<Row className="align-items-center">
					<Col><p>Kerry Kapaku comes from a long history of training, performing, and teaching ballet and contemporary dance professionally, and holds a BFA in Dance Performance and Education from The Ohio State University. Kerry was introduced to the wonderful world of lindy hop while attending OSU, and has been teaching, competing, and performing across the US since. In the classroom, Kerry strives to cultivate an upbeat and encouraging environment for all of her students, while emphasizing the importance of efficient muscular engagement, physical and musical connection, and self expression.  Kerry holds competitive titles from events such as ILHC, Nevermore Jazz Ball, and Lindy Focus, Camp Hollywood, and more. At home, Kerry serves as an instructor for Naptown Stomp, and is the owner and director of DanceWorks Indy, Indianapolis's only dance and fitness studio for adults of all ages and abilities.</p></Col>
					<Col className="col-4"><img src="/images/instructors/Kerry.jpg" width="100%" object-fit="scale-down"/></Col>
				</Row>
				<h2>Aliceann Talley & Kemper Talley</h2>
				<Row className="align-items-center">
					<Col className="col-4"><img src="/images/instructors/Aliceann.jpg" width="100%" object-fit="scale-down"/></Col>
					<Col>
						<p>Aliceann is as passionate about her students as she is about her dancing, and is constantly working on new ways to help them learn - whether through teaching weekly lessons, hosting house practices, or coaching/choreographing for her performance team.  She believes learning to dance should be fun, and that you don't need to master a particular set of moves in order to become a swing dancer.</p>
						<p>Aliceann started dancing in 2001 before eventually falling in love with swing dancing, Carolina shag, and other vernacular dances.  She strives to grow swing dancing on and off the dance floor through her classes, organizing lessons in Knoxville, and her event, Hard Knox.</p>
					</Col>
				</Row>
				<Row className="align-items-center">
					<Col>
						<p>Kemper has been studying and dancing Lindy Hop since the summer of 2011. He is as passionate about the music and history of the swing era as he is about the dance, and his passion for teaching comes from his lifelong obsession with learning and growing. Kemper believes dancing is a form of personal expression, and hopes to help others better express themselves by empowering them with technical skills and providing an environment of exploration.</p>
						<p>Outside of teaching, you can find Kemper behind the DJ booth, organizing events, or leading his swing band, the Jump Shop Sextet.</p>
					</Col>
					<Col className="col-4"><img src="/images/instructors/Kemper.jpg" width="100%" object-fit="scale-down"/></Col>
				</Row>
			</>
		)
		}
	</Page>
)

export default Instructors

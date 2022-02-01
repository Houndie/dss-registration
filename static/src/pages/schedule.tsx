import React from 'react'
import Page from '../components/Page'

export default () => (
	<Page title="Schedule">
		{ () => (
			<>
			<style>{`
				table, tr, td {
					border: 1px solid black;
					padding: 2px;
				}
			`}</style>
			<h3>Friday Evening</h3>
			<table>
				<tr><td>8:00p</td><td>Doors Open at the Baum Opera House</td></tr>
				<tr><td>8:00p</td><td>DJed Music</td></tr>
				<tr><td>9:00p</td><td>Live Band: Dave Greer and His Classic Jazz Stompers</td></tr>
				<tr><td>3rd Band Break (est 11:00p)</td><td>Solo Jazz Prelims and Finals</td></tr>
				<tr><td>12:00a</td><td>End of dance...everyone go to late night!</td></tr>
			</table>
			<br/>
			<h3>Friday Late Night</h3>
			<table>
				<tr><td>12:00a</td><td>Doors Open at Elegance in Dance</td></tr>
				<tr><td>12:00a</td><td>DJed Music</td></tr>
				<tr><td>3:00a</td><td>End of dance...get rest before tomorrow!</td></tr>
			</table>
			<br/>
			<h3>Saturday Afternoon</h3>
			<table>
				<tr><td>9:00a</td><Tdc colSpan={3}>Doors Open at The Baum Opera House</Tdc></tr>
				<tr><td>9:20a</td><Tdc colSpan={3}>Level Test</Tdc></tr>
				<tr><td></td><Tdc><b>Level 1 (Baum Opera House Upstairs)</b></Tdc><Tdc><b>Level 2 (Baum Opera House Downstairs)</b></Tdc><Tdc><b>Level 3 (B&B Riverfront Hall)</b></Tdc></tr>
				<tr><td>10:00a</td><Tdc>John & Kerry</Tdc><Tdc>Gaby & Jon</Tdc><Tdc>Aliceann & Kemper</Tdc></tr>
				<tr><td>11:00a</td><Tdc colSpan={3}>Midmorning Break</Tdc></tr>
				<tr><td>11:15a</td><Tdc>Gaby & Jon</Tdc><Tdc>Aliceann & Kemper</Tdc><Tdc>John & Kerry</Tdc></tr>
				<tr><td>12:15p</td><Tdc colSpan={3}>Lunch Break</Tdc></tr>
				<tr><td>1:30p</td><Tdc>Aliceann & Kemper</Tdc><Tdc>John & Kerry</Tdc><Tdc>Gaby & Jon</Tdc></tr>
				<tr><td>2:30p</td><Tdc colSpan={3}>Afternoon Break 1</Tdc></tr>
				<tr><td>2:45p</td><Tdc>Aliceann & Kemper</Tdc><Tdc>John & Kerry</Tdc><Tdc>Gaby & Jon</Tdc></tr>
				<tr><td>3:45p</td><Tdc colSpan={3}>Afternoon Break 2</Tdc></tr>
				<tr><td></td><Tdc><b>Elective (Baum Opera House Upstairs)</b></Tdc><Tdc><b>Elective (Baum Opera House Downstairs)</b></Tdc><Tdc><b>Elective (B&B Riverfront Hall)</b></Tdc></tr>
				<tr><td>4:00p</td><Tdc>Gaby & Jon</Tdc><Tdc>John & Kerry</Tdc><Tdc>Aliceann & Kemper</Tdc></tr>
				<tr><td>5:00p</td><Tdc colSpan={3}>Dinner Break</Tdc></tr>
				<tr><td>5:00p</td><Tdc colSpan={3}>Team Floor Trials (Ballroom Closed)</Tdc></tr>
			</table>
			<br/>
			<h3>Saturday Evening</h3>
			<table>
				<tr><td>7:00p</td><td>Ballroom Opens</td></tr>
				<tr><td>7:30p</td><td>Mix and Match Prelims</td></tr>
				<tr><td>8:00p</td><td>DJed Music</td></tr>
				<tr><td>9:30p</td><td>Team Competition</td></tr>
				<tr><td>11:45p</td><td>Mix and Match Finals</td></tr>
				<tr><td>12:00a</td><td>End of dance...everyone go to late night!</td></tr>
			</table>
			<br/>
			<h3>Saturday Late Night</h3>
			<table>
				<tr><td>12:00a</td><td>Doors Open at Elegance in Dance</td></tr>
				<tr><td>12:00a</td><td>DJed Music</td></tr>
				<tr><td>1:00a</td><td>Awards</td></tr>
				<tr><td>3:00a</td><td>Josh Forbes Power Hour</td></tr>
				<tr><td>4:00a</td><td>End of dance</td></tr>
			</table>
			<br/>
			<h3>Sunday Afternoon</h3>
			<table>
				<tr><td>1:00p</td><td>Doors Open at Elegance and Dance</td></tr>
				<tr><td>1:00p</td><td>Class 1 (All Levels)</td></tr>
				<tr><td>2:00p</td><td>Class 2 (All Levels)</td></tr>
				<tr><td>3:00p</td><td>DJed Music</td></tr>
				<tr><td>6:00p</td><td>End of event.  Go Home (sorry)</td></tr>
			</table>
			</>
		)}
	</Page>
) 

type Tdcprops = {
	children: React.ReactNode
	colSpan?: number
}

const Tdc = ({children, ...props}: Tdcprops) => (
	<td style={{"textAlign": "center"}} {...props}>{children}</td>
)

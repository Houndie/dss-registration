import React from 'react'
import Page from '../components/Page'

export default () => (
	<Page title="Tenative Schedule">
		<p>Details subject to change as we get closer to the event!</p>
		<style>{`
				table, tr, td {
					border: 1px solid black;
					padding: 2px;
				}
			`}</style>
		<h3>Friday Evening</h3>
		<table>
			<tr><td>8:00p</td><Tdc colSpan={2}>Doors Open at the Baum Opera House</Tdc></tr>
			<tr><td></td><Tdc><b>Baum Opera House Upstairs</b></Tdc><Tdc><b>Baum Opera House Downstairs</b></Tdc></tr>
			<tr><td>8:30p</td><Tdc>DJed Music</Tdc><td /></tr>
			<tr><td>9:00p</td><Tdc>Live Band</Tdc><Tdc rowSpan={2}>DJ'd Band Breaks</Tdc></tr>
			<tr><td>3rd Band Break (est 11:00p)</td><Tdc>Solo Jazz Prelims and Finals</Tdc></tr>
			<tr><td>12:00a</td><Tdc colSpan={2}>End of dance...everyone go to late night!</Tdc></tr>
		</table>
		<br />
		<h3>Friday Late Night</h3>
		<table>
			<tr><td>12:00a</td><td>Doors Open at Elegance in Dance</td></tr>
			<tr><td>12:00a</td><td>DJed Music</td></tr>
			<tr><td>3:00a</td><td>End of dance...get rest before tomorrow!</td></tr>
		</table>
		<br />
		<h3>Saturday Afternoon</h3>
		<table>
			<tr><td>9:00a</td><Tdc colSpan={3}>Doors Open at The Baum Opera House</Tdc></tr>
			<tr><td></td><Tdc><b>Baum Opera House Upstairs</b></Tdc><Tdc><b>Star City Social Hall</b></Tdc><Tdc><b>Baum Opera House Downstairs</b></Tdc></tr>
			<tr><td>9:20a</td><Tdc>Level 2 Test</Tdc><td /><Tdc>Level 3 Test</Tdc></tr>
			<tr><td></td><Tdc><b>Level 1 Track</b></Tdc><Tdc><b>Level 2 Track</b></Tdc><Tdc><b>Level 3 Track</b></Tdc></tr>
			<tr><td>10:00a</td><Tdc>Ana Lisa Sutherland & Jon Tigert</Tdc><Tdc>Irina Amzashvili & Rafal Pustelny</Tdc><Tdc>Doug Bae & Val Kutchko</Tdc></tr>
			<tr><td>11:00a</td><Tdc colSpan={3}>Midmorning Break</Tdc></tr>
			<tr><td>11:15a</td><Tdc>Irina Amzashvili & Rafal Pustelny</Tdc><Tdc>Doug Bae & Val Kutchko</Tdc><Tdc>Ana Lisa Sutherland & Jon Tigert</Tdc></tr>
			<tr><td>12:15p</td><Tdc colSpan={3}>Lunch Break</Tdc></tr>
			<tr><td>12:15p</td><Tdc colSpan={3}>Team Floor Trials (Ballroom Closed)</Tdc></tr>
			<tr><td>1:30p</td><Tdc>Doug Bae & Val Kutchko</Tdc><Tdc>Ana Lisa Sutherland & Jon Tigert</Tdc><Tdc>Irina Amzashvili & Rafal Pustelny</Tdc></tr>
			<tr><td>2:30p</td><Tdc colSpan={3}>Afternoon Break 1</Tdc></tr>
			<tr><td>2:45p</td><Tdc>Doug Bae & Val Kutchko</Tdc><Tdc>Ana Lisa Sutherland & Jon Tigert</Tdc><Tdc>Irina Amzashvili & Rafal Pustelny</Tdc></tr>
			<tr><td>3:45p</td><Tdc colSpan={3}>Afternoon Break 2</Tdc></tr>
			<tr><td></td><Tdc><b>Elective</b></Tdc><Tdc><b>Elective</b></Tdc><Tdc><b>Elective</b></Tdc></tr>
			<tr><td>4:00p</td><Tdc>Ana Lisa Sutherland & Jon Tigert</Tdc><Tdc>Doug Bae & Val Kutchko</Tdc><Tdc>Irina Amzashvili & Rafal Pustelny</Tdc></tr>
			<tr><td>5:00p</td><Tdc colSpan={3}>Dinner Break</Tdc></tr>
			<tr><td>5:00p</td><Tdc colSpan={3}>Team Floor Trials (Ballroom Closed)</Tdc></tr>
		</table>
		<br />
		<h3>Saturday Evening</h3>
		<table>
			<tr><td>7:00p</td><Tdc colSpan={2}>Ballroom Opens</Tdc></tr>
			<tr><td></td><Tdc><b>Baum Opera House Upstairs</b></Tdc><Tdc><b>Baum Opera House Downstairs</b></Tdc></tr>
			<tr><td>7:30p</td><Tdc>Mix and Match Prelims</Tdc><td /></tr>
			<tr><td>8:00p</td><Tdc rowSpan={2}>DJed Music</Tdc><Tdc>DJed Music</Tdc></tr>
			<tr><td>9:00p</td><Tdc rowSpan={2}>Closed for Team Competition Green Room</Tdc></tr>
			<tr><td>9:30p</td><Tdc>Team Competition</Tdc></tr>
			<tr><td>After Team Competition (est 10:30)</td><Tdc>DJed Music</Tdc><Tdc rowSpan={2}>DJed Music</Tdc></tr>
			<tr><td>11:40p</td><Tdc>Mix and Match Finals</Tdc></tr>
			<tr><td>12:00a</td><Tdc colSpan={2}>End of dance...everyone go to late night!</Tdc></tr>
		</table>
		<br />
		<h3>Saturday Late Night</h3>
		<table>
			<tr><td>12:00a</td><td>Doors Open at Elegance in Dance</td></tr>
			<tr><td>12:00a</td><td>DJed Music</td></tr>
			<tr><td>1:00a</td><td>Awards</td></tr>
			<tr><td>3:00a</td><td>Josh Forbes Power Hour</td></tr>
			<tr><td>4:00a</td><td>End of dance</td></tr>
		</table>
		<br />
		<h3>Sunday Afternoon</h3>
		<table>
			<tr><td>1:00p</td><td>Doors Open at Elegance and Dance</td></tr>
			<tr><td>1:00p</td><td>Class 1 (All Levels)</td></tr>
			<tr><td>2:00p</td><td>Class 2 (All Levels)</td></tr>
			<tr><td>3:00p</td><td>DJed Music</td></tr>
			<tr><td>6:00p</td><td>End of event.  Go Home (sorry)</td></tr>
		</table>
	</Page>
)

type Tdcprops = {
	children: React.ReactNode
	colSpan?: number
	rowSpan?: number
}

const Tdc = ({ children, ...props }: Tdcprops) => (
	<td style={{ "textAlign": "center" }} {...props}>{children}</td>
)

import React from "react"
import Page from "../components/Page"
import Table from "react-bootstrap/Table"

export default () => (
	<Page title="Pricing">
		{() => (
			<>
				<h2>Main Event Pricing</h2>
				<p>The tiered pricing structure in the below table is for a Full Weekend Pass. The Full Weekend Pass includes access to all dances and workshop classes. It does not include competition fees (see competition entry fee prices below).</p>
				<Table>
					<thead>
						<tr><th>Registration</th><th>Price</th><th>Availability</th></tr>
					</thead>
					<tbody>
						<tr><td>1st-25th Registrants</td><td>$65</td><td> </td></tr>
						<tr><td>26st-50th Registrants</td><td>$75</td><td> </td></tr>
						<tr><td>51st-75th Registrants</td><td>$85</td><td> </td></tr>
						<tr><td>76st-100th Registrants</td><td>$95</td><td> </td></tr>
						<tr><td>101+ Registrants</td><td>$105</td><td> </td></tr>
					</tbody>
				</Table>
				<h2>Other price options are as follows:</h2>
				<ul>
					<li>Dance Only Pass (no workshops): $45.00</li>
					<li>Student Discount: -$5 on a Full Weekend Pass!</li>
				</ul>
				<h2>A la Carte: Sold at the door only</h2>
				<ul>
					<li>Full Weekend Pass — $110.00</li>
					<li>Dance Only Pass — $50.00</li>
					<li>Friday Night Dance — $20.00</li>
					<li>Friday Late Night Dance — $5.00</li>
					<li>Saturday Night  and Late Night Dance — $25.00</li>
					<li>Sunday Afternoon Dance — $5.00</li>
					<li>Saturday Night Dance Spectator Pass — $5.00</li>
					<li>Individual Workshop — $10.00</li>
				</ul>
				<h2>Competitions:</h2>
				<ul>
					<li>Team Competition Entry Fee — $55.00/team. Deadline 1 February 2023</li>
					<li>Mix and Match Entry Fee — $5.00</li>
					<li>Solo Charleston Entry Fee — $5.00</li>
				</ul>
				<h2>Jackets</h2>
				<img src="/images/2023_jacket.png" height="300px"/>
				<p>We are going with jackets (instead of tshirts) this year!</p>
				<p>All jackets will be $30</p>
				<p>See style and sizing information <a href="https://www.sportswearcollection.com/p/bella/3939?site=QDXTDOOVIQ">here</a>.</p>
			</>
		)}
	</Page>
)

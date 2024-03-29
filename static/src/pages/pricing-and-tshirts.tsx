import React from "react"
import Page from "../components/Page"
import Table from "react-bootstrap/Table"

export default () => (
	<Page title="Pricing">
		<h2><a href="https://dayton-swing-smackdown-2024.dancecamps.org">Registration is open now!</a></h2>
		<h2>Main Event Pricing</h2>
		<p>The tiered pricing structure in the below table is for a Full Weekend Pass. The Full Weekend Pass includes access to all dances and workshop classes. It does not include competition fees (see competition entry fee prices below).</p>
		<p>This year we are capping our ticket sales at 200 full weekend passes!  Make sure to sign up early!</p>
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
			<li>Friday Night Dance and Late Night Dance — $25.00</li>
			<li>Saturday Night and Late Night Dance — $25.00</li>
			<li>Sunday Afternoon Dance — $5.00</li>
			<li>Saturday Night Dance Spectator Pass — $5.00</li>
			<li>Individual Workshop — $10.00</li>
		</ul>
		<h2>Competitions:</h2>
		<ul>
			<li>Team Competition Entry Fee — $55.00/team. Deadline 1 February 2024</li>
			<li>Mix and Match Entry Fee — $5.00</li>
			<li>Solo Charleston Entry Fee — $5.00</li>
		</ul>
		<h2>T-Shirts</h2>
		<img src="/images/tshirt.png" width="200px" />
		<p>All t-shirts will be $17</p>
		<h3>T-Shirt Sizing Information</h3>
		<h4>Unisex</h4>
		<p>Measurements</p>
		<Table>
			<thead>
				<tr><th> </th><th>XS</th><th>S</th><th>M</th><th>L</th><th>XL</th><th>2XL</th><th>3XL</th></tr>
			</thead>
			<tbody>
				<tr><td>Chest Width</td><td>17</td><td>18.5</td><td>20</td><td>21.5</td><td>23</td><td>24.5</td><td>26</td></tr>
				<tr><td>Total Length</td><td>27</td><td>28</td><td>29</td><td>30</td><td>31</td><td>32</td><td>33</td></tr>
			</tbody>
		</Table>
		<h4>Bella</h4>
		<p>Measurements</p>
		<Table>
			<thead>
				<tr><th></th><th>S</th><th>M</th><th>L</th><th>XL</th><th>2XL</th></tr>
			</thead>
			<tbody>
				<tr><td>Shirt Length</td><td>26 1/4</td><td>26 3/4</td><td>27 3/8</td><td>28</td><td>28 5/8</td></tr>
				<tr><td>Shirt Length Tolerance</td><td>3/8</td><td>3/8</td><td>3/8</td><td>3/8</td><td>3/8</td></tr>
				<tr><td>Shirt Width</td><td>16</td><td>16 3/4</td><td>17 3/4</td><td>18 3/4</td><td>19 3/4</td></tr>
				<tr><td>Shirt Width Tolerance</td><td>1/2</td><td>1/2</td><td>1/2</td><td>1/2</td><td>1/2</td></tr>
			</tbody>
		</Table>
	</Page>
)

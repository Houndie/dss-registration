export default (intCost) => {
	let dollar = intCost.toString()
	while(dollar.length < 3) {
		dollar = "0" + dollar;
	}
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

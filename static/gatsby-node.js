const fs = require("fs")
const path = require("path")

exports.onPostBuild = () => {
	const infoPath = "public/info"
	if (!fs.existsSync(infoPath)) {
		fs.mkdirSync(infoPath)
	}
	fs.writeFileSync(`${infoPath}/health.json`, JSON.stringify({Healthiness: "Healthy"}))
	fs.writeFileSync(`${infoPath}/version.json`, JSON.stringify({Version: process.env.GATSBY_VERSION}))
}

exports.createPages = ({actions}) => {
	actions.createPage({
		path: "/registration/:id",
		matchPath: "/registration/:id",
		component: path.resolve("./src/programmatic_pages/UserRegistration.tsx")
	})
}

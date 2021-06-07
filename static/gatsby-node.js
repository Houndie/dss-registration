const fs = require("fs")

exports.onPostBuild = () => {
	const infoPath = "public/info"
	if (!fs.existsSync(infoPath)) {
		fs.mkdirSync(infoPath)
	}
	fs.writeFileSync(`${infoPath}/health.json`, JSON.stringify({Healthiness: "Healthy"}))
	fs.writeFileSync(`${infoPath}/version.json`, JSON.stringify({Version: process.env.GATSBY_VERSION}))
}

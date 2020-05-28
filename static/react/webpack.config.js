module.exports = {
	entry: './src/registration_form.jsx',
	module: {
		rules: [
			{
				test: /\.(js|jsx)$/,
				exclude: /node_modules/,
				use: ['babel-loader']
			}
		]
	},
	resolve: {
		extensions: ['*', '.js', '.jsx']
	},
	output: {
		path: __dirname + '/output',
		publicPath: '/',
		filename: 'bundle.js'
	},
};

module.exports = {
	entry: './react/registration_form.jsx',
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
		path: __dirname + '/static/js/react',
		publicPath: '/',
		filename: 'bundle.js'
	},
};

module.exports = {
	presets: ['module:metro-react-native-babel-preset'],
	env: {
		production: {
			plugins: ['transform-remove-console'],
		},
	},
	plugins: [
		[
			'module-resolver',
			{
				alias: {
					'@berty-tech': './packages/',
				},
			},
		],
		'react-native-reanimated/plugin',
	],
}

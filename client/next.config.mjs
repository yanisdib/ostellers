/** @type {import('next').NextConfig} */
const nextConfig = {
	images: {
		remotePatterns: [
			{
				protocol: 'https',
				hostname: 'ostellers-media-storage.s3.fr-par.scw.cloud',
				port: '',
				pathname: '/content/images/**',
			},
		],
	},
};

export default nextConfig;

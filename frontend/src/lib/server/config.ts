import { dev } from '$app/environment';
import { env } from '$env/dynamic/private';

const hostName = env.DETA_SPACE_APP_HOSTNAME;
const apiKey = env.DETA_API_KEY ?? '';

const serverApiUrl = dev ? 'http://localhost:8080' : `https://${hostName}/api`;

const serverApiKey = dev ? '' : apiKey;
const fetchHeaders = {
	'x-api-key': serverApiKey
};

export { serverApiUrl, serverApiKey, fetchHeaders };
